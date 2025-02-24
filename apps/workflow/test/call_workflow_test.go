package test

import (
	"github.com/aiagt/aiagt/kitex_gen/workflowsvc"
	"github.com/aiagt/aiagt/pkg/tests"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/aiagt/aiagt/pkg/workflow"
	"github.com/aiagt/aiagt/rpc"
	"testing"
)

func TestWorkflowServiceImpl_CallWorkflow(t *testing.T) {
	ctx := tests.InitTesting()

	input := workflow.Object{
		"query": "hello",
	}

	tests.RpcCallWrap(rpc.WorkflowCli.CallWorkflow(ctx, &workflowsvc.CallWorkflowReq{
		WorkflowId: 1,
		Request:    utils.FirstResult(input.JSON()),
	}))
}
