namespace go workflowsvc

struct CallWorkflowReq {
    1: required i64 workflow_id
    2: required binary request
}

struct CallWorkflowResp {
    1: required binary response
}

service WorkflowService {
    CallWorkflowResp CallWorkflow(1: CallWorkflowReq req)
}