// The MIT License (MIT)

// Copyright (c) <year> <copyright holders>

// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"fmt"
	"github.com/peterh/liner"
	"github.com/zhemao/glisp/interpreter"
	"io"
	"log"
	"os"
	"os/user"
)

var term *liner.State
var ierr error

//var history *os.File

func shutdown() {
	if history, err := os.Create(getHistoryPath()); err != nil {
		log.Print("Error writing history: ", err)
	} else {
		term.WriteHistory(history)
		history.Close()
	}
	err := term.Close()
	if err != nil {
		log.Fatal("Error closing term: ", err)
	}
}

func getHistoryPath() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Error identifying home dir: %s", err.Error())
	}
	return u.HomeDir + "/.glisp_shell_demo_history"
}

func main() {

	defer shutdown()

	env := glisp.NewGlisp()

	//var ierr error

	ps1 := "glisp> "
	//ps2 := "...      "
	prompt := &ps1

	term = liner.NewLiner()
	if history, err := os.Open(getHistoryPath()); err == nil {
		term.ReadHistory(history)
		history.Close()
	}

	for {
		line, err := term.Prompt(*prompt)
		if err != nil {
			if err != io.EOF {
				ierr = err
			} else {
				ierr = nil
			}
			break
		}
		if line == "" || line == ";" {
			prompt = &ps1
			break
		}

		term.AppendHistory(line)

		//err = env.LoadString(line)
		res, err := env.EvalString(line)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			env.Clear()
		} else {
			fmt.Printf("> %s\n", res.SexpString())
		}
	}
}
