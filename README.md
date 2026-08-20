# enum

Enums implementation for Golang. To define an enum named `MyEnum`, declare the following three layers according to the naming convention.

1. Enum element: `MyEnum`. A struct embedding `e.EnumElem`, used to declare and extend custom attributes.
2. Enum collection: `_MyEnum`. A struct embedding `e.Enum`, used to declare and extend enum elements and custom methods.
3. Enum variable: `MyEnum_`. A variable created via `e.NewEnum`, used to initialize the enum collection.

👉 [🇨🇳 切换至中文说明 ↓](#chinese-version)

# Usage & Examples

```go
package main

import (
	"fmt"

	e "github.com/jishaocong0910/enum"
)

// Declare enum element
type ImageType struct {
	e.EnumElem
	MIME string // Custom field
}

// Declare enum collection
type _ImageType struct {
	e.Enum[ImageType]
	JPG, // Custom enum elements
	PNG,
	GIF ImageType
}

// Custom method
func (i _ImageType) GetByMime(mine string) ImageType {
	for _, el := range i.Elems() {
		if el.MIME == mine {
			return el
		}
	}
	return i.UNDEFINED
}

// Create enum variable
var ImageType_ = e.NewEnum(_ImageType{
	JPG: ImageType{MIME: "image/jpeg"},
	PNG: ImageType{MIME: "image/png"},
	GIF: ImageType{MIME: "image/gif"},
})

func main() {
	fmt.Println(ImageType_.JPG.Name())             // print the enum element name
	fmt.Println(ImageType_.JPG)                    // same as ImageType_.JPG.Name()
	fmt.Println(ImageType_.GetByName("JPG"))       // built-in lookup method by name
	fmt.Println(ImageType_.GetByNameCI("png"))     // built-in case-insensitive lookup method by name
	fmt.Println(ImageType_.GetByMime("image/gif")) // custom lookup method

	// check if an enum element exists
	img := ImageType_.GetByName("BMP")
	fmt.Println(img.IsPresent())
	fmt.Println(img.IsUndefined())

	// check if enum elements are equal
	img2 := ImageType_.GetByMime("image/webp")
	fmt.Println(ImageType_.JPG.Is(img2))

	// undefined enum element
	var img3 ImageType           // the zero value of an enum element is an undefined element
	img4 := ImageType_.UNDEFINED // built-in UNDEFINED element
	fmt.Println(img3.IsUndefined())
	fmt.Println(img4.Is(img3))

	// Use ID() return value for comparison in switch
	i4 := ImageType_.GetByMime("image/gif")
	switch i4.ID() {
	case ImageType_.JPG.ID():
		fmt.Println("is jpg")
	case ImageType_.PNG.ID():
		fmt.Println("is png")
	case ImageType_.GIF.ID():
		fmt.Println("is gif")
	default:
		fmt.Println("unknown image type")
	}
}
```

---

<div id="chinese-version"></div>

# enum

在Golang中实现枚举功能。 假设要定义一个名称为`MyEnum`的枚举，请按照命名规律声明以下三层概念。

1. 枚举元素：`MyEnum`。结构体，内嵌`e.EnumElem`，用于声明、扩展自定义属性。
2. 枚举集合：`_MyEnum`。结构体，内嵌`e.Enum`，用于声明、扩展枚举元素和自定义方法。
3. 枚举变量：`MyEnum_`。变量，通过函数`e.NewEnum`创建，用于初始化枚举集合。

👉 [Back to English ↑](#enum)

# 用法&例子

```go
package main

import (
	"fmt"

	e "github.com/jishaocong0910/enum"
)

// 声明枚举元素
type ImageType struct {
	e.EnumElem
	MIME string // 自定义字段
}

// 声明枚举集合
type _ImageType struct {
	e.Enum[ImageType]
	JPG, // 自定义枚举元素
	PNG,
	GIF ImageType
}

// 自定义方法
func (i _ImageType) GetByMime(mine string) ImageType {
	for _, el := range i.Elems() {
		if el.MIME == mine {
			return el
		}
	}
	return i.UNDEFINED
}

// 创建枚举变量
var ImageType_ = e.NewEnum(_ImageType{
	JPG: ImageType{MIME: "image/jpeg"},
	PNG: ImageType{MIME: "image/png"},
	GIF: ImageType{MIME: "image/gif"},
})

func main() {
	fmt.Println(ImageType_.JPG.Name())             // 获取枚举元素名称
	fmt.Println(ImageType_.JPG)                    // 与ImageType_.JPG.Name()相同
	fmt.Println(ImageType_.GetByName("JPG"))       // 内置名称查找方法
	fmt.Println(ImageType_.GetByNameCI("png"))     // 内置名称查找方法（忽略大小写）
	fmt.Println(ImageType_.GetByMime("image/gif")) // 自定义查找方法

	// 判断枚举元素是否存在
	img := ImageType_.GetByName("BMP")
	fmt.Println(img.IsPresent())
	fmt.Println(img.IsUndefined())

	// 判断枚举元素是否相等
	img2 := ImageType_.GetByMime("image/webp")
	fmt.Println(ImageType_.JPG.Is(img2))

	// 未定义枚举元素
	var img3 ImageType           // 枚举元素零值本身是一个未定义元素
	img4 := ImageType_.UNDEFINED // 枚举集合内置了一个UNDEFINED元素
	fmt.Println(img3.IsUndefined())
	fmt.Println(img4.Is(img3))

	// switch中使用ID()方法返回值进行比较
	img5 := ImageType_.GetByMime("image/gif")
	switch img5.ID() {
	case ImageType_.JPG.ID():
		fmt.Println("is jpg")
	case ImageType_.PNG.ID():
		fmt.Println("is png")
	case ImageType_.GIF.ID():
		fmt.Println("is gif")
	default:
		fmt.Println("unknown image type")
	}
}
```