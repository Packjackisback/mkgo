package engine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ModelMesh struct {
	BaseMesh
}

/*
*   WARNING THIS IS A SUPER HACKY WAY OF READING OBJ FILES, 
*		I DON'T KNOW HOW IT WORKS AND AM TO AFRAID TO MAKE IT
*		BETTER. THANK YOU TO https://github.com/sheenobu/go-obj
*/



// load the mesh
func NewModelFromOBJ(filepath string, opts ...MeshOption) (*ModelMesh, error) {
	vertices, indices, err := loadOBJ(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load OBJ file: %v", err)
	}

	model := &ModelMesh{BaseMesh: NewBaseMesh()}
	model.SetVertexData(vertices, indices)

	for _, opt := range opts {
		opt(&model.BaseMesh)
	}

	return model, nil
}

// parse obj into vertices and indices
func loadOBJ(filepath string) ([]float32, []uint32, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var positions [][]float32
	var normals [][]float32
	var vertices []float32
	var indices []uint32
	var indexCounter uint32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "v": //vertice
			if len(parts) >= 4 {
				x, _ := strconv.ParseFloat(parts[1], 32)
				y, _ := strconv.ParseFloat(parts[2], 32)
				z, _ := strconv.ParseFloat(parts[3], 32)
				positions = append(positions, []float32{float32(x), float32(y), float32(z)})
			}

		case "vn": //vertice normal
			if len(parts) >= 4 {
				x, _ := strconv.ParseFloat(parts[1], 32)
				y, _ := strconv.ParseFloat(parts[2], 32)
				z, _ := strconv.ParseFloat(parts[3], 32)
				normals = append(normals, []float32{float32(x), float32(y), float32(z)})
			}

		case "f": // face
			if len(parts) >= 4 {
				faceVertices := parts[1:]

				if len(faceVertices) == 4 {
					for _, i := range []int{0, 1, 2} {
						vertex, _ := parseFaceVertex(faceVertices[i], positions, normals)
						vertices = append(vertices, vertex...)
						indices = append(indices, indexCounter)
						indexCounter++
					}
					for _, i := range []int{0, 2, 3} {
						vertex, _ := parseFaceVertex(faceVertices[i], positions, normals)
						vertices = append(vertices, vertex...)
						indices = append(indices, indexCounter)
						indexCounter++
					}
				} else if len(faceVertices) == 3 {
					for _, i := range []int{0, 1, 2} {
						vertex, _ := parseFaceVertex(faceVertices[i], positions, normals)
						vertices = append(vertices, vertex...)
						indices = append(indices, indexCounter)
						indexCounter++
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return vertices, indices, nil
}

// parse face vertex strings
func parseFaceVertex(vertexStr string, positions, normals [][]float32) ([]float32, []float32) {
	parts := strings.Split(vertexStr, "/")

	var vertex []float32
	var normal []float32

	if len(parts) >= 1 && parts[0] != "" {
		if posIdx, err := strconv.Atoi(parts[0]); err == nil {
			if posIdx > 0 && posIdx <= len(positions) {
				vertex = append(vertex, positions[posIdx-1]...)
			}
		}
	}

	if len(parts) >= 3 && parts[2] != "" {
		if normIdx, err := strconv.Atoi(parts[2]); err == nil {
			if normIdx > 0 && normIdx <= len(normals) {
				normal = normals[normIdx-1]
				vertex = append(vertex, normal...)
			}
		}
	}

	if len(vertex) == 3 {
		vertex = append(vertex, 0.0, 1.0, 0.0)
	}

	return vertex, normal
}

func LoadModel(filepath string, opts ...MeshOption) (*ModelMesh, error) {
	if strings.HasSuffix(strings.ToLower(filepath), ".obj") {
		return NewModelFromOBJ(filepath, opts...)
	}

	return nil, fmt.Errorf("unsupported file format: %s", filepath)
}
