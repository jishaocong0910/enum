<a id="english"></a>

# enum

[English](#english) | [中文](#chinese)

An enum implementation for Go.

To define an enum named `MyEnum`, declare the following three parts according to the naming convention:

1. Enum value type: `MyEnum`. A struct that embeds `enum.Val` and can contain custom fields.
2. Enum type: `_MyEnum`. A struct that embeds `enum.Typ`, declares the enum values, and can have custom methods.
3. Enum variable: `MyEnum_`. A variable created by initializing the enum type with `enum.NewType`, used to access the enum values.

# Quick Start

```go
package main

import (
	"fmt"

	"github.com/jishaocong0910/enum"
)

// Enum value type
type ImageType struct {
	enum.Val
	DefaultMIME string // Custom field
}

// Enum type
type _ImageType struct {
	enum.Typ[ImageType]
	JPG,
	PNG,
	GIF ImageType
}

// Custom lookup method
func (i _ImageType) OfDefaultMIME(mime string) ImageType {
	for _, el := range i.Values() {
		if el.DefaultMIME == mime {
			return el
		}
	}
	return i.UNDEFINED
}

// Create the enum variable
var ImageType_ = enum.NewType(_ImageType{
	JPG: ImageType{DefaultMIME: "image/jpeg"},
	PNG: ImageType{DefaultMIME: "image/png"},
	GIF: ImageType{DefaultMIME: "image/gif"},
})

func main() {
	// Print the enum value's name
	fmt.Println(ImageType_.JPG.Name) // Get the enum value's name
	fmt.Println(ImageType_.JPG)      // Prints the same value as Name
	// Output:
	// JPG
	// JPG

	fmt.Println()
	// Look up enum values
	fmt.Println(ImageType_.OfName("JPG"))              // Built-in lookup by name
	fmt.Println(ImageType_.OfNameIgnoreCase("png"))    // Built-in case-insensitive lookup by name
	fmt.Println(ImageType_.OfDefaultMIME("image/gif")) // Custom lookup method
	// Output:
	// JPG
	// PNG
	// GIF

	fmt.Println()
	// Compare enum values
	img := ImageType_.PNG
	img2 := ImageType_.JPG
	fmt.Println(ImageType_.JPG.Is(img))
	fmt.Println(ImageType_.JPG.Is(img2))
	// Output:
	// false
	// true

	fmt.Println()
	// Check for undefined enum values
	img3 := ImageType_.UNDEFINED     // Built-in undefined enum value
	var img4 ImageType               // The zero value of the enum value type is also undefined
	img5 := ImageType_.OfName("BMP") // Returns an undefined value if the name is not found
	fmt.Println(img5.Is(img3))
	fmt.Println(img5.Is(img4))
	fmt.Println(img5.IsUndefined())
	fmt.Println(img5.IsPresent())
	// Output:
	// true
	// true
	// true
	// false

	fmt.Println()
	// Compare enum value IDs in a switch
	img6 := ImageType_.OfDefaultMIME("image/gif")
	switch img6.ID {
	case ImageType_.JPG.ID:
		fmt.Println("jpg")
	case ImageType_.PNG.ID:
		fmt.Println("png")
	case ImageType_.GIF.ID:
		fmt.Println("gif")
	default:
		fmt.Println("unknown image format")
	}
	// Output:
	// gif
}
```

---

<a id="chinese"></a>

# enum

[English](#english) | [中文](#chinese)

在 Go 中实现枚举功能。 

假设要定义一个名为 `MyEnum` 的枚举，需要按照以下命名规则声明三个部分：

1. 枚举值类型：`MyEnum`。一个内嵌 `enum.Val` 的结构体，可在其中声明自定义字段。
2. 枚举类型：`_MyEnum`。一个内嵌 `enum.Typ` 的结构体，用于声明枚举值，也可以为其定义自定义方法。
3. 枚举变量：`MyEnum_`。通过 `enum.NewType` 初始化枚举类型并创建的变量，用于访问枚举值。

# 快速开始

```go
package main

import (
	"fmt"

	"github.com/jishaocong0910/enum"
)

// 枚举值类型
type ImageType struct {
	enum.Val
	DefaultMIME string // 自定义字段
}

// 枚举类型
type _ImageType struct {
	enum.Typ[ImageType]
	JPG,
	PNG,
	GIF ImageType
}

// 自定义查找方法
func (i _ImageType) OfDefaultMIME(mime string) ImageType {
	for _, el := range i.Values() {
		if el.DefaultMIME == mime {
			return el
		}
	}
	return i.UNDEFINED
}

// 创建枚举变量
var ImageType_ = enum.NewType(_ImageType{
	JPG: ImageType{DefaultMIME: "image/jpeg"},
	PNG: ImageType{DefaultMIME: "image/png"},
	GIF: ImageType{DefaultMIME: "image/gif"},
})

func main() {
	// 打印枚举值名称
	fmt.Println(ImageType_.JPG.Name) // 获取枚举值名称
	fmt.Println(ImageType_.JPG)      // 输出与Name相同
	// Output:
	// JPG
	// JPG

	fmt.Println()
	// 枚举值查找
	fmt.Println(ImageType_.OfName("JPG"))              // 内置名称查找方法
	fmt.Println(ImageType_.OfNameIgnoreCase("png"))    // 内置名称查找方法（忽略大小写）
	fmt.Println(ImageType_.OfDefaultMIME("image/gif")) // 自定义查找方法
	// Output:
	// JPG
	// PNG
	// GIF

	fmt.Println()
	// 判断枚举值是否相等
	img := ImageType_.PNG
	img2 := ImageType_.JPG
	fmt.Println(ImageType_.JPG.Is(img))
	fmt.Println(ImageType_.JPG.Is(img2))
	// Output:
	// false
	// true

	fmt.Println()
	// 判断未定义枚举值
	img3 := ImageType_.UNDEFINED     // 内置的未定义枚举值
	var img4 ImageType               // 枚举值类型的零值也是一个未定义枚举值
	img5 := ImageType_.OfName("BMP") // 查找不到时返回未定义值
	fmt.Println(img5.Is(img3))
	fmt.Println(img5.Is(img4))
	fmt.Println(img5.IsUndefined())
	fmt.Println(img5.IsPresent())
	// Output:
	// true
	// true
	// true
	// false

	fmt.Println()
	// switch中使用枚举值的ID属性进行比较
	img6 := ImageType_.OfDefaultMIME("image/gif")
	switch img6.ID {
	case ImageType_.JPG.ID:
		fmt.Println("jpg")
	case ImageType_.PNG.ID:
		fmt.Println("png")
	case ImageType_.GIF.ID:
		fmt.Println("gif")
	default:
		fmt.Println("unknown image format")
	}
	// Output:
	// gif
}
```
