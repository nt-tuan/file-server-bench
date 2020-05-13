package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var staticPrefix = "http://104.248.152.202/images/static"
var resizePrefix = "http://104.248.152.202/images/size"
var originPrefix = "https://beta.phatthanhcafe.com"

func getFiles(folder string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(folder)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func getOrigin(file string) (err error) {
	var res *http.Response
	res, err = http.Get(originPrefix + "/" + file)
	if err != nil {
		return
	}
	if res == nil {
		return errors.New("No responses")
	}

	if res.StatusCode != 200 {
		return errors.New("No success")
	}
	return
}

func getStatic(file string) (err error) {
	var res *http.Response
	res, err = http.Get(staticPrefix + "/" + file)
	if err != nil {
		return
	}
	if res == nil {
		return errors.New("No responses")
	}

	if res.StatusCode != 200 {
		return errors.New("No success")
	}
	return
}

func getResizeWidth(file string, width uint, height uint) (err error) {
	var res *http.Response
	var url = fmt.Sprintf("%v/%v/%v/%v", resizePrefix, width, height, file)
	res, err = http.Get(url)
	if err != nil {
		return
	}
	if res == nil {
		return errors.New("No responses")
	}

	if res.StatusCode != 200 {
		return errors.New("No success")
	}
	return
}
