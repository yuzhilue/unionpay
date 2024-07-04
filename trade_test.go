package unionpay_test

import (
	"fmt"
	"github.com/yuzhilue/unionpay/pkg"
	"testing"

	"github.com/yuzhilue/unionpay"
)

// 统一下单
func TestUnifiedTradeNative(t *testing.T) {
	conf := unionpay.PayConfig{
		MchID:       "QRA29045311KKR1",
		Key:         "979da4cfccbae7923641daa5dd7047c2",
		SignType:    "MD5",
		NotifyUrl:   "/",
		MchCreateIP: "127.0.0.1",
		OutTradeNo:  "517201871001475",
		NonceStr:    pkg.GenerateRandomString(),
	}
	body := unionpay.PayBody{
		Attach:   "附加信息",
		Body:     "商品描述",
		TotalFee: "1",
	}
	conf.SetPayVersion()
	da, err := conf.UnifiedTradeNative(body)
	if err != nil {
		t.Error(err)
		fmt.Println("Error ", err)
	}

	fmt.Println("sccess ", da)
}

// 测试设置pay版本
func TestPayConfig_SetPayVersion(t *testing.T) {
	conf := unionpay.PayConfig{}
	conf.SetPayVersion()
	fmt.Println(conf)
}

// 测试查询订单
func TestUnifiedTradeQuery(t *testing.T) {
	conf := unionpay.PayConfig{
		MchID:       "QRA29045311KKR1",
		Key:         "979da4cfccbae7923641daa5dd7047c2",
		SignType:    "MD5",
		MchCreateIP: "127.0.0.1",
		OutTradeNo:  "517201871001475",
		NonceStr:    pkg.GenerateRandomString(),
	}
	query, err := conf.UnifiedTradeQuery(unionpay.PayBody{
		Body: "测试商品",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(query)
}

// 关闭订单
func TestUnifiedTradeCloset(t *testing.T) {

}

// 申请退款

// 查询退款
