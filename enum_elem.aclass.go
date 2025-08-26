package e

type enumElem_ interface {
	enumElem_() *EnumElem__
}

type EnumElem__ struct {
	fieldName string
}

func (this *EnumElem__) enumElem_() *EnumElem__ {
	return this
}

// String 返回枚举值字符串形式，与枚举集合中的字段名相同，因此具有唯一性。
func (this *EnumElem__) String() string {
	if this != nil {
		return this.fieldName
	}
	return "<undefined>"
}

// IsUndefined 是否未定义的枚举
func (this *EnumElem__) IsUndefined() bool {
	return this == nil
}
