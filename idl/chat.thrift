namespace go chatsvc

include './base.thrift'

enum MessageType {
    TEXT
    IMAGE
    FILE
}

enum MessageRole {
    USER
    ASSISTANT
    SYSTEM
}

union MessageContent {
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

struct MessageItem {
    1: required MessageType type
    2: required MessageContent content
}

struct ChatReq {
    1: required i64 conversation_id  // 为空时自动创建 conversation
    2: required i64 app_id
    3: required list<MessageItem> message
    4: required bool stream
}

struct ChatResp {
    1: required list<MessageItem> message
    2: required i64 conversation_id
}

struct Conversation {
    1: required i64 id
    2: required i64 user_id
    3: required i64 app_id
    4: required string title
    5: required base.Time created_at
    6: required base.Time updated_at
}

struct Message {
    1: required i64 id
    2: required i64 conversation_id
    3: required MessageRole role
    4: required list<MessageItem> message
    5: required base.Time created_at
    6: required base.Time updated_at
}

struct UpdateMessageReq {
    1: required i64 id
    2: required list<MessageItem> message
}

struct ListMessageReq {
    1: required base.PaginationReq pagination
    2: required i64 conversation_id
}

struct ListMessageResp {
    1: required base.PaginationResp pagination
    2: required list<Message> messages
}

struct UpdateConversationReq {
    1: required i64 id
    2: required string title
}

struct ListConversationReq {
    1: required base.PaginationReq pagination
    2: required i64 user_id
    3: required i64 app_id
}

struct ListConversationResp {
    1: required base.PaginationResp pagination
    2: required list<Conversation> conversations
}

service ChatService {
    ChatResp Chat(1: ChatReq req) (streaming.mode="server")

    base.Empty UpdateMessage(1: UpdateMessageReq req)
    base.Empty DeleteMessage(1: base.IDReq req)
    ListMessageResp ListMessage(1: ListMessageReq req)

    base.Empty UpdateConversation(1: UpdateConversationReq req)
    base.Empty DeleteConversation(1: base.IDReq req)
    Conversation GetConversationByID(1: base.IDReq req)
    ListConversationResp ListConversation(1: ListConversationReq req)
}
