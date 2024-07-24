package util

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"time"
)

func AbsEqual[T any](a T, b T) bool {
	return reflect.DeepEqual(a, b)
}

func Decimal[T float32 | float64](v T, dp int) (float64, error) {
	value, err := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(dp)+"f", v), 64)

	if err != nil {
		return 0, err
	}

	return value, nil
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func NumberToString[T Number](v T, dp int) string {
	typeof := reflect.ValueOf(v).Kind().String()
	str := ""

	if dp == 0 {
		dp = -1
	}

	switch typeof {
	case "int64":
		str = strconv.FormatInt(int64(v), 10)
	case "float32":
		str = strconv.FormatFloat(float64(v), 'f', dp, 32)
	case "float64":
		str = strconv.FormatFloat(float64(v), 'f', dp, 64)
	default:
		str = strconv.Itoa(int(v))
	}

	return str
}

func StringToNumber[T Number](v string, numberType string, dp int) (T, error) {
	d, err := strconv.Atoi(v)

	if err != nil {
		return 0, err
	}

	switch numberType {
	case "int32":
		return T(int32(d)), nil
	case "int64":
		return T(int64(d)), nil
	case "float32":
		return T(float32(d)), nil
	case "float64":
		return T(float64(d)), nil
	default:
		return T(int32(d)), nil
	}
}

func MapHasKey[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]

	return ok
}

func StructToJSON(s any) (string, error) {
	typeof := reflect.ValueOf(s).Kind().String()
	// ptr
	if typeof != "struct" {
		return "{}", nil
	}

	jsonByte, err := json.Marshal(s)

	if err != nil {
		return "{}", err
	}

	return string(jsonByte), nil
}

func JSONToStruct[T any](j string, s *T) (any, error) {
	typeof := reflect.ValueOf(*s).Kind().String()

	if typeof != "struct" {
		return nil, nil
	}

	err := json.Unmarshal([]byte(j), s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func MapToStruct[T any](m map[string]any, s *T) (*T, error) {
	err := mapstructure.Decode(m, &s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func StructToMap(s any) (map[string]any, error) {
	m := make(map[string]any)
	jsonData, err := StructToJSON(s)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonData), &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func SliceUnique[T comparable](s []T) (map[T]int, []T) {
	tmp := make(map[T]int)
	var ns []T

	for i := 0; i < len(s); i++ {
		if MapHasKey(tmp, s[i]) {
			tmp[s[i]] += 1
		} else {
			tmp[s[i]] = 1
		}
	}

	for k := range tmp {
		ns = append(ns, k)
	}

	return tmp, ns
}

func DateTimeFormat(t time.Time, ds string, ts string) string {
	return t.Format("2006" + ds + "01" + ds + "02 15" + ts + "04" + ts + "05")
}
