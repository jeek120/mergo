package mergo_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jeek120/mergo"
)

type transformer struct {
	m map[reflect.Type]func(dst, src reflect.Value) error
}

func (s *transformer) Transformer(t reflect.Type) func(dst, src reflect.Value) error {
	if fn, ok := s.m[t]; ok {
		return fn
	}
	return nil
}

type foo struct {
	S   string
	Bar *bar
}

type bar struct {
	I int
	S map[string]string
}

func (b *bar) String() string {
	return fmt.Sprintf("{I: %d, S: %+v}", b.I, b.S)
}

func TestEqual(t *testing.T) {
	old := foo{}
	a := foo{}
	b := foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}

	want := true
	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo"}
	a = foo{S: "foo"}
	b = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100}}
	a = foo{S: "foo", Bar: &bar{I: 100}}
	b = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1"}}}
	b = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v3"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo", Bar: &bar{I: 101, S: map[string]string{"k1": "v1", "k2": "v2"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo1", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	want = false

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo"}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo", Bar: &bar{I: 100}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("合并的hasChange结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("change结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}

	old = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	a = foo{S: "foo", Bar: &bar{I: 100, S: map[string]string{"k1": "v1", "k2": "v2"}}}
	b = foo{Bar: &bar{S: map[string]string{"k1": "v1"}}}

	if hasChange, err := mergo.Merge(&a, &b); err != nil || hasChange != want {
		if err != nil {
			t.Error(err)
		}
		if want != hasChange {
			t.Errorf("change结果是%v,希望是%v，before:%+v, after:%+v", hasChange, want, old, a)
		}
	}
}

func TestMergeWithTransformerNilStruct(t *testing.T) {
	a := foo{S: "foo"}
	b := foo{Bar: &bar{I: 2, S: map[string]string{"foo": "bar"}}}

	if _, err := mergo.Merge(&a, &b, mergo.WithOverride, mergo.WithTransformers(&transformer{
		m: map[reflect.Type]func(dst, src reflect.Value) error{
			reflect.TypeOf(&bar{}): func(dst, src reflect.Value) error {
				// Do sthg with Elem
				t.Log(dst.Elem().FieldByName("I"))
				t.Log(src.Elem())
				return nil
			},
		},
	})); err != nil {
		t.Error(err)
	}

	if a.S != "foo" {
		t.Errorf("b not merged in properly: a.S.Value(%s) != expected(%s)", a.S, "foo")
	}

	if a.Bar == nil {
		t.Errorf("b not merged in properly: a.Bar shouldn't be nil")
	}
}

func TestMergeNonPointer(t *testing.T) {
	dst := bar{
		I: 1,
	}
	src := bar{
		I: 2,
		S: map[string]string{
			"a": "1",
		},
	}
	want := mergo.ErrNonPointerAgument

	if _, got := mergo.Merge(dst, src); got != want {
		t.Errorf("want: %S, got: %S", want, got)
	}
}

func TestMapNonPointer(t *testing.T) {
	dst := make(map[string]bar)
	src := map[string]bar{
		"a": {
			I: 2,
			S: map[string]string{
				"a": "1",
			},
		},
	}
	want := mergo.ErrNonPointerAgument
	if _, got := mergo.Merge(dst, src); got != want {
		t.Errorf("want: %S, got: %S", want, got)
	}
}
