package proto

type LoginRsp struct {
	ErrCode int64     `json:"err_code"`
	ErrMsg  string    `json:"err_msg"`
	Data    LoginData `json:"data"`
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
