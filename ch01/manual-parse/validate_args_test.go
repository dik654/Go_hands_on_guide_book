package main

import (
	// 지정한 에러 생성용
	"errors"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	// 익명구조체 슬라이스
	tests := []struct {
		c   config
		err error
	}{
		// 테스트케이스
		{
			c:   config{},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	for _, tc := range tests {
		// 실제 실행
		err := validateArgs(tc.c)
		// 테스트 케이스에서 예상한 에러와 실제 에러 비교해서 다르다면 에러표시
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		// 테스트 케이스에서는 통과인데 실제에선 에러가 날 경우
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
	}
}
