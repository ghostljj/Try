package main

import (
	"fmt"
	. "github.com/ghostljj/Try"
	"github.com/ghostljj/Try/test/myerrors"
	"log"
)

//变量错误 例子2

func main() {
	err := Try(func() {
		boo, err := NameCheck4("")
		if err != nil {
			panic(err)
		}
		log.Println(boo, "不可能输出拉")
	})
	//catch 部分
	switch e := err.(type) {
	case myerrors.ConnectionError:
		code := myerrors.ErrCode(e)
		fmt.Println(code) // ==> 返回 Null_ERROR  就是 String() 方法返回的值
		switch code {
		case myerrors.ErrCodeNo:
			log.Println("自定义的错误条件1", err)
		case myerrors.ErrCodeNull:
			log.Println("自定义的错误条件2", err)
		}
	case error:
		log.Println("默认error对象", err)
	}
	log.Println("接着输出哈")
}

func NameCheck4(name string) (bool, error) {
	if name == "" {
		//直接使用变量,这个最好,可以判断 是否 ConnectionError 类型 (good~~!!)
		err := myerrors.ConnectionError(myerrors.ErrCodeNull)
		return false, err
	}
	return true, nil
}
