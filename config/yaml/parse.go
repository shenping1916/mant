package yaml

import (
	"fmt"
)

const (
	TypeString = iota
	TypeInt
	TypeInt16
	TypeInt32
	TypeInt64
	TypeUint
	TypeUint16
	TypeUint32
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeBool
	TypeMap
	TypeStringArray
	TypeIntArray
	TypeInt16Array
	TypeInt32Array
	TypeInt64Array
	TypeUintArray
	TypeUint16Array
	TypeUint32Array
	TypeUint64Array
	TypeFloat32Array
	TypeFloat64Array
)

func (y *Yaml) ExistKey(key string) (interface{}, error) {
	if value, ok := y.Data[key]; ok {
		return value, nil
	}

	return nil, fmt.Errorf("key: [%s] isn't exist\n", key)
}

func (y *Yaml) GetString(key string) (string, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return "", err
		}

		r := y.Reflect(value, TypeString)
		if r.Error() != nil {
			return "", r.Error()
		}

		return r.ToString()
	}

	return "", nil
}

func (y *Yaml) GetStringArray(key string) ([]string, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []string{}, err
		}

		r := y.Reflect(value, TypeStringArray)
		if r.Error() != nil {
			return []string{}, r.Error()
		}

		return r.ToStringArray()
	}

	return []string{}, nil
}

func (y *Yaml) GetInt(key string) (int, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return 0, err
		}

		r := y.Reflect(value, TypeInt)
		if r.Error() != nil {
			return 0, r.Error()
		}

		return r.ToInt()
	}

	return int(0), nil
}

func (y *Yaml) GetIntArray(key string) ([]int, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []int{}, err
		}

		r := y.Reflect(value, TypeIntArray)
		if r.Error() != nil {
			return []int{}, r.Error()
		}

		return r.ToIntArray()
	}

	return []int{}, nil
}

func (y *Yaml) GetInt16(key string) (int16, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return int16(0), err
		}

		r := y.Reflect(value, TypeInt16)
		if r.Error() != nil {
			return int16(0), r.Error()
		}

		return r.ToInt16()
	}

	return int16(0), nil
}

func (y *Yaml) GetInt16Array(key string) ([]int16, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []int16{}, err
		}

		r := y.Reflect(value, TypeInt16Array)
		if r.Error() != nil {
			return []int16{}, r.Error()
		}

		return r.ToInt16Array()
	}

	return []int16{}, nil
}

func (y *Yaml) GetInt32(key string) (int32, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return int32(0), err
		}

		r := y.Reflect(value, TypeInt32)
		if r.Error() != nil {
			return int32(0), r.Error()
		}

		return r.ToInt32()
	}

	return int32(0), nil
}

func (y *Yaml) GetInt32Array(key string) ([]int32, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []int32{}, err
		}

		r := y.Reflect(value, TypeInt32Array)
		if r.Error() != nil {
			return []int32{}, r.Error()
		}

		return r.ToInt32Array()
	}

	return []int32{}, nil
}

func (y *Yaml) GetInt64(key string) (int64, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return int64(0), err
		}

		r := y.Reflect(value, TypeInt64)
		if r.Error() != nil {
			return int64(0), r.Error()
		}

		return r.ToInt64()
	}

	return int64(0), nil
}

func (y *Yaml) GetInt64Array(key string) ([]int64, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []int64{}, err
		}

		r := y.Reflect(value, TypeInt64Array)
		if r.Error() != nil {
			return []int64{}, r.Error()
		}

		return r.ToInt64Array()
	}

	return []int64{}, nil
}

func (y *Yaml) GetUint(key string) (uint, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return uint(0), err
		}

		r := y.Reflect(value, TypeUint)
		if r.Error() != nil {
			return uint(0), r.Error()
		}

		return r.ToUint()
	}

	return uint(0), nil
}

func (y *Yaml) GetUintArray(key string) ([]uint, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []uint{}, err
		}

		r := y.Reflect(value, TypeUintArray)
		if r.Error() != nil {
			return []uint{}, r.Error()
		}

		return r.ToUintArray()
	}

	return []uint{}, nil
}

func (y *Yaml) GetUint16(key string) (uint16, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return uint16(0), err
		}

		r := y.Reflect(value, TypeUint16)
		if r.Error() != nil {
			return uint16(0), r.Error()
		}

		return r.ToUint16()
	}

	return uint16(0), nil
}

func (y *Yaml) GetUint16Array(key string) ([]uint16, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []uint16{}, err
		}

		r := y.Reflect(value, TypeUint16Array)
		if r.Error() != nil {
			return []uint16{}, r.Error()
		}

		return r.ToUint16Array()
	}

	return []uint16{}, nil
}

func (y *Yaml) GetUint32(key string) (uint32, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return uint32(0), err
		}

		r := y.Reflect(value, TypeUint32)
		if r.Error() != nil {
			return uint32(0), r.Error()
		}

		return r.ToUint32()
	}

	return uint32(0), nil
}

func (y *Yaml) GetUint32Array(key string) ([]uint32, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []uint32{}, err
		}

		r := y.Reflect(value, TypeUint32Array)
		if r.Error() != nil {
			return []uint32{}, r.Error()
		}

		return r.ToUint32Array()
	}

	return []uint32{}, nil
}

func (y *Yaml) GetUint64(key string) (uint64, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return uint64(0), err
		}

		r := y.Reflect(value, TypeUint64)
		if r.Error() != nil {
			return uint64(0), r.Error()
		}

		return r.ToUint64()
	}

	return uint64(0), nil
}

func (y *Yaml) GetUint64Array(key string) ([]uint64, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []uint64{}, err
		}

		r := y.Reflect(value, TypeUint64Array)
		if r.Error() != nil {
			return []uint64{}, r.Error()
		}

		return r.ToUint64Array()
	}

	return []uint64{}, nil
}

func (y *Yaml) GetFloat32(key string) (float32, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return float32(0), err
		}

		r := y.Reflect(value, TypeFloat32)
		if r.Error() != nil {
			return float32(0), r.Error()
		}

		return r.ToFloat32()
	}

	return float32(0), nil
}

func (y *Yaml) GetFloat32Array(key string) ([]float32, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []float32{}, err
		}

		r := y.Reflect(value, TypeFloat32Array)
		if r.Error() != nil {
			return []float32{}, r.Error()
		}

		return r.ToFloat32Array()
	}

	return []float32{}, nil
}

func (y *Yaml) GetFloat64(key string) (float64, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return float64(0), err
		}

		r := y.Reflect(value, TypeFloat64)
		if r.Error() != nil {
			return float64(0), r.Error()
		}

		return r.ToFloat64()
	}

	return float64(0), nil
}

func (y *Yaml) GetFloat64Array(key string) ([]float64, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return []float64{}, err
		}

		r := y.Reflect(value, TypeFloat64Array)
		if r.Error() != nil {
			return []float64{}, r.Error()
		}

		return r.ToFloat64Array()
	}

	return []float64{}, nil
}

func (y *Yaml) GetBool(key string) (bool, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return false, err
		}

		r := y.Reflect(value, TypeBool)
		if r.Error() != nil {
			return false, r.Error()
		}

		return r.ToBool()
	}

	return false, nil
}

func (y *Yaml) GetMap(key string) (map[string]interface{}, error) {
	if key != "" {
		value, err := y.ExistKey(key)
		if err != nil {
			return map[string]interface{}{}, err
		}

		r := y.Reflect(value, TypeMap)
		if r.Error() != nil {
			return map[string]interface{}{}, r.Error()
		}

		return r.ToMap()
	}

	return map[string]interface{}{}, nil
}
