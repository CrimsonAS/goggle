package sg

type FactoryFunction func(index int) Node

type Repeater struct {
	New   FactoryFunction
	Model int // interface{}?
}

// interface assertions
var _ Renderable = (*Repeater)(nil)

func (this *Repeater) Render(w Windowable) Node {
	c := []Node{}
	for i := 0; i < this.Model; i++ {
		c = append(c, this.New(i))
	}
	return &ParentNode{
		Children: c,
	}
}
