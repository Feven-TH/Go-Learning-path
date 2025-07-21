package models

type PromoteAdmin struct {
	TargetUserID string `json:"target_user_id" binding:"required"`
}
