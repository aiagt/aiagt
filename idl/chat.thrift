namespace go chatsvc

include './base.thrift'

enum MessageType {
    TEXT
    IMAGE
    FILE
    FUNCTION
    FUNCTION_CALL
    TOOL
    TOOL_CALL
}

enum MessageRole {
    USER
    ASSISTANT
    SYSTEM
    FUNCTION
    TOOL
}

union MessageContentValue {
    1: MessageContentValueText text
    2: MessageContentValueImage image
    3: MessageContentValueFile file
    4: MessageContentValueFunc func
    5: MessageContentValueFuncCall func_call
    6: MessageContentValueTool tool
    7: MessageContentValueToolCall tool_call
}

struct MessageContentValueText {
    1: required string text
}

struct MessageContentValueImage {
    1: required string url
}

struct MessageContentValueFile {
    1: required string url
    2: required string type
}

struct MessageContentValueFunc {
    1: required string name
    2: required string content  // JSON
}

struct MessageContentValueFuncCall {
    1: required string name
    2: required string arguments  // JSON
}

struct MessageContentValueTool {
    1: required string id
    2: required string name
    3: required string content
}

struct MessageContentValueToolCall {
    1: required string id
    2: required string name
    3: required string arguments
}

struct MessageContent {
    1: required MessageType type
    2: required MessageContentValue content
}

struct ChatReq {
    1: optional i64 conversation_id  // Automatically create conversation when empty
    2: required i64 app_id
    3: required list<MessageContent> messages
}

struct ChatResp {
    1: required list<Message> messages
    2: required i64 conversation_id
    3: optional string conversation_title
}

struct ChatRespMessage {
    1: required MessageRole role
    2: required MessageContent content
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
    4: required MessageContent content
    5: required base.Time created_at
    6: required base.Time updated_at
}

struct UpdateMessageReq {
    1: required i64 id (go.tag='path:"id"')
    2: required MessageContent message
}

struct DeleteMessageReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional i8 more (go.tag='query:"more"')
}

struct ListMessageReq {
    1: required base.PaginationReq pagination
    2: required i64 conversation_id (go.tag='query:"conversation_id"')
}

struct ListMessageResp {
    1: required base.PaginationResp pagination
    2: required list<Message> messages
}

struct UpdateConversationReq {
    1: required i64 id (go.tag='path:"id"')
    2: required string title
}

struct ListConversationReq {
    1: required base.PaginationReq pagination
    2: required i64 app_id (go.tag='query:"app_id"')
}

struct ListConversationResp {
    1: required base.PaginationResp pagination
    2: required list<Conversation> conversations
}

struct InitDevelopReq {
    1: required i64 app_id
}

struct InitDevelopResp {
    1: required Conversation conversation
    2: required list<Message> messages
}

service ChatService {
    ChatResp Chat(1: ChatReq req) (streaming.mode="server")

    base.Empty UpdateMessage(1: UpdateMessageReq req)
    base.Empty DeleteMessage(1: DeleteMessageReq req)
    ListMessageResp ListMessage(1: ListMessageReq req)

    base.Empty UpdateConversation(1: UpdateConversationReq req)
    base.Empty DeleteConversation(1: base.IDReq req)
    Conversation GetConversationByID(1: base.IDReq req)
    ListConversationResp ListConversation(1: ListConversationReq req)
    InitDevelopResp InitDevelop(1: InitDevelopReq req)
}
