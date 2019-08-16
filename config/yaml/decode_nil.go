package yaml

type Nil struct {
}

func (n *Nil) Error() error {
	return nil
}

func (n *Nil) ToString() (string, error) {
	return "", nil
}

func (n *Nil) ToStringArray() ([]string, error) {
	return []string{}, nil
}

func (n *Nil) ToInt() (int, error) {
	return 0, nil
}

func (n *Nil) ToIntArray() ([]int, error) {
	return []int{}, nil
}

func (n *Nil) ToInt16() (int16, error) {
	return int16(0), nil
}

func (n *Nil) ToInt16Array() ([]int16, error) {
	return []int16{}, nil
}

func (n *Nil) ToInt32() (int32, error) {
	return int32(0), nil
}

func (n *Nil) ToInt32Array() ([]int32, error) {
	return []int32{}, nil
}

func (n *Nil) ToInt64() (int64, error) {
	return int64(0), nil
}

func (n *Nil) ToInt64Array() ([]int64, error) {
	return []int64{}, nil
}

func (n *Nil) ToUint() (uint, error) {
	return uint(0), nil
}

func (n *Nil) ToUintArray() ([]uint, error) {
	return []uint{}, nil
}

func (n *Nil) ToUint16() (uint16, error) {
	return uint16(0), nil
}

func (n *Nil) ToUint16Array() ([]uint16, error) {
	return []uint16{}, nil
}

func (n *Nil) ToUint32() (uint32, error) {
	return uint32(0), nil
}

func (n *Nil) ToUint32Array() ([]uint32, error) {
	return []uint32{}, nil
}

func (n *Nil) ToUint64() (uint64, error) {
	return uint64(0), nil
}

func (n *Nil) ToUint64Array() ([]uint64, error) {
	return []uint64{}, nil
}

func (n *Nil) ToFloat32() (float32, error) {
	return float32(0), nil
}

func (n *Nil) ToFloat32Array() ([]float32, error) {
	return []float32{}, nil
}

func (n *Nil) ToFloat64() (float64, error) {
	return float64(0), nil
}

func (n *Nil) ToFloat64Array() ([]float64, error) {
	return []float64{}, nil
}

func (n *Nil) ToBool() (bool, error) {
	return false, nil
}

func (n *Nil) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
