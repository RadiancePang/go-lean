package tools

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
}

const (
	//实时间隔，单位s
	RETRY_INTERVAL = 1
)

//带重试机制
func HttpRequestByGet(url string, retry int) ([]byte, error) {

	var resp *http.Response
	var err error

	for i := 0; i < retry; i++ {
		resp, err = http.Get(url)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(RETRY_INTERVAL+i) * time.Second)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	return result, nil
}

//带重试机制
func HttpRequestByPost(req []byte, url string, retry int) ([]byte, error) {

	var resp *http.Response
	var err error

	for i := 0; i < retry; i++ {
		body := bytes.NewBuffer(req)
		resp, err = http.Post(url, "application/json;charset=utf-8", body)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(RETRY_INTERVAL+i) * time.Second)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//通过识别url是否以https，来自动选择相应的调用方式
func HttpGet(url string, retry int) ([]byte, error) {
	if strings.HasPrefix(url, "https://") {
		return httpsGet(url, retry)
	} else {
		return HttpRequestByGet(url, retry)
	}
}

func HttpPost(req []byte, url string, retry int) ([]byte, error) {
	if strings.HasPrefix(url, "https://") {
		return httpsPost(req, url, retry)
	} else {
		return HttpRequestByPost(req, url, retry)
	}
}

//带重试机制
func httpsGet(url string, retry int) ([]byte, error) {

	var resp *http.Response
	var err error

	for i := 0; i < retry; i++ {
		resp, err = client.Get(url)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(RETRY_INTERVAL+i) * time.Second)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//带重试机制
func httpsPost(req []byte, url string, retry int) ([]byte, error) {
	var resp *http.Response
	var err error

	for i := 0; i < retry; i++ {
		body := bytes.NewBuffer(req)
		resp, err = client.Post(url, "application/json;charset=utf-8", body)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(RETRY_INTERVAL+i) * time.Second)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
