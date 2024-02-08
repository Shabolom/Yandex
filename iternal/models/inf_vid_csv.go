package models

type InfVidCsv struct {
	VideoID             string `csv:"video_id"`
	TrendingDate        string `csv:"trending_date"`
	Title               string `csv:"title"`
	ChannelTitle        string `csv:"channel_title"`
	CategoryId          int    `csv:"category_id"`
	PublishTime         string `csv:"publish_time"`
	Tags                string `csv:"tags"`
	views               int    `csv:"views"`
	Likes               int    `csv:"likes"`
	Dislikes            int    `csv:"dislikes"`
	CommentCount        int    `csv:"comment_count"`
	ThumbnailLink       string `csv:"thumbnail_link"`
	CommentsDisabled    bool   `csv:"comments_disabled"`
	RatingsDisabled     bool   `csv:"ratings_disabled"`
	VideoErrorOrRemoved bool   `csv:"video_error_or_removed"`
	Description         string `csv:"description"`
}
