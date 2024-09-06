package handler

import (
	"github.com/aiagt/aiagt/app/plugin/dal/db"
)

// PluginServiceImpl implements the last service interface defined in the IDL.
type PluginServiceImpl struct {
	pluginDao *db.PluginDao
}

func NewPluginService(d *db.PluginDao) *PluginServiceImpl {
	initServiceBusiness(1)

	return &PluginServiceImpl{pluginDao: d}
}
