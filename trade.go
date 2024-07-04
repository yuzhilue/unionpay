package unionpay

import (
	"errors"

	"github.com/yuzhilue/unionpay/pkg"
)

// 扫码支付-统一下单
func (da PayConfig) UnifiedTradeNative(body PayBody) (*PayResponse, error) {
	parameters := map[string]string{
		"attach":        body.Attach,
		"body":          body.Body,
		"mch_create_ip": da.MchCreateIP,
		"mch_id":        da.MchID,
		"nonce_str":     da.NonceStr,
		"notify_url":    da.NotifyUrl,
		"out_trade_no":  da.OutTradeNo,
		"service":       "unified.trade.native",
		"sign_type":     da.SignType,
		"total_fee":     body.TotalFee,
		"version":       da.Version,
	}
	sign := pkg.CreateMd5(parameters, da.Key)

	date := map[string]string{
		"attach":        body.Attach,
		"body":          body.Body,
		"mch_create_ip": da.MchCreateIP,
		"mch_id":        da.MchID,
		"nonce_str":     da.NonceStr,
		"notify_url":    da.NotifyUrl,
		"out_trade_no":  da.OutTradeNo,
		"service":       "unified.trade.native",
		"sign_type":     da.SignType,
		"total_fee":     body.TotalFee,
		"version":       da.Version,
		"sign":          sign,
	}
	xmlString := ToXML(date)

	// 请求数据
	resdate, err := NewHttpClient(xmlString)
	if err != nil {
		return nil, err
	}
	// 判断数据
	if resdate.Status.Value == "0" {
		return resdate, nil
	} else {
		return nil, errors.New("error creating")
	}
}

// 查询订单
func (da PayConfig) UnifiedTradeQuery(body PayBody) (*PayResponse, error) {
	//parameters := map[string]string{
	//	"attach":    body.Attach,
	//	"body":      body.Body,
	//	"total_fee": body.TotalFee,
	//}
	parameters := map[string]string{
		"mch_id":       da.MchID,
		"nonce_str":    da.NonceStr,
		"out_trade_no": da.OutTradeNo,
		"service":      "unified.trade.query",
		"sign_type":    da.SignType,
		"version":      da.Version,
	}
	sign := pkg.CreateMd5(parameters, da.Key)
	date := map[string]string{
		"mch_id":       da.MchID,
		"nonce_str":    da.NonceStr,
		"out_trade_no": da.OutTradeNo,
		"service":      "unified.trade.query",
		"sign_type":    da.SignType,
		"version":      da.Version,
		"sign":         sign,
	}
	xmlString := ToXML(date)

	resp, err := NewHttpClient(xmlString)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 关闭订单
func (da PayConfig) UnifiedTradeClose(body PayBody) (*PayResponse, error) {
	parameters := map[string]string{
		"mch_id":       da.MchID,
		"nonce_str":    da.NonceStr,
		"out_trade_no": da.OutTradeNo,
		"service":      "unified.trade.close",
		"sign_type":    da.SignType,
		"version":      da.Version,
	}
	sign := pkg.CreateMd5(parameters, da.Key)
	date := map[string]string{
		"mch_id":       da.MchID,
		"nonce_str":    da.NonceStr,
		"out_trade_no": da.OutTradeNo,
		"service":      "unified.trade.close",
		"sign_type":    da.SignType,
		"version":      da.Version,
		"sign":         sign,
	}
	xmlString := ToXML(date)

	resp, err := NewHttpClient(xmlString)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 申请退款
func (da PayConfig) UnifiedTradeRefund() {

}

// 查询退款
func (da PayConfig) UnifiedTradeRefundquery() {}

// 支付通知
func (da PayConfig) UnifiedTradeNotify() {

}
