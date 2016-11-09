package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	//"os"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
	unknowIndex =3
)

//func main()  {
//	lissajous(os.Stdout)
//}

func Lissajous(out io.Writer, r *http.Request)  {
	//size := 100
	const (
		cycles = 5
		res = 0.001
		//size = 200
		nframes = 64
		delay = 8
	)
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 100
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), 3)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
