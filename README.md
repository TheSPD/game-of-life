# game-of-life

I am trying to learn webAssembly, while at the same time, trying to build something interesting. As a result, I have not paid much attention to the design and put forth a quick and dirty

## Pre-requisites

Go should be installed (not required but good to have). Follow the install instructions from [here](https://golang.org/doc/install).  
This application has been developed for 1.11 and then ported to work with 1.13.

## How to run

If you have go installed, then you might use the below command to run a HTTP server that serves the files in "static" directory. You should be in the root directory of this application to run this app.

1. Run the HTTP Server using the command `go run server.go -listen :<port> -dir static`
2. Open your browser and go to `localhost:<port>`

## Development

This section tries to help you if you decide to develop if you decide to.

### Structure

static/  
    |__index.html  
    |__style.css  
main.go

`index.html` is the plain HTML to display the grid and two buttons "Start" and "Stop"  
`style.css` has the CSS to help display the grid  
`main.go` has all the methods to play the game of life

To start the development, copy the `wasm_exec.js` provided by your go runtime into the static directory `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static/`  
To build the wasm file, use the command `GOOS=js GOARCH=wasm go build -o static/lib.wasm main.go`
