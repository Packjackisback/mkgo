
package engine

import (
	"math"
)

type Circle struct {
	BaseMesh
}

func NewCircle(opts ...MeshOption) *Circle {
	c := &Circle{BaseMesh: NewBaseMesh()}

	for _, opt := range opts {
		opt(&c.BaseMesh)
	}

	const segments = 32
	vertices := []float32{
		0.0, 0.0, 0.0, 1.0, 1.0, 1.0,
	}

	angleStep := 2.0 * math.Pi / segments
	for i := 0; i <= segments; i++ {
		angle := float64(i) * angleStep
		x := float32(math.Cos(angle) * 0.5)
		y := float32(math.Sin(angle) * 0.5)
		vertices = append(vertices, x, y, 0.0, 1.0, 1.0, 1.0)
	}

	indices := []uint32{}
	for i := 1; i <= segments; i++ {
		indices = append(indices, 0, uint32(i), uint32(i+1))
	}

	c.SetVertexData(vertices, indices)

	return c
}

