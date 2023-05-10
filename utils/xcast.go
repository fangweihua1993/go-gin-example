// @Description
// @Author yixia
// @Copyright 2021 sndks.com. All rights reserved.
// @LastModify 2021/1/14 5:21 下午

package utils

// copied from spf13/cast

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ToBool
//
//	@Description 将空接口转换为 bool 忽略错误
//	@Param i
//	@Return bool
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToTime casts an empty interface to time ignoring error
//
//	@Description
//	@param i
//	@return time.Time
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToDuration casts an empty interface to duration ignoring error
//
//	@Description
//	@param i
//	@return time.Duration
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToFloat64 casts an empty interface to float64 ignoring error
//
//	@Description
//	@param i
//	@return float64
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToInt64 casts an empty interface to int64 ignoring error
//
//	@Description
//	@param i
//	@return int64
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt casts an empty interface to int ignoring error
//
//	@Description
//	@param i
//	@return int
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

// ToString casts an empty interface to string ignoring error
//
//	@Description
//	@param i
//	@return string
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

// ToStringMapString casts an empty interface to map[string]string ignoring error
//
//	@Description
//	@param i
//	@return map[string]string
func ToStringMapString(i interface{}) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStringMapStringSlice casts an empty interface to map[string][]string ignoring error
//
//	@Description
//	@param i
//	@return map[string][]string
func ToStringMapStringSlice(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

// ToStringMapBool casts an empty interface to map[string]bool ignoring error
func ToStringMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStringMap casts an empty interface to string map ignoring error
func ToStringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

// ToSlice casts an empty interface []interface ignoring error
//
//	@Description
//	@param i
//	@return []interface{}
func ToSlice(i interface{}) []interface{} {
	v, _ := ToSliceE(i)
	return v
}

// ToStringSlice casts an empty interface to []string ignoring error
//
//	@Description
//	@param i
//	@return []string
func ToStringSlice(i interface{}) []string {
	v, _ := ToStringSliceE(i)
	return v
}

// ToSliceStringMap casts an empty interface to []map[string]interface{} ignoring error
//
//	@Description
//	@param i
//	@return []map[string]interface{}
func ToSliceStringMap(i interface{}) []map[string]interface{} {
	v, _ := ToSliceStringMapE(i)
	return v
}

// ToIntSlice casts an empty interface to []int ignoring error
//
//	@Description
//	@param i
//	@return []int
func ToIntSlice(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToTimeE casts an empty interface to time.Time.
//
//	@Description
//	@param i
//	@return tim
//	@return err
func ToTimeE(i interface{}) (tim time.Time, err error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Time:
		return s, nil
	case string:
		d, e := StringToDate(s)
		if e == nil {
			return d, nil
		}
		return time.Time{}, fmt.Errorf("could not parse Date/Time format: %v\n", e)
	default:
		return time.Time{}, fmt.Errorf("unable to Cast %#v to Time\n", i)
	}
}

// ToDurationE casts an empty interface to time.Duration.
//
//	@Description
//	@param i
//	@return d
//	@return err
func ToDurationE(i interface{}) (d time.Duration, err error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int64:
		d = time.Duration(s)
		return
	case float64:
		d = time.Duration(s)
		return
	case string:
		d, err = time.ParseDuration(s)
		return
	default:
		err = fmt.Errorf("unable to Cast %#v to Duration\n", i)
		return
	}
}

// ToBoolE casts an empty interface to a bool.
//
//	@Description
//	@param i
//	@return bool
//	@return error
func ToBoolE(i interface{}) (bool, error) {
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("unable to Cast %#v to bool", i)
	}
}

// ToFloat64E casts an empty interface to a float64.
//
//	@Description
//	@param i
//	@return float64
//	@return error
func ToFloat64E(i interface{}) (float64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case int:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return float64(v), nil
		}
		return 0.0, fmt.Errorf("unable to Cast %#v to float", i)
	default:
		return 0.0, fmt.Errorf("unable to Cast %#v to float", i)
	}
}

// ToInt64E casts an empty interface to an int64.
//
//	@Description
//	@param i
//	@return int64
//	@return error
func ToInt64E(i interface{}) (int64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return s, nil
	case int:
		return int64(s), nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to Cast %#v to int64", i)
	case float64:
		return int64(s), nil
	case bool:
		if bool(s) {
			return int64(1), nil
		}
		return int64(0), nil
	case nil:
		return int64(0), nil
	default:
		return int64(0), fmt.Errorf("unable to Cast %#v to int64", i)
	}
}

// ToIntE casts an empty interface to an int.
//
//	@Description
//	@param i
//	@return int
//	@return error
func ToIntE(i interface{}) (int, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to Cast %#v to int", i)
	case float64:
		return int(s), nil
	case bool:
		if bool(s) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to Cast %#v to int", i)
	}
}

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToStringE casts an empty interface to a string.
//
//	@Description
//	@param i
//	@return string
//	@return error
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64), nil
	case int64:
		return strconv.FormatInt(i.(int64), 10), nil
	case int:
		return strconv.FormatInt(int64(i.(int)), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to Cast %#v to string", i)
	}
}

