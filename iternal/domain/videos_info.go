package domain

type VideoInfo struct {
	VideoID             string `gorm:"primaryKey; column:video_id; type:text"`
	TrendingDate        string `gorm:"column:trending_date; type:text"`
	Title               string `gorm:"column:title; type:text"`
	ChannelTitle        string `gorm:"column:channel_title; type:text"`
	CategoryId          int    `gorm:"column:category_id; type:int"`
	PublishTime         string `gorm:"column:publish_time; type:text"`
	Tags                string `gorm:"column:tags; type:text"`
	views               int    `gorm:"column:views; type:int"`
	Likes               int    `gorm:"column:likes; type:int"`
	Dislikes            int    `gorm:"column:dislikes; type:int"`
	CommentCount        int    `gorm:"column:comment_count; type:int"`
	ThumbnailLink       string `gorm:"column:thumbnail_link; type:text"`
	CommentsDisabled    bool   `gorm:"column:comments_disabled; type:bool"`
	RatingsDisabled     bool   `gorm:"column:ratings_disabled; type:bool"`
	VideoErrorOrRemoved bool   `gorm:"column:video_error_or_removed; type:bool"`
	Description         string `gorm:"column:description; type:text"`
}
