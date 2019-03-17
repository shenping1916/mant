package yaml

import (
	"fmt"
	"reflect"
)

type Decoder interface {
	Error() error
	ToString() (string, error)
	ToStringArray() ([]string, error)
	ToInt() (int, error)
	ToIntArray() ([]int, error)
	ToInt16() (int16, error)
	ToInt16Array() ([]int16, error)
	ToInt32() (int32, error)
	ToInt32Array() ([]int32, error)
	ToInt64() (int64, error)
	ToInt64Array() ([]int64, error)
	ToUint() (uint, error)
	ToUintArray() ([]uint, error)
	ToUint16() (uint16, error)
	ToUint16Array() ([]uint16, error)
	ToUint32() (uint32, error)
	ToUint32Array() ([]uint32, error)
	ToUint64() (uint64, error)
	ToUint64Array() ([]uint64, error)
	ToFloat32() (float32, error)
	ToFloat32Array() ([]float32, error)
	ToFloat64() (float64, error)
	ToFloat64Array() ([]float64, error)
	ToBool() (bool, error)
	ToMap() (map[string]interface{}, error)
}

func (y *Yaml) Reflect(data interface{}, _type int) Decoder {
	if data == nil {
		return &Nil{}
	}

	//value := reflect.ValueOf(data)
	//switch value.Kind() {
	//case reflect.String:
	//	return &String{value.Interface().(string)}
	//case reflect.Map:
	//	return &Map{value.Interface().(map[string]interface{})}
	//case reflect.Slice:
	//	return &slice{value.Interface()}
	//case reflect.Int:
	//case reflect.Int16:
	//case reflect.Int32:
	//case reflect.Int64:
	//case reflect.Uint:
	//case reflect.Uint16:
	//case reflect.Uint32:
	//case reflect.Uint64:
	//// TODO: continue
	//case reflect.Bool:
	//	fmt.Println("bool")
	//	return &Bool{value.Interface().(bool)}
	//}

	switch _type {
	case TypeString,
		TypeInt,
		TypeInt16,
		TypeInt32,
		TypeInt64,
		TypeUint,
		TypeUint16,
		TypeUint32,
		TypeUint64,
		TypeFloat32,
		TypeFloat64,
		TypeBool,
		TypeMap:
		return &decode{data.(string)}
	case TypeStringArray,
		TypeIntArray,
		TypeInt16Array,
		TypeInt32Array,
		TypeInt64Array,
		TypeUintArray,
		TypeUint16Array,
		TypeUint32Array,
		TypeUint64Array,
		TypeFloat32Array,
		TypeFloat64Array:
		return &slice{data.(List)}
	}

	return &Invaild{fmt.Errorf("unsupported type: %v", reflect.TypeOf(data))}
}
