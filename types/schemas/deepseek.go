package schemas

type DeepSeekResponseFormat struct {
	Type string `json:"type"`
}

// 创意类写作/诗歌创作1.5;翻译1.3;通用对话1.3;数据抽取/分析1.0;代码生成/数学解题0.0
type DeepSeekReq struct {
	Model          string                 `json:"model"`
	Messages       []DeepSeekMessage      `json:"messages"`
	Stream         bool                   `json:"stream"`
	MaxTokens      int                    `json:"max_tokens"`
	Temperature    float64                `json:"temperature"`
	ResponseFormat DeepSeekResponseFormat `json:"response_format"`
}

type DeepSeekRes struct {
	ID                string           `json:"id"`
	Object            string           `json:"object"`
	Created           int64            `json:"created"`
	Model             string           `json:"model"`
	Choices           []DeepSeekChoice `json:"choices"`
	Usage             DeepSeekUsage    `json:"usage"`
	SystemFingerprint string           `json:"system_fingerprint"`
}

type DeepSeekChoice struct {
	Index        int               `json:"index"`
	Message      DeepSeekMessage   `json:"message"`
	LogProbs     *DeepSeekLogProbs `json:"logprobs,omitempty"`
	FinishReason string            `json:"finish_reason"`
}

type DeepSeekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekLogProbs struct {
	Tokens        []string             `json:"tokens,omitempty"`
	TokenLogProbs []float64            `json:"token_logprobs,omitempty"`
	TopLogProbs   []map[string]float64 `json:"top_logprobs,omitempty"`
}

type DeepSeekUsage struct {
	PromptTokens          int `json:"prompt_tokens"`
	CompletionTokens      int `json:"completion_tokens"`
	TotalTokens           int `json:"total_tokens"`
	PromptCacheHitTokens  int `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"`
}
