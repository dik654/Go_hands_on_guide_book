package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int
	name     string
}

var errInvalidPosArgSpecified = errors.New("More than one position argument specified")

func getName(r io.Reader, w io.Writer) (string, error) {
	// 인수로 들어온 값을 읽는 scanner 준비
	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when done.\n"
	// writer에 위의 이름을 넣어달라는 요청문구 쓰기
	fmt.Fprintf(w, msg)
	// 콘솔에 유저가 이름을 작성하기를 대기
	scanner.Scan()
	// 에러처리
	if err := scanner.Err(); err != nil {
		return "", err
	}
	// 들어온 이름 저장
	name := scanner.Text()
	// 이름을 아예 안넣었다면 에러처리
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}

func greetUser(c config, w io.Writer) {
	// Sprint는 문자열을 생성해서 메모리에 저장(화면에 출력x)
	msg := fmt.Sprintf("Nice to meet you %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		// 파일에 쓰기
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	// config 구조체에 name이 지정되지않은 경우
	if len(c.name) == 0 {
		// 콘솔화면에 이름 작성 요청
		c.name, err = getName(r, w)
		// 에러 버블링(getName에서 일어난 에러가 타고 올라감)
		if err != nil {
			return err
		}
	}
	// 화면에 greet 내용 numTimes만큼 반복해서 출력
	greetUser(c, w)
	return nil
}

func validateArgs(c config) error {
	// numTimes값의 유효성 검사
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	// config 구조체 선언
	c := config{}
	// 오류가 발생해도 다음 플래그 처리
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	// fs의 출력 대상을 io.Writer로 설정
	fs.SetOutput(w)
	fs.Usage = func() {
		var usageString = `
A greeter application which prints the name you entered a specified number of times.

Usage of %s: <options> [name]`
		// 먼저 파일에 usageString 출력
		fmt.Fprintf(w, usageString, fs.Name())
		// 빈 줄 두번 뛰기
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		// 프로그램의 플래그 정보 출력 (greeter)
		fs.PrintDefaults()
	}
	// 기본값이 0인 n flag 선언 (인수로 -n 3 처럼 주면 numTimes에 3이 할당된다)
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	// 시작 인수 파싱
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	// 인수 개수가 1개 이상이면 에러
	if fs.NArg() > 1 {
		return c, errInvalidPosArgSpecified
	}
	// 인수가 1개면 그 인수를 name에 할당
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	// 구조체 리턴
	return c, nil
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		if errors.Is(err, errInvalidPosArgSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
