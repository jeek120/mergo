package mergo_test

import (
	"testing"
	"time"

	"github.com/jeek120/mergo"
)

type testStruct struct {
	time.Duration
}

func TestIssue50Merge(t *testing.T) {
	to := testStruct{}
	from := testStruct{}

	if _, err := mergo.Merge(&to, from); err != nil {
		t.Fail()
	}
}
