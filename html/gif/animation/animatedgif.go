package animation

import (
	"image"
	"image/color"
	"image/gif"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"io"
	"io/ioutil"
	"log"
)
// Generate an animated gif
func GenerateAnimation(text string, fontfile string, out io.Writer) {
	const (
		nframes = 64 // number of animation frames
		delay =8 // delay between frames in 10ms units
	)
	var xsize int = 30 + 30 * len(text)
	var ysize int = 200
	
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	var palette = make([]color.Color,0)
	// generate palette
	for i := 0; i < nframes; i++ {
		palette = append(palette, color.RGBA{R:0,G:0,B:uint8(4*i), A:255})
	}
	anim := gif.GIF{LoopCount: nframes}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, xsize, ysize)
		img := image.NewPaletted(rect, palette)
		for x := 0; x < xsize; x++ {
			for y := 0; y < ysize; y++ {
				img.SetColorIndex(x, y, uint8(i))
			}
		}
		WriteText(text, f, img)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
// Write a text in the image
func WriteText(text string, f *truetype.Font, img *image.Paletted){
	// Initialize the context.
	fg:= image.Black
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(64)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fg)
	// Draw the text.
	pt := freetype.Pt(40, 120)
	c.DrawString(text, pt)
}