package vimeo

import "gopkg.in/mgo.v2/bson"

//VideoService interacts with the video services provided by the vimeo video API.

type videoURI struct {
	URI string `json:"uri"`
}

type responseURI struct {
	Total   int        `json:"total"`
	Page    int        `json:"page"`
	Perpage int        `json:"per_page"`
	Paging  paging     `json:"paging"`
	Data    []videoURI `json:"data"`
}

type paging struct {
	Next     string  `json:"next"`
	Previous *paging `json:"previous"`
	First    string  `json:"first"`
	Last     string  `json:"last"`
}

//VideoDetails is a generic struct for kinds of video data
type VideoDetails struct {
	ID                     bson.ObjectId `json:"id" bson:"_id"`
	Taxonomy               string        `json:"taxonomy" bson:"taxonomy"`
	SearchTerm             string        `json:"search_term" bson:"search_term"`
	SearchType             string        `json:"search_type" bson:"search_type"`
	SearchLevel            int           `json:"search_level" bson:"search_level"`
	ParentVideoID          string        `json:"parent_video_id" bson:"parent_video_id"`
	YouplusID              string        `json:"youplus_id" bson:"youplus_id"`
	YouplusSourceURL       string        `json:"youplus_source_url" bson:"youplus_source_url"`
	Source                 string        `json:"source" bson:"source"`
	SourceID               string        `json:"source_id" bson:"source_id"`
	SourceURL              string        `json:"source_url" bson:"source_url"`
	SourceCountry          string        `json:"source_country" bson:"source_country"`
	Title                  string        `json:"title" bson:"title"`
	Description            string        `json:"description" bson:"description"`
	Caption                string        `json:"caption" bson:"caption"`
	VideoThumbnail         string        `json:"thumbnail_image_url" bson:"thumbnail_image_url"`
	VideoLanguage          string        `json:"video_language" bson:"video_language"`
	VideoDuration          uint64        `json:"video_duration" bson:"video_duration"`
	YoutubeCategory        string        `json:"yt_category" bson:"yt_category"`
	VideoTags              string        `json:"tags" bson:"tags"`
	ChannelID              string        `json:"channel_id" bson:"channel_id"`
	ChannelTitle           string        `json:"channel_title" bson:"channel_title"`
	ChannelThumbnail       string        `json:"channel_thumbnail" bson:"channel_thumbnail"`
	ChannelVideoCount      uint64        `json:"channel_video_count" bson:"channel_video_count"`
	ChannelSubscriberCount uint64        `json:"channel_subscriber_count" bson:"channel_subscriber_count"`
	ChannelViewCount       uint64        `json:"channel_view_count" bson:"channel_view_count"`
	ChannelCommentCount    uint64        `json:"channel_comment_count" bson:"channel_comment_count"`
	ChannelCreatedAt       string        `json:"channel_created_at" bson:"channel_created_at"`
	AvgViewsPerDay         int           `json:"avg_views_per_day" bson:"avg_views_per_day"`
	PublishedAt            string        `json:"published_at" bson:"published_at"`
	ViewCount              uint64        `json:"view_count" bson:"view_count"`
	LikeCount              uint64        `json:"like_count" bson:"like_count"`
	DislikeCount           uint64        `json:"dislike_count" bson:"dislike_count"`
	CommentCount           uint64        `json:"comment_count" bson:"comment_count"`
	FavouriteCount         uint64        `json:"favourite_count" bson:"favourite_count"`
	IsSponsored            bool          `json:"is_sponsored" bson:"is_sponsored"`
	IsDownloaded           bool          `json:"is_downloaded" bson:"is_downloaded"`
	IsTranscribed          bool          `json:"is_transcribed" bson:"is_transcribed"`
	LastCrawledAt          string        `json:"last_crawled_at" bson:"last_crawled_at"`
	UpdatedAt              string        `json:"updated_at" bson:"updated_at"`
}
