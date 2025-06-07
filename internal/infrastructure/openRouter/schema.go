package openRouter

// ResponseScheme represents the main API response object.
type ResponseScheme struct {
	ID                string         `json:"id"`
	Choices           []Choice       `json:"choices"`
	Created           int64          `json:"created"`
	Model             string         `json:"model"`
	Object            string         `json:"object"` // "chat.completion" | "chat.completion.chunk"
	SystemFingerprint *string        `json:"system_fingerprint,omitempty"`
	Usage             *ResponseUsage `json:"usage,omitempty"`
}

// ResponseUsage provides token count information for the response.
type ResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Choice is a union type for NonChatChoice, NonStreamingChoice, and StreamingChoice.
type Choice struct {
	// ChoiceType helps distinguish which choice struct to use
	// (not marshaled; for internal use)
	ChoiceType string `json:"-"`

	// NonChatChoice fields
	FinishReason *string `json:"finish_reason,omitempty"`
	Text         *string `json:"text,omitempty"`

	// NonStreamingChoice fields
	NativeFinishReason *string        `json:"native_finish_reason,omitempty"`
	Message            *ChoiceMessage `json:"message,omitempty"`

	// StreamingChoice fields
	Delta *DeltaMessage `json:"delta,omitempty"`

	Error *ErrorResponse `json:"error,omitempty"`
}

// ChoiceMessage for NonStreamingChoice
type ChoiceMessage struct {
	Content   *string    `json:"content"`
	Role      string     `json:"role"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

// DeltaMessage for StreamingChoice
type DeltaMessage struct {
	Content   *string    `json:"content"`
	Role      *string    `json:"role,omitempty"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

type ErrorResponse struct {
	Code     int                    `json:"code"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"` // always "function"
	Function FunctionCall `json:"function"`
}

// FunctionCall is a placeholder; define this according to your API's spec.
type FunctionCall struct {
	// Add fields as needed
}
