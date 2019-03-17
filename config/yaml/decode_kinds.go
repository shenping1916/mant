package yaml

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type decode struct {
	val string
}

func (d *decode) Error() error {
	return nil
}

func (d *decode) ToString() (string, error) {
	return d.val, nil
}

func (d *decode) ToStringArray() ([]string, error) {
	if d.val == "" {
		return []string{}, nil
	}

	return []string{d.val}, nil
}

func (d *decode) ToInt() (int, error) {
	_int64, err := d.ToInt64()
	if err != nil {
		return 0, err
	}

	return int(_int64), nil
}

func (d *decode) ToIntArray() ([]int, error) {
	_int, err := d.ToInt()
	if err != nil {
		return []int{}, err
	}

	return []int{_int}, nil
}

func (d *decode) ToInt16() (int16, error) {
	_int64, err := d.ToInt64()
	if err != nil {
		return int16(0), err
	}

	return int16(_int64), nil
}

func (d *decode) ToInt16Array() ([]int16, error) {
	_int16, err := d.ToInt16()
	if err != nil {
		return []int16{}, err
	}

	return []int16{_int16}, nil
}

func (d *decode) ToInt32() (int32, error) {
	_int64, err := d.ToInt64()
	if err != nil {
		return int32(0), err
	}

	return int32(_int64), nil
}

func (d *decode) ToInt32Array() ([]int32, error) {
	_int32, err := d.ToInt32()
	if err != nil {
		return []int32{}, err
	}

	return []int32{_int32}, nil
}

func (d *decode) ToInt64() (int64, error) {
	if d.val == "" {
		return int64(0), nil
	}

	var multiple int64
	startPos := 0
	endPos := 0

	if d.val[0] == '-' {
		startPos = 1
		multiple = -1
	} else {
		multiple = 1
	}

	for i := startPos; i < len(d.val); i++ {
		if d.val[i] >= '0' || d.val[0] <= '9' {
			endPos = i + 1
		} else {
			break
		}
	}

	r, err := strconv.ParseInt(d.val[startPos:endPos], 10, 64)
	if err != nil {
		return 0, err
	}

	return r * multiple, nil
}

func (d *decode) ToInt64Array() ([]int64, error) {
	_int64, err := d.ToInt64()
	if err != nil {
		return []int64{}, err
	}

	return []int64{_int64}, nil
}

func (d *decode) ToUint() (uint, error) {
	_uint64, err := d.ToUint64()
	if err != nil {
		return 0, err
	}

	return uint(_uint64), nil
}

func (d *decode) ToUintArray() ([]uint, error) {
	_uint, err := d.ToUint()
	if err != nil {
		return []uint{}, err
	}

	return []uint{_uint}, nil
}

func (d *decode) ToUint16() (uint16, error) {
	_uint64, err := d.ToUint64()
	if err != nil {
		return uint16(0), err
	}

	return uint16(_uint64), nil
}

func (d *decode) ToUint16Array() ([]uint16, error) {
	_uint16, err := d.ToUint16()
	if err != nil {
		return []uint16{}, err
	}

	return []uint16{_uint16}, nil
}

func (d *decode) ToUint32() (uint32, error) {
	_uint64, err := d.ToUint64()
	if err != nil {
		return uint32(0), err
	}

	return uint32(_uint64), nil
}

func (d *decode) ToUint32Array() ([]uint32, error) {
	_uint32, err := d.ToUint32()
	if err != nil {
		return []uint32{}, err
	}

	return []uint32{_uint32}, nil
}

func (d *decode) ToUint64() (uint64, error) {
	if d.val == "" {
		return uint64(0), nil
	}

	startPos := 0
	endPos := 0

	if d.val[0] == '-' {
		return uint64(0), errors.New("the number is illegal and can't be negative")
	}

	for i := startPos; i < len(d.val); i++ {
		if d.val[i] >= '0' || d.val[0] <= '9' {
			endPos = i + 1
		} else {
			break
		}
	}

	return strconv.ParseUint(d.val[startPos:endPos], 10, 64)
}

func (d *decode) ToUint64Array() ([]uint64, error) {
	_uint64, err := d.ToUint64()
	if err != nil {
		return []uint64{}, err
	}

	return []uint64{_uint64}, nil
}

func (d *decode) ToFloat32() (float32, error) {
	_float64, err := d.ToFloat64()
	if err != nil {
		return float32(0), err
	}

	return float32(_float64), nil
}

func (d *decode) ToFloat32Array() ([]float32, error) {
	_float32, err := d.ToFloat32()
	if err != nil {
		return []float32{}, err
	}

	return []float32{_float32}, nil
}

func (d *decode) ToFloat64() (float64, error) {
	if d.val == "" {
		return float64(0), nil
	}

	if d.val[0] != '+' && d.val[0] != '-' && (d.val[0] > '9' || d.val[0] < '0') {
		return float64(0), errors.New("illegal float")
	}

	endPos := 1
	for i := 1; i < len(d.val); i++ {
		if d.val[i] == '.' || d.val[i] == 'e' || d.val[i] == 'E' || d.val[i] == '+' || d.val[i] == '-' {
			endPos = i + 1
			continue
		}

		if d.val[i] >= '0' && d.val[i] <= '9' {
			endPos = i + 1
		} else {
			endPos = i
			break
		}
	}

	return strconv.ParseFloat(d.val[:endPos], 64)
}

func (d *decode) ToFloat64Array() ([]float64, error) {
	_float64, err := d.ToFloat64()
	if err != nil {
		return []float64{}, err
	}

	return []float64{_float64}, nil
}

func (d *decode) ToBool() (bool, error) {
	if d.val != "" {
		switch d.val {
		case "true", "True", "TRUE":
			return true, nil
		case "false", "False", "FALSE", "", "0":
			return false, nil
		default:
			return false, fmt.Errorf("invalid bool value: %s", d.val)
		}
	}

	return false, nil
}

func (d *decode) ToMap() (map[string]interface{}, error) {
	if d.val == "" {
		return nil, nil
	}

	var m map[string]interface{}
	err := json.Unmarshal([]byte(d.val), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