// ToStringMapStringE casts an empty interface to a stringmap
//
//	@Description
//	@param i
//	@return map[string]string
//	@return error
func ToStringMapStringE(i interface{}) (map[string]string, error) {
	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[string]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToString(val)
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			m[ToString(k)] = ToString(val)
		}
		return m, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToString(val)
		}
		return m, nil
	default:
		return m, fmt.Errorf("unable to Cast %#v to map[string]string", i)
	}
}

// ToStringMapStringSliceE cast an empty interface to a string array map
//
//	@Description
//	@param i
//	@return map[string][]string
//	@return error
func ToStringMapStringSliceE(i interface{}) (map[string][]string, error) {
	var m = map[string][]string{}

	switch v := i.(type) {
	case map[string][]string:
		return v, nil
	case map[string][]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[string]string:
		for k, val := range v {
			m[ToString(k)] = []string{val}
		}
	case map[string]interface{}:
		for k, val := range v {
			//m[ToString(k)] = []string{ToString(val)}
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}][]string:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}][]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToStringSlice(val)
		}
		return m, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			key, err := ToStringE(k)
			if err != nil {
				return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
			}
			value, err := ToStringSliceE(val)
			if err != nil {
				return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
			}
			m[key] = value

		}
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
	}
	return m, nil
}

// ToStringMapBoolE casts an empty interface to a map[string]bool.
//
//	@Description
//	@param i
//	@return map[string]bool
//	@return error
func ToStringMapBoolE(i interface{}) (map[string]bool, error) {
	var m = map[string]bool{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToBool(val)
		}
		return m, nil
	case map[string]interface{}:
		for k, val := range v {
			m[ToString(k)] = ToBool(val)
		}
		return m, nil
	case map[string]bool:
		return v, nil
	default:
		return m, fmt.Errorf("unable to Cast %#v to map[string]bool", i)
	}
}

// ToStringMapE casts an empty interface to a map[string]interface{}.
//
//	@Description
//	@param i
//	@return map[string]interface{}
//	@return error
func ToStringMapE(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = val
		}
		return m, nil
	case map[string]interface{}:
		return v, nil
	case map[string]string:
		for k, v := range v {
			m[k] = v
		}
		return m, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]interface{}", i)
	}
}

// ToSliceE casts an empty interface to a []interface{}.
//
//	@Description
//	@param i
//	@return []interface{}
//	@return error
func ToSliceE(i interface{}) ([]interface{}, error) {
	var s = make([]interface{}, 0)

	switch v := i.(type) {
	case []interface{}:
		s = append(s, v...)
		return s, nil
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("Unable to Cast %#v of type %v to []interface{}", i, reflect.TypeOf(i))
	}
}

// ToSliceStringMapE casts an empty interface to a []interface{}.
//
//	@Description
//	@param i
//	@return []map[string]interface{}
//	@return error
func ToSliceStringMapE(i interface{}) ([]map[string]interface{}, error) {
	var s = make([]map[string]interface{}, 0)

	switch v := i.(type) {
	case []interface{}:
		for _, u := range v {
			s = append(s, ToStringMap(u))
		}
		return s, nil
	case []map[string]interface{}:
		s = append(s, v...)
		return s, nil
	default:
		return s, fmt.Errorf("Unable to Cast %#v of type %v to []map[string]interface{}", i, reflect.TypeOf(i))
	}
}

// ToStringSliceE casts an empty interface to a []string.
//
//	@Description
//	@param i
//	@return []string
//	@return error
func ToStringSliceE(i interface{}) ([]string, error) {
	var a = make([]string, 0)

	switch v := i.(type) {
	case []interface{}:
		for _, u := range v {
			a = append(a, ToString(u))
		}
		return a, nil
	case []string:
		return v, nil
	case string:
		return strings.Fields(v), nil
	case interface{}:
		str, err := ToStringE(v)
		if err != nil {
			return a, fmt.Errorf("Unable to Cast %#v to []string", i)
		}
		return []string{str}, nil
	default:
		return a, fmt.Errorf("Unable to Cast %#v to []string", i)
	}
}

// ToIntSliceE casts an empty interface to a []int.
//
//	@Description
//	@param i
//	@return []int
//	@return error
func ToIntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, fmt.Errorf("Unable to Cast %#v to []int", i)
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("Unable to Cast %#v to []int", i)
			}
			a[j] = val

		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("Unable to Cast %#v to []int", i)
	}
}

// StringToDate casts an empty interface to a time.Time.
//
//	@Description
//	@param s
//	@return time.Time
//	@return error
func StringToDate(s string) (time.Time, error) {
	return parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05Z07:00",
		"02 Jan 06 15:04 MST",
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
	})

}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
}
