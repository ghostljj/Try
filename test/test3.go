package main

import (
	"fmt"
	. "github.com/ghostljj/Try"
	"github.com/ghostljj/Try/test/myerrors"
	"log"
)

//变量错误 例子1

func main() {
	err := Try(func() {
		boo, err := NameCheck3("")
		if err != nil {
			panic(err)
		}
		log.Println(boo, "不可能输出拉")
	})
	//catch 部分
	switch e := err.(type) {
	case myerrors.MyStringError:
		fmt.Println(e) // ==> MyStringError "NameCheck3 名字不能空"  就是 Error() 方法返回的值
		log.Println("自定义的错误条件", err)
	case error:
		log.Println("默认error对象", err)
	}
	log.Println("接着输出哈")
}

func NameCheck3(name string) (bool, error) {
	if name == "" {
		//直接使用变量,这个最好,可以判断 是否 MyStringError 类型 (good~~!!)
		err := myerrors.MyStringError("NameCheck3 名字不能空")
		return false, err
	}
	return true, nil
}
