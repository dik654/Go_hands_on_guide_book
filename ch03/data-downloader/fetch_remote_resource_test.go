package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	// 테스트용 http 서버 생성
	ts := httptest.NewServer(
		// 모든 요청에 아래 함수를 실행
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// Hello World를 리턴
				fmt.Fprint(w, "Hello World")
			}))
	return ts
}
func TestFetchRemoteResource(t *testing.T) {
	// 테스트용 http 서버 주소 받기
	ts := startTestHTTPServer()
	defer ts.Close()
	expected := "Hello World"
	// 테스트용 http 서버에 GET요청
	data, err := fetchRemoteResource(ts.URL)

	if err != nil {
		t.Fatal(err)
	}
	// 테스트용 http 서버가 Hello World를 리턴하는지 체크
	if expected != string(data) {
		t.Errorf("Expected response to be: %s, Got: %s", expected, data)
	}
}