namespace go appsvc

include './base.thrift'
include './plugin.thrift'
include './user.thrift'
include './openai.thrift'

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
    12: required list<i64> tool_ids
    13: optional list<plugin.PluginTool> tools
    14: required string logo
    15: required i64 author_id
    16: optional user.User author
    17: required list<i64> label_ids
    18: optional list<AppLabel> labels
    19: required ModelConfig model_config
    20: required base.Time created_at
    21: required base.Time updated_at
    22: optional base.Time published_at
}

struct ModelConfig {
    1: optional i32 max_tokens
    2: optional double temperature
    3: optional double top_p
    4: optional i32 n = 1
    5: optional bool stream = true
    6: optional list<string> stop
    7: optional double presence_penalty
    8: optional openai.ChatCompletionResponseFormat response_format
    9: optional i32 seed
    10: optional double frequency_penalty
    11: optional map<string, i32> logit_bias
    12: optional bool logprobs
    13: optional i32 top_logprobs
    14: optional string user
    15: optional openai.StreamOptions stream_options
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
    11: required list<i64> tool_ids
    12: required string logo
    13: required list<i64> label_ids
    14: required list<string> label_texts
    15: required ModelConfig model_config
}

struct UpdateAppReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional string name
    3: optional string description
    4: optional string description_md
    5: optional i64 model_id
    6: optional bool enable_image
    7: optional bool enable_file
    8: optional string version
    9: optional bool is_private
    10: optional string home_page
    11: optional list<string> preset_questions
    12: optional list<i64> tool_ids
    13: optional string logo
    14: optional list<i64> label_ids
    15: optional list<string> label_texts
    16: optional ModelConfig model_config
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
    2: required list<App> apps
}

struct PublishAppReq {
    1: required i64 id
    2: required string version
}

struct AppLabel {
    1: required i64 id
    2: required string text
    3: required base.Time created_at
}

struct ListAppLabelReq {
    1: required base.PaginationReq pagination
    2: optional string text (go.tag='query:"text"')
}

struct ListAppLabelResp {
    1: required base.PaginationResp pagination
    2: required list<AppLabel> labels
}

service AppService {
    base.Empty CreateApp(1: CreateAppReq req)
    base.Empty UpdateApp(1: UpdateAppReq req)
    base.Empty DeleteApp(1: base.IDReq req)
    App GetAppByID(1: base.IDReq req)
    ListAppResp ListApp(1: ListAppReq req)

    base.Empty PublishApp(1: PublishAppReq req)

    ListAppLabelResp ListAppLabel(1: ListAppLabelReq req)
}