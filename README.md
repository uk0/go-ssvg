## go-svg 


### 实现了一个svg的基本实现，可以通过参数创建图像以及文字


```go

package ssvg

import (
	"testing"
)

func TestWriteSvg(t *testing.T) {
	svg := new(Svg)
	svg.Add(&Line{X1: 0, Y1: 0, X2: 100, Y2: 200})
	svg.Add(&Line{X1: -100, Y1: -100, X2: 100, Y2: 200})
	svg.Add(&Text{X: 50, Y: -50, Text: "Test"})
	svg.WriteFile("test.svg", 220)
}


```

### demo

![test.svg](test.svg)
