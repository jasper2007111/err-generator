package main

var goFrame = `
// Code generated by err-generator.
// source: @filename
// DO NOT EDIT!

package @packagename

import (
	"strconv"
)

// 错误码
type @classname int

const (
@key-vals
)

var @classnameMaps = map[int32]string{
@val-strs
}

func (e  @classname) String() string {
	s, ok := @classnameMaps[int32(e)]
	if ok {
		return s
	}
	return strconv.Itoa(int(e))
}
`

var goKeyVals = `	%s @classname = %d
`
var goValStrs = `	%d:"%s",
`
