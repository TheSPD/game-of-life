package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

var quitGame chan bool

const maxX = 10
const maxY = 10

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func toggle(i []js.Value) {
	address := i[0].String()
	class := js.Global().Get("document").Call("getElementById", address).Get("className").String()

	var newClass string
	if strings.Contains(class, "on") {
		newClass = strings.Replace(class, "on", "off", -1)
	} else if strings.Contains(class, "off") {
		newClass = strings.Replace(class, "off", "on", -1)
	}

	js.Global().Get("document").Call("getElementById", address).Set("className", newClass)
}

func getCellInt(x, y int) js.Value {
	return js.Global().Get("document").Call("getElementById", fmt.Sprintf("cell-%d-%d", x, y))
}

func isOnInt(cell js.Value) bool {
	if cell.String() == "null" {
		return false
	}

	class := cell.Get("className").String()

	if strings.Contains(class, "on") {
		return true
	}

	return false
}

func turnOnInt(cell js.Value) {
	class := cell.Get("className").String()
	newClass := strings.Replace(class, "off", "on", -1)
	cell.Set("className", newClass)
}

func turnOffInt(cell js.Value) {
	class := cell.Get("className").String()
	newClass := strings.Replace(class, "on", "off", -1)
	cell.Set("className", newClass)
}

func gameOfLifeInt() {
	board := make([][]bool, maxY)
	for y := range board {
		board[y] = make([]bool, maxX)
	}

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			board[x][y] = gameOfLifeSingleInt(x, y)
		}
	}

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if board[x][y] {
				turnOnInt(getCellInt(x, y))
			} else {
				turnOffInt(getCellInt(x, y))
			}
		}
	}

}

func getAliveNeighbourCount(x, y int) int {
	nCell := getCellInt(x, y-1)
	neCell := getCellInt(x+1, y-1)
	eCell := getCellInt(x+1, y)
	seCell := getCellInt(x+1, y+1)
	sCell := getCellInt(x, y+1)
	swCell := getCellInt(x-1, y+1)
	wCell := getCellInt(x-1, y)
	nwCell := getCellInt(x-1, y-1)

	neighbours := []js.Value{nCell, neCell, eCell, seCell, sCell, swCell, wCell, nwCell}

	aliveNeighbourCount := 0
	for _, neighour := range neighbours {
		if isOnInt(neighour) {
			aliveNeighbourCount++
		}
	}

	return aliveNeighbourCount
}

func gameOfLifeSingleInt(x, y int) bool {
	aliveNeighbourCount := getAliveNeighbourCount(x, y)

	if isOnInt(getCellInt(x, y)) {
		if aliveNeighbourCount < 2 || aliveNeighbourCount > 3 {
			return false
		}
		return true
	} else if aliveNeighbourCount == 3 {
		return true
	}

	return false
}

func onInt(x, y int) {
	address := fmt.Sprintf("cell-%d-%d", x, y)

	class := js.Global().Get("document").Call("getElementById", address).Get("className").String()

	newClass := strings.Replace(class, "off", "on", -1)

	js.Global().Get("document").Call("getElementById", address).Set("className", newClass)
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
			gameOfLifeInt()
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}

func stopGame(i []js.Value) {
	quitGame <- true
}

func registerCallbacks() {
	js.Global().Set("toggle", js.NewCallback(toggle))
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
