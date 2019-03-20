// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports net/http/cookiejar'. DO NOT EDIT.

import (
	"net/http/cookiejar"
	"reflect"
)

func init() {
	Value["net/http/cookiejar"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(cookiejar.New),

		// type definitions
		"Jar":              reflect.ValueOf((*cookiejar.Jar)(nil)),
		"Options":          reflect.ValueOf((*cookiejar.Options)(nil)),
		"PublicSuffixList": reflect.ValueOf((*cookiejar.PublicSuffixList)(nil)),
	}
}