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
    12: required list<plugin.Plugin> plugins
    13: required string logo
    14: required user.User author
    15: required list<string> labels
    16: required ModelConfig model_config
    17: required base.Time created_at
    18: required base.Time updated_at
    19: optional base.Time published_at
}

struct ModelConfig {
    1: optional double temperature
    2: optional double top_p
    3: optional i32 n = 1
    4: optional bool stream = true
    5: optional double presence_penalty
    6: optional openai.ChatCompletionResponseFormat response_format
    7: optional i32 seed
    8: optional double frequency_penalty
    9: optional map<string, i32> logit_bias
    10: optional bool logprobs
    11: optional i32 top_logprobs
    12: optional openai.StreamOptions stream_options
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
    13: required list<i64> label_ids
    14: required list<string> new_label_texts
    15: required ModelConfig model_config
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
    14: required list<i64> label_ids
    15: required list<string> new_label_texts
    16: required ModelConfig model_config
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

struct AppLabel {
    1: required i64 id
    2: required string text
    3: required base.Time created_at
}

struct ListAppLabelReq {
    1: required base.PaginationReq pagination
    2: required string text
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

    base.Empty PublishApp(1: base.IDReq req)

    ListAppLabelResp ListAppLabel(1: ListAppLabelReq req)
}