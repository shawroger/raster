# raster

`raster `是一个将图片转化为点阵的库

有两个子 package，分别是 `raster/image` 与 `raster/util`


# image
`import "github.com/shawroger/raster/image"`

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



## <a name="BlackChecker">type</a> BlackChecker
``` go
type BlackChecker struct{}

```
BlackChecker 黑色检查器










### <a name="BlackChecker.Is">func</a> (BlackChecker) Is 
``` go
func (b BlackChecker) Is(color Color) bool
```
Is 黑色检查器实现接口 Checker.Is




### <a name="BlackChecker.RGBA">func</a> (BlackChecker) RGBA 
``` go
func (b BlackChecker) RGBA() [4]int
```
RGBA 黑色检查器实现接口 Checker.RGBA




### <a name="BlackChecker.Weigh">func</a> (BlackChecker) Weigh 
``` go
func (b BlackChecker) Weigh(color Color) float64
```
Weigh 黑色检查器实现接口 Checker.Weigh




## <a name="Checker">type</a> Checker
``` go
type Checker interface {
    Is(color Color) bool
    Weigh(color Color) float64
    RGBA() [4]int
}
```
Checker 检查颜色










## <a name="Color">type</a> Color
``` go
type Color [4]int
```
Color 像素颜色数组
[r, b, g, a]










## <a name="ColorList">type</a> ColorList
``` go
type ColorList [][]Color
```
ColorList 像素颜色数组列表
长度 * 宽度 * [r, b, g, a]










### <a name="ColorList.Print">func</a> (ColorList) Print 
``` go
func (c ColorList) Print()
```
Print 控制台输出 RGBA 数据




## <a name="Content">type</a> Content
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







### <a name="From">func</a> From
``` go
func From(file *os.File) Content
```
From 读取图片信息


### <a name="FromPath">func</a> FromPath
``` go
func FromPath(filePath string) Content
```
FromPath 从路径直接读取图片信息





## <a name="Raster">type</a> Raster
``` go
type Raster [][]float64
```
Raster 点阵数据







### <a name="Rasterize">func</a> Rasterize
``` go
func Rasterize(c Content, checker Checker) Raster
```
Rasterize 返回图片点阵数组





### <a name="Raster.Print">func</a> (Raster) Print
``` go
func (r Raster) Print()
```
Print 控制台输出点阵数据


 
# utils
`import "github.com/shawroger/raster/utils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func DealError(err error)](#DealError)


## <a name="DealError">func</a> DealError
``` go
func DealError(err error)
```
DealError 全局处理错误

 