//
// Package cryptoit
// @Author: feymanlee@gmail.com
// @Description:
// @File:  utils.go
// @Date: 2023-04-23 12:24:28
//
package cryptoit

import "fmt"

func AnyToByte(src interface{}) []byte {
	if src == nil {
		return nil
	}
	switch value := src.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		return []byte(fmt.Sprint(src))
	}
}
