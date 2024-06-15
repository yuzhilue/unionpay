package unionpay

import "encoding/xml"

const (
	QrAPI = "https://qra.95516.com/pay/gateway"
)

// ScannedRequest 提交订单请求
type PayConfig struct {
	XMLName      xml.Name `xml:"xml"`
	Service      string   `xml:"service"`
	Version      string   `xml:"version,omitempty"`
	Charset      string   `xml:"charset,omitempty"`
	SignType     string   `xml:"sign_type,omitempty"`
	MchID        string   `xml:"mch_id"` // 商户id
	OutTradeNo   string   `xml:"out_trade_no"`
	DeviceInfo   string   `xml:"device_info,omitempty"`
	Body         string   `xml:"body"`
	GoodsDetail  string   `xml:"goods_detail,omitempty"`
	SubAppID     string   `xml:"sub_appid,omitempty"`
	Attach       string   `xml:"attach,omitempty"`
	PnrInsIdCd   string   `xml:"pnrInsIdCd,omitempty"`
	NeedReceipt  string   `xml:"need_receipt,omitempty"`
	TotalFee     string   `xml:"total_fee"`
	MchCreateIP  string   `xml:"mch_create_ip"`
	AuthCode     string   `xml:"auth_code"`
	TimeStart    string   `xml:"time_start,omitempty"`
	TimeExpire   string   `xml:"time_expire,omitempty"`
	OpUserID     string   `xml:"op_user_id,omitempty"`
	OpShopID     string   `xml:"op_shop_id,omitempty"`
	OpDeviceID   string   `xml:"op_device_id"`
	GoodsTag     string   `xml:"goods_tag,omitempty"`
	TerminalInfo string   `xml:"terminal_info"`
	NonceStr     string   `xml:"nonce_str"`
	Sign         string   `xml:"sign"`
	SignAgentNo  string   `xml:"sign_agentno,omitempty"`
	GroupNo      string   `xml:"groupno,omitempty"`
	Key          string   `xml:"_"`
}

type CDATAText struct {
	Value string `xml:",cdata"`
}

// 返回数据
type ScannedResponse struct {
	XMLName     xml.Name `xml:"xml"`
	Version     string   `xml:"version"`
	Charset     string   `xml:"charset"`
	SignType    string   `xml:"sign_type"`
	SignAgentno string   `xml:"sign_agentno"`
	Groupno     string   `xml:"groupno"`
	Status      string   `xml:"status"`
	Message     string   `xml:"message"`
	Code        string   `xml:"code"`
	NeedQuery   string   `xml:"need_query"`
}
