package partitioner

import "fmt"

type Partitioner struct {
	min int
	max int
}

func New(max int) Partitioner {
	return Partitioner{
		min: 0,
		max: max,
	}
}

func (p Partitioner) Over() Partitioner {
	return Partitioner{
		min: p.min + ((p.max - p.min + 1) / 2),
		max: p.max,
	}
}

func (p Partitioner) Under() Partitioner {
	return Partitioner{
		min: p.min,
		max: p.max/2 + (p.min / 2),
	}
}

func (p Partitioner) GetResult() (int, error) {
	if p.min != p.max {
		return 0, fmt.Errorf("partitioner not fully reduced, min %d max %d", p.min, p.max)
	}
	return p.min, nil
}
