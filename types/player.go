package types

const SPEED float32 = 1.0

type Player struct {
	position *Position
	to       *Position
}

func NewPlayer() *Player {
	return &Player{
		position: &Position{},
		to:       &Position{},
	}
}

func (p *Player) MoveTowards(to *Position) {
	p.to = to
}

func (p *Player) UpdatePosition() {
	direction := Position{}

	if p.position.x < p.to.x {
		direction.x += SPEED
	} else if p.position.x > p.to.x {
		direction.x -= SPEED
	}

	if p.position.z < p.to.z {
		direction.z += SPEED
	} else if p.position.z > p.to.z {
		direction.z -= SPEED
	}

	p.position.Move(direction)
}

func (p Player) IsMoving() bool {
	return p.position.x != p.to.x || p.position.z != p.to.z
}

func (p Player) ToDTO(id int) *PlayerDTO {
	return &PlayerDTO{
		Id: id,
		Position: PositionDTO{
			X: p.position.x,
			Z: p.position.z,
		},
	}
}