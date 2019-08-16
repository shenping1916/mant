package yaml

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type String struct {
	val string
}

func (s *String) Error() error {
	return nil
}

func (s *String) ToString() (string, error) {
	return s.val, nil
}

func (s *String) ToStringArray() ([]string, error) {
	if s.val == "" {
		return []string{}, nil
	}

	return []string{s.val}, nil
}

func (s *String) ToInt() (int, error) {
	_int64, err := s.ToInt64()
	if err != nil {
		return 0, err
	}

	return int(_int64), nil
}

func (s *String) ToIntArray() ([]int, error) {
	_int, err := s.ToInt()
	if err != nil {
		return []int{}, err
	}

	return []int{_int}, nil
}

func (s *String) ToInt16() (int16, error) {
	_int64, err := s.ToInt64()
	if err != nil {
		return int16(0), err
	}

	return int16(_int64), nil
}

func (s *String) ToInt16Array() ([]int16, error) {
	_int16, err := s.ToInt16()
	if err != nil {
		return []int16{}, err
	}

	return []int16{_int16}, nil
}

func (s *String) ToInt32() (int32, error) {
	_int64, err := s.ToInt64()
	if err != nil {
		return int32(0), err
	}

	return int32(_int64), nil
}

func (s *String) ToInt32Array() ([]int32, error) {
	_int32, err := s.ToInt32()
	if err != nil {
		return []int32{}, err
	}

	return []int32{_int32}, nil
}

func (s *String) ToInt64() (int64, error) {
	if s.val == "" {
		return int64(0), nil
	}

	var multiple int64
	startPos := 0
	endPos := 0

	if s.val[0] == '-' {
		startPos = 1
		multiple = -1
	} else {
		multiple = 1
	}

	for i := startPos; i < len(s.val); i++ {
		if s.val[i] >= '0' || s.val[0] <= '9' {
			endPos = i + 1
		} else {
			break
		}
	}

	r, err := strconv.ParseInt(s.val[startPos:endPos], 10, 64)
	if err != nil {
		return 0, err
	}

	return r * multiple, nil
}

func (s *String) ToInt64Array() ([]int64, error) {
	_int64, err := s.ToInt64()
	if err != nil {
		return []int64{}, err
	}

	return []int64{_int64}, nil
}

func (s *String) ToUint() (uint, error) {
	_uint64, err := s.ToUint64()
	if err != nil {
		return 0, err
	}

	return uint(_uint64), nil
}

func (s *String) ToUintArray() ([]uint, error) {
	_uint, err := s.ToUint()
	if err != nil {
		return []uint{}, err
	}

	return []uint{_uint}, nil
}

func (s *String) ToUint16() (uint16, error) {
	_uint64, err := s.ToUint64()
	if err != nil {
		return uint16(0), err
	}

	return uint16(_uint64), nil
}

func (s *String) ToUint16Array() ([]uint16, error) {
	_uint16, err := s.ToUint16()
	if err != nil {
		return []uint16{}, err
	}

	return []uint16{_uint16}, nil
}

func (s *String) ToUint32() (uint32, error) {
	_uint64, err := s.ToUint64()
	if err != nil {
		return uint32(0), err
	}

	return uint32(_uint64), nil
}

func (s *String) ToUint32Array() ([]uint32, error) {
	_uint32, err := s.ToUint32()
	if err != nil {
		return []uint32{}, err
	}

	return []uint32{_uint32}, nil
}

func (s *String) ToUint64() (uint64, error) {
	if s.val == "" {
		return uint64(0), nil
	}

	startPos := 0
	endPos := 0

	if s.val[0] == '-' {
		return uint64(0), errors.New("the number is illegal and can't be negative")
	}

	for i := startPos; i < len(s.val); i++ {
		if s.val[i] >= '0' || s.val[0] <= '9' {
			endPos = i + 1
		} else {
			break
		}
	}

	return strconv.ParseUint(s.val[startPos:endPos], 10, 64)
}

func (s *String) ToUint64Array() ([]uint64, error) {
	_uint64, err := s.ToUint64()
	if err != nil {
		return []uint64{}, err
	}

	return []uint64{_uint64}, nil
}

func (s *String) ToFloat32() (float32, error) {
	_float64, err := s.ToFloat64()
	if err != nil {
		return float32(0), err
	}

	return float32(_float64), nil
}

func (s *String) ToFloat32Array() ([]float32, error) {
	_float32, err := s.ToFloat32()
	if err != nil {
		return []float32{}, err
	}

	return []float32{_float32}, nil
}

func (s *String) ToFloat64() (float64, error) {
	if s.val == "" {
		return float64(0), nil
	}

	if s.val[0] != '+' && s.val[0] != '-' && (s.val[0] > '9' || s.val[0] < '0') {
		return float64(0), errors.New("illegal float")
	}

	endPos := 1
	for i := 1; i < len(s.val); i++ {
		if s.val[i] == '.' || s.val[i] == 'e' || s.val[i] == 'E' || s.val[i] == '+' || s.val[i] == '-' {
			endPos = i + 1
			continue
		}

		if s.val[i] >= '0' && s.val[i] <= '9' {
			endPos = i + 1
		} else {
			endPos = i
			break
		}
	}

	return strconv.ParseFloat(s.val[:endPos], 64)
}

func (s *String) ToFloat64Array() ([]float64, error) {
	_float64, err := s.ToFloat64()
	if err != nil {
		return []float64{}, err
	}

	return []float64{_float64}, nil
}

func (s *String) ToBool() (bool, error) {
	if s.val != "" {
		switch s.val {
		case "true", "True", "TRUE":
			return true, nil
		case "false", "False", "FALSE", "", "0":
			return false, nil
		default:
			return false, fmt.Errorf("invalid bool value: %s", s.val)
		}
	}

	return false, nil
}

func (s *String) ToMap() (map[string]interface{}, error) {
	if s.val == "" {
		return nil, nil
	}

	var m map[string]interface{}
	err := json.Unmarshal([]byte(s.val), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
