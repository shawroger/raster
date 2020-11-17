package image

import "fmt"

// Raster 点阵数据
type Raster [][]float64

// Print 控制台输出点阵数据
func (r Raster) Print() {
	for _, result := range r {
		fmt.Println(result)
	}
}

// Rasterize 返回图片点阵数组
func Rasterize(c Content, checker ColorChecker) Raster {

	// 定义点阵数组
	var r Raster
	for _, row := range c.Color {
		// 临时变量，储存一行数据
		var t []float64

		for _, color := range row {
			// 判断颜色
			var flag = checker.Weigh(color)
			t = append(t, float64(flag))
		}

		r = append(r, t)
	}

	return r
}
