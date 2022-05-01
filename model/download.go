package model

// 根据文件的uuid下载文件, 需要鉴权
type DownloadByUUID struct{}

// 通过分享下载
type DownloadBySharingReq struct {
	ShareID  string `json:"share_id" form:"share_id"`
	Password string `json:"password" form:"password"`
}
