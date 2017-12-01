package sg

type Rotateable interface {
	GetRotation() float32
	SetRotation(r float32)
}

type RotationNode struct {
	Rotation float32
	Children []Node
}

func (this *RotationNode) GetRotation() float32 {
	return this.Rotation
}

func (this *RotationNode) SetRotation(r float32) {
	this.Rotation = r
}

func (this *RotationNode) GetChildren() []Node {
	return this.Children
}
