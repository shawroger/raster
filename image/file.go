package image

import (
	"fmt"
	"image"

	// 引入不同图片类型解析包
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/shawroger/raster/utils"
)

// Color 像素颜色数组
// [r, b, g, a]
type Color [4]int

// ColorList 像素颜色数组列表
// 长度 * 宽度 * [r, b, g, a]
type ColorList [][]Color

// Print 控制台输出 RGBA 数据
func (c ColorList) Print() {
	for _, row := range c {
		for _, item := range row {
			fmt.Println(item)
		}
	}
}

// Content 图片信息
// Color 颜色序列
// width 图片宽度
// Height 图片高度
type Content struct {
	Color  ColorList
	Width  int
	Height int
}

// From 读取图片信息
func From(file *os.File) Content {
	// 异步关闭文件
	defer file.Close()

	// 解析图片文件
	img, _, err := image.Decode(file)
	utils.DealError(err)

	// 获取图片大小
	size := img.Bounds().Max

	// 最终的像素RGBA数组
	var c ColorList

	// 将所有颜色输入 c
	for y := 0; y < size.Y; y++ {
		var t []Color
		for x := 0; x < size.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// 获取整数形式的 0~255 的 RGBA 值
			v := [4]int{int(r) >> 8, int(g) >> 8, int(b) >> 8, int(a) >> 8}
			t = append(t, v)
		}
		c = append(c, t)
	}

	return Content{c, size.X, size.Y}
}

// FromPath 从路径直接读取图片信息
func FromPath(filePath string) Content {

	// 打开图片文件
	file, err := os.Open(filePath)

	//处理异常
	utils.DealError(err)

	return From(file)

}
