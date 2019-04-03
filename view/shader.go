package view

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var vertexShader = `
#version 330

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * view * model * vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

var simpleVS = `
#version 330

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

in vec3 vert;

void main() {
    gl_Position = projection * view * model * vec4(vert, 1);
}
` + "\x00"

var simpleFS = `
#version 330
out vec4 outputColor;

void main() {
    outputColor = vec4(1.0, 0.0, 0.0, 1.0);
}
` + "\x00"

var spriteShader *Shader
var simpleShader *Shader

func GetSpriteShader() *Shader {
	return spriteShader
}

func GetSimpleShader() *Shader {
	return simpleShader
}

type Shader struct {
	program  uint32
	Uniforms map[string]int32
}

func InitShader() {
	spriteShader = &Shader{
		program:  newProgram(vertexShader, fragmentShader),
		Uniforms: map[string]int32{},
	}
	spriteShader.Uniforms["projection"] = gl.GetUniformLocation(spriteShader.program, gl.Str("projection\x00"))
	spriteShader.Uniforms["view"] = gl.GetUniformLocation(spriteShader.program, gl.Str("view\x00"))
	spriteShader.Uniforms["model"] = gl.GetUniformLocation(spriteShader.program, gl.Str("model\x00"))
	spriteShader.Uniforms["vert"] = gl.GetAttribLocation(spriteShader.program, gl.Str("vert\x00"))
	spriteShader.Uniforms["vertTexCoord"] = gl.GetAttribLocation(spriteShader.program, gl.Str("vertTexCoord\x00"))

	simpleShader = &Shader{
		program:  newProgram(simpleVS, simpleFS),
		Uniforms: map[string]int32{},
	}
	simpleShader.Uniforms["projection"] = gl.GetUniformLocation(simpleShader.program, gl.Str("projection\x00"))
	simpleShader.Uniforms["view"] = gl.GetUniformLocation(simpleShader.program, gl.Str("view\x00"))
	simpleShader.Uniforms["model"] = gl.GetUniformLocation(simpleShader.program, gl.Str("model\x00"))
	simpleShader.Uniforms["vert"] = gl.GetAttribLocation(simpleShader.program, gl.Str("vert\x00"))
}

func (s *Shader) GetUnitform(str string) int32 {
	i := gl.GetAttribLocation(s.program, gl.Str(str+"\x00"))
	fmt.Println(i)
	return i
}

func NewShader(vertexShaderSource, fragmentShaderSource string) *Shader {
	return &Shader{
		program: newProgram(vertexShaderSource, fragmentShaderSource),
	}
}

func (s *Shader) GetProgram() uint32 {
	return s.program
}

func newProgram(vertexShaderSource, fragmentShaderSource string) uint32 {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

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

		panic(err)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
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

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
