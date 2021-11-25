package multi_party_login

import (
	"context"
	"fmt"
	"github.com/MichealJl/multi-party-login/proto"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	driver := GetDriver(QQ)
	driver.SetAppId("101867382")
	ret, err := driver.Login(ctx,proto.ReqQQLoginParams{
		AccessToken: "a",
		OpenId:      "a",
	})
	fmt.Println(ret, err)
}