package WebshellDetector

import (
	"bufio"
	"io"

	"github.com/glaslos/ssdeep"
)

/*
WebshellDetector - Refactor version 1
Date	0822
Author	Twice
Intro	Match some PHP code from webshell samples using ssdeep
*/

type sampleMatcher struct {
	hashes []string
}

func newSampleMatcher(sampleHashFile io.Reader) (*sampleMatcher, error) {

	matcher := sampleMatcher{nil}
	reader := bufio.NewReader(sampleHashFile)
	for line, _, err := reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		matcher.hashes = append(matcher.hashes, string(line))
	}

	return &matcher, nil
}

func (matcher *sampleMatcher) Match(src []byte) (bool, error) {
	hash, err := ssdeep.FuzzyBytes(src) // 计算"检测对象"的模糊哈希
	if err != nil {
		return false, err
	}

	for _, h := range matcher.hashes {
		score, err := ssdeep.Distance(hash, h) // 和"webshell模糊哈希列表"做相似度比较
		if err != nil {
			return false, err
		}
		if score > 90 {
			return true, nil
		}
	}

	return false, nil
}
