package number

import (
	"fmt"
	"math"
	"testing"
)

// 在 Go 中直接定义常量表达式，xxx/0，编译时不报错，运行时 panic
// 但是如果是变量计算值，分母出现 0.0，并不会 panic（除以 0 依旧 panic），会得到一个特殊的 float64 类型的值
// 这个值通过 fmt 打印出来是 +Inf（为什么能打印出 +Inf，详见 go\src\strconv\ftoa.go/85）
func TestInf(t *testing.T) {
	// i1, i2 := 1, 0
	// i := i1 / i2 // 直接 panic
	// fmt.Println(i)

	f1 := 1.0
	f := f1 / 0
	fmt.Println(f)
	fmt.Println(math.IsInf(f, 0))
}
