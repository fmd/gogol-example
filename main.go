package main

import (
    "github.com/fmd/bronson"
    "fmt"
)

func main() {
    bronson.Init(bronson.WindowOpts{Title: "Bottom Dollar", Width: 640, Height: 480})
    defer bronson.Cleanup()

    bronson.OnKeyDown(bronson.K_A, doStuff)
    bronson.OnKeyUp(bronson.K_A, doMoreStuff)

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