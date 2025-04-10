package general

type Message struct {
	MessageID       int  `json:"message_id"`
	MessageThreadID *int `json:"message_thread_id,omitempty"`
	// From                        *User                      `json:"from,omitempty"`
	SenderChat       *Chat `json:"sender_chat,omitempty"`
	SenderBoostCount *int  `json:"sender_boost_count,omitempty"`
	// SenderBusinessBot          *User                      `json:"sender_business_bot,omitempty"`
	Date                 int     `json:"date"`
	BusinessConnectionID *string `json:"business_connection_id,omitempty"`
	Chat                 Chat    `json:"chat"`
	// ForwardOrigin               *MessageOrigin            `json:"forward_origin,omitempty"`
	IsTopicMessage     *bool    `json:"is_topic_message,omitempty"`
	IsAutomaticForward *bool    `json:"is_automatic_forward,omitempty"`
	ReplyToMessage     *Message `json:"reply_to_message,omitempty"`
	// ExternalReply               *ExternalReplyInfo         `json:"external_reply,omitempty"`
	// Quote                       *TextQuote                 `json:"quote,omitempty"`
	// ReplyToStory                *Story                     `json:"reply_to_story,omitempty"`
	// ViaBot                      *User                      `json:"via_bot,omitempty"`
	EditDate            *int    `json:"edit_date,omitempty"`
	HasProtectedContent *bool   `json:"has_protected_content,omitempty"`
	IsFromOffline       *bool   `json:"is_from_offline,omitempty"`
	MediaGroupID        *string `json:"media_group_id,omitempty"`
	AuthorSignature     *string `json:"author_signature,omitempty"`
	Text                *string `json:"text,omitempty"`
	// Entities                    []MessageEntity            `json:"entities,omitempty"`
	// LinkPreviewOptions          *LinkPreviewOptions        `json:"link_preview_options,omitempty"`
	// EffectID                    *string                    `json:"effect_id,omitempty"`
	// Animation                   *Animation                 `json:"animation,omitempty"`
	// Audio                       *Audio                     `json:"audio,omitempty"`
	// Document                    *Document                  `json:"document,omitempty"`
	// PaidMedia                   *PaidMediaInfo             `json:"paid_media,omitempty"`
	// Photo                       []PhotoSize                `json:"photo,omitempty"`
	// Sticker                     *Sticker                   `json:"sticker,omitempty"`
	// Story                       *Story                     `json:"story,omitempty"`
	// Video                       *Video                     `json:"video,omitempty"`
	// VideoNote                   *VideoNote                 `json:"video_note,omitempty"`
	// Voice                       *Voice                     `json:"voice,omitempty"`
	Caption *string `json:"caption,omitempty"`
	// CaptionEntities             []MessageEntity            `json:"caption_entities,omitempty"`
	// ShowCaptionAboveMedia       *bool                      `json:"show_caption_above_media,omitempty"`
	// HasMediaSpoiler             *bool                      `json:"has_media_spoiler,omitempty"`
	// Contact                     *Contact                   `json:"contact,omitempty"`
	// Dice                        *Dice                      `json:"dice,omitempty"`
	// Game                        *Game                      `json:"game,omitempty"`
	// Poll                        *Poll                      `json:"poll,omitempty"`
	// Venue                       *Venue                     `json:"venue,omitempty"`
	// Location                    *Location                  `json:"location,omitempty"`
	// NewChatMembers              []User                     `json:"new_chat_members,omitempty"`
	// LeftChatMember              *User                      `json:"left_chat_member,omitempty"`
	NewChatTitle *string `json:"new_chat_title,omitempty"`
	// NewChatPhoto                []PhotoSize                `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       *bool `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      *bool `json:"group_chat_created,omitempty"`
	SupergroupChatCreated *bool `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    *bool `json:"channel_chat_created,omitempty"`
	// MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID   *int64 `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID *int64 `json:"migrate_from_chat_id,omitempty"`
	// PinnedMessage               *MaybeInaccessibleMessage  `json:"pinned_message,omitempty"`
	// Invoice                     *Invoice                   `json:"invoice,omitempty"`
	// SuccessfulPayment           *SuccessfulPayment         `json:"successful_payment,omitempty"`
	// RefundedPayment             *RefundedPayment           `json:"refunded_payment,omitempty"`
	// UsersShared                 *UsersShared               `json:"users_shared,omitempty"`
	// ChatShared                  *ChatShared                `json:"chat_shared,omitempty"`
	ConnectedWebsite *string `json:"connected_website,omitempty"`
	// WriteAccessAllowed          *WriteAccessAllowed        `json:"write_access_allowed,omitempty"`
	// PassportData                *PassportData              `json:"passport_data,omitempty"`
	// ProximityAlertTriggered     *ProximityAlertTriggered   `json:"proximity_alert_triggered,omitempty"`
	// BoostAdded                  *ChatBoostAdded            `json:"boost_added,omitempty"`
	// ChatBackgroundSet           *ChatBackground            `json:"chat_background_set,omitempty"`
	// ForumTopicCreated           *ForumTopicCreated         `json:"forum_topic_created,omitempty"`
	// ForumTopicEdited            *ForumTopicEdited          `json:"forum_topic_edited,omitempty"`
	// ForumTopicClosed            *ForumTopicClosed          `json:"forum_topic_closed,omitempty"`
	// ForumTopicReopened          *ForumTopicReopened        `json:"forum_topic_reopened,omitempty"`
	// GeneralForumTopicHidden     *GeneralForumTopicHidden   `json:"general_forum_topic_hidden,omitempty"`
	// GeneralForumTopicUnhidden   *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`
	// GiveawayCreated             *GiveawayCreated           `json:"giveaway_created,omitempty"`
	// Giveaway                    *Giveaway                  `json:"giveaway,omitempty"`
	// GiveawayWinners             *GiveawayWinners           `json:"giveaway_winners,omitempty"`
	// GiveawayCompleted           *GiveawayCompleted         `json:"giveaway_completed,omitempty"`
	// VideoChatScheduled          *VideoChatScheduled        `json:"video_chat_scheduled,omitempty"`
	// VideoChatStarted            *VideoChatStarted          `json:"video_chat_started,omitempty"`
	// VideoChatEnded              *VideoChatEnded            `json:"video_chat_ended,omitempty"`
	// VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	// WebAppData                  *WebAppData                `json:"web_app_data,omitempty"`
	// ReplyMarkup                 *InlineKeyboardMarkup      `json:"reply_markup,omitempty"`
} // Note: Definitions for all referenced types like User, Chat, etc. should be implemented separately.
