package wechat

import (
	"github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppID string
	AppSecret string
}
//该Resolver是实现在Wechat.Srvice的指针上，而不是结构体本身，所以在传递给使用者的时候需要传递wecaht.Service的指针，而不是解耦提本身
//因为本身并未实现OpenIDResolver接口
func (s *Service) Resolver(code string) (token string,err error){

	//accessInfo, err := weapp.Login("wx851020fe449a84e2", "26c95bc945fd0d7ccf7c473ef5a5e7f8", code)
	accessInfo, err := weapp.Login(s.AppID, s.AppSecret, code)
	if err != nil {
		return "", err
	}

	if err = accessInfo.GetResponseError();err != nil {
		return "", err
	}

	//accessStr, err := json.Marshal(accessInfo)
	return accessInfo.OpenID, err
}