package main

import (
	"fmt"
	"os"
	"testing"
)

var files []string = []string{
	"static/ap.png",
	"static/biepi-mce-1g.png",
	"static/drum.png",
}

func setup() (err error) {
	//files, err = getFiles("files")
	return
}

func BenchmarkOrigin(b *testing.B) {
	errCount := 0
	//len := len(files)
	for i := 0; i < b.N; i++ {
		for _, file := range files {
			if err := getOrigin(file); err != nil {
				errCount++
				fmt.Println(file)
			}
		}
	}
	//fmt.Printf("Errors: %v/%v\n", errCount, b.N*len)
}

func BenchmarkStatic(b *testing.B) {
	// run the Fib function b.N times
	errCount := 0
	//len := len(files)
	for i := 0; i < b.N; i++ {
		for _, file := range files {
			if err := getStatic(file); err != nil {
				errCount++
				fmt.Println(file)
			}
		}
	}
	//fmt.Printf("Errors: %v/%v\n", errCount, b.N*len)
}

func TestResize(t *testing.T) {
	for _, file := range files {
		if err := getResizeWidth(file, 200, 0); err != nil {
			url := fmt.Sprintf("%v/%v/%v/%v", resizePrefix, 200, 0, file)
			fmt.Println(url)
			t.Error(err)
			return
		}
	}
}

func BenchmarkResize(b *testing.B) {
	// run the Fib function b.N times
	widths := []uint{600}
	errCount := 0
	//len := len(files)
	//lenWidth := 10
	for i := 0; i < b.N; i++ {
		for _, w := range widths {
			for _, file := range files {
				if err := getResizeWidth(file, w, 0); err != nil {
					errCount++
				}
			}
		}
	}
	//fmt.Printf("Errors: %v/%v\n", errCount, b.N*len*lenWidth)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
