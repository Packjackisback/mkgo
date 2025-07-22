package engine


import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh interface {
	Render(shader *Shader)	         //   These will all be satisfied by including BaseMesh
	Delete()                         //
	GetTransform() *Transform        //
}

type BaseMesh struct {
	vao        uint32
	vbo        uint32
	ebo        uint32
	hasEBO     bool
	vertexCount int32
	indexCount  int32
	color       mgl32.Vec3
	transform   Transform
}

func NewBaseMesh() BaseMesh {
	return BaseMesh{
		transform: NewTransform(),
	}
}

func (m *BaseMesh) GetTransform() *Transform {
	return &m.transform
}

func (m *BaseMesh) SetPosition(x, y, z float32) {
	m.transform.Position = mgl32.Vec3{x, y, z}
}

func (m *BaseMesh) SetScale(x, y, z float32) {
	m.transform.Scale = mgl32.Vec3{x, y, z}
}

func (m *BaseMesh) SetRotation(angle float32, x, y, z float32) {
	m.transform.RotationAngle = angle
	m.transform.RotationAxis =  mgl32.Vec3{x, y, z}.Normalize()
}

func (m *BaseMesh) SetColor(x, y, z float32) {
    m.color = mgl32.Vec3{x, y, z}
}

func (m *BaseMesh) GetColor() (x, y, z float32) {
    return m.color.X(), m.color.Y(), m.color.Z()
}

func (m *BaseMesh) Render(shader *Shader) {
	shader.Use()

	shader.SetVec3("uColor", m.color)
	shader.SetMat4("uModel", m.transform.GetModelMatrix())

	gl.BindVertexArray(m.vao)
	if m.hasEBO {
		gl.DrawElements(gl.TRIANGLES, m.indexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, m.vertexCount)
	}
	gl.BindVertexArray(0)
}

func (m *BaseMesh) SetVertexData(vertices []float32, indices []uint32) {
	m.hasEBO = len(indices) > 0
	m.vertexCount = int32(len(vertices) / 6) 
	m.indexCount = int32(len(indices))

	gl.GenVertexArrays(1, &m.vao)
	gl.GenBuffers(1, &m.vbo)

	gl.BindVertexArray(m.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	stride := int32(6 * 4)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	if m.hasEBO {
		gl.GenBuffers(1, &m.ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	}

	gl.BindVertexArray(0)
}

func (m *BaseMesh) Delete() {
	if m.vao != 0 {
		gl.DeleteVertexArrays(1, &m.vao)
		m.vao = 0
	}
	if m.vbo != 0 {
		gl.DeleteBuffers(1, &m.vbo)
		m.vbo = 0
	}
	if m.ebo != 0 {
		gl.DeleteBuffers(1, &m.ebo)
		m.ebo = 0
	}
}

