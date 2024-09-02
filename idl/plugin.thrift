namespace go pluginsvc

include './base.thrift'

struct Plugin {
    1: required i64 id
    2: required string name
    3: required string description
    4: required i64 author_id
    5: required bool is_private
    6: required string home_page
    7: required string enable_secret
    8: required list<PluginSecret> secrets
    9: required list<string> labels
    10: required list<PluginTool> tools
    11: required i64 created_at
    12: required i64 updated_at
    13: required i64 published_at
}

struct PluginSecret {
    1: required string name
    2: required string description
}

struct PluginTool {
    1: required i64 id
    2: required string name
    3: required string description
    4: required i64 plugin_id
    5: required string request_type
    6: required string response_type
    7: required string api
    8: optional i64 import_model_id
    9: required i64 created_at
    10: required i64 updated_at
    11: required i64 tested_at
}

struct CreatePluginReq {
    1: required string name
    2: required string description
    3: required bool is_private
    4: required string home_page
    5: required string enable_secret
    6: required list<PluginSecret> secrets
    7: required list<string> labels
    8: required list<i64> tool_ids  // 工具列表（主要用于插件复制）
}

struct UpdatePluginReq {
    1: required i64 id
    2: required string name
    3: required string description
    4: required bool is_private
    5: required string home_page
    6: required string enable_secret
    7: required list<PluginSecret> secrets
    8: required list<string> labels
    9: required list<i64> tool_ids  // 工具列表（主要用于插件复制）
}

struct ListPluginReq {
    1: required base.PaginationReq pagination
    2: optional i64 author_id
    3: optional string name
    4: optional string description
}

struct ListPluginResp {
    1: required list<Plugin> plugins
    2: required base.PaginationResp pagination
}

struct CreatePluginToolReq {
    1: required string name
    2: required string description
    3: required i64 plugin_id
    4: required string request_type
    5: required string response_type
    6: required string api
    7: optional i64 import_model_id
}

struct UpdatePluginToolReq {
    1: required i64 id
    2: required string name
    3: required string description
    4: required i64 plugin_id
    5: required string request_type
    6: required string response_type
    7: required string api
    8: optional i64 import_model_id
}

struct ListPluginToolResp {
    1: required list<Plugin> plugins;
    2: required base.PaginationResp pagination;
}

struct CallPluginToolReq {
    1: required i64 plugin_id
    2: required i64 tool_id
    3: optional map<string, string> secrets
    4: required binary request
}

struct CallPluginToolResp {
    1: required i64 code
    2: required string msg
    3: required binary response
}

struct TestPluginToolResp {
    1: required bool code
    2: required string msg
    3: required binary response
}

service PluginService {
    base.Empty Create(1: CreatePluginReq req)
    base.Empty Update(1: UpdatePluginReq req)
    ListPluginResp List(1: base.PaginationReq req)
    Plugin GetByID(1: base.IDReq req)
    base.Empty Delete(1: base.IDReq req)

    base.Empty CreateTool(1: CreatePluginToolReq req)
    base.Empty UpdateTool(1: UpdatePluginToolReq req)
    ListPluginToolResp ListTool(1: base.IDReq req)
    PluginTool GetToolByID(1: base.IDReq req)
    base.Empty DeleteTool(1: base.IDReq req)

    CallPluginToolResp CallPluginTool(1: CallPluginToolReq req)
    TestPluginToolResp TestPluginTool(1: CallPluginToolReq req)
}
