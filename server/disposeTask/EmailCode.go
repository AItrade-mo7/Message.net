package disposeTask

import (
	"fmt"
)

func UpdateEmailCode(to []string, code string) {
	fmt.Println("存储验证码", to, code)
}
