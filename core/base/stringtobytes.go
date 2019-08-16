package base

import (
	"reflect"
	"unsafe"
)

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, sh.Len}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

//func StringToBytes(s string) []byte {
//	x := (*[2]uintptr)(unsafe.Pointer(&s))
//	h := [3]uintptr{x[0], x[1], x[1]}
//	return *(*[]byte)(unsafe.Pointer(&h))
//}
