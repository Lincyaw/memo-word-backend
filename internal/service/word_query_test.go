package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordQuery(t *testing.T) {
	ans := QueryWord("example")
	assert.NotEmpty(t, ans, "")
	fmt.Println(ans)
}

func TestGetWordByTime(t *testing.T) {
	ans := GetWordByTime(10)
	fmt.Println(ans)
}
