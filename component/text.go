package component

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/toshusai/bui"
	"github.com/toshusai/bui/view"
)

type Text struct {
	parent  *view.Object
	Text    string
	program uint32
	Color   bui.Color
	Font    *view.Font
	Size    float32
}

func NewText() *Text {
	return &Text{}
}

func (t *Text) Init() {
	t.program = t.Font.Program
}

func (t *Text) Update() {
	x := t.parent.Position.X()
	indices := []rune(t.Text)

	if len(indices) == 0 {
		return
	}

	lowChar := rune(32)

	//setup blending mode
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Activate corresponding render state
	gl.UseProgram(t.program)
	//set text color
	gl.Uniform4f(gl.GetUniformLocation(t.program, gl.Str("textColor\x00")), t.Color.R, t.Color.G, t.Color.B, t.Color.A)
	//set screen resolution
	//resUniform := gl.GetUniformLocation(f.program, gl.Str("resolution\x00"))
	//gl.Uniform2f(resUniform, float32(2560), float32(1440))

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindVertexArray(t.Font.Vao)

	// Iterate through all characters in string
	for i := range indices {

		//get rune
		runeIndex := indices[i]

		//skip runes that are not in font chacter range
		if int(runeIndex)-int(lowChar) > len(t.Font.FontChar) || runeIndex < lowChar {
			fmt.Printf("%c %d\n", runeIndex, runeIndex)
			continue
		}

		//find rune in fontChar list
		ch := t.Font.FontChar[runeIndex-lowChar]

		//calculate position and size for current rune
		xpos := x + float32(ch.BearingH)*t.Size
		ypos := t.parent.Position.Y() - float32(ch.Height-ch.BearingV)*t.Size
		w := float32(ch.Width) * t.Size
		h := float32(ch.Height) * t.Size

		//set quad positions
		var x1 = xpos
		var x2 = xpos + w
		var y1 = ypos
		var y2 = ypos + h

		//setup quad array
		var vertices = []float32{
			//  X, Y, Z, U, V
			// Front
			x1, y1, 0.0, 0.0,
			x2, y1, 1.0, 0.0,
			x1, y2, 0.0, 1.0,
			x1, y2, 0.0, 1.0,
			x2, y1, 1.0, 0.0,
			x2, y2, 1.0, 1.0}

		// Render glyph texture over quad
		gl.BindTexture(gl.TEXTURE_2D, ch.TextureID)
		// Update content of VBO memory
		gl.BindBuffer(gl.ARRAY_BUFFER, t.Font.Vbo)

		//BufferSubData(target Enum, offset int, data []byte)
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertices)*4, gl.Ptr(vertices)) // Be sure to use glBufferSubData and not glBufferData
		// Render quad
		gl.DrawArrays(gl.TRIANGLES, 0, 16)

		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
		// Now advance cursors for next glyph (note that advance is number of 1/64 pixels)
		x += float32((ch.Advance >> 6)) * t.Size // Bitshift by 6 to get value in pixels (2^6 = 64 (divide amount of 1/64th pixels by 64 to get amount of pixels))

	}
	//clear opengl textures and programs
	gl.BindVertexArray(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.UseProgram(0)
	gl.Disable(gl.BLEND)
}

func (t *Text) SetParent(obj *view.Object) {
	t.parent = obj
}
