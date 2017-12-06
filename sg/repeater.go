package sg

type FactoryFunction func(index int) Node

type Repeater struct {
	X     float32
	Y     float32
	New   FactoryFunction
	Model int // interface{}?
}

// interface assertions
var _ Renderable = (*Repeater)(nil)
var _ Positionable = (*Repeater)(nil)

func (this *Repeater) Position() Vec2 {
	return Vec2{this.X, this.Y}
}

func (this *Repeater) SetPosition(pos Vec2) {
	this.X, this.Y = pos.X, pos.Y
}

func (this *Repeater) Render(w Windowable) Node {
	c := []Node{}
	for i := 0; i < this.Model; i++ {
		c = append(c, this.New(i))
	}
	return &ParentNode{
		Children: c,
	}
}
