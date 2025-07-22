package engine

type Rectangle struct {
	BaseMesh
}

func NewRectangle() *Rectangle {
	r := &Rectangle{
		BaseMesh: NewBaseMesh(),
	}

	vertices := []float32{
		// Positions       // Colors
		 0.5,  0.5, 0.0,   1.0, 0.0, 0.0, // top right
		 0.5, -0.5, 0.0,   0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0,   0.0, 0.0, 1.0, // bottom left
		-0.5,  0.5, 0.0,   1.0, 1.0, 0.0, // top left
	}

	indices := []uint32{
		0, 1, 3, // first 
		1, 2, 3, // second
	}

	r.SetColor(1.0, 1.0, 1.0) // white
	r.SetVertexData(vertices, indices)
	return r
}

