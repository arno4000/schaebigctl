package bot

import "time"

type ThreadMessage []struct {
	ID              string    `json:"id"`
	Type            int       `json:"type"`
	Content         string    `json:"content"`
	ChannelID       string    `json:"channel_id"`
	Author          Author    `json:"author"`
	Attachments     []any     `json:"attachments"`
	Embeds          []any     `json:"embeds"`
	Mentions        []any     `json:"mentions"`
	MentionRoles    []any     `json:"mention_roles"`
	Pinned          bool      `json:"pinned"`
	MentionEveryone bool      `json:"mention_everyone"`
	Tts             bool      `json:"tts"`
	Timestamp       time.Time `json:"timestamp"`
	EditedTimestamp any       `json:"edited_timestamp"`
	Flags           int       `json:"flags"`
	Components      []any     `json:"components"`
	Position        int       `json:"position"`
}
type Author struct {
	ID                   string `json:"id"`
	Username             string `json:"username"`
	Avatar               string `json:"avatar"`
	Discriminator        string `json:"discriminator"`
	PublicFlags          int    `json:"public_flags"`
	Flags                int    `json:"flags"`
	Banner               any    `json:"banner"`
	AccentColor          any    `json:"accent_color"`
	GlobalName           string `json:"global_name"`
	AvatarDecorationData any    `json:"avatar_decoration_data"`
	BannerColor          any    `json:"banner_color"`
}
