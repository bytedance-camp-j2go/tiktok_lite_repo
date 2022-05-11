package init

import (
	"fmt"
	"os"
	"testing"
)

// 多版本测试
func TestMultiVersionViper(t *testing.T) {
	env := os.Getenv("GO_ENV")
	fmt.Println("GO_ENV >>", env)
	if env != "" {
		fmt.Println(env)
	} else {
		fmt.Println("def")
	}
}
