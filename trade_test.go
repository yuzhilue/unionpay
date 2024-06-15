package unionpay_test

import (
	"fmt"
	"testing"

	"github.com/yuzhilue/unionpay"
)

func TestUnifiedTradeNative(t *testing.T) {
	conf := unionpay.PayConfig{
		MchID:       "QRA29045311KKR1",
		Key:         "979da4cfccbae7923641daa5dd7047c2",
		SignType:    "MD5",
		NotifyUrl:   "/",
		MchCreateIP: "127.0.0.1",
		OutTradeNo:  "517201871001475",
		NonceStr:    "838603111",
	}
	body := unionpay.PayBody{
		Attach:   "附加信息",
		Body:     "商品描述",
		TotalFee: "1",
	}
	conf.SetPayVersion()
	da, err := conf.UnifiedTradeNative(body)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("sccess ", da)
}

func TestPayConfig_SetPayVersion(t *testing.T) {
	conf := unionpay.PayConfig{}
	conf.SetPayVersion()
	fmt.Println(conf)
}
