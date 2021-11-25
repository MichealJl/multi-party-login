package proto

type ReqMpWechatLoginParams struct {
	Code string `json:"code" valid:"required"`
}

type MpUserInfo struct {
	OpenID    string    `json:"openId"`
	UnionID   string    `json:"unionId,omitempty"`
	NickName  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarURL string    `json:"avatarUrl"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark,omitempty"`
}

type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       Watermark `json:"watermark"`
}

type Watermark struct {
	AppId     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}
