package sprigmath

import (
	"math"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

//
// Copied from sprig, BSD license
//

// toFloat64 converts 64-bit floats
func toFloat64(v interface{}) (float64, error) {
	if str, ok := v.(string); ok {
		iv, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, errors.Errorf("cannot convert %v to float64", v)
		}
		return iv, nil
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return float64(val.Int()), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return float64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Bool:
		if val.Bool() == true {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, errors.Errorf("cannot convert %v to float64", v)
	}
}

func toInt(v interface{}) (int, error) {
	//It's not optimal. Bud I don't want duplicate toInt64 code.
	vv, err := toInt64(v)
	return int(vv), err
}

// toInt64 converts integer types to 64-bit integers
func toInt64(v interface{}) (int64, error) {
	if str, ok := v.(string); ok {
		iv, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, errors.Errorf("cannot convert %v to int64", v)
		}
		return iv, nil
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int(), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		tv := val.Uint()
		if tv <= math.MaxInt64 {
			return int64(tv), nil
		}
		return math.MaxInt64, errors.Errorf("%v is too big", tv)
	case reflect.Float32, reflect.Float64:
		return int64(val.Float()), nil
	case reflect.Bool:
		if val.Bool() == true {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, errors.Errorf("cannot convert %v to int64", v)
	}
}

// converts to either an int64 or a float64
func toNumber(v interface{}) (interface{}, error) {
	if str, ok := v.(string); ok {
		iv, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return iv, nil
		}

		fv, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return fv, nil
		}

		return nil, errors.Errorf("%v is not a float64 or int64", v)
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int(), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		tv := val.Uint()
		if tv <= math.MaxInt64 {
			return int64(tv), nil
		}
		return math.MaxInt64, errors.Errorf("%v is too big", tv)
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Bool:
		if val.Bool() == true {
			return 1, nil
		}
		return 0, nil
	default:
		return nil, errors.Errorf("cannot convert %v to float64 or int64", v)
	}
}
