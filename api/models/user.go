package models

type TgUser struct {
	UserId                  *int   `json:"id" bson:"_id"`
	IsBot                   bool   `json:"is_bot" bson:"is_bot"`
	FirstName               string `json:"first_name" bson:"first_name"`
	LastName                string `json:"last_name" bson:"last_name"`
	Username                string `json:"username" bson:"username"`
	LanguageCode            string `json:"language_code" bson:"language_code"`
	IsPremium               bool   `json:"is_premium" bson:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu" bson:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups" bson:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages" bson:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries" bson:"supports_inline_queries"`
}
