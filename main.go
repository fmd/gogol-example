package main

import (
    "github.com/fmd/gogol"
    "fmt"
)

func main() {
    gogol.Init(gogol.WindowOpts{Title: "Gogol Example", Width: 640, Height: 480})
    defer gogol.Cleanup()

    gogol.OnKeyDown(gogol.K_RETURN, doStuff)
    gogol.OnKeyUp(gogol.K_RETURN, doMoreStuff)
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

func doStuff(code gogol.KeyCode) {
    fmt.Println("Doing stuff!")
}

func doMoreStuff(code gogol.KeyCode) {
    fmt.Println("Done stuff!")
}

func doMouseStuff(code gogol.KeyCode) {
    fmt.Println("Done mouse stuff!")
}