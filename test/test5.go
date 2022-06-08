package main

import (
	"fmt"
	. "github.com/ghostljj/Try"
	myerrors "github.com/ghostljj/Try/test/myerrors"
	"log"
)

//结构体错误 例子

func main() {
	err := Try(func() {
		boo, err := NameCheck5("")
		if err != nil {
			panic(err)
		}
		log.Println(boo, "不可能输出拉")
	})
	//catch 部分
	switch e := err.(type) {
	case *ConnectionErrorStruct: //示例3 使用结构体，做条件
		code := e.EC
		fmt.Println(code) // ==> 返回 Null_ERROR  就是 String() 方法返回的值
		switch code {
		case myerrors.ErrCodeNo:
			log.Println("自定义的错误条件1", err)
		case myerrors.ErrCodeNull:
			log.Println("自定义的错误条件2", err)
		}
	case error:
		log.Println("默认error对象", e)
	}
	log.Println("接着输出哈")
}

//ConnectionErrorStruct 使用结构体，做错误，方法不好
type ConnectionErrorStruct struct {
	EC myerrors.ErrCode
}

//Error 结构体继承error的方法
func (e ConnectionErrorStruct) Error() string {
	return fmt.Sprintf("Connection Error: %s", myerrors.ErrCode(e.EC))
}

func NameCheck5(name string) (bool, error) {
	if name == "" {
		// 注意error这里必须是地址&引用 ，interface本身就职指针来的
		// 返回一个Struct，带Error错误函数的
		err := &ConnectionErrorStruct{EC: myerrors.ErrCodeNull}
		return false, err
	}
	return true, nil
}
