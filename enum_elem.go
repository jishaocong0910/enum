package e

const undefined = "<undefined>"

type iEnumElem interface {
	iEnumElem()
	getFieldName() string

	String() string
	IsUndefined() bool
}

type EnumElem struct {
	fieldName string
}

func (el EnumElem) iEnumElem() {}

func (el EnumElem) getFieldName() string {
	if el.fieldName == "" {
		return undefined
	}
	return el.fieldName
}

// String 返回枚举值字符串形式，与枚举集合中的字段名相同，因此具有唯一性。
func (el EnumElem) String() string {
	return el.getFieldName()
}

// Is 判断是否存在指定枚举值
func (el EnumElem) Is(targets ...any) bool {
	if !el.IsUndefined() {
		for _, t := range targets {
			if e, ok := t.(iEnumElem); ok {
				if el.getFieldName() == e.getFieldName() {
					return true
				}
			}
		}
	} else {
		for _, t := range targets {
			if e, ok := t.(iEnumElem); ok {
				if e.IsUndefined() {
					return true
				}
			}
		}
	}
	return false
}

// Not 与Is方法相反
func (el EnumElem) Not(targets ...any) bool {
	return !el.Is(targets...)
}

// IsUndefined 是否未定义的枚举
func (el EnumElem) IsUndefined() bool {
	return el.getFieldName() == undefined
}
