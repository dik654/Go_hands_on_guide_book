package pkgquery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestPackageServer() *httptest.Server { 
	pkgData := `[ 
{"name": "package1", "version": "1.1"}, 
{"name": "package2", "version": "1.0"} 
]` 
	// 서버에 요청이 들어오면 항상 아래와 같은 함수 실행
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
		// 헤더에서 컨텐츠 타입을 json으로 설정
		w.Header().Set("Content-Type", "application/json") 
		// 직렬화된 pkgData 리턴
		fmt.Fprint(w, pkgData) 
	})) 

	return ts 
}
func TestFetchPackageData(t *testing.T) {
	ts := startTestPackageServer()
	// 테스트가 끝나면 테스트서버 닫기
	defer ts.Close()
	// 테스트 서버 url에서 데이터 가져오기
	packages, err := fetchPackageData(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	// 2개의 json 데이터가 들어왔는지 체크
	if len(packages) != 2 {
		t.Fatalf("Expected 2 packages, Got back: %d", len(packages))
	}
}