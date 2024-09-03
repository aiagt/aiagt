namespace go appsvc

include './base.thrift'
include './plugin.thrift'
include './user.thrift'

struct App {
    1: required i64 id
    2: required string name
    3: required string description
    4: required string description_md
    5: required i64 model_id
    6: required bool enable_image
    7: required bool enable_file
    8: required string version
    9: required bool is_private
    10: required string home_page
    11: required list<string> preset_questions
    12: required list<plugin.Plugin> plugins
    13: required string logo
    14: required user.User author
    15: required list<string> labels
    16: required i64 created_at
    17: required i64 updated_at
    18: required i64 published_at
}

struct CreateAppReq {
    1: required string name
    2: required string description
    3: required string description_md
    4: required i64 model_id
    5: required bool enable_image
    6: required bool enable_file
    7: required string version
    8: required bool is_private
    9: required string home_page
    10: required list<string> preset_questions
    11: required list<i64> plugin_ids
    12: required string logo
    13: required list<string> labels
}

struct UpdateAppReq {
    1: required i64 id
    2: required string name
    3: required string description
    4: required string description_md
    5: required i64 model_id
    6: required bool enable_image
    7: required bool enable_file
    8: required string version
    9: required bool is_private
    10: required string home_page
    11: required list<string> preset_questions
    12: required list<i64> plugin_ids
    13: required string logo
    14: required list<string> labels
}

struct ListAppReq {
    1: required base.PaginationReq pagination
    2: optional i64 author_id
    3: optional string name
    4: optional string description
    5: optional list<string> labels
}

struct ListAppResp {
    1: required base.PaginationResp pagination
    2: required list<App> plugins
}

service AppService {
    base.Empty CreateApp(1: CreateAppReq req)
    base.Empty UpdateApp(1: UpdateAppReq req)
    base.Empty DeleteApp(1: base.IDReq req)
    App GetAppByID(1: base.IDReq req)
    ListAppResp ListApp(1: ListAppReq req)
}