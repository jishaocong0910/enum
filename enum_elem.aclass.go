package e

type enumElem_ interface {
	enumElem_()
	setFieldName(string)
	getFieldName() string
	String() string
	IsUndefined() bool
}

type EnumElem__ struct {
	fieldName string
}

func (this *EnumElem__) enumElem_() {}

func (this *EnumElem__) setFieldName(fieldName string) {
	this.fieldName = fieldName
}

func (this *EnumElem__) getFieldName() string {
	if this != nil {
		return this.fieldName
	}
	return ""
}

// String 返回枚举值字符串形式，与枚举集合中的字段名相同，因此具有唯一性。
func (this *EnumElem__) String() string {
	if this != nil {
		return this.fieldName
	}
	return "<undefined>"
}

// Is 判断是否存在指定枚举值
func (this *EnumElem__) Is(targets ...any) bool {
	if !this.IsUndefined() {
		for _, t := range targets {
			if e, ok := t.(enumElem_); ok {
				if this.fieldName == e.getFieldName() {
					return true
				}
			}
		}
	} else {
		for _, t := range targets {
			if e, ok := t.(enumElem_); ok {
				if e.IsUndefined() {
					return true
				}
			}
		}
	}
	return false
}

// Not 与Is方法相反
func (this *EnumElem__) Not(targets ...any) bool {
	return !this.Is(targets...)
}

// IsUndefined 是否未定义的枚举
func (this *EnumElem__) IsUndefined() bool {
	return this == nil
}
