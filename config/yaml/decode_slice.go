package yaml

import (
	"fmt"
)

type slice struct {
	val List
}

func (s *slice) Error() error {
	return nil
}

func (s *slice) ToString() (string, error) {
	return "", nil
}

func (s *slice) ToStringArray() ([]string, error) {
	if s.val != nil {
		var data = make([]string, 0, len(s.val))
		for _, value := range s.val {
			data = append(data, fmt.Sprint(value))
		}

		return data, nil
	}
	return []string{}, nil
}

func (s *slice) ToInt() (int, error) {
	return 0, nil
}

func (s *slice) ToIntArray() ([]int, error) {
	if s.val != nil {
		var data = make([]int, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToInt()
				if err != nil {
					return []int{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []int{}, nil
}

func (s *slice) ToInt16() (int16, error) {
	return int16(0), nil
}

func (s *slice) ToInt16Array() ([]int16, error) {
	if s.val != nil {
		var data = make([]int16, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToInt16()
				if err != nil {
					return []int16{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []int16{}, nil
}

func (s *slice) ToInt32() (int32, error) {
	return int32(0), nil
}

func (s *slice) ToInt32Array() ([]int32, error) {
	if s.val != nil {
		var data = make([]int32, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToInt32()
				if err != nil {
					return []int32{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []int32{}, nil
}

func (s *slice) ToInt64() (int64, error) {
	return int64(0), nil
}

func (s *slice) ToInt64Array() ([]int64, error) {
	if s.val != nil {
		var data = make([]int64, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToInt64()
				if err != nil {
					return []int64{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []int64{}, nil
}

func (s *slice) ToUint() (uint, error) {
	return uint(0), nil
}

func (s *slice) ToUintArray() ([]uint, error) {
	if s.val != nil {
		var data = make([]uint, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToUint()
				if err != nil {
					return []uint{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []uint{}, nil
}

func (s *slice) ToUint16() (uint16, error) {
	return uint16(0), nil
}

func (s *slice) ToUint16Array() ([]uint16, error) {
	if s.val != nil {
		var data = make([]uint16, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToUint16()
				if err != nil {
					return []uint16{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []uint16{}, nil
}

func (s *slice) ToUint32() (uint32, error) {
	return uint32(0), nil
}

func (s *slice) ToUint32Array() ([]uint32, error) {
	if s.val != nil {
		var data = make([]uint32, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToUint32()
				if err != nil {
					return []uint32{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []uint32{}, nil
}

func (s *slice) ToUint64() (uint64, error) {
	return uint64(0), nil
}

func (s *slice) ToUint64Array() ([]uint64, error) {
	if s.val != nil {
		var data = make([]uint64, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToUint64()
				if err != nil {
					return []uint64{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []uint64{}, nil
}

func (s *slice) ToFloat32() (float32, error) {
	return float32(0), nil
}

func (s *slice) ToFloat32Array() ([]float32, error) {
	if s.val != nil {
		var data = make([]float32, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToFloat32()
				if err != nil {
					return []float32{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []float32{}, nil
}

func (s *slice) ToFloat64() (float64, error) {
	return float64(0), nil
}

func (s *slice) ToFloat64Array() ([]float64, error) {
	if s.val != nil {
		var data = make([]float64, 0, len(s.val))
		for _, value := range s.val {
			switch value := value.(type) {
			case string:
				d := &decode{val: value}
				i, err := d.ToFloat64()
				if err != nil {
					return []float64{}, err
				}

				data = append(data, i)
			}
		}

		return data, nil
	}
	return []float64{}, nil
}

func (s *slice) ToBool() (bool, error) {
	return false, nil
}

func (s *slice) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
