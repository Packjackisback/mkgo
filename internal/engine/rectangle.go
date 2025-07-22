package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Rectangle struct {
	BaseMesh
	vao, vbo, ebo uint32
}

func NewRectangle() *Rectangle {
	r := &Rectangle{
		BaseMesh: NewBaseMesh(),
	}

	vertices := []float32{
		// Positions        // Colors
		 0.5,  0.5, 0.0,    1.0, 0.0, 0.0, // top right
		 0.5, -0.5, 0.0,    0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0,    0.0, 0.0, 1.0, // bottom left
		-0.5,  0.5, 0.0,    1.0, 1.0, 0.0, // top left
	}

	indices := []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)
	gl.GenBuffers(1, &r.ebo)

	gl.BindVertexArray(r.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	stride := int32(6 * 4) // 6 floats per vertex (3 pos + 3 color)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.BindVertexArray(0) // Unbind VAO

	return r
}
func (r *Rectangle) Render(shader *Shader) {
	shader.Use()
	gl.BindVertexArray(r.vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)
}

func (r *Rectangle) Delete() {
	gl.DeleteVertexArrays(1, &r.vao)
	gl.DeleteBuffers(1, &r.vbo)
	gl.DeleteBuffers(1, &r.ebo)
}

func (r *Rectangle) GetTransform() *Transform {
	return &r.transform
}
