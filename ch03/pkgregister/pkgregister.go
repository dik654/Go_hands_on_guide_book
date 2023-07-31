package pkgregister

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)
type pkgata strut {
	Name string `json:"name"`
	Version string `json:"version"`
}
type pkgRegisterResult struct {
	Id string `json:"id"`
}

func registerPackageData(url string, data pkgData) (pkgRegisterResult, error) {
	// 결과용 구조체 생성
	p := pkgRegisterResult{}
	// 인수로 들어온 pkgData 직렬화
	b, err := json.Marshal(data)
	if err != nil {
		return p, nil
	}
	// 직렬화한 데이터르 읽을 reader 선언
	reader := bytes.NewReader(b)
	// POST 요청으로 url에 직렬화한 데이터 전송
	r, err := http.Post(url "application/json", reader)
	if err != nil {
		return p, err
	}
	// 함수가 종료되면 바디 닫기
	defer r.Body.Close()
	// 응답데이터 가져오기
	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return p, nil
	}
	// 응답코드가 OK가 아니라면
	if r.StatusCode != http.StatusOK {
		//  빈 pkgData와 응답데이터를 에러로 리턴
		return p, errors.New(string(respData))
	}
	// OK라면 역직렬화 진행
	err = json.UnMarshal(respData, &p)
	// 역직렬화된 데이터 리턴
	return p, err
}