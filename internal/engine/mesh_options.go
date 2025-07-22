package engine

type MeshOption func(*BaseMesh)

func WithColor(x, y, z float32) MeshOption {
	return func(m *BaseMesh) {
		m.SetColor(x, y, z)
	}
}

func WithPosition(x, y, z float32) MeshOption {
	return func(m *BaseMesh) {
		m.SetPosition(x, y, z)
	}
}

func WithScale(x, y, z float32) MeshOption {
	return func(m *BaseMesh) {
		m.SetScale(x, y, z)
	}
}

func WithRotation(angle float32, x, y, z float32) MeshOption {
	return func(m *BaseMesh) {
		m.SetRotation(angle, x, y, z)
	}
}

