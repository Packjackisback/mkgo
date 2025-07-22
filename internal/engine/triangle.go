package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Triangle struct {
	BaseMesh
	vao, vbo uint32 
}

func NewTriangle() *Triangle {
	r := &Triangle{
		BaseMesh: NewBaseMesh(),
	}

	vertices := []float32{
		// Position        // Color
		 0.0,  0.5, 0.0,   1.0, 0.0, 0.0, // Top
		-0.5, -0.5, 0.0,   0.0, 1.0, 0.0, // Bottom-left
		 0.5, -0.5, 0.0,   0.0, 0.0, 1.0, // Bottom-right
	}

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))   // position
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4)) // color
	gl.EnableVertexAttribArray(1)

	gl.BindVertexArray(0)

	return r
}

func (t *Triangle) Render(shader *Shader) {
	shader.Use()
	gl.BindVertexArray(t.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)
}

func (t *Triangle) Delete() {
	gl.DeleteVertexArrays(1, &t.vao)
	gl.DeleteBuffers(1, &t.vbo)
}

