package main

import (
    "github.com/ying32/govcl/vcl"
)

func main() {

    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateFormFromFile("Form1.gfm", &Form1)
    vcl.Application.Run()

}

