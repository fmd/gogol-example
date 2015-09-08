package main

import (
    "github.com/fmd/gogol"
    "fmt"
)

func main() {
    gogol.Init(gogol.WindowOpts{Title: "Bottom Dollar", Width: 640, Height: 480})
    defer gogol.Cleanup()

    gogol.OnKeyDown(gogol.K_RETURN, doStuff)
    gogol.OnKeyUp(gogol.K_RETURN, doMoreStuff)
    gogol.OnMouseDown(gogol.M_LEFT, doMouseStuff)

    gogol.CreateQuad(gogol.QuadOpts{Width: 25.0, Height: 25.0})

    for !gogol.ShouldQuit() {
        gogol.ProcessOneFrame()
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