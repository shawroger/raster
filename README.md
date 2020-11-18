# How to use 

## Download

``` bash
go get github.com/shawroger/raster
```

## Import 

``` go
package main

import (
	"github.com/shawroger/raster/image"
)

const url = "{your-image-file}"

func main() {
	c := image.FromPath(url)
	r := image.Rasterize(c, image.BlackChecker{})
	r.Print()
}

```