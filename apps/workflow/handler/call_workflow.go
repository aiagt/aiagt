package handler

import (
	"context"
	"github.com/aiagt/aiagt/apps/workflow/model"
	"github.com/aiagt/aiagt/apps/workflow/pkg/wfutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/aiagt/aiagt/pkg/workflow"
	"github.com/pkg/errors"
)

// CallWorkflow implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) CallWorkflow(ctx context.Context, req *workflowsvc.CallWorkflowReq) (resp *workflowsvc.CallWorkflowResp, err error) {
	nodes, err := s.workflowNodeDao.GetByWorkflowID(ctx, req.WorkflowId)
	if err != nil {
		return nil, bizCallWorkflow.NewErr(err).Log(ctx, "get workflow error")
	}

	wf, err := s.buildWorkflow(ctx, nodes)
	if err != nil {
		return nil, bizCallWorkflow.NewErr(err).Log(ctx, "build workflow error")
	}

	reqObj, err := workflow.NewJSONObject(req.Request)
	if err != nil {
		return nil, bizCallWorkflow.NewErr(err).Log(ctx, "decode workflow request error")
	}

	respObj, err := wf.Run(ctx, reqObj)
	if err != nil {
		return nil, bizCallWorkflow.NewErr(err).Log(ctx, "run workflow error")
	}

	respBody, err := respObj.JSON()
	if err != nil {
		return nil, bizCallWorkflow.NewErr(err).Log(ctx, "encode workflow response error")
	}

	resp = &workflowsvc.CallWorkflowResp{
		Response: respBody,
	}

	return
}

// buildWorkflow orchestrate workflows using topological sorting
func (s *WorkflowServiceImpl) buildWorkflow(ctx context.Context, nodes []*model.WorkflowNode) (*workflow.Workflow, error) {
	var (
		stack    []*model.WorkflowNode
		nodeMap  = make(map[int64]*model.WorkflowNode)
		inDegree = make(map[int64]int)
	)

	// calculate the in-degree for each node
	for _, node := range nodes {
		nodeMap[node.ID] = node
		for _, nextID := range node.NextIDs {
			inDegree[nextID]++
		}
	}

	// find all nodes with an in-degree of 0
	for _, node := range nodes {
		if inDegree[node.ID] == 0 {
			stack = append(stack, node)
		}
	}

	var (
		wf         = workflow.NewWorkflow()
		nodeLength int
	)

	for len(stack) > 0 {
		nodeLength++

		// if the number of nodes is more than the original node, a loop exists
		if nodeLength > len(nodes) {
			return nil, errors.New("there are loops in the workflow")
		}

		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// build workflow node
		n, err := s.buildWorkflowNode(ctx, currentNode)
		if err != nil {
			return nil, errors.Wrap(err, "build workflow node error")
		}

		wf.AppendNode(n)

		// subtract the in-degree of the next node by 1
		for _, nextID := range currentNode.NextIDs {
			if nextNode, ok := nodeMap[nextID]; ok {
				// in-degree minus 1
				inDegree[nextID]--

				// if in-degree is 0, add to stack
				if inDegree[nextID] == 0 {
					stack = append(stack, nextNode)
				}
			}
		}
	}

	return wf, nil
}

func (s *WorkflowServiceImpl) buildWorkflowNode(ctx context.Context, node *model.WorkflowNode) (*workflow.Node, error) {
	switch node.Type {
	case model.WorkflowNodeTypeStart:
		return workflow.NewStartNode(), nil
	case model.WorkflowNodeTypeEnd:
		return workflow.NewEndNode(node.InputMapper), nil
	}

	n := workflow.Node{
		Name:         node.Name,
		InputMapper:  node.InputMapper,
		OutputSchema: node.OutputSchema,
		BatchField:   node.BatchField,
	}

	switch node.Type {
	case model.WorkflowNodeTypeLLM:
		params := node.NodeParams.LLM

		llmModel, err := s.modelCli.GetModelByID(ctx, &base.IDReq{Id: params.ModelID})
		if err != nil {
			return nil, err
		}

		apiKey, err := s.modelCli.GetAPIKeyByModel(ctx, &modelsvc.GetAPIKeyByModelReq{ModelId: utils.PtrOf(params.ModelID)})
		if err != nil {
			return nil, err
		}

		n.Runner = workflow.NewLLMNodeRunner(
			apiKey.BaseUrl,
			apiKey.ApiKey,
			llmModel.ModelKey,
			params.SystemPrompt,
			params.UserPrompt,
			node.OutputSchema.Properties,
		)
	case model.WorkflowNodeTypePlugin:
		params := node.NodeParams.Plugin

		const maxSecrets = 100

		listSecretResp, err := s.userCli.ListSecret(ctx, &usersvc.ListSecretReq{
			Pagination: &base.PaginationReq{PageSize: maxSecrets},
			PluginId:   utils.PtrOf(params.PluginID),
		})
		if err != nil {
			return nil, errors.Wrap(err, "list secrets error")
		}

		n.Runner = wfutil.NewPluginNodeRunner(
			params.PluginID,
			params.ToolID,
			hmap.FromSliceEntries(listSecretResp.Secrets, func(t *usersvc.Secret) (string, string, bool) { return t.Name, t.Value, true }),
			s.pluginCli,
		)
	}

	return &n, nil
}
