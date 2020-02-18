package model

import (
	"log"
)

//	Insert 插入操作

// CheckErr 用来校验error对象是否为空
func CheckErr(err error,msg string)  {
	if nil != err {
		log.Panicln(msg,err)
	}
}