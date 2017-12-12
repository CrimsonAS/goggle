package main

import (
	"github.com/CrimsonAS/goggle/renderer2/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg2"
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
		// ### I do not like user code calling render functions at all. Avoid.
		w.Render(MainWindowRender(nil, &sg2.RenderState{Window: w}))
	}

	r.Quit()
}
