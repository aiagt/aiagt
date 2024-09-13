namespace go modelsvc

include './base.thrift'
include './openai.thrift'

// Both internal and external calls are made using temporary tokens, but internally, the tokens can be generated directly without external awareness, whereas externally, the temporary token needs to be provided to the user.
//
// Internally, interfaces are called directly through RPC, and tokens are generated in the background without needing to expose HTTP interfaces externally.
//
// For external access, make sure to isolate the interfaces properly. Only temporary access interfaces should be exposed externally, with each interface route being different each time. Additionally, the token must always have an expiration time.


struct ChatReq {
    1: required string token
    2: required i64 model_id
    3: required openai.ChatCompletionRequest openai_req
}

struct ChatResp {
    1: required openai.ChatCompletionStreamResponse openai_resp
}

struct GenTokenReq {
    1: required i64 app_id
    2: optional i64 plugin_id
    3: required i64 conversation_id
    4: required i32 call_limit
}

struct GenTokenResp {
    1: required string token
    2: required base.Time expired_at
}

service ModelService {
    GenTokenResp GenToken(1: GenTokenReq req)
    ChatResp Chat(1: ChatReq req) (streaming.mode="server")
}