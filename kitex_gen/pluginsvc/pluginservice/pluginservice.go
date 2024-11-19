// Code generated by Kitex v0.10.0. DO NOT EDIT.

package pluginservice

import (
	"context"
	"errors"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreatePlugin": kitex.NewMethodInfo(
		createPluginHandler,
		newPluginServiceCreatePluginArgs,
		newPluginServiceCreatePluginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdatePlugin": kitex.NewMethodInfo(
		updatePluginHandler,
		newPluginServiceUpdatePluginArgs,
		newPluginServiceUpdatePluginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeletePlugin": kitex.NewMethodInfo(
		deletePluginHandler,
		newPluginServiceDeletePluginArgs,
		newPluginServiceDeletePluginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPluginByID": kitex.NewMethodInfo(
		getPluginByIDHandler,
		newPluginServiceGetPluginByIDArgs,
		newPluginServiceGetPluginByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPluginByKey": kitex.NewMethodInfo(
		getPluginByKeyHandler,
		newPluginServiceGetPluginByKeyArgs,
		newPluginServiceGetPluginByKeyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListPlugin": kitex.NewMethodInfo(
		listPluginHandler,
		newPluginServiceListPluginArgs,
		newPluginServiceListPluginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListPluginByTools": kitex.NewMethodInfo(
		listPluginByToolsHandler,
		newPluginServiceListPluginByToolsArgs,
		newPluginServiceListPluginByToolsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"PublishPlugin": kitex.NewMethodInfo(
		publishPluginHandler,
		newPluginServicePublishPluginArgs,
		newPluginServicePublishPluginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateTool": kitex.NewMethodInfo(
		createToolHandler,
		newPluginServiceCreateToolArgs,
		newPluginServiceCreateToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateTool": kitex.NewMethodInfo(
		updateToolHandler,
		newPluginServiceUpdateToolArgs,
		newPluginServiceUpdateToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteTool": kitex.NewMethodInfo(
		deleteToolHandler,
		newPluginServiceDeleteToolArgs,
		newPluginServiceDeleteToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetToolByID": kitex.NewMethodInfo(
		getToolByIDHandler,
		newPluginServiceGetToolByIDArgs,
		newPluginServiceGetToolByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListPluginTool": kitex.NewMethodInfo(
		listPluginToolHandler,
		newPluginServiceListPluginToolArgs,
		newPluginServiceListPluginToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AllPluginTool": kitex.NewMethodInfo(
		allPluginToolHandler,
		newPluginServiceAllPluginToolArgs,
		newPluginServiceAllPluginToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListPluginLabel": kitex.NewMethodInfo(
		listPluginLabelHandler,
		newPluginServiceListPluginLabelArgs,
		newPluginServiceListPluginLabelResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CallPluginTool": kitex.NewMethodInfo(
		callPluginToolHandler,
		newPluginServiceCallPluginToolArgs,
		newPluginServiceCallPluginToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"TestPluginTool": kitex.NewMethodInfo(
		testPluginToolHandler,
		newPluginServiceTestPluginToolArgs,
		newPluginServiceTestPluginToolResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	pluginServiceServiceInfo                = NewServiceInfo()
	pluginServiceServiceInfoForClient       = NewServiceInfoForClient()
	pluginServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return pluginServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return pluginServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return pluginServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PluginService"
	handlerType := (*pluginsvc.PluginService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "pluginsvc",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.10.0",
		Extra:           extra,
	}
	return svcInfo
}

func createPluginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceCreatePluginArgs)
	realResult := result.(*pluginsvc.PluginServiceCreatePluginResult)
	success, err := handler.(pluginsvc.PluginService).CreatePlugin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceCreatePluginArgs() interface{} {
	return pluginsvc.NewPluginServiceCreatePluginArgs()
}

func newPluginServiceCreatePluginResult() interface{} {
	return pluginsvc.NewPluginServiceCreatePluginResult()
}

func updatePluginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceUpdatePluginArgs)
	realResult := result.(*pluginsvc.PluginServiceUpdatePluginResult)
	success, err := handler.(pluginsvc.PluginService).UpdatePlugin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceUpdatePluginArgs() interface{} {
	return pluginsvc.NewPluginServiceUpdatePluginArgs()
}

func newPluginServiceUpdatePluginResult() interface{} {
	return pluginsvc.NewPluginServiceUpdatePluginResult()
}

func deletePluginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceDeletePluginArgs)
	realResult := result.(*pluginsvc.PluginServiceDeletePluginResult)
	success, err := handler.(pluginsvc.PluginService).DeletePlugin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceDeletePluginArgs() interface{} {
	return pluginsvc.NewPluginServiceDeletePluginArgs()
}

func newPluginServiceDeletePluginResult() interface{} {
	return pluginsvc.NewPluginServiceDeletePluginResult()
}

func getPluginByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceGetPluginByIDArgs)
	realResult := result.(*pluginsvc.PluginServiceGetPluginByIDResult)
	success, err := handler.(pluginsvc.PluginService).GetPluginByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceGetPluginByIDArgs() interface{} {
	return pluginsvc.NewPluginServiceGetPluginByIDArgs()
}

