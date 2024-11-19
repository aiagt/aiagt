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

struct Model {
    1: required i64 id
    2: required string name
    3: required string description
    4: required string source
    5: required string model_key
    6: required string logo
    7: required string input_price
    8: required string output_price
}

struct CreateModelReq {
    1: required string name
    2: required string description
    3: required string source
    4: required string model_key
    5: required string logo
    6: required string input_price
    7: required string output_price
}

struct UpdateModelReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional string name
    3: optional string description
    4: optional string source
    5: optional string model_key
    6: optional string logo
    7: optional string input_price
    8: optional string output_price
}

struct ListModelReq {
    1: required base.PaginationReq pagination
    2: optional string name (go.tag='query:"name"')
    3: optional string source (go.tag='query:"source"')
}

struct ListModelResp {
    1: required base.PaginationResp pagination
    2: required list<Model> models
}

service ModelService {
    GenTokenResp GenToken(1: GenTokenReq req)
    ChatResp Chat(1: ChatReq req) (streaming.mode="server")

    base.Empty CreateModel(1: CreateModelReq req)
    base.Empty UpdateModel(1: UpdateModelReq req)
    base.Empty DeleteModel(1: base.IDReq req)
    Model GetModelByID(1: base.IDReq req)
    ListModelResp ListModel(1: ListModelReq req)
}
