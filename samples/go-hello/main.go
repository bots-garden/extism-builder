package main

import (
    "github.com/extism/go-pdk"
)

//export hello
func hello() {
    input := pdk.Input()

    message := "🤗 Hello " + string(input)
    
    mem := pdk.AllocateString(message)
    pdk.OutputMemory(mem)
}

func main() {}