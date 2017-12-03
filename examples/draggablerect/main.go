package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
)

func main() {
	r, err := sdlsoftware.NewRenderer()
	if err != nil {
		panic(err)
	}

	w, err := r.CreateWindow()
	defer w.Destroy()

	if err != nil {
		panic(err)
	}

	thing := &MainWindow{}

	for r.IsRunning() {
		r.ProcessEvents()
		w.Render(thing)
	}

	r.Quit()
}
