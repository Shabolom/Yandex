package models

type RequestSearch struct {
	VideoID             string `json:"video_id"`
	TrendingDate        string `json:"trending_date"`
	Title               string `json:"title"`
	ChannelTitle        string `json:"channel_title"`
	CategoryId          int    `json:"category_id"`
	PublishTime         string `json:"publish_time"`
	Tags                string `json:"tags"`
	Views               int    `json:"views"`
	Likes               int    `json:"likes"`
	Dislikes            int    `json:"dislikes"`
	CommentCount        int    `json:"comment_count"`
	ThumbnailLink       string `json:"thumbnail_link"`
	CommentsDisabled    bool   `json:"comments_disabled"`
	RatingsDisabled     bool   `json:"ratings_disabled"`
	VideoErrorOrRemoved bool   `json:"video_error_or_removed"`
	Description         string `json:"description"`
}
