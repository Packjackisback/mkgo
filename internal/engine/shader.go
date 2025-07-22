package engine

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const vertexShaderSource = `
#version 410 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 color;

uniform mat4 model;

out vec3 vertexColor;

void main() {
    gl_Position = model * vec4(position, 1.0);
    vertexColor = color;
}` + "\x00"

const fragmentShaderSource = `
#version 410 core

uniform vec3 uColor;

out vec4 fragColor;

void main() {
    fragColor = vec4(uColor, 1.0);
}
` + "\x00"

type Shader struct {
	program uint32
}

func NewShader() (*Shader, error) {
	fmt.Println("Creating shader program...")
	
	// Compile vertex shader
	fmt.Println("Compiling vertex shader...")
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	defer gl.DeleteShader(vertexShader)
	fmt.Println("Vertex shader compiled successfully")

	fmt.Println("Compiling fragment shader...")
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}
	defer gl.DeleteShader(fragmentShader)
	fmt.Println("Fragment shader compiled successfully")

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}
	
	fmt.Printf("Shader program linked successfully (ID: %d)\n", program)
	return &Shader{program: program}, nil
}

func (s *Shader) Use() {
	gl.UseProgram(s.program)
}

func (s *Shader) Delete() {
	gl.DeleteProgram(s.program)
}

func (s *Shader) SetMat4(name string, matrix mgl32.Mat4) {
	location := gl.GetUniformLocation(s.program, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}

func (s *Shader) SetVec3(name string, vec mgl32.Vec3) {
    location := gl.GetUniformLocation(s.program, gl.Str(name+"\x00"))
    gl.Uniform3f(location, vec.X(), vec.Y(), vec.Z())
}


func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile shader: %v", log)
	}

	return shader, nil
}


