package proto

type Conf struct {
	AppID  string
	Secret string
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

type LoginRsp struct {
	ErrCode    int64     `json:"err_code"`
	ErrMsg     string    `json:"err_msg"`
	Data       LoginData `json:"data"`
	OriginData string    `json:"origin_data"` // 请求的原始返回数据
}

type LoginData struct {
	OpenID      string `json:"open_id"`     // 第三方平台唯一标识
	UnionID     string `json:"union_id"`    // 用户在开放平台的唯一标识符（如：微信开放平台）
	SessionKey  string `json:"session_key"` // 会话密钥
	AccessToken string `json:"access_token"`
}

type GetUserInfoRsp struct {
	OpenID    string `json:"openId"`
	UnionID   string `json:"unionId,omitempty"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`
}

type GetPhoneRsp struct {
	PhoneNumber     string `json:"phoneNumber"`     //用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string `json:"countryCode"`     //区号
}
