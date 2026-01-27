package e

import (
	"reflect"
	"strings"
	"unsafe"
)

type iEnum[EL iEnumElem] interface {
	iEnum(EL)
}

type Enum[EL iEnumElem] struct {
	elems        []EL
	fieldNameMap map[string]EL
}

func (e Enum[EL]) iEnum(EL) {}

// Elems 返回所有枚举值
func (e Enum[EL]) Elems() []EL {
	return e.elems
}

// Strings 返回所有枚举值字符串形式
func (e Enum[EL]) Strings() []string {
	var arr []string
	for _, el := range e.Elems() {
		arr = append(arr, el.getFieldName())
	}
	return arr
}

// Undefined 返回一个未定义的枚举值
func (e Enum[EL]) Undefined() EL {
	var v EL
	return v
}

// OfString 查找字符串对应枚举值
func (e Enum[EL]) OfString(str string) (el EL) {
	if v, ok := e.fieldNameMap[str]; ok {
		return v
	}
	return
}

// OfStringCI 查找字符串对应枚举值，不区分大小写
func (e Enum[EL]) OfStringCI(str string) (el EL) {
	for _, v := range e.elems {
		if strings.EqualFold(v.getFieldName(), str) {
			return v
		}
	}
	return
}

func NewEnum[E iEnum[EL], EL iEnumElem](e E) E {
	t := reflect.TypeOf(&e).Elem()
	v := reflect.ValueOf(&e).Elem()
	elType := reflect.TypeOf((*EL)(nil)).Elem()
	els := make([]EL, 0, v.NumField()-1)

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		if !tf.Type.AssignableTo(elType) {
			continue
		}

		if tf.IsExported() {
			vf.FieldByName("EnumElem").Set(reflect.ValueOf(EnumElem{fieldName: tf.Name}))
			els = append(els, vf.Interface().(EL))
		} else {
			*(*EnumElem)(unsafe.Pointer(vf.FieldByName("EnumElem").UnsafeAddr())) = EnumElem{fieldName: tf.Name}
			els = append(els, *(*EL)(unsafe.Pointer(vf.UnsafeAddr())))
		}
	}

	fieldNameMap := make(map[string]EL, len(els))
	for _, elem := range els {
		fieldNameMap[elem.getFieldName()] = elem
	}

	*(*Enum[EL])(unsafe.Pointer(v.FieldByName("Enum").UnsafeAddr())) = Enum[EL]{
		elems:        els,
		fieldNameMap: fieldNameMap,
	}

	return v.Interface().(E)
}