func newPluginServiceGetPluginByIDResult() interface{} {
	return pluginsvc.NewPluginServiceGetPluginByIDResult()
}

func getPluginByKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceGetPluginByKeyArgs)
	realResult := result.(*pluginsvc.PluginServiceGetPluginByKeyResult)
	success, err := handler.(pluginsvc.PluginService).GetPluginByKey(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceGetPluginByKeyArgs() interface{} {
	return pluginsvc.NewPluginServiceGetPluginByKeyArgs()
}

func newPluginServiceGetPluginByKeyResult() interface{} {
	return pluginsvc.NewPluginServiceGetPluginByKeyResult()
}

func listPluginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceListPluginArgs)
	realResult := result.(*pluginsvc.PluginServiceListPluginResult)
	success, err := handler.(pluginsvc.PluginService).ListPlugin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceListPluginArgs() interface{} {
	return pluginsvc.NewPluginServiceListPluginArgs()
}

func newPluginServiceListPluginResult() interface{} {
	return pluginsvc.NewPluginServiceListPluginResult()
}

func listPluginByToolsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceListPluginByToolsArgs)
	realResult := result.(*pluginsvc.PluginServiceListPluginByToolsResult)
	success, err := handler.(pluginsvc.PluginService).ListPluginByTools(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceListPluginByToolsArgs() interface{} {
	return pluginsvc.NewPluginServiceListPluginByToolsArgs()
}

func newPluginServiceListPluginByToolsResult() interface{} {
	return pluginsvc.NewPluginServiceListPluginByToolsResult()
}

func publishPluginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServicePublishPluginArgs)
	realResult := result.(*pluginsvc.PluginServicePublishPluginResult)
	success, err := handler.(pluginsvc.PluginService).PublishPlugin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServicePublishPluginArgs() interface{} {
	return pluginsvc.NewPluginServicePublishPluginArgs()
}

func newPluginServicePublishPluginResult() interface{} {
	return pluginsvc.NewPluginServicePublishPluginResult()
}

func createToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceCreateToolArgs)
	realResult := result.(*pluginsvc.PluginServiceCreateToolResult)
	success, err := handler.(pluginsvc.PluginService).CreateTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceCreateToolArgs() interface{} {
	return pluginsvc.NewPluginServiceCreateToolArgs()
}

func newPluginServiceCreateToolResult() interface{} {
	return pluginsvc.NewPluginServiceCreateToolResult()
}

func updateToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceUpdateToolArgs)
	realResult := result.(*pluginsvc.PluginServiceUpdateToolResult)
	success, err := handler.(pluginsvc.PluginService).UpdateTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceUpdateToolArgs() interface{} {
	return pluginsvc.NewPluginServiceUpdateToolArgs()
}

func newPluginServiceUpdateToolResult() interface{} {
	return pluginsvc.NewPluginServiceUpdateToolResult()
}

func deleteToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceDeleteToolArgs)
	realResult := result.(*pluginsvc.PluginServiceDeleteToolResult)
	success, err := handler.(pluginsvc.PluginService).DeleteTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceDeleteToolArgs() interface{} {
	return pluginsvc.NewPluginServiceDeleteToolArgs()
}

func newPluginServiceDeleteToolResult() interface{} {
	return pluginsvc.NewPluginServiceDeleteToolResult()
}

func getToolByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceGetToolByIDArgs)
	realResult := result.(*pluginsvc.PluginServiceGetToolByIDResult)
	success, err := handler.(pluginsvc.PluginService).GetToolByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceGetToolByIDArgs() interface{} {
	return pluginsvc.NewPluginServiceGetToolByIDArgs()
}

func newPluginServiceGetToolByIDResult() interface{} {
	return pluginsvc.NewPluginServiceGetToolByIDResult()
}

func listPluginToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceListPluginToolArgs)
	realResult := result.(*pluginsvc.PluginServiceListPluginToolResult)
	success, err := handler.(pluginsvc.PluginService).ListPluginTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceListPluginToolArgs() interface{} {
	return pluginsvc.NewPluginServiceListPluginToolArgs()
}

func newPluginServiceListPluginToolResult() interface{} {
	return pluginsvc.NewPluginServiceListPluginToolResult()
}

func allPluginToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceAllPluginToolArgs)
	realResult := result.(*pluginsvc.PluginServiceAllPluginToolResult)
	success, err := handler.(pluginsvc.PluginService).AllPluginTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceAllPluginToolArgs() interface{} {
	return pluginsvc.NewPluginServiceAllPluginToolArgs()
}

func newPluginServiceAllPluginToolResult() interface{} {
	return pluginsvc.NewPluginServiceAllPluginToolResult()
}

func listPluginLabelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceListPluginLabelArgs)
	realResult := result.(*pluginsvc.PluginServiceListPluginLabelResult)
	success, err := handler.(pluginsvc.PluginService).ListPluginLabel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceListPluginLabelArgs() interface{} {
	return pluginsvc.NewPluginServiceListPluginLabelArgs()
}

