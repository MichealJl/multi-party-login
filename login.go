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
}

type (
	Platform    int8
	LoginDriver interface {
		SetAppId(appId string)
		SetSecret(secret string)
		Login(ctx context.Context, code string) (*proto.LoginRsp, error)
	}
)

func GetDriver(platform Platform) LoginDriver {
	return loginDriverMap[platform]
}
