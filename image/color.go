package image

// Checker 检查颜色
type Checker interface {
	Is(color Color) bool
	Weigh(color Color) float64
	RGBA() [4]int
}

// BlackChecker 黑色检查器
type BlackChecker struct{}

// Is 黑色检查器实现接口 Checker.Is
func (b BlackChecker) Is(color Color) bool {
	f := true
	r := b.RGBA()
	for index := range r {
		if r[index] != color[index] {
			f = false
		}
	}

	return f
}

// Weigh 黑色检查器实现接口 Checker.Weigh
func (b BlackChecker) Weigh(color Color) float64 {
	f := float64(0)
	if b.Is(color) {
		f = 1
	}

	return f
}

// RGBA 黑色检查器实现接口 Checker.RGBA
func (b BlackChecker) RGBA() [4]int {
	return [4]int{0, 0, 0, 255}
}
