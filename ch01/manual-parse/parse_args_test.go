package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	// test해볼 입력 arg
	args []string
	// 나와야하는 에러
	err error
	// main.go에 있는 config 구조체
	config
}

func TestParseArgs(t *testing.T) {
	// 테스트 케이스들
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    errors.New("Invalid number of arguments"),
			config: config{printUsage: false, numTimes: 0},
		},
	}
	// 테스트 케이스별 테스트
	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		// %v는 기본 형식을 문자열로 변환하는 포맷지정자
		// 실제 에러와 test case에서 예상한 에러가 다를 경우
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be : %v, got: %v\n", tc.err, err)
		}
		// 실제는 에러가 나타나지만 test case에서는 에러가 없는 경우
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		// -h, --help을 넣었을 때 test case와 서로 다른 경우
		if c.printUsage != tc.config.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.config.printUsage, c.printUsage)
		}
		// cmd에서 인수로 넣은 횟수가 test case와 서로 다르게 나올 때
		if c.numTimes != tc.config.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.config.numTimes, c.numTimes)
		}
	}
}
