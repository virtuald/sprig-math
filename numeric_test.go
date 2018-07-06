package sprigmath

import (
	"testing"

	"github.com/pkg/errors"
)

func testFloat(expected float64, init interface{}, errstr string) error {
	v, err := toFloat64(init)
	if err = testError(errstr, err); err != nil {
		return err
	}

	if errstr == "" && !floatEquals(v, expected) {
		return errors.Errorf("Expected %v, got %v", expected, v)
	}

	return nil
}

func TestToFloat64(t *testing.T) {
	var err error

	if err = testFloat(102, int8(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, int(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, int32(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, int16(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, int64(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, "102", ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, "bob", "cannot convert bob to float64"); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, uint16(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102, uint64(102), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(102.1234, float64(102.1234), ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(0, false, ""); err != nil {
		t.Error(err)
	}
	if err = testFloat(1, true, ""); err != nil {
		t.Error(err)
	}
}

func testInt64(expected int64, init interface{}, errstr string) error {
	v, err := toInt64(init)
	if err = testError(errstr, err); err != nil {
		return err
	}

	if errstr == "" && v != expected {
		return errors.Errorf("Expected %v, got %v", expected, v)
	}

	return nil
}

func TestToInt64(t *testing.T) {
	var err error

	if err = testInt64(102, int8(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, int(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, int32(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, int16(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, int64(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, "102", ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, "bob", "cannot convert bob to int64"); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, uint16(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, uint64(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(102, float64(102.1234), ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(0, false, ""); err != nil {
		t.Error(err)
	}
	if err = testInt64(1, true, ""); err != nil {
		t.Error(err)
	}
}

func testInt(expected int, init interface{}, errstr string) error {
	v, err := toInt(init)
	if err = testError(errstr, err); err != nil {
		return err
	}

	if errstr == "" && v != expected {
		return errors.Errorf("Expected %v, got %v", expected, v)
	}

	return nil
}

func TestToInt(t *testing.T) {
	var err error

	if err = testInt(102, int8(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, int(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, int32(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, int16(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, int64(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, "102", ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, "bob", "cannot convert bob to int64"); err != nil {
		t.Error(err)
	}
	if err = testInt(102, uint16(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, uint64(102), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(102, float64(102.1234), ""); err != nil {
		t.Error(err)
	}
	if err = testInt(0, false, ""); err != nil {
		t.Error(err)
	}
	if err = testInt(1, true, ""); err != nil {
		t.Error(err)
	}
}

func TestToNumber(t *testing.T) {
	if v, err := toNumber(5); err == nil {
		switch vv := v.(type) {
		case int64:
			if err = testInt64(5, vv, ""); err != nil {
				t.Error(err)
			}
		default:
			t.Errorf("Expected int64, got %T", v)
		}
	} else {
		t.Error(err)
	}

	if v, err := toNumber(5.5); err == nil {
		switch vv := v.(type) {
		case float64:
			if err = testFloat(5.5, vv, ""); err != nil {
				t.Error(err)
			}
		default:
			t.Errorf("Expected float64, got %T", v)
		}
	} else {
		t.Error(err)
	}

	if v, err := toNumber("5"); err == nil {
		switch vv := v.(type) {
		case int64:
			if err = testInt64(5, vv, ""); err != nil {
				t.Error(err)
			}
		default:
			t.Errorf("Expected int64, got %T", v)
		}
	} else {
		t.Error(err)
	}

	if v, err := toNumber("5.5"); err == nil {
		switch vv := v.(type) {
		case float64:
			if err = testFloat(5.5, vv, ""); err != nil {
				t.Error(err)
			}
		default:
			t.Errorf("Expected float64, got %T", v)
		}
	} else {
		t.Error(err)
	}
}
