# Try
golang 的 try catch

### 例子1

```
package main

import (
	. "github.com/ghostljj/Try"
	"log"
)

func main() {
	err := Try(func() {
		c := calculation()
		log.Println(c, "不可能输出拉")
	})
	log.Println(err)
	log.Println("接着输出哈")
}
func calculation() int {
	a := 100
	b := 0
	c := a / b
	return c
}

```


### 例子2 手动中断，抛出异常

```
package main

import (
	. "github.com/ghostljj/Try"
	"log"
)

func main() {
	err := Try(func() {
		panic("手动中断，下面不继续")
		log.Println(c, "不可能输出拉")
	})
	log.Println(err)
	log.Println("接着输出哈")
}

```