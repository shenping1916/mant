package yaml

type Map struct {
	val map[string]interface{}
}

func (m *Map) Error() error {
	return nil
}

func (m *Map) ToString() (string, error) {
	return "", nil
}

func (m *Map) ToStringArray() ([]string, error) {
	return []string{}, nil
}

func (m *Map) ToInt() (int, error) {
	return 0, nil
}

func (m *Map) ToIntArray() ([]int, error) {
	return []int{}, nil
}

func (m *Map) ToInt16() (int16, error) {
	return int16(0), nil
}

func (m *Map) ToInt16Array() ([]int16, error) {
	return []int16{}, nil
}

func (m *Map) ToInt32() (int32, error) {
	return int32(0), nil
}

func (m *Map) ToInt32Array() ([]int32, error) {
	return []int32{}, nil
}

func (m *Map) ToInt64() (int64, error) {
	return int64(0), nil
}

func (m *Map) ToInt64Array() ([]int64, error) {
	return []int64{}, nil
}

func (m *Map) ToUint() (uint, error) {
	return uint(0), nil
}

func (m *Map) ToUintArray() ([]uint, error) {
	return []uint{}, nil
}

func (m *Map) ToUint16() (uint16, error) {
	return uint16(0), nil
}

func (m *Map) ToUint16Array() ([]uint16, error) {
	return []uint16{}, nil
}

func (m *Map) ToUint32() (uint32, error) {
	return uint32(0), nil
}

func (m *Map) ToUint32Array() ([]uint32, error) {
	return []uint32{}, nil
}

func (m *Map) ToUint64() (uint64, error) {
	return uint64(0), nil
}

func (m *Map) ToUint64Array() ([]uint64, error) {
	return []uint64{}, nil
}

func (m *Map) ToFloat32() (float32, error) {
	return float32(0), nil
}

func (m *Map) ToFloat32Array() ([]float32, error) {
	return []float32{}, nil
}

func (m *Map) ToFloat64() (float64, error) {
	return float64(0), nil
}

func (m *Map) ToFloat64Array() ([]float64, error) {
	return []float64{}, nil
}

func (m *Map) ToBool() (bool, error) {
	return false, nil
}

func (m *Map) ToMap() (map[string]interface{}, error) {
	return m.val, nil
}
