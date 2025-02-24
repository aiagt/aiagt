package handler

import (
	"github.com/aiagt/aiagt/apps/workflow/dal/db"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
)

// WorkflowServiceImpl implements the last service interface defined in the IDL.
type WorkflowServiceImpl struct {
	workflowDao     *db.WorkflowDao
	workflowNodeDao *db.WorkflowNodeDao
	modelCli        modelsvc.Client
	pluginCli       pluginsvc.Client
	userCli         usersvc.Client
}

func NewWorkflowServiceImpl(workflowDao *db.WorkflowDao, workflowNodeDao *db.WorkflowNodeDao, modelCli modelsvc.Client, pluginCli pluginsvc.Client, userCli usersvc.Client) *WorkflowServiceImpl {
	initServiceBusiness(6)

	return &WorkflowServiceImpl{
		workflowDao:     workflowDao,
		workflowNodeDao: workflowNodeDao,
		modelCli:        modelCli,
		pluginCli:       pluginCli,
		userCli:         userCli,
	}
}
