package dbcustom

import (
	"fmt"
	"reflect"
	"strings"
)

/**
  此处为编写公共通用方法
*/

var (
	CurlyBracesLeft = 123
	SquareBracketsLeft = 91
	ParenthesesLeft = 40
	AngleBracketsLeft = 60
	DiagonalLeft = 47
  // 对称
	CurlyBracesRight = 125
	SquareBracketsRight = 93
	ParenthesesRight = 41
	AngleBracketsRight = 62
	DiagonalRight = 92
)


// 字符串反序
func StringReverse(str string) string {
	// 字符串转字节
	bytes := []byte(str)
	for i := 0; i < len(str)/2; i++ {
		// 定义一个变量存放从后往前的值
		tmp := bytes[len(str)-i-1]
		// 从后往前的值跟从前往后的值调换
		bytes[len(str)-i-1] = bytes[i]
		// 从前往后的值跟从后往前的值进行调换
		bytes[i] = tmp
	}
	return string(bytes)
}

//获取字符串 转义对称符号后 字符串
func GetSymmetrySymbol(str string) string {
	var bye = []byte(str)
	reBye := make([]byte,len(bye))
	for k,v := range bye{
		var temp byte
		switch v {
			case byte(CurlyBracesLeft): temp = byte(CurlyBracesRight)
			case byte(SquareBracketsLeft): temp = byte(SquareBracketsRight)
			case byte(ParenthesesLeft): temp = byte(ParenthesesRight)
			case byte(AngleBracketsLeft): temp = byte(AngleBracketsRight)
			case byte(DiagonalLeft): temp = byte(DiagonalRight)
			case byte(CurlyBracesRight): temp = byte(CurlyBracesLeft)
			case byte(SquareBracketsRight): temp = byte(SquareBracketsLeft)
			case byte(ParenthesesRight): temp = byte(ParenthesesLeft)
			case byte(AngleBracketsRight): temp = byte(AngleBracketsLeft)
			case byte(DiagonalRight): temp = byte(DiagonalLeft)
			default:
				temp = v
			}
		reBye[k] = temp
	}
	return string(reBye)
}

//以cp为轴args两边对称 字符串
func CopySymbol(cp interface{}, args ...string) string {
	strSuf := fmt.Sprint(strings.Join(args,""),cp)
	var strPre string
	for i := 0; i < len(args); i++{
		strPre = fmt.Sprint(strPre,GetSymmetrySymbol(StringReverse(args[len(args)-i-1])))
	}
	return fmt.Sprint(strSuf,strPre)
}

// interface类型 判空, 反射传入field().interface()
// 常量：default value is true
// func、interface、ptr nil value is true
// struct has no value is true
// map、slice、channel len==0 is true
func IsNULL(i interface{}) bool {
	if i == nil{
		return true
	}
	vof := reflect.ValueOf(i)
	switch vof.Kind() {
		case reflect.Map,reflect.Slice :
			return vof.Len() == 0
		default: return vof.IsZero()
	}
}
