package main

import (
    "os"

		"github.com/hennedo/escpos"
)

func main() {
    f, err := os.OpenFile("/dev/ttys003", os.O_RDWR, 0)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    //w := io.NewWriter(f)
    p := escpos.New(f)
		p.Bold(true).Size(2, 2).Write("Fichi the shark!")
		err = p.Print()
		if err != nil {
			panic(err)
		}

    //p.Init()
}
