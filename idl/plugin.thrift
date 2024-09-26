namespace go pluginsvc

include './base.thrift'
include './user.thrift'

struct Plugin {
    1: required i64 id
    2: required i64 key
    3: required string name
    4: required string description
    5: required string description_md
    6: required i64 author_id
    7: optional user.User author
    8: required bool is_private
    9: required string home_page
    10: required bool enable_secret
    11: required list<PluginSecret> secrets
    12: required list<i64> label_ids
    13: optional list<PluginLabel> labels
    14: required list<i64> tool_ids
    15: optional list<PluginTool> tools
    16: required string logo
    17: required base.Time created_at
    18: required base.Time updated_at
    19: optional base.Time published_at
}

struct PluginSecret {
    1: required string name
    2: required string description
    3: required string acquire_method
    4: required string link
}

struct PluginTool {
    1: required i64 id
    2: required string name
    3: required string description
    4: required i64 plugin_id
    5: required binary request_type
    6: required binary response_type
    7: required string api_url
    8: optional i64 import_model_id
    9: required base.Time created_at
    10: required base.Time updated_at
    11: optional base.Time tested_at
}

struct PluginLabel {
    1: required i64 id
    2: required string text
    3: required base.Time created_at
}

struct ListPluginLabelReq {
    1: required base.PaginationReq pagination
    2: optional string text (go.tag='query:"text"')
}

struct ListPluginLabelResp {
    1: required base.PaginationResp pagination
    2: required list<PluginLabel> labels
}

struct CreatePluginReq {
    1: required i64 key
    2: required string name
    3: required string description
    4: required string description_md
    5: required bool is_private
    6: required string home_page
    7: required bool enable_secret
    8: required list<PluginSecret> secrets
    9: required list<i64> label_ids
    10: required list<string> label_texts
    11: required list<i64> tool_ids  // Tool list (mainly used for plug-in copying)
    12: required string logo
}

struct UpdatePluginReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional i64 key
    3: optional string name
    4: optional string description
    5: optional string description_md
    6: optional bool is_private
    7: optional string home_page
    8: optional bool enable_secret
    9: optional list<PluginSecret> secrets
    10: optional list<i64> label_ids
    11: optional list<string> label_texts
    12: optional list<i64> tool_ids
    13: optional string logo
}

struct ListPluginReq {
    1: required base.PaginationReq pagination
    2: optional i64 author_id
    3: optional string name
    4: optional string description
    5: optional list<i64> labels
}

struct ListPluginResp {
    1: required base.PaginationResp pagination
    2: required list<Plugin> plugins
}

struct CreatePluginToolReq {
    1: required string name
    2: required string description
    3: required i64 plugin_id
    4: required binary request_type
    5: required binary response_type
    6: required string api_url
    7: optional i64 import_model_id
}

struct UpdatePluginToolReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional string name
    3: optional string description
    4: optional i64 plugin_id
    5: optional binary request_type
    6: optional binary response_type
    7: optional string api_url
    8: optional i64 import_model_id
}

struct ListPluginToolReq {
    1: required base.PaginationReq pagination
    2: optional i64 plugin_id
    3: optional list<i64> tool_ids
}

struct ListPluginToolResp {
    1: required list<PluginTool> tools;
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
    base.Empty CreatePlugin(1: CreatePluginReq req)
    base.Empty UpdatePlugin(1: UpdatePluginReq req)
    base.Empty DeletePlugin(1: base.IDReq req)
    Plugin GetPluginByID(1: base.IDReq req)
    ListPluginResp ListPlugin(1: ListPluginReq req)

    base.Empty PublishPlugin(1: base.IDReq req)

    base.Empty CreateTool(1: CreatePluginToolReq req)
    base.Empty UpdateTool(1: UpdatePluginToolReq req)
    base.Empty DeleteTool(1: base.IDReq req)
    PluginTool GetToolByID(1: base.IDReq req)
    ListPluginToolResp ListPluginTool(1: ListPluginToolReq req)

    ListPluginLabelResp ListPluginLabel(1: ListPluginLabelReq req)

    CallPluginToolResp CallPluginTool(1: CallPluginToolReq req)
    TestPluginToolResp TestPluginTool(1: CallPluginToolReq req)
}
