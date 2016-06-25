package goutil

import (
	"strconv"
	"errors"
	"fmt"
)

func ToInt64(v interface{}) (intValue int64, err error) {
	switch v.(type) {
	case string:
		if iV, err := strconv.Atoi(v.(string)); err == nil {
			return int64(iV), nil
		} else {
			return 0, err
		}

	case int:
		return int64(v.(int)), nil
	case int8:
		return int64(v.(int8)), nil
	case int16:
		return int64(v.(int16)), nil
	case int32:
		return int64(v.(int32)), nil
	case int64:
		return int64(v.(int64)), nil
	case float32:
		return int64(v.(float32)), nil
	case float64:
		return int64(v.(float64)), nil
	default:
		return 0, errors.New("无效的类型")
	}

}

//转字符串
func ToString(v interface{}) (res string) {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return fmt.Sprintf("%d", v)
	case int8:
		return fmt.Sprintf("%d", v)
	case int16:
		return fmt.Sprintf("%d", v)
	case int32:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case uint:
		return fmt.Sprintf("%d", v)
	case float32:
		return fmt.Sprintf("%.2f", v)
	case float64:
		return fmt.Sprintf("%2f", v)

	default:
		return "error"
	}

}
