package main

import (
    "github.com/fmd/gogol"
    "fmt"
)

func main() {
    gogol.Init(gogol.WindowOpts{Title: "Gogol Example", Width: 640, Height: 480})
    defer gogol.Cleanup()

    gogol.OnKeyDown(gogol.K_RETURN, hideStuff)
    gogol.OnKeyUp(gogol.K_RETURN, showStuff)
    gogol.OnMouseDown(gogol.M_LEFT, doMouseStuff)

    p := gogol.NewQuad(gogol.QuadOpts{Width: 25.0, Height: 25.0})
    g := gogol.NewQuad(gogol.QuadOpts{Width: 25.0, Height: 25.0})

    g.Translate(100.0, 0.0)
    p.AddChild(g)

    for !gogol.ShouldQuit() {
        gogol.ProcessOneFrame()
        p.Rotate(0.1)
    }
}

func hideStuff(code gogol.KeyCode) {
    gogol.G.Renderer.RenderList.GetLayer("default").Hide()
}

func showStuff(code gogol.KeyCode) {
    gogol.G.Renderer.RenderList.GetLayer("default").Show()
}

func doMouseStuff(code gogol.KeyCode) {
    fmt.Println("Done mouse stuff!")
}