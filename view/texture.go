package view

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

// Texture !
type Texture struct {
	texture uint32
	width   int
	height  int
}

// NewTexture create a new texture
func NewTexture(file string) (*Texture, error) {
	texture := &Texture{}
	imgFile, err := os.Open(file)
	if err != nil {
		return texture, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return texture, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return texture, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var textureID uint32
	width := rgba.Rect.Size().X
	height := rgba.Rect.Size().Y
	gl.GenTextures(1, &textureID)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(width),
		int32(height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	texture.texture = textureID
	texture.width = width
	texture.height = height
	return texture, nil
}