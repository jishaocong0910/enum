package e

type enumElem_ interface {
	enumElem_()
	setFieldName(string)
	getFieldName() string
	String() string
}

type EnumElem__ struct {
	fieldName string
}

func (this *EnumElem__) enumElem_() {}
func (this *EnumElem__) setFieldName(fieldName string) {
	this.fieldName = fieldName
}
func (this *EnumElem__) getFieldName() string {
	return this.fieldName
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
