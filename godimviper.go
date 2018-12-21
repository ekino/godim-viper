// Copyright 2018 ekino.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package godimviper

import (
	"fmt"
	"reflect"
	"time"

	"github.com/spf13/viper"
)

var viperConfig *viper.Viper

// SetViperConfig set the Viper configured
func SetViperConfig(vc *viper.Viper) {
	viperConfig = vc
}

// ViperForGodim is a mapper between Viper and Godim for configuration
func ViperForGodim(key string, val reflect.Value) (res interface{}, err error) {
	if viperConfig == nil {
		err = fmt.Errorf("Viper not setted")
		return
	}
	e := fmt.Errorf("Type Not Supported")
	kind := val.Kind()
	typ := val.Type()
	switch kind {
	case reflect.String:
		res = viperConfig.GetString(key)
	case reflect.Int:
		res = viperConfig.GetInt(key)
	case reflect.Int32:
		res = viperConfig.GetInt32(key)
	case reflect.Int64:
		var d time.Duration
		if typ == reflect.TypeOf(d) {
			res = viperConfig.GetDuration(key)
		} else {
			res = viperConfig.GetInt64(key)
		}
	case reflect.Bool:
		res = viperConfig.GetBool(key)
	case reflect.Float64:
		res = viperConfig.GetFloat64(key)
	case reflect.Slice:
		res = viperConfig.GetStringSlice(key)
	case reflect.Map:
		t := val.Type().Elem()
		switch t.Kind() {
		case reflect.String:
			res = viperConfig.GetStringMapString(key)
		case reflect.Interface:
			res = viperConfig.GetStringMap(key)
		case reflect.Slice:
			if t.Elem().Kind() == reflect.String {
				res = viperConfig.GetStringMapStringSlice(key)
			} else {
				err = e
			}
		default:
			err = e
		}
	case reflect.Struct:
		t := val.Type()
		if t == reflect.TypeOf(time.Time{}) {
			res = viperConfig.GetTime(key)
		}
	default:
		err = e
	}
	return
}
