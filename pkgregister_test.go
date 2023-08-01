package pkgregister

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 서버 내부로직 분리
func packageRegHandler(w http.ResponseWriter, r *http.Request) {
	// POST요청이 맞다면
	if r.Method == "POST" {
		// 빈 pkgData 구조체 생성
		p := pkgData{}

		// 빈 pkgRegisterResult 구조체 생성
		d := pkgRegisterResult{}
		defer r.Body.Close()
		// 직렬화되어있는 데이터 읽어오기
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 데이터 역직렬화
		err = json.Unmarshal(data, &p)
		// 데이터가 json형식에 맞게 존재하는지 체크
		if err != nil || len(p.Name) == 0 || len(p.Version) == 0 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// pkgRegisterResult에 맞게 데이터 조작
		d.ID = p.Name + "-" + p.Version 
		// pkgRegisterResult타입에 맞게 json으로 직렬화해서
		jsonData, err := json.Marshal(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		서서
		// json으로 헤더 설정 후
		w.Header().Set("Content-Type", "application/json")
		// response로 직렬화된 데이터 쓰기
		fmt.Fprint(w, string(jsonData))
	} else {
		http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
		return
	}
}

// 서버 실행과 내부로직 분리
func startTestPackageServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(packageRegHandler))
	return ts
}

func TestRegisterPackageData(t *testing.T) {
	// 테스트 서버 실행
	ts := startTestPackageServer()
	defer ts.Close()
	// 테스트 입력
	p := pkgData{
		Name:    "mypackage",
		Version: "0.1",
	}
	// 돌아온 pkgRegisterResult 직렬화 데이터 역직렬화
	resp, err := registerPackageData(ts.URL, p)
	if err != nil {
		t.Fatal(err)
	}
	// 예상한 결과와 동일한지 체크
	if resp.ID != "mypackage-0.1" {
		t.Errorf("Expected package id to be mypackage-0.1, Got: %s", resp.ID)
	}
}

func TestRegisterEmptyPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()
	p := pkgData{}
	resp, err := registerPackageData(ts.URL, p)
	if err == nil {
		t.Fatal("Expected error to be non-nil, got nil")
	}
	if len(resp.ID) != 0 {
		t.Errorf("Expected package ID to be empty, got: %s", resp.ID)
	}
}
