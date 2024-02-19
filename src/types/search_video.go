package types

type SearchVideo struct {
	// The id of the video
	Id string `json:"id"`

	// The provider of the video
	Provider string `json:"provider"`

	// The URL of the video
	Url string `json:"url"`

	// The title of the video
	Title string `json:"title"`

	// The thumbnail of the video
	Thumbnail string `json:"thumbnail"`

	// The embed URL of the video
	EmbedUrl string `json:"embed_url"`

	// The duration of the video in seconds
	Duration int `json:"duration"`

	// The view count of the video
	ViewCount int `json:"view_count"`

	// The user who uploaded the video
	// note: This field is optional
	Uploader *string `json:"uploader"`
}
