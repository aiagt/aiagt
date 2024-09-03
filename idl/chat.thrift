namespace go chatsvc

include './base.thrift'

enum MessageType {
    TEXT
    IMAGE
    FILE
}

union MessaegContent {
    1: MessageContentText text
    2: MessageContentImage image
    3: MessageContentFile file
}

struct MessageContentText {
    1: required string text
}

struct MessageContentImage {
    1: required string key
    2: required string url
}

struct MessageContentFile {
    1: required string key
    2: required string url
    3: required string type
}

struct ChatMessage {
    1: required MessageType type
    2: required MessaegContent content
}

struct ChatReq {
    1: required i64 conversation_id  // 为空时自动创建 conversation
    2: required i64 app_id
    3: required list<ChatMessage> messages
    4: required bool stream
    // 每个token都具有额度限制，每个user&app都有对应的长期token，此外有些临时使用的token，token应该在后端生成和存储
    // TODO: call_token 应该在 ModelChat 接口中使用，所以应该在这里的 Chat 接口内生成再传递到 ModelChat，而不是接收外部参数
    // 5: required string call_token
}

struct ChatResp {
    1: required i64 conversation
    2: required ChatMessage message
}

// TODO: message & conversation manager

struct Message {}

struct CreateMessageReq {}

struct UpdateMessageReq {}

struct ListMessageReq {}

struct ListMessageResp {}

service ChatService {
    ChatResp Chat(1: ChatReq req) (streaming.mode="server")

//    base.Empty CreateMessage(1: CreateMessageReq req)  // create 应该对外吗？
//    base.Empty UpdateMessage(1: UpdateMessageReq req)  // 支持 update 吗？应该对外吗？
    base.Empty DeleteMessage(1: base.IDReq req)
    ListMessageResp ListMessage(1: ListMessageReq req)
}
