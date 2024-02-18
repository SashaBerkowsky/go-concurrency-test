package types

type Position struct {
	x float32
	z float32
}

func NewPosition(x, z float32) *Position {
	return &Position{
		x: x,
		z: z,
	}
}

func (p *Position) Move(amount Position) {
	p.x += amount.x
	p.z += amount.z
}
