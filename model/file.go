package model

type FileListReq struct {
	Username *string `json:"username,omitempty" form:"username,omitempty"`
	IsPublic *bool  `json:"is_public,omitempty" form:"is_public,omitempty"`
	Filename *string `json:"filename,omitempty" form:"filename,omitempty"`
}

type FileListResp struct {
	Files []FileInfo `json:"files"`
}

type FileInfo struct {
	Filename string `json:"filename"`
	Username string `json:"username"`
	Size     int64  `json:"size"`
}
