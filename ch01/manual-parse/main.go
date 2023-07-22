package main

import (
	// 버퍼를 이용한 reader/writer
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}
	// 인수로 들어온 문자열을 숫자로
	numTimes, err = strconv.Atoi(args[0])
	// 에러 처리
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

func printUsage(w io.Writer) {
	fmt.Fprint(w, usageString)
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	// writer에 msg쓰기
	fmt.Fprintf(w, msg)
	// reader에 전송
	scanner := bufio.NewScanner(r)
	// 개행 전까지 데이터 읽고 반환 (버퍼에 bytes로 읽어짐)
	scanner.Scan()
	// 에러처리
	if err := scanner.Err(); err != nil {
		return "", err
	}
	// 읽은 데이터를 문자열로 변환
	name := scanner.Text()
	// 인수가 안들어왔으면 에러처리
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func main() {
	// 앞에서 선언한 parseArgs 함수를 이용해서 cmd에서 넣은 인수 파싱
	// Args[1:]인 이유는 0은 프로그램명이기 때문
	c, err := parseArgs(os.Args[1:])
	// 에러가 발생했다면
	if err != nil {
		// 콘솔에 에러 작성
		fmt.Fprintln(os.Stdout, err)
		// Usage: ./application <integer> [-h|-help]
		printUsage(os.Stdout)
		// 종료
		os.Exit(1)
	}
	// 인수로 들어온 값 논리적인 검증 (c > 0)
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	// 인수를 바탕으로 프로그램 실행
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

}
