package main

import (
    "github.com/fmd/bronson"
    "fmt"
)

func main() {
    bronson.Init(bronson.WindowOpts{Title: "Bottom Dollar", Width: 640, Height: 480})
    defer bronson.Cleanup()

    bronson.OnKeyDown(bronson.K_RETURN, doStuff)
    bronson.OnKeyUp(bronson.K_RETURN, doMoreStuff)
    bronson.OnMouseDown(bronson.M_LEFT, doMouseStuff)

    for !bronson.ShouldQuit() {
        bronson.ProcessOneFrame()
    }
}

func doStuff(code bronson.KeyCode) {
    fmt.Println("Doing stuff!")
}

func doMoreStuff(code bronson.KeyCode) {
    fmt.Println("Done stuff!")
}

func doMouseStuff(code bronson.KeyCode) {
    fmt.Println("Done mouse stuff!")
}