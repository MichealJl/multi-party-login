package proto

type ReqQQLoginParams struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"open_id"`
}

type QQRsp struct {
	Ret          int64  `json:"ret"`
	Msg          string `json:"msg"`
	Nickname     string `json:"nickname"`
	Figureurl    string `json:"figureurl"`
	Figureurl1   string `json:"figureurl_1"`
	Figureurl2   string `json:"figureurl_2"`
	FigureurlQq1 string `json:"figureurl_qq_1"`
	FigureurlQq2 string `json:"figureurl_qq_2"`
	Gender       string `json:"gender"`
}
