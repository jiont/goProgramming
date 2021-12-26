package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var palette = []color.Color{color.White, color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 打印url的PATH
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "method:[%s], url:[%s] , proto:[%s]", r.Method, r.URL, r.Proto)
	//for k, v := range r.Header {
	//	fmt.Fprintf(w, "Headr[%q] = %q \n", k, v)
	//}
	//
	//fmt.Fprintf(w, "Host = %q \n", r.Host)
	//fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//
	//for k, v := range r.Form {
	//	fmt.Fprintf(w, "Form [%q] = %q \n ", k, v)
	//}

	if _, ok := r.Form["cycles"]; ok {
		cyclesInt, err := strconv.Atoi(r.Form.Get("cycles"))
		if err != nil {
			lissajous(w, 5)
		}
		lissajous(w, cyclesInt)
	}

	lissajous(w, 5)

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nFrames = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < math.Pi*2*float64(cycles); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
