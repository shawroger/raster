package utils

import "log"

// DealError 全局处理错误
func DealError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
