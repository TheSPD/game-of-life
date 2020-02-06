package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

func add(i []js.Value) {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
}

func subtract(i []js.Value) {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
}

func toggle(i []js.Value) {
	x, _ := strconv.Atoi(i[0].String())
	y, _ := strconv.Atoi(i[1].String())

	address := fmt.Sprintf("cell-%d-%d", x, y)

	class := js.Global().Get("document").Call("getElementById", address).Get("className").String()

	var newClass string
	if strings.Contains(class, "on") {
		newClass = strings.Replace(class, "on", "off", -1)
	} else if strings.Contains(class, "off") {
		newClass = strings.Replace(class, "off", "on", -1)
	}

	fmt.Println(class)
	fmt.Println(newClass)
	js.Global().Get("document").Call("getElementById", address).Set("className", newClass)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
	js.Global().Set("toggle", js.NewCallback(toggle))
}

func main() {
	c := make(chan struct{}, 0)
	println("Go Webassembly Initialized")
	registerCallbacks()

	<-c
}
