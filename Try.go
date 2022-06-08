//Try catch
//
//我就要使用Try catch
//defer 在 return 执行之后执行的，抛出异常马上执行。
//panic 可以传送东西给recover
//recover 可以接受很多东西，数字 字符串，对象，什么都可以。

package Try

import (
	"errors"
	"fmt"
	"reflect"
)

//Try catch 无敌好用 返回error 对象
func Try(runFun func()) (tryErr error) {
	tryErr = nil
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				tryErr = err.(error)
			case string:
				tryErr = errors.New(err.(string))
			default:
				tryErr = errors.New(fmt.Sprintf("未知错误类型 value:%v type:", err, reflect.TypeOf(err)))
			}
		}
	}()
	runFun()
	return
}
