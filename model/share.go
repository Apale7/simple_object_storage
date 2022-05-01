package model

type ShareFileReq struct {
	FileID   uint    `json:"file_id"`  // 文件链接的id
	NeedPwd  bool    `json:"need_pwd"` // 是否需要提取码
	Password *string `json:"password"` // 提取码
	Duration int64   `json:"duration"` // 分享码有效期，单位秒
}

type ShareFileRes struct {
	ShareID  string  `json:"share_id"`           // 分享码
	Password *string `json:"password,omitempty"` // 提取码
}
