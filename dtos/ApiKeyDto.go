package dtos

type ApiKeyDto struct {
	Limited   bool  `json:"limited"`
	TimeLimit int32 `json:"timeLimit,omitempty"`
}
