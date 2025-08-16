# enum

在Golang中实现枚举功能。

> 本项目使用了**go-object**风格思想实现。
>
> https://github.com/jishaocong0910/go-object

# 用法

```go
package main

import (
	"fmt"
	e "github.com/jishaocong0910/enum"
)

// 声明枚举元素
type ImageType struct {
	*e.EnumElem__
	MIME string
}

// 声明枚举集合
type _ImageType struct {
	*e.Enum__[ImageType]
	JPG,
	PNG,
	GIF ImageType
}

// 自定义元素查找
func (i _ImageType) OfMime(mine string) ImageType {
	for _, e := range i.Elems() {
		if e.MIME == mine {
			return e
		}
	}
	return i.Undefined()
}

// 创建枚举变量
var ImageType_ = e.NewEnum[ImageType](_ImageType{
	JPG: ImageType{MIME: "image/jpeg"},
	PNG: ImageType{MIME: "image/png"},
	GIF: ImageType{MIME: "image/gif"},
})

func main() {
	fmt.Println(ImageType_.JPG.ID())                           // ID为枚举集合中的变量名
	fmt.Println(ImageType_.OfId("JPG").ID())                   // 内置的查找方法
	fmt.Println(ImageType_.OfIdIgnoreCase("png").ID())         // 内置的查找方法
	fmt.Println(ImageType_.OfMime("image/jpeg").ID())          // 自定义查找方法
	fmt.Println(ImageType_.OfId("BMP").Undefined())            // 判断枚举元素是否存在
	fmt.Println(ImageType_.Is(ImageType_.JPG, ImageType_.PNG)) // 判断枚举元素是否相等

	// switch块中使用枚举元素的ID判断
	i := ImageType_.OfMime("image/webp")
	switch i.ID() {
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