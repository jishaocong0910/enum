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
    MIME string // Custom attribute
}

// Declare enum collection
type _ImageType struct {
    e.Enum[ImageType]
    JPG, // Custom enum elements
    PNG,
    GIF ImageType
}

// Custom method
func (i _ImageType) OfMime(mine string) ImageType {
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
    fmt.Println(ImageType_.JPG)                 // Outputs the string form, same as the field name in the enum collection
    fmt.Println(ImageType_.OfString("JPG"))     // Built-in lookup method
    fmt.Println(ImageType_.OfStringCI("png"))   // Built-in case-insensitive lookup method
    fmt.Println(ImageType_.OfMime("image/gif")) // Custom lookup method

    // Check if an enum element exists
    i := ImageType_.OfString("BMP")
    fmt.Println(i.IsPresent())

    // Check if enum elements are equal
    i2 := ImageType_.OfMime("image/webp")
    fmt.Println(ImageType_.JPG.Is(i2))

    // Undefined enum element
    var i3 ImageType // The zero value of an enum element itself is an undefined element
    fmt.Println(i3.IsUndefined()) // Opposite of IsPresent
    fmt.Println(ImageType_.UNDEFINED.Is(i3)) // The enum collection has a built-in UNDEFINED element
    fmt.Println(i3.Is(ImageType_.JPG))

    // Use String() return value for comparison in switch (direct comparison of enum elements compares all fields, which is less efficient)
    i4 := ImageType_.OfMime("image/jpeg")
    switch i4.String() {
    case ImageType_.JPG.String():
        fmt.Println("is jpg")
    case ImageType_.PNG.String():
        fmt.Println("is png")
    case ImageType_.GIF.String():
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
    MIME string // 自定义属性
}

// 声明枚举集合
type _ImageType struct {
    e.Enum[ImageType]
    JPG, // 自定义枚举元素
    PNG,
    GIF ImageType
}

// 自定义方法
func (i _ImageType) OfMime(mine string) ImageType {
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
    fmt.Println(ImageType_.JPG)                 // 将输出字符串形式，与枚举集合中的字段名相同
    fmt.Println(ImageType_.OfString("JPG"))     // 内置的查找方法
    fmt.Println(ImageType_.OfStringCI("png"))   // 内置的查找方法
    fmt.Println(ImageType_.OfMime("image/gif")) // 自定义查找方法

    // 判断枚举元素是否存在
    i := ImageType_.OfString("BMP")
    fmt.Println(i.IsPresent())

    // 判断枚举元素是否相等
    i2 := ImageType_.OfMime("image/webp")
    fmt.Println(ImageType_.JPG.Is(i2))

    // 未定义枚举元素
    var i3 ImageType // 枚举元素零值本身是一个未定义元素
    fmt.Println(i3.IsUndefined()) // 与IsPresent相反
    fmt.Println(ImageType_.UNDEFINED.Is(i3)) // 枚举集合内置了一个UNDEFINED元素
    fmt.Println(i3.Is(ImageType_.JPG))

    // switch中使用String()方法返回值进行比较（直接使用枚举元素会逐字段比较所有属性，效率较低）
    i4 := ImageType_.OfMime("image/jpeg")
    switch i4.String() {
    case ImageType_.JPG.String():
        fmt.Println("is jpg")
    case ImageType_.PNG.String():
        fmt.Println("is png")
    case ImageType_.GIF.String():
        fmt.Println("is gif")
    default:
        fmt.Println("unknown image type")
    }
}
```