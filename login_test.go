package multi_party_login

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	driver := GetDriver(MpWechat)
	driver.SetAppId("appId")
	driver.SetSecret("secret")
	ret, err := driver.Login(ctx,"code")
	fmt.Println(ret, err)
}