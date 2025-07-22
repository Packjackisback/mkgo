package engine

type Triangle struct {
	BaseMesh
}

func NewTriangle() *Triangle {
	t := &Triangle{BaseMesh: NewBaseMesh()}

	vertices := []float32{
		// Position        // Color
		 0.0,  0.5, 0.0,   1.0, 0.0, 0.0, // Top
		-0.5, -0.5, 0.0,   0.0, 1.0, 0.0, // Bottom-left
		 0.5, -0.5, 0.0,   0.0, 0.0, 1.0, // Bottom-right
	}

	t.SetVertexData(vertices, nil)
	t.SetColor(1, 1, 1) 
	return t
}


