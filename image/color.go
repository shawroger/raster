package image

// ColorChecker 检查颜色
type ColorChecker interface {
	Is(color Color) bool
	Weigh(color Color) float64
	RGBA() [4]int
}

// blackChecker 黑色检查器
type blackChecker struct{ ColorChecker, rgba [4]int }

// Is 黑色检查器实现接口 ColorChecker.Is
func (p blackChecker) Is(color Color) bool {
	flag := true
	for index := range p.rgba {
		if p.rgba[index] != color[index] {
			flag = false
		}
	}

	return flag
}

// Weigh 黑色检查器实现接口 ColorChecker.Weigh
func (p blackChecker) Weigh(color Color) float64 {
	flag := float64(0)
	if p.Is(color) {
		flag = 1
	}

	return flag
}

// RGBA 黑色检查器实现接口 ColorChecker.RGBA
func (p blackChecker) RGBA() [4]int {
	return p.rgba
}

// BlackChecker 黑色检查器
var BlackChecker ColorChecker = blackChecker{rgba: [4]int{0, 0, 0, 255}}
