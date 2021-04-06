package proxy

import "fmt"

// Image 模型接口
type Image interface {
	Display()
}

// RealImage 原来的image类
type RealImage struct {
	FileName string
}

// NewRealImage 实例化RealImage
func NewRealImage(filename string) *RealImage {
	return &RealImage{
		FileName: filename,
	}
}

// Display RealImage实现Image接口的Display方法
func (ri *RealImage) DisPlay() {
	fmt.Printf("Displaying %s.\n", ri.FileName)
}

// MyProxyImage 代理Image类
type MyProxyImage struct {
	RealImg  *RealImage
	FileName string
}

// NewMyProxyImage 实例化代理Image类
func NewMyProxyImage(fileName string) *MyProxyImage {
	return &MyProxyImage{
		FileName: fileName,
	}
}

// DisPlay 实现Image接口函数
func (pi *MyProxyImage) Display() {
	if pi.RealImg == nil {
		pi.RealImg = NewRealImage(pi.FileName)
	}
	pi.RealImg.DisPlay()
}
