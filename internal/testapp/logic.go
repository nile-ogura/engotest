package logic

import (
	"fmt"
)

type Human struct {
	Name string
}

func (h *Human) Move(dt float64) {
	fmt.Printf("%.2f", dt)
}

func NewHuman(name string) *Human {
	return &Human{name}
}
