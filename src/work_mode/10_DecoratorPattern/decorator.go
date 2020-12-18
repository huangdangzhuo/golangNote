package decorator

// Shape 模型接口
type Shape interface {
	Draw()
}

// Circle 圆形类
type Circle struct {
	
}

// NewCircle 实例化圆形
func NewCircle() *Circle {
	return &Circle{}
}



