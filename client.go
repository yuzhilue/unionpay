package unionpay

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

// 创建支付配置
func NewPayClient() *PayConfig {
	return nil
}

func DefaultPayClient() *Client {
	return &Client{HttpClient: &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     true,
			ForceAttemptHTTP2:     true,
		},
	}}
}

// 创建http客户端
func NewHttpClient(da string) (*PayResponse, error) {
	var resp PayResponse
	res, err := DefaultPayClient().HttpClient.Post(QrAPI, "application/xml", bytes.NewBuffer([]byte(da)))
	if err != nil {
		return nil, errors.New("http client error: " + err.Error())
	}
	defer res.Body.Close()
	xmldate, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("read response body failed!")
	}
	err = xml.Unmarshal(xmldate, &resp)
	if err != nil {
		return nil, errors.New("XML parsing error")
	}
	return &resp, nil
}

// 设置APi版本
func (da *PayConfig) SetPayVersion() *PayConfig {
	da.Version = "2.0"
	return da
}

// 设置加密方式
func (da *PayConfig) SetPaySignType(signType string) *PayConfig {
	if signType == "SM3" {
		da.SignType = "SM3"

	} else {
		da.SignType = "MD5"
	}
	return da
}

func (da *PayConfig) SetCharset(notifyUrl string) *PayConfig {
	da.Charset = "UTF-8"
	return da
}

type XMLData struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

// 转换格式
func ToXML(data map[string]string) string {
	buf := new(bytes.Buffer)
	buf.WriteString("<xml>")
	for k, v := range data {
		xmlData := XMLData{
			XMLName: xml.Name{Local: k},
			Value:   v,
		}
		xmlBytes, err := xml.Marshal(xmlData)
		if err != nil {
			fmt.Println("XML marshaling error:", err)
			return ""
		}
		buf.Write(xmlBytes)
	}
	buf.WriteString("</xml>")
	return buf.String()
}

// 解析xml
func PareseXML(da string) *PayResponse {
	var data PayResponse
	err := xml.Unmarshal([]byte(da), &data)
	if err != nil {
		return nil
	}
	return &data
}
