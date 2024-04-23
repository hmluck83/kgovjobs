package main

import (
	"fmt"

	"github.com/hmluck83/kgovjobs/notifier"
	"github.com/hmluck83/kgovjobs/retriever"
)

func main() {
    eF, err := retriever.Retrieve()
    if err != nil {
        fmt.Println(err)
        return
    }

    for i := range *eF {
        fmt.Println((*eF)[i].String())
    }

    notifier.Send(eF)
}