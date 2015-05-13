package main

import (
	"fmt"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	e := &Equipment{Name: "Tester 1", Status: "running", Duration: 750 * time.Hour}
	want := "Tester 1 is running for 750h0m0s"
	got := e.Stringer()
	if want != got {
		t.Errorf("format not correct\nwant:%s\ngot :%s", want, got)
	}
}

func TestAge(t *testing.T) {
	e := new(Equipment)
	e.RegistDate = time.Date(2013, 1, 1, 8, 0, 0, 0, time.Local)
	fmt.Println(e.Age())

}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
