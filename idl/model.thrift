namespace go modelsvc

include './base.thrift'
include './openai.thrift'

// 无论内部还是外部都是通过临时token来调用，但是内部可以直接在内部生成token，对外无感知，外部则需要把临时token给用户
// 内部直接通过RPC来调用接口，token直接在后台生成，不需要对外开放HTTP接口
// 网关注意做接口隔离，对外只能开放临时访问的接口，每次的接口路由都是不一样的，并且token一定是具有时效性的

struct ChatReq {
    1: required string token
    2: required i64 model_id
    3: required openai.ChatCompletionRequest openai_req
}

struct ChatResp {
    1: required openai.ChatCompletionResponse openai_resp
}

struct GenTokenReq {
    1: required i64 app_id
    2: required i64 user_id
    3: required i64 plugin_id
    4: required i64 conversation_id
    5: required i32 call_limit
}

struct GenTokenResp {
    1: required string token
    2: required base.Time expired_at
}

service ModelService {
    GenTokenResp GenToken(1: GenTokenReq req)
    ChatResp Chat(1: ChatReq req) (streaming.mode="server")
}