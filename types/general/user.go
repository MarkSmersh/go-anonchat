package general

type User struct {
	ID                      int     `json:"id"`                                    // Required
	IsBot                   bool    `json:"is_bot"`                                // Required
	FirstName               string  `json:"first_name"`                            // Required
	LastName                *string `json:"last_name,omitempty"`                   // Optional
	Username                *string `json:"username,omitempty"`                    // Optional
	LanguageCode            *string `json:"language_code,omitempty"`               // Optional
	IsPremium               *bool   `json:"is_premium,omitempty"`                  // Optional
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`    // Optional
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`             // Optional
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"` // Optional
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`     // Optional
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`     // Optional
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`            // Optional
}