func newPluginServiceListPluginLabelResult() interface{} {
	return pluginsvc.NewPluginServiceListPluginLabelResult()
}

func callPluginToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceCallPluginToolArgs)
	realResult := result.(*pluginsvc.PluginServiceCallPluginToolResult)
	success, err := handler.(pluginsvc.PluginService).CallPluginTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceCallPluginToolArgs() interface{} {
	return pluginsvc.NewPluginServiceCallPluginToolArgs()
}

func newPluginServiceCallPluginToolResult() interface{} {
	return pluginsvc.NewPluginServiceCallPluginToolResult()
}

func testPluginToolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pluginsvc.PluginServiceTestPluginToolArgs)
	realResult := result.(*pluginsvc.PluginServiceTestPluginToolResult)
	success, err := handler.(pluginsvc.PluginService).TestPluginTool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPluginServiceTestPluginToolArgs() interface{} {
	return pluginsvc.NewPluginServiceTestPluginToolArgs()
}

func newPluginServiceTestPluginToolResult() interface{} {
	return pluginsvc.NewPluginServiceTestPluginToolResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreatePlugin(ctx context.Context, req *pluginsvc.CreatePluginReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceCreatePluginArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceCreatePluginResult
	if err = p.c.Call(ctx, "CreatePlugin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePlugin(ctx context.Context, req *pluginsvc.UpdatePluginReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceUpdatePluginArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceUpdatePluginResult
	if err = p.c.Call(ctx, "UpdatePlugin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePlugin(ctx context.Context, req *base.IDReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceDeletePluginArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceDeletePluginResult
	if err = p.c.Call(ctx, "DeletePlugin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPluginByID(ctx context.Context, req *base.IDReq) (r *pluginsvc.Plugin, err error) {
	var _args pluginsvc.PluginServiceGetPluginByIDArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceGetPluginByIDResult
	if err = p.c.Call(ctx, "GetPluginByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPluginByKey(ctx context.Context, req *pluginsvc.GetPluginByKeyReq) (r *pluginsvc.Plugin, err error) {
	var _args pluginsvc.PluginServiceGetPluginByKeyArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceGetPluginByKeyResult
	if err = p.c.Call(ctx, "GetPluginByKey", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListPlugin(ctx context.Context, req *pluginsvc.ListPluginReq) (r *pluginsvc.ListPluginResp, err error) {
	var _args pluginsvc.PluginServiceListPluginArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceListPluginResult
	if err = p.c.Call(ctx, "ListPlugin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListPluginByTools(ctx context.Context, req *pluginsvc.ListPluginByToolsReq) (r *pluginsvc.ListPluginByToolsResp, err error) {
	var _args pluginsvc.PluginServiceListPluginByToolsArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceListPluginByToolsResult
	if err = p.c.Call(ctx, "ListPluginByTools", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishPlugin(ctx context.Context, req *base.IDReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServicePublishPluginArgs
	_args.Req = req
	var _result pluginsvc.PluginServicePublishPluginResult
	if err = p.c.Call(ctx, "PublishPlugin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceCreateToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceCreateToolResult
	if err = p.c.Call(ctx, "CreateTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceUpdateToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceUpdateToolResult
	if err = p.c.Call(ctx, "UpdateTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteTool(ctx context.Context, req *base.IDReq) (r *base.Empty, err error) {
	var _args pluginsvc.PluginServiceDeleteToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceDeleteToolResult
	if err = p.c.Call(ctx, "DeleteTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetToolByID(ctx context.Context, req *base.IDReq) (r *pluginsvc.PluginTool, err error) {
	var _args pluginsvc.PluginServiceGetToolByIDArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceGetToolByIDResult
	if err = p.c.Call(ctx, "GetToolByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListPluginTool(ctx context.Context, req *pluginsvc.ListPluginToolReq) (r *pluginsvc.ListPluginToolResp, err error) {
	var _args pluginsvc.PluginServiceListPluginToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceListPluginToolResult
	if err = p.c.Call(ctx, "ListPluginTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AllPluginTool(ctx context.Context, req *pluginsvc.AllPluginToolReq) (r []*pluginsvc.PluginTool, err error) {
	var _args pluginsvc.PluginServiceAllPluginToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceAllPluginToolResult
	if err = p.c.Call(ctx, "AllPluginTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListPluginLabel(ctx context.Context, req *pluginsvc.ListPluginLabelReq) (r *pluginsvc.ListPluginLabelResp, err error) {
	var _args pluginsvc.PluginServiceListPluginLabelArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceListPluginLabelResult
	if err = p.c.Call(ctx, "ListPluginLabel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (r *pluginsvc.CallPluginToolResp, err error) {
	var _args pluginsvc.PluginServiceCallPluginToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceCallPluginToolResult
	if err = p.c.Call(ctx, "CallPluginTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (r *pluginsvc.TestPluginToolResp, err error) {
	var _args pluginsvc.PluginServiceTestPluginToolArgs
	_args.Req = req
	var _result pluginsvc.PluginServiceTestPluginToolResult
	if err = p.c.Call(ctx, "TestPluginTool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
