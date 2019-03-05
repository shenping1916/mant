package yaml

type Invaild struct {
	err error
}

func (i *Invaild) Error() error {
	return i.err
}

func (i *Invaild) ToString() (string, error) {
	return "", nil
}

func (i *Invaild) ToStringArray() ([]string, error) {
	return []string{}, nil
}

func (i *Invaild) ToInt() (int, error) {
	return 0, nil
}

func (i *Invaild) ToIntArray() ([]int, error) {
	return []int{}, nil
}

func (i *Invaild) ToInt16() (int16, error) {
	return int16(0), nil
}

func (i *Invaild) ToInt16Array() ([]int16, error) {
	return []int16{}, nil
}

func (i *Invaild) ToInt32() (int32, error) {
	return int32(0), nil
}

func (i *Invaild) ToInt32Array() ([]int32, error) {
	return []int32{}, nil
}

func (i *Invaild) ToInt64() (int64, error) {
	return int64(0), nil
}

func (i *Invaild) ToInt64Array() ([]int64, error) {
	return []int64{}, nil
}

func (i *Invaild) ToUint() (uint, error) {
	return uint(0), nil
}

func (i *Invaild) ToUintArray() ([]uint, error) {
	return []uint{}, nil
}

func (i *Invaild) ToUint16() (uint16, error) {
	return uint16(0), nil
}

func (i *Invaild) ToUint16Array() ([]uint16, error) {
	return []uint16{}, nil
}

func (i *Invaild) ToUint32() (uint32, error) {
	return uint32(0), nil
}

func (i *Invaild) ToUint32Array() ([]uint32, error) {
	return []uint32{}, nil
}

func (i *Invaild) ToUint64() (uint64, error) {
	return uint64(0), nil
}

func (i *Invaild) ToUint64Array() ([]uint64, error) {
	return []uint64{}, nil
}

func (i *Invaild) ToFloat32() (float32, error) {
	return float32(0), nil
}

func (i *Invaild) ToFloat32Array() ([]float32, error) {
	return []float32{}, nil
}

func (i *Invaild) ToFloat64() (float64, error) {
	return float64(0), nil
}

func (i *Invaild) ToFloat64Array() ([]float64, error) {
	return []float64{}, nil
}

func (i *Invaild) ToBool() (bool, error) {
	return false, nil
}

func (i *Invaild) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
