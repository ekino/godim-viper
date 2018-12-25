// Copyright 2018 ekino.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package godimviper

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/spf13/viper"
)

var yamlTest = []byte(`
Astring: hereitis
Abool: true
AsliceOfString:
- one
- two
age: 42
AFloat: 24.1
AMapOfString:
 key1: one
 key2: two
AMapOfInt:
 k1: 1
 k2: 2
AMapOfSliceString:
 first:
 - banana
 - apple
 second:
 - whiskey
 - rhum
dur: 4200
time: 123456789
i8: 5
`)

type all struct {
	l   []string
	m   map[string]string
	n   map[string]interface{}
	o   map[string][]string
	s   string
	i   int
	i32 int32
	i64 int64
	b   bool
	f   float64
	d   time.Duration
	t   time.Time
	i8  int8
}

func TestConfigFunc(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(bytes.NewBuffer(yamlTest))
	if err != nil {
		t.Fatalf("unable to read config array %+v", err)
	}
	a := &all{}
	val := reflect.ValueOf(a).Elem()

	_, err = ViperForGodim("Astring", val.FieldByName("s"))

	if err == nil {
		t.Fatalf("Must have an initialiszation error ")
	}
	SetViperConfig(v)

	// String
	s, err := ViperForGodim("Astring", val.FieldByName("s"))

	if s != "hereitis" {
		t.Fatalf("wrong result, waiting for hereitis, got %s", s)
	}

	// int
	i, err := ViperForGodim("age", val.FieldByName("i"))

	if i != 42 {
		t.Fatalf("wrong result, waiting for 42, got %s", i)
	}

	i32, err := ViperForGodim("age", val.FieldByName("i32"))
	var i42 int32
	i42 = 42
	if i32 != i42 {
		t.Fatalf("wrong result, waiting for 42, got %s", i32)
	}

	i64, err := ViperForGodim("age", val.FieldByName("i64"))
	var i642 int64
	i642 = 42
	if i64 != i642 {
		t.Fatalf("wrong result, waiting for 42, got %s", i64)
	}

	ok, err := ViperForGodim("Abool", val.FieldByName("b"))
	if !(ok.(bool)) {
		t.Fatalf("wrong bool result, got false")
	}

	f, err := ViperForGodim("AFloat", val.FieldByName("f"))
	var f64 float64
	f64 = 24.1
	if f != f64 {
		t.Fatalf("Wrong result, waiting for 24,1 and got %s", f)
	}
	var l []string

	ll, err := ViperForGodim("AsliceOfString", val.FieldByName("l"))

	l = ll.([]string)
	if l[0] != "one" || l[1] != "two" {
		t.Fatalf("Wrong result, waiting for one and two, got %+v", l)
	}

	mm, err := ViperForGodim("AMapOfString", val.FieldByName("m"))
	m := mm.(map[string]string)

	if m["key1"] != "one" || m["key2"] != "two" {
		t.Fatalf("Wrong result, waiting for one and two, got %+v", mm)
	}

	nn, err := ViperForGodim("AMapOfInt", val.FieldByName("n"))
	n := nn.(map[string]interface{})

	if n["k1"] != 1 || n["k2"] != 2 {
		t.Fatalf("Wrong result, waiting for 1 and 2, got %+v", mm)
	}

	oo, err := ViperForGodim("AMapOfSliceString", val.FieldByName("o"))
	o := oo.(map[string][]string)

	s1 := o["first"]
	s2 := o["second"]
	if s1[0] != "banana" || s1[1] != "apple" {
		t.Fatalf("Wrong result in first entry, waiting for banana and apple, got %+v", s1)
	}
	if s2[0] != "whiskey" || s2[1] != "rhum" {
		t.Fatalf("Wrong result in second entry, waiting for whiskey and rhum, got %+v", s2)
	}

	d, err := ViperForGodim("dur", val.FieldByName("d"))
	var dd time.Duration
	dd = 4200
	if d != dd {
		t.Fatalf("Wrong result, waiting for 4200 and got %s", d)
	}

	tt, err := ViperForGodim("time", val.FieldByName("t"))
	ti := time.Unix(123456789, 0)
	if ti != tt {
		t.Fatalf("Wrong result, waiting for 123456789, got %s", tt)
	}

}

func TestNotSupported(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(bytes.NewBuffer(yamlTest))
	if err != nil {
		t.Fatalf("unable to read config array %+v", err)
	}
	a := &all{}
	val := reflect.ValueOf(a).Elem()

	_, err = ViperForGodim("i8", val.FieldByName("i8"))
	if err == nil {
		t.Fatalf("Must not be supported without test change")
	}
}
