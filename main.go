package main

import (
    "github.com/fmd/gogol"
    "fmt"
)

func main() {
    gogol.Init(gogol.WindowOpts{Title: "Gogol Example", Width: 640, Height: 480})
    defer gogol.Cleanup()

    gogol.OnKeyDown(gogol.K_Q, hideLayer("inner"))
    gogol.OnKeyUp(gogol.K_Q, showLayer("inner"))

    gogol.OnKeyDown(gogol.K_W, hideLayer("outer"))
    gogol.OnKeyUp(gogol.K_W, showLayer("outer"))

    gogol.OnMouseDown(gogol.M_LEFT, doMouseStuff)

    q1 := gogol.NewQuad(gogol.QuadOpts{Width: 25.0, Height: 25.0})
    q2 := gogol.NewQuad(gogol.QuadOpts{Width: 25.0, Height: 25.0})
    q3 := gogol.NewQuad(gogol.QuadOpts{Width: 25.0, Height: 25.0})

    q1.AddChild(q2)
    q2.AddChild(q3)

    q2.Translate(100.0, 0.0)
    q3.Translate(50.0, 0.0)

    gogol.GetLayer("base").AppendLayer("inner")
    gogol.GetLayer("inner").AppendLayer("outer")

    q2.AddToLayer("inner")
    q3.AddToLayer("outer")

    for !gogol.ShouldQuit() {
        gogol.ProcessOneFrame()
        q1.Rotate(0.1)
        q2.Rotate(0.1)
        q3.Rotate(0.1)
    }
}

func hideLayer(name string) (func(code gogol.KeyCode)) {
    return func(code gogol.KeyCode) {
        gogol.GetLayer(name).Hide()
    }
}

func showLayer(name string) (func(code gogol.KeyCode)) {
    return func(code gogol.KeyCode) {
        gogol.GetLayer(name).Show()
    }
}

func doMouseStuff(code gogol.KeyCode) {
    fmt.Println("Done mouse stuff!")
}