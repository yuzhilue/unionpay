package unionpay

const (
	QrAPI                   = "https://qra.95516.com/pay/gateway" // api
	UNIFIEDTRADENATIVE      = "unified.trade.native"              // 扫码支付-提交订单
	UNIFIEDTRADEQUERY       = "unified.trade.query"               // 扫码支付-查询订单
	UNIFIEDTRADECLOSE       = "unified.trade.close"               // 扫码支付-关闭订单
	UNIFIEDTRADEREFUND      = "unified.trade.refund"              // 扫码支付-申请退款
	UNIFIEDTRADEREFUNDQUERY = "unified.trade.refundquery"         // 扫码支付-查询退款
	ERRMESSAGE              = "SIGNATURE_ERROR"                   // 签名错误,参数签名结果不正确
)

type PayConfig struct {
	Version      string `xml:"version,omitempty"`      // API版本
	Charset      string `xml:"charset,omitempty"`      // 字符串编码
	SignType     string `xml:"sign_type,omitempty"`    // 加密方式
	MchID        string `xml:"mch_id"`                 // 商户id
	OutTradeNo   string `xml:"out_trade_no"`           //  商户订单号
	DeviceInfo   string `xml:"device_info,omitempty"`  // 终端设备号
	SubAppID     string `xml:"sub_appid,omitempty"`    // 微信公众平台基本配置中的AppID(应用ID)，传入后支付成功可返回对应公众号下的用户openid
	PnrInsIdCd   string `xml:"pnrInsIdCd,omitempty"`   // 银联服务商机构标识码
	NeedReceipt  string `xml:"need_receipt,omitempty"` // 需要和微信公众平台的发票功能联合，传入true时，微信支付成功消息和支付详情页将出现开票入口[新增need_receipt【适用于微信】]
	MchCreateIP  string `xml:"mch_create_ip"`          // 订单生成的机器 IP
	OpUserID     string `xml:"op_user_id,omitempty"`   // 收银员帐号,默认为商户号
	OpShopID     string `xml:"op_shop_id,omitempty"`   // 门店编号
	GoodsTag     string `xml:"goods_tag,omitempty"`    // 商品标记，微信平台配置的商品标记，用于优惠券或者满减使用
	TerminalInfo string `xml:"terminal_info"`          // 商户侧受理终端信息，字段详细说明参考统一被扫其他扩展字段说明。
	NonceStr     string `xml:"nonce_str"`              // 随机字符串
	NotifyUrl    string `xml:"notify_url"`             // 通知地址
	Sign         string `xml:"sign"`                   // 签名
	SignAgentNo  string `xml:"sign_agentno,omitempty"` // 授权交易的服务商机构代码，商户授权给服务商交易的情况下必填，签名使用服务商的密钥
	GroupNo      string `xml:"groupno,omitempty"`      // 连锁商户为其下门店发交易的情况必填，签名使用连锁商户的密钥
	Key          string `xml:"_"`                      // 密钥
}

type PayBody struct {
	Attach     string `xml:"attach,omitempty"` // 商户附加信息，可做扩展参数
	Body       string `xml:"body"`             // 商户及商品名称
	TotalFee   string `xml:"total_fee"`        // 金额
	TimeStart  string `xml:"time_start"`       // 订单生成时间 格式为yyyyMMddHHmmss 时区为GMT+8 beijing 订单生成时间与超时时间需要同时传入才会生效。
	TimeExpire string `xml:"time_expire"`      // 订单失效时间 格式为yyyyMMddHHmmss

}

type PayResponse struct {
	Charset    CDATAText `xml:"charset"`
	CodeImgURL CDATAText `xml:"code_img_url"` // 支付二维码
	CodeURL    CDATAText `xml:"code_url"`     // 支付连接
	Code       CDATAText `xml:"code"`         // 网关返回码
	MchID      CDATAText `xml:"mch_id"`       // 商户ID
	NonceStr   CDATAText `xml:"nonce_str"`    // 随机字符串
	ResultCode CDATAText `xml:"result_code"`  // 业务结果 0表示成功 非0表示失败
	Sign       CDATAText `xml:"sign"`         // 签名
	SignType   CDATAText `xml:"sign_type"`    // 签名信息
	Message    CDATAText `xml:"message"`      // 提示信息
	DeviceInfo CDATAText `xml:"device_info"`  // 终端设备号
	Status     CDATAText `xml:"status"`       // 状态
	UUID       CDATAText `xml:"uuid"`         // 唯一识别码
	ErrCode    CDATAText `xml:"err_code"`     // 错误代码
	Version    CDATAText `xml:"version"`
}

type CDATAText struct {
	Value string `xml:",cdata"`
}
