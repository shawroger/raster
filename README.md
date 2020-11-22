# image
`import "github.com\shawroger\raster\image"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type BlackChecker](#BlackChecker)
  * [func (b BlackChecker) Is(color Color) bool](#BlackChecker.Is)
  * [func (b BlackChecker) RGBA() [4]int](#BlackChecker.RGBA)
  * [func (b BlackChecker) Weigh(color Color) float64](#BlackChecker.Weigh)
* [type Checker](#Checker)
* [type Color](#Color)
* [type ColorList](#ColorList)
  * [func (c ColorList) Print()](#ColorList.Print)
* [type Content](#Content)
  * [func From(file *os.File) Content](#From)
  * [func FromPath(filePath string) Content](#FromPath)
* [type Raster](#Raster)
  * [func Rasterize(c Content, checker Checker) Raster](#Rasterize)
  * [func (r Raster) Print()](#Raster.Print)


#### <a name="pkg-files">Package files</a>
[color.go](/src/target/color.go) [file.go](/src/target/file.go) [raster.go](/src/target/raster.go)






## <a name="BlackChecker">type</a> [BlackChecker](/src/target/color.go?s=174:200#L11)
``` go
type BlackChecker struct{}

```
BlackChecker 黑色检查器










### <a name="BlackChecker.Is">func</a> (BlackChecker) [Is](/src/target/color.go?s=250:292#L14)
``` go
func (b BlackChecker) Is(color Color) bool
```
Is 黑色检查器实现接口 Checker.Is




### <a name="BlackChecker.RGBA">func</a> (BlackChecker) [RGBA](/src/target/color.go?s=645:680#L37)
``` go
func (b BlackChecker) RGBA() [4]int
```
RGBA 黑色检查器实现接口 Checker.RGBA




### <a name="BlackChecker.Weigh">func</a> (BlackChecker) [Weigh](/src/target/color.go?s=475:523#L27)
``` go
func (b BlackChecker) Weigh(color Color) float64
```
Weigh 黑色检查器实现接口 Checker.Weigh




## <a name="Checker">type</a> [Checker](/src/target/color.go?s=42:137#L4)
``` go
type Checker interface {
    Is(color Color) bool
    Weigh(color Color) float64
    RGBA() [4]int
}
```
Checker 检查颜色










## <a name="Color">type</a> [Color](/src/target/file.go?s=233:250#L18)
``` go
type Color [4]int
```
Color 像素颜色数组
[r, b, g, a]










## <a name="ColorList">type</a> [ColorList](/src/target/file.go?s=328:352#L22)
``` go
type ColorList [][]Color
```
ColorList 像素颜色数组列表
长度 * 宽度 * [r, b, g, a]










### <a name="ColorList.Print">func</a> (ColorList) [Print](/src/target/file.go?s=394:420#L25)
``` go
func (c ColorList) Print()
```
Print 控制台输出 RGBA 数据




## <a name="Content">type</a> [Content](/src/target/file.go?s=611:680#L37)
``` go
type Content struct {
    Color  ColorList
    Width  int
    Height int
}

```
Content 图片信息
Color 颜色序列
width 图片宽度
Height 图片高度







### <a name="From">func</a> [From](/src/target/file.go?s=712:744#L44)
``` go
func From(file *os.File) Content
```
From 读取图片信息


### <a name="FromPath">func</a> [FromPath](/src/target/file.go?s=1389:1427#L74)
``` go
func FromPath(filePath string) Content
```
FromPath 从路径直接读取图片信息





## <a name="Raster">type</a> [Raster](/src/target/raster.go?s=64:87#L8)
``` go
type Raster [][]float64
```
Raster 点阵数据







### <a name="Rasterize">func</a> [Rasterize](/src/target/raster.go?s=256:305#L18)
``` go
func Rasterize(c Content, checker Checker) Raster
```
Rasterize 返回图片点阵数组



### <a name="Raster.Print">func</a> (Raster) [Print](/src/target/raster.go?s=129:152#L11)
``` go
func (r Raster) Print()
```
Print 控制台输出点阵数据


# utils
`import "E:\Software\gopath\src\github.com\shawroger\raster\utils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func DealError(err error)](#DealError)


#### <a name="pkg-files">Package files</a>
[common.go](/src/target/common.go)





## <a name="DealError">func</a> [DealError](/src/target/common.go?s=66:91#L6)
``` go
func DealError(err error)
```
DealError 全局处理错误
