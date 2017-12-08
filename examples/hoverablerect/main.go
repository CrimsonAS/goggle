package main

import (
	"log"

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

	for r.IsRunning() {
		r.ProcessEvents()
		log.Printf("%+v", MainWindowRender(nil, nil, w))
		w.Render(MainWindowRender(nil, nil, w))
	}

	r.Quit()
}
