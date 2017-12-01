package sg

type Scaleable interface {
	GetScale() float32
	SetScale(s float32)
}

type ScaleNode struct {
	Scale    float32
	Children []Node
}

func (this *ScaleNode) GetScale() float32 {
	return this.Scale
}

func (this *ScaleNode) SetScale(r float32) {
	this.Scale = r
}

func (this *ScaleNode) GetChildren() []Node {
	return this.Children
}
