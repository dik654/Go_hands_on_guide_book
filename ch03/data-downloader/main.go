package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchRemoteResource(url string) ([]byte, error) {
	// 인수로 들어온 url로 GET요청
	r, err := http.Get(url)
	// GET요청 중 에러처리
	if err != nil {
		return nil, err
	}
	// 함수 종료시 body 닫기
	defer r.Body.Close()
	// 응답의 body 리턴
	return io.ReadAll(r.Body)
}

func main() {
	// 인수가 2개가 아니라면
	if len(os.Args) != 2 {
		// 인수를 리턴하고 문제점을 반환
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data form")
		os.Exit(1)
	}
	// url을 인수로하여 GET요청 후 body 받기
	body, err := fetchRemoteResource(os.Args[1])
	// GET요청 중 에러 리턴
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	// 받은 body 화면에 리턴
	fmt.Fprintf(os.Stdout, "%s\n", body)
}