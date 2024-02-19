package types

type Video struct {
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

	// The user who uploaded the video
	Uploader string `json:"uploader"`

	// The duration of the video in seconds
	Duration int `json:"duration"`

	// The view count of the video
	ViewCount int `json:"view_count"`

	// The tags of the video
	Tags []string `json:"tags"`

	// The categories of the video
	Categories []string `json:"categories"`

	// The pornstars in the video
	Pornstars []string `json:"pornstars"`

	// The video's likes
	Likes int `json:"likes"`

	// The video's dislikes
	Dislikes int `json:"dislikes"`
}
