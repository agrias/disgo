package discord

import "github.com/DisgoOrg/snowflake"

//Attachment is used for files sent in a Message
type Attachment struct {
	ID        snowflake.Snowflake `json:"id,omitempty"`
	Filename  string              `json:"filename,omitempty"`
	Size      int                 `json:"size,omitempty"`
	URL       string              `json:"url,omitempty"`
	ProxyURL  string              `json:"proxy_url,omitempty"`
	Height    *int                `json:"height,omitempty"`
	Width     *int                `json:"width,omitempty"`
	Ephemeral bool                `json:"ephemeral,omitempty"`
}
