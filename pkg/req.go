package pkg

import "errors"

// 状态码处理
func CodeL(status string) error {
	if status == "0" {
		return nil
	} else if status == "400" {
		return errors.New("签名错误")
	}
	return nil
}
