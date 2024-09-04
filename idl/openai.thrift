namespace go openai

enum ToolType {
  FUNCTION = 1
}

enum ChatMessagePartType {
  TEXT = 1,
  IMAGE_URL = 2
}

enum ImageURLDetail {
  HIGH = 1,
  LOW = 2,
  AUTO = 3
}

enum ChatCompletionResponseFormatType {
  JSON_OBJECT = 1,
  JSON_SCHEMA = 2,
  TEXT = 3
}

enum FinishReason {
  STOP = 1,
  LENGTH = 2,
  FUNCTION_CALL = 3,
  TOOL_CALLS = 4,
  CONTENT_FILTER = 5,
  NULL_REASON = 6
}

struct Hate {
  1: required bool filtered,
  2: optional string severity
}

struct SelfHarm {
  1: required bool filtered,
  2: optional string severity
}

struct Sexual {
  1: required bool filtered,
  2: optional string severity
}

struct Violence {
  1: required bool filtered,
  2: optional string severity
}

struct ContentFilterResults {
  1: required Hate hate,
  2: required SelfHarm self_harm,
  3: required Sexual sexual,
  4: required Violence violence
}

struct PromptAnnotation {
  1: required i32 prompt_index,
  2: required ContentFilterResults content_filter_results
}

struct ChatMessageImageURL {
  1: required string url,
  2: required ImageURLDetail detail
}

struct ChatMessagePart {
  1: required ChatMessagePartType type,
  2: optional string text,
  3: optional ChatMessageImageURL image_url
}

struct FunctionCall {
  1: optional string name,
  2: optional string arguments
}

struct ToolCall {
  1: optional i32 index,
  2: required string id,
  3: required ToolType type,
  4: required FunctionCall function
}

struct ChatCompletionMessage {
  1: required string role,
  2: optional string content,
  3: optional list<ChatMessagePart> multi_content,
  4: optional string name,
  5: optional FunctionCall function_call,
  6: optional list<ToolCall> tool_calls,
  7: optional string tool_call_id
}

struct FunctionDefinition {
  1: required string name,
  2: optional string description,
  3: optional bool strict,
  4: optional string parameters
}

struct Tool {
  1: required ToolType type,
  2: optional FunctionDefinition function
}

struct ToolChoice {
  1: required ToolType type,
  2: optional string function_name
}

struct StreamOptions {
  1: optional bool include_usage
}

struct ChatCompletionResponseFormatJSONSchema {
  1: required string name,
  2: optional string description,
  3: required string schema,
  4: optional bool strict
}

struct ChatCompletionResponseFormat {
  1: required ChatCompletionResponseFormatType type,
  2: optional ChatCompletionResponseFormatJSONSchema json_schema
}

struct ChatCompletionRequest {
  1: required string model,
  2: required list<ChatCompletionMessage> messages,
  3: optional i32 max_tokens,
  4: optional double temperature,
  5: optional double top_p,
  6: optional i32 n,
  7: optional bool stream,
  8: optional list<string> stop,
  9: optional double presence_penalty,
  10: optional ChatCompletionResponseFormat response_format,
  11: optional i32 seed,
  12: optional double frequency_penalty,
  13: optional map<string, i32> logit_bias,
  14: optional bool logprobs,
  15: optional i32 top_logprobs,
  16: optional string user,
  17: optional list<FunctionDefinition> functions,
  18: optional string function_call, // JSON Schema
  19: optional list<Tool> tools,
  20: optional ToolChoice tool_choice,
  21: optional StreamOptions stream_options,
  22: optional bool parallel_tool_calls // Simplified, assuming it's a boolean flag
}

struct ChatCompletionChoice {
  1: required i32 index,
  2: required ChatCompletionMessage message,
  3: optional string finish_reason,
  4: optional double finish_probability
}

struct Usage {
  1: required i32 prompt_tokens,
  2: required i32 completion_tokens,
  3: required i32 total_tokens
}

struct ChatCompletionResponse {
  1: required string id,
  2: required string object,
  3: required i64 created,
  4: required string model,
  5: required list<ChatCompletionChoice> choices,
  6: optional Usage usage
}