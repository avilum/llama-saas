package main

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"

	"github.com/valyala/fasthttp"
)

func main() {
	LLAMA_MAIN := "/path/to/llama.cpp/main"
	LLAMA_MODEL_PATH := "/path/to/7B/ggml-model-f16.bin"
	LLAMA_NUM_THREADS := "16"
	MAX_OUTPUT_LENGTH := int(256)

	handler := func(ctx *fasthttp.RequestCtx) {
		modelInput := ctx.QueryArgs().Peek("text")

		cmd := exec.Command(LLAMA_MAIN, "-m", LLAMA_MODEL_PATH, "-t", LLAMA_NUM_THREADS, "-n", strconv.Itoa(MAX_OUTPUT_LENGTH), "-p", fmt.Sprintf("%s", modelInput))
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			return
		}
		if err := cmd.Start(); err != nil {
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Set("Content-Type", "text/plain; charset=utf-8")
		ctx.Response.Header.Set("Cache-Control", "no-cache")
		ctx.Response.SetBodyStream(io.Reader(stdout), -1)
	}

	port := 8080
	addr := ":" + strconv.Itoa(port)
	if err := fasthttp.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
