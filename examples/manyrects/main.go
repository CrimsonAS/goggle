package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/CrimsonAS/goggle/renderer2/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg2"
)

const profile = true

func main() {
	if profile {
		go http.ListenAndServe(":6060", nil)
	}
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
		w.Render(ManyRectRender(nil, &sg2.RenderState{Window: w}))
	}

	r.Quit()
}
