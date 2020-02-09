package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

var quitGame chan bool

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

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

func on(i []js.Value) {
	x, _ := strconv.Atoi(i[0].String())
	y, _ := strconv.Atoi(i[1].String())

	onInt(x, y)
}

func onInt(x, y int) {
	address := fmt.Sprintf("cell-%d-%d", x, y)

	class := js.Global().Get("document").Call("getElementById", address).Get("className").String()

	newClass := strings.Replace(class, "off", "on", -1)

	js.Global().Get("document").Call("getElementById", address).Set("className", newClass)
}

func off(i []js.Value) {
	x, _ := strconv.Atoi(i[0].String())
	y, _ := strconv.Atoi(i[1].String())

	offInt(x, y)
}

func offInt(x, y int) {
	address := fmt.Sprintf("cell-%d-%d", x, y)

	class := js.Global().Get("document").Call("getElementById", address).Get("className").String()

	newClass := strings.Replace(class, "on", "off", -1)

	js.Global().Get("document").Call("getElementById", address).Set("className", newClass)
}

func startGame(i []js.Value) {
	minDelay := 100
	baseDelay := 1000
	speedSettings := 10
	speed, _ := strconv.Atoi(i[0].String())
	speed = min(speedSettings, speed)
	delay := baseDelay - (baseDelay/speedSettings)*speed + minDelay

	go gameInt(delay)
}

func gameInt(delay int) {
	for {
		select {
		case <-quitGame:
			return
		default:
			onInt(3, 4)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			offInt(3, 4)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}

func stopGame(i []js.Value) {
	quitGame <- true
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
	js.Global().Set("toggle", js.NewCallback(toggle))
	js.Global().Set("on", js.NewCallback(on))
	js.Global().Set("off", js.NewCallback(off))
	js.Global().Set("startGame", js.NewCallback(startGame))
	js.Global().Set("stopGame", js.NewCallback(stopGame))
}

func main() {
	c := make(chan struct{}, 0)
	quitGame = make(chan bool)
	println("Go Webassembly Initialized")
	registerCallbacks()

	<-c
}
