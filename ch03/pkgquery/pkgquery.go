package pkgquery

import (
	"encoding/json"
	"io"
	"net/http"
)

type pkgData struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

func fetchPackageData(url string) ([]pkgData, error) {
	var packages []pkgData
	// 인수로 들어온 url로 GET요청
	r, err := http.Get(url)
	// GET 요청 중 에러가 생기면 공백의턴 구조체와 에러 리턴
	if err != nil {
		return nil, err
	}
	// 함수가 종료되면 body닫기
	defer r.Body.Close()
	// 데이터 타입이 application/json이 아니라면 pkgData 구조체 리턴
	if r.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}// 직렬화 되어있는 바디를 모두 읽은 뒤
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return packages, err
	}
	// 역직렬화 진행
	err = json.Unmarshal(data, &packages)
	return packages, err
}