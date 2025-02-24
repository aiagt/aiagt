package test

import (
	"github.com/aiagt/aiagt/common/tests"
	"github.com/aiagt/aiagt/kitex_gen/workflowsvc"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/aiagt/aiagt/pkg/workflow"
	"github.com/aiagt/aiagt/rpc"
	"testing"
)

var ctx = tests.InitTesting()

func TestCallWorkflow(t *testing.T) {
	input := workflow.Object{
		"query": "hello",
	}

	tests.RpcCallWrap(rpc.WorkflowCli.CallWorkflow(ctx, &workflowsvc.CallWorkflowReq{
		WorkflowId: 1,
		Request:    utils.FirstResult(input.JSON()),
	}))
}
