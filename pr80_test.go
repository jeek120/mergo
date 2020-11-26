package mergo_test

import (
	"testing"

	"github.com/jeek120/mergo"
)

type mapInterface map[string]interface{}

func TestMergeMapsEmptyString(t *testing.T) {
	a := mapInterface{"S": ""}
	b := mapInterface{"S": "foo"}
	if _, err := mergo.Merge(&a, b); err != nil {
		t.Error(err)
	}
	if a["S"] != "foo" {
		t.Errorf("b not merged in properly: a.S.Value(%s) != expected(%s)", a["S"], "foo")
	}
}
