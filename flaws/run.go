package flaws

import (
	"math/rand"
)

type RunOption interface {
	CanRun() bool
}

type RunAlways struct{}

func (RunAlways) CanRun() bool {
	return true
}

type RunPerc struct {
	Percentage int
}

func (r RunPerc) CanRun() bool {
	return rand.Intn(100) < r.Percentage
}
