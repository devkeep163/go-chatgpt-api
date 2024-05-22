package chatgpt

import (
	"github.com/google/uuid"
)

type CreateConversationRequest struct {
	Action                     string    	`json:"action"`
	Messages                   []Message 	`json:"messages,omitempty"`
	Model                      string    	`json:"model"`
	ParentMessageID            string    	`json:"parent_message_id,omitempty"`
	ConversationID             string    	`json:"conversation_id,omitempty"`
	TimezoneOffsetMin          int       	`json:"timezone_offset_min"`
	ForceUseSse                bool      	`json:"force_use_sse"`
	HistoryAndTrainingDisabled bool      	`json:"history_and_training_disabled"`
	AutoContinue               bool      	`json:"auto_continue"`
	Suggestions                []string  	`json:"suggestions"`
	WebsocketRequestId         string    	`json:"websocket_request_id"`
	Conversation_mode          interface{}  `json:"conversation_mode,omitempty"`
	PluginIDs                  []string  	`json:"plugin_ids,omitempty"`
	PluginData                 interface{} 	`json:"plugin_data,omitempty"`
	Force_nulligen             bool      	`json:"force_nulligen,omitempty"`
	Force_paragen              bool      	`json:"force_paragen,omitempty"`
	Force_paragen_model_slug   string    	`json:"force_paragen_model_slug,omitempty"`
	Force_rate_limit           bool      	`json:"force_rate_limit,omitempty"`
}

func (c *CreateConversationRequest) AddMessage(role string, content string) {
	c.Messages = append(c.Messages, Message{
		ID:       uuid.New().String(),
		Author:   Author{Role: role},
		Content:  Content{ContentType: "text", Parts: []interface{}{content}},
		Metadata: map[string]string{},
	})
}

type Message struct {
	Author   Author      `json:"author"`
	Content  Content     `json:"content"`
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata"`
}

type Author struct {
	Role string `json:"role"`
	Nane interface{} `json:"name,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

type Content struct {
	ContentType string        `json:"content_type"`
	Parts       []interface{} `json:"parts"`
}

type CreateConversationResponse struct {
	Message struct {
		ID     string `json:"id"`
		Author struct {
			Role     string      `json:"role"`
			Name     interface{} `json:"name"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"author"`
		CreateTime float64     `json:"create_time"`
		UpdateTime interface{} `json:"update_time"`
		Content    struct {
			ContentType string   `json:"content_type"`
			Parts       []string `json:"parts"`
		} `json:"content"`
		Status   string  `json:"status"`
		EndTurn  bool    `json:"end_turn"`
		Weight   float64 `json:"weight"`
		Metadata struct {
			MessageType   string `json:"message_type"`
			ModelSlug     string `json:"model_slug"`
			FinishDetails struct {
				Type string `json:"type"`
			} `json:"finish_details"`
		} `json:"metadata"`
		Recipient string `json:"recipient"`
	} `json:"message"`
	ConversationID string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
}

type GetModelsResponse struct {
	Models []struct {
		Slug         string   `json:"slug"`
		MaxTokens    int      `json:"max_tokens"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Tags         []string `json:"tags"`
		Capabilities struct {
		} `json:"capabilities"`
		EnabledTools []string `json:"enabled_tools,omitempty"`
	} `json:"models"`
	Categories []struct {
		Category             string `json:"category"`
		HumanCategoryName    string `json:"human_category_name"`
		SubscriptionLevel    string `json:"subscription_level"`
		DefaultModel         string `json:"default_model"`
		CodeInterpreterModel string `json:"code_interpreter_model"`
		PluginsModel         string `json:"plugins_model"`
	} `json:"categories"`
}

type ChatGPTWSSResponse struct {
	WssUrl         string `json:"wss_url"`
	ConversationId string `json:"conversation_id,omitempty"`
	ResponseId     string `json:"response_id,omitempty"`
}

type WSSMsgResponse struct {
	SequenceId int                `json:"sequenceId"`
	Type       string             `json:"type"`
	From       string             `json:"from"`
	DataType   string             `json:"dataType"`
	Data       WSSMsgResponseData `json:"data"`
}

type WSSSequenceAckMessage struct {
	Type       string `json:"type"`
	SequenceId int    `json:"sequenceId"`
}

type WSSMsgResponseData struct {
	Type           string `json:"type"`
	Body           string `json:"body"`
	MoreBody       bool   `json:"more_body"`
	ResponseId     string `json:"response_id"`
	ConversationId string `json:"conversation_id"`
}

type ChatRequire struct {
	Token  string    `json:"token"`
	Proof  ProofWork `json:"proofofwork,omitempty"`
	Arkose struct {
		Required bool   `json:"required"`
		DX       string `json:"dx,omitempty"`
	} `json:"arkose"`
}

type FileInfo struct {
	DownloadURL string `json:"download_url"`
	Status      string `json:"status"`
}

type DalleContent struct {
	AssetPointer string `json:"asset_pointer"`
	Metadata     struct {
		Dalle struct {
			Prompt string `json:"prompt"`
		} `json:"dalle"`
	} `json:"metadata"`
}