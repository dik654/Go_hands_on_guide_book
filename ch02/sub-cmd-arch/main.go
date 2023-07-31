package main

import (
	"github.com/dik654/Go_hands_on_guide_book/ch02/sub-cmd-arch/cmd"
	"errors"
	"fmt"
	"io"
	"os"
)

var errInvalidSubCommand = errors.New("Invalid sub-command specified")

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	// 인수가 없는경우 커스텀 에러 리턴
	if len(args) < 1{
		err = errInvalidSubCommand
	} else {
		// 첫 번째 인수로 접근 방법 지정, 나머지는 인수로 리턴
		switch args[0] {
		case "http":
			err = cmd.HandleHttp(w, args[1:])
		case "grpc":
			err = cmd.HandleGrpc(w, args[1:])
		case "-h":
			printUsage(w)
		case "-help":
			printUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}
	// 만약 커스텀 에러나 서버가 지정되지않은 경우 에러와 사용법을 리턴
	if errors.Is(err, cmd.ErrNoServerSpecified) || errors.Is(err, errInvalidSubCommand) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}