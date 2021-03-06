package multi_party_login

import (
	"context"
	"github.com/MichealJl/multi-party-login/driver"
	"github.com/MichealJl/multi-party-login/proto"
)

const (
	MpWechat Platform = iota + 1 // 微信小程序
	MpQQ                         // qq小程序
	MpAlipay                     // 阿里小程序
	QQ                           // qq
	Wechat                       // 微信
)

var loginDriverMap = map[Platform]LoginDriver{
	MpWechat: driver.GetMpWechat(),
	MpQQ:     driver.GetMpQQ(),
	QQ:       driver.GetQQ(),
}

type (
	Platform    int8
	LoginDriver interface {
		SetConf(*proto.Conf)
		Login(ctx context.Context, params interface{}) (*proto.LoginRsp, error)
		GetPhoneInfo(encryptData, iv, sessionKey string) (*proto.PhoneInfo, error)
	}
)

func GetDriver(platform Platform) LoginDriver {
	return loginDriverMap[platform]
}
