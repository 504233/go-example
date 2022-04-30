package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// 四则运算
	// 函数式编程实现
	var ja = FourFun("+")
	var jian = FourFun("-")
	var cheng = FourFun("*")
	var chu = FourFun("/")

	res1, _ := ja(2, 3)
	res2, _ := jian(3, 3)
	res3, _ := cheng(4, 3)
	res4, _ := chu(5, 3)
	fmt.Println("加法", res1)
	fmt.Println("减法", res2)
	fmt.Println("乘法", res3)
	fmt.Println("除法", res4)
}

//四则运算
func FourFun(str string) func(interface{}, interface{}) (float64, error) {
	var res float64
	var err error

	return func(a interface{}, b interface{}) (float64, error) {
		af := SetType(a)
		bf := SetType(b)
		switch str {
		case "+":
			res = af + bf
		case "-":
			res = af - bf
		case "*":
			res = af * bf
		case "/":
			if af != 0 {
				res = af / bf
			} else {
				err = errors.New("除数不能为空")
			}
		}
		return res, err
	}
}

//对参数格式进行校验
func SetType(inter interface{}) float64 {
	var res float64
	switch inter.(type) {
	case string:
		strconv.ParseFloat(inter.(string), 64)
	case int:
		res = float64(inter.(int))
	case float64:
		res = inter.(float64)
	}
	return res
}