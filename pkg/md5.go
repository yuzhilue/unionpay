package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func CreateMd5(parameters map[string]string, key string) string {
	signPars := ""
	keys := make([]string, 0, len(parameters))
	for k := range parameters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := parameters[k]
		if v != "" && k != "sign" {
			signPars += k + "=" + v + "&"
		}
	}
	signPars += "key=" + key
	fmt.Println(signPars)
	hash := md5.Sum([]byte(signPars))
	sign := strings.ToUpper(hex.EncodeToString(hash[:]))

	return sign
}
