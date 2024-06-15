# Unionpay
Go unionpay(银联支付) SDK 非官方

## 示例

```go
package main

import (
    "fmt"
    "github.com/yuzhilue/unionpay"
)

func main() {
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
	conf.UnifiedTradeNative(body)
}
```
## 使用方法
```shell
go get -u github.com/yuzhilue/unionpay
```

## 功能
**扫码支付**
 - [x]  统一下单(unified.trade.native)
 - [ ]  查询订单(unified.trade.query)
 - [ ]  关闭订单(unified.trade.close)
 - [ ]  申请退款(unified.trade.refund)
 - [ ]  查询退款(unified.trade.refundquery)

**被扫支付**
 - [ ]  提交被扫支付(unified.trade.micropay)

**微信APP支付**
 - [ ]  统一下单(pay.weixin.raw.app)

**微信公众号&小程序支付**
 - [ ]  统一下单(unified.trade.native)

**支付宝服务窗支付**
 - [ ]  统一下单(unified.trade.native)

**银联JS支付**
 - [ ]  支付下单(pay.unionpay.jspay)
