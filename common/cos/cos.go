package cos

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var (
	Cli *cos.Client
)

const (
	AvatarDir     = "avatar"
	AppLogoDir    = "app_logo"
	PluginLogoDir = "plugin_logo"
)

func InitCos(cosURL, secretID, secretKey string) {
	u, _ := url.Parse(cosURL)
	b := &cos.BaseURL{BucketURL: u}
	Cli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
}
