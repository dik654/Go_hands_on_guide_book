package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func handleCmdA(w io.Writer, args []string) error {
	var v string
	// cmd-a 플래그 생성
	fs := flag.NewFlagSet("cmd-a", flag.ContinueOnError)
	// 출력을 io.Writer로 지정
	fs.SetOutput(w)
	// cmd-a 플래그에 기본값이 argument-value인 -verb 생성
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	// cmd-a 아래 인수로 들어온 args 파싱
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Executing command A")
	return nil
}

func handleCmdB(w io.Writer, args []string) error {
	var v string
	// cmd-b 플래그 선언
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	// 출력을 io.Writer로 지정
	fs.SetOutput(w)
	// cmd-b 플래그에 기본값이 argument-value인 -verb 생성
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	// cmd-b 아래 인수로 들어온 args 파싱
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Executing command B")
	return nil
}

func printUsage(w io.Writer) {
	// 전체적인 사용법 출력
	fmt.Fprintf(w, "Usage: %s [cmd-a|cmd-b] -h\n", os.Args[0])
	// cmd-a, cmd-b일 때의 설명도 출력
	handleCmdA(w, []string{"-h"})
	handleCmdB(w, []string{"-h"})
}

func main() {
	var err error
	// 인수가 2개보다 적으면 사용법 출력 후 종료
	if len(os.Args) < 2 {
		printUsage(os.Stdout)
		os.Exit(1)
	}
	switch os.Args[1] {
	// 첫 번째 인수가 cmd-a일 경우 handleCmdA함수를 실행시키고
	// 첫 번째 인수로 들어온 cmd-a를 제외한 인수들을 넘긴다
	case "cmd-a":
		err = handleCmdA(os.Stdout, os.Args[2:])
	case "cmd-b":
		err = handleCmdB(os.Stdout, os.Args[2:])
	// 첫 번째 인수가 cmd-a, cmd-b가 아닐 경우 사용법 출력
	default:
		printUsage(os.Stdout)
	}
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}