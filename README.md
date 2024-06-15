# Unionpay
Go unionpay SDK 非官方

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
https://up.95516.com/open/openapi/doc?index_1=1&index_2=1&chapter_1=235&chapter_2=259
## 使用方法
```shell
go get -u github.com/yuzhilue/unionpay
```

## 功能
**扫码支付**
 - [x]  统一下单(unified.trade.native)
 - [x]  查询订单(unified.trade.query)
 - [x]  关闭订单(unified.trade.close)
 - [x]  申请退款(unified.trade.refund)
 - [x]  查询退款(unified.trade.refundquery)
**被扫支付**
**微信APP支付**
**微信公众号&小程序支付**
**支付宝服务窗支付**
**银联JS支付**