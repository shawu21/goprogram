// 加密密码
package helper

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Getmd5(password string) string {
	w := md5.New()
	_, err := io.WriteString(w, password)
	if err != nil {
		panic(err)
	}
	password = fmt.Sprintf("%x", w.Sum(nil))
	return password
}
