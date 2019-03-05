package yaml

type Bool struct {
	val bool
}

func (b *Bool) Error() error {
	return nil
}

func (b *Bool) ToString() (string, error) {
	return "", nil
}

func (b *Bool) ToStringArray() ([]string, error) {
	return []string{}, nil
}

func (b *Bool) ToInt() (int, error) {
	return 0, nil
}

func (b *Bool) ToIntArray() ([]int, error) {
	return []int{}, nil
}

func (b *Bool) ToInt16() (int16, error) {
	return int16(0), nil
}

func (b *Bool) ToInt16Array() ([]int16, error) {
	return []int16{}, nil
}

func (b *Bool) ToInt32() (int32, error) {
	return int32(0), nil
}

func (b *Bool) ToInt32Array() ([]int32, error) {
	return []int32{}, nil
}

func (b *Bool) ToInt64() (int64, error) {
	return int64(0), nil
}

func (b *Bool) ToInt64Array() ([]int64, error) {
	return []int64{}, nil
}

func (b *Bool) ToUint() (uint, error) {
	return uint(0), nil
}

func (b *Bool) ToUintArray() ([]uint, error) {
	return []uint{}, nil
}

func (b *Bool) ToUint16() (uint16, error) {
	return uint16(0), nil
}

func (b *Bool) ToUint16Array() ([]uint16, error) {
	return []uint16{}, nil
}

func (b *Bool) ToUint32() (uint32, error) {
	return uint32(0), nil
}

func (b *Bool) ToUint32Array() ([]uint32, error) {
	return []uint32{}, nil
}

func (b *Bool) ToUint64() (uint64, error) {
	return uint64(0), nil
}

func (b *Bool) ToUint64Array() ([]uint64, error) {
	return []uint64{}, nil
}

func (b *Bool) ToFloat32() (float32, error) {
	return float32(0), nil
}

func (b *Bool) ToFloat32Array() ([]float32, error) {
	return []float32{}, nil
}

func (b *Bool) ToFloat64() (float64, error) {
	return float64(0), nil
}

func (b *Bool) ToFloat64Array() ([]float64, error) {
	return []float64{}, nil
}

func (b *Bool) ToBool() (bool, error) {
	return b.val, nil
}

func (b *Bool) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
