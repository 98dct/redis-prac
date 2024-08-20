package main

import (
	"reflect"
	"testing"
	"unsafe"
)

var testStr = "NUhuJCDHhhel world你好"
var testBytes = []byte(testStr)

func str2sliv1(s string) []byte {
	return []byte(s)
}

func str2Sliv2(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func sli2strv1(b []byte) string {
	return string(b)
}

func sli2strv2(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 3.58ns
func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str2sliv1(testStr)
	}
}

// 0.236ns
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str2Sliv2(testStr)
	}
}

// 3.017ns
func Benchmark3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sli2strv1(testBytes)
	}
}

// 0.238ns
func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sli2strv2(testBytes)
	}
}
