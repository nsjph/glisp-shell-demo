package main

import (
	_ "bufio"
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
	return u.HomeDir + "/.goshell_history"
}

func main() {

	defer shutdown()

	env := glisp.NewGlisp()

	//var ierr error

	ps1 := "goshell> "
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
