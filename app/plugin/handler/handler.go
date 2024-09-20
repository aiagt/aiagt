package handler

import (
	"github.com/aiagt/aiagt/app/plugin/dal/db"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
)

// PluginServiceImpl implements the last service interface defined in the IDL.
type PluginServiceImpl struct {
	pluginDao *db.PluginDao
	labelDao  *db.LabelDao
	toolDao   *db.ToolDao

	userCli usersvc.Client
}

func NewPluginService(pluginDao *db.PluginDao, labelDao *db.LabelDao, toolDao *db.ToolDao, userCli usersvc.Client) *PluginServiceImpl {
	initServiceBusiness(2)

	return &PluginServiceImpl{pluginDao: pluginDao, labelDao: labelDao, toolDao: toolDao, userCli: userCli}
}
