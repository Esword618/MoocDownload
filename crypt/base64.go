/*
 * @Author: Esword
 * @Description:
 * @FileName:  base64
 * @Version: 1.0.0
 * @Date: 2022-07-16 13:08
 */

package crypt

import (
	"encoding/base64"
)

func Base64StdDecode(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
