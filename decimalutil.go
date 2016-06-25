package goutil

import (
	"strconv"
	"errors"
	"math/big"
	"fmt"
)

//现金加减乘除
func DecimalAdd(first, second string, precision uint) (res string, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}
	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	resFloatNum := big.NewFloat(float64(0))
	resFloatNum = resFloatNum.Add(firBig, secBig)
	resFloatNum.SetPrec(big.MaxPrec)
	formStr := "%." + fmt.Sprintf("%d", precision) + "f"
	formValue, _ := resFloatNum.Float64()
	return fmt.Sprintf(formStr, formValue), nil
	return resFloatNum.String(), nil

}
//现金加减乘除
func DecimalSub(first, second string, precision uint) (res string, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}
	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	resFloatNum := big.NewFloat(float64(0))
	resFloatNum = resFloatNum.Sub(firBig, secBig)
	resFloatNum.SetPrec(big.MaxPrec)
	formStr := "%." + fmt.Sprintf("%d", precision) + "f"
	formValue, _ := resFloatNum.Float64()
	return fmt.Sprintf(formStr, formValue), nil

}
//现金加减乘除
func DecimalMul(first, second string, precision uint) (res string, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	resFloatNum := big.NewFloat(float64(0))
	resFloatNum = resFloatNum.Mul(firBig, secBig)
	resFloatNum.SetPrec(big.MaxPrec)
	formStr := "%." + fmt.Sprintf("%d", precision) + "f"
	formValue, _ := resFloatNum.Float64()
	return fmt.Sprintf(formStr, formValue), nil
}
//精确除法
func DecimalDiv(first, second string, precision uint) (res string, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second || "0" == second {
		return "", errors.New("除数不能为0")
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return "", errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	resFloatNum := big.NewFloat(float64(0))
	resFloatNum = resFloatNum.Quo(firBig, secBig)
	resFloatNum.SetPrec(big.MaxPrec)
	formStr := "%." + fmt.Sprintf("%d", precision) + "f"
	formValue, _ := resFloatNum.Float64()
	return fmt.Sprintf(formStr, formValue), nil
}
//两个数字比较小于
func DecimalLess(first, second string) (resb bool, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	res := firBig.Cmp(secBig)
	if res == -1 {
		return true, nil
	} else {
		return false, nil
	}
}
//两个数字比较 小于 等于
func DecimalLessEqual(first, second string) (resb bool, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	res := firBig.Cmp(secBig)
	if res < 1 {
		return true, nil
	} else {
		return false, nil
	}
}

//两个数字比较 等于
func DecimalEqual(first, second string) (resb bool, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	res := firBig.Cmp(secBig)
	if res == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

//两个数字比较大于
func DecimalGrant(first, second string) (resb bool, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	res := firBig.Cmp(secBig)
	if res == 1 {
		return true, nil
	} else {
		return false, nil
	}
}


//两个数字比较大于等于
func DecimalGrantEqual(first, second string) (resb bool, err error) {
	if "" == first {
		first = "0"
	}
	if "" == second {
		second = "0"
	}

	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + first)
	}
	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return false, errors.New("无法解析的字符串：" + second)
	}
	firBig := big.NewFloat(firstNum)
	secBig := big.NewFloat(secondNum)
	res := firBig.Cmp(secBig)
	if res >= 0 {
		return true, nil
	} else {
		return false, nil
	}
}