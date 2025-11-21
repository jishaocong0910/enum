package e

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type enum_[E enumElem_] interface {
	enum_()
	setElems([]E)
	setFieldNameMap(map[string]E)
}

type Enum__[E enumElem_] struct {
	elems        []E
	fieldNameMap map[string]E
}

func (this *Enum__[E]) enum_() {}

func (this *Enum__[E]) setElems(elems []E) {
	this.elems = elems
}

func (this *Enum__[E]) setFieldNameMap(fieldNameMap map[string]E) {
	this.fieldNameMap = fieldNameMap
}

// Elems 返回所有枚举值
func (this *Enum__[E]) Elems() []E {
	var result []E
	if this != nil {
		result = this.elems
	}
	return result
}

// Elems 返回所有枚举值字符串形式
func (this *Enum__[E]) Strings() []string {
	var strs []string
	for _, el := range this.Elems() {
		strs = append(strs, el.getFieldName())
	}
	return strs
}

// Undefined 返回一个未定义的枚举值
func (this *Enum__[E]) Undefined() E {
	var v E
	return v
}

// OfString 查找字符串对应枚举值
func (this *Enum__[E]) OfString(str string) (e E) {
	if this != nil {
		if v, ok := this.fieldNameMap[str]; ok {
			e = v
		}
	}
	return
}

// OfStringCI 查找字符串对应枚举值，不区分大小写
func (this *Enum__[E]) OfStringCI(str string) (e E) {
	if this != nil {
		for _, v := range this.elems {
			if strings.EqualFold(v.getFieldName(), str) {
				return v
			}
		}
	}
	return
}

// Is 判断是否存在指定枚举值
func (this *Enum__[E]) Is(source E, targets ...E) bool {
	if this != nil {
		for _, t := range targets {
			if t.String() == source.String() {
				return true
			}
		}
	}
	return false
}

// Not 与Is方法相反
func (this *Enum__[E]) Not(source E, targets ...E) bool {
	return !this.Is(source, targets...)
}

func NewEnum[E enumElem_, ES enum_[E]](e ES) ES {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	if t.Kind() != reflect.Struct {
		panic("parameter \"e\" must be a struct value")
	}
	t = reflect.TypeOf(&e).Elem()
	v = reflect.ValueOf(&e).Elem()
	expectedType := reflect.TypeOf((*E)(nil)).Elem()
	v.FieldByName("Enum__").Set(reflect.ValueOf(&Enum__[E]{}))

	var elems []E
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		actualType := tf.Type
		if actualType.Kind() == reflect.Pointer {
			actualType = actualType.Elem()
		}
		if !actualType.AssignableTo(expectedType) {
			continue
		}
		if vf.Kind() != reflect.Struct {
			panic(fmt.Sprintf("%s.%s must be a struct value", t.String(), tf.Name))
		}

		var elem E
		evField := vf.FieldByName("EnumElem__")
		if !tf.IsExported() {
			reflect.NewAt(evField.Type(), unsafe.Pointer(evField.UnsafeAddr())).Elem().Set(reflect.ValueOf(&EnumElem__{}))
			elem = reflect.NewAt(vf.Type(), unsafe.Pointer(vf.UnsafeAddr())).Elem().Interface().(E)
		} else {
			evField.Set(reflect.ValueOf(&EnumElem__{}))
			elem = vf.Interface().(E)
		}

		elem.setFieldName(tf.Name)
		elems = append(elems, elem)
	}

	fieldNameMap := make(map[string]E, len(elems))
	for _, elem := range elems {
		fieldNameMap[elem.getFieldName()] = elem
	}

	e.setElems(elems)
	e.setFieldNameMap(fieldNameMap)

	return v.Interface().(ES)
}
