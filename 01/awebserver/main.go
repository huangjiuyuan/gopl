package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	var palette = []color.Color{color.White, color.Black}
	const (
		// Angular resolution
		res = 0.001
		// Image canvas covers [size..+size]
		size = 100
		// Number of animation frames
		nframes = 64
		// Delay between frames in 10ms units
		delay = 8
		// First color in palette
		whiteIndex = 0
		// Next color in palette
		blackIndex = 1
	)
	parsedCycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil {
		parsedCycles = 5
	}
	cycles := float64(parsedCycles)
	// Relative frequency of y oscillator
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	// Phase difference
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
