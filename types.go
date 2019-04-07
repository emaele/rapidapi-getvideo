package getvideo

// Client contains the apikey
type Client struct {
	apikey string
}

// Video is the result of a newRequest and contains all the info about the video requested
type Video struct {
	Status       bool     `json:"status"`
	Message      string   `json:"message"`
	Description  string   `json:"description"`
	Uploader     string   `json:"uploader"`
	URL          string   `json:"url"`
	ID           string   `json:"id"`
	IsPlaylist   bool     `json:"is_playlist"`
	Site         string   `json:"site"`
	Title        string   `json:"title"`
	LikeCount    int      `json:"like_count"`
	DislikeCount int      `json:"dislike_count"`
	ViewCount    int      `json:"view_count"`
	Duration     int      `json:"duration"`
	UploadDate   string   `json:"upload_date"`
	Tags         []string `json:"tags"`
	UploaderURL  string   `json:"uploader_url"`
	Thumbnail    string   `json:"thumbnail"`
	Streams      []stream `json:"streams"`
}

type stream struct {
	URL            string      `json:"url"`
	Format         string      `json:"format"`
	FormatNote     string      `json:"format_note"`
	Extension      string      `json:"extension"`
	VideoCodec     string      `json:"video_codec"`
	AudioCodec     string      `json:"audio_codec"`
	Height         int         `json:"height,omitempty"`
	Width          int         `json:"width,omitempty"`
	Fps            interface{} `json:"fps,omitempty"`
	FmtID          string      `json:"fmt_id"`
	Filesize       int         `json:"filesize"`
	FilesizePretty string      `json:"filesize_pretty"`
	HasAudio       bool        `json:"has_audio"`
	HasVideo       bool        `json:"has_video"`
	IsHd           bool        `json:"is_hd"`
}
