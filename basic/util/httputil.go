package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var simpleHTTPClientOnce sync.Once
var simpleHTTPClient *SimpleHTTPClient

// SimpleHTTPClient 简易 http client
type SimpleHTTPClient struct {
}

// JSONPost 发送 json post
func (client *SimpleHTTPClient) JSONPost(url string, headers map[string]string, data interface{}) ([]byte, error) {
	bts, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bts))
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	httpClient := &http.Client{Timeout: time.Second * 30}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

// FormPost 表单post
func (client *SimpleHTTPClient) FormPost(queryUrl string, headers map[string]string, paramsMap map[string]string) ([]byte, error) {
	params := url.Values{}
	for k, v := range paramsMap {
		params.Set(k, v)
	}

	req, err := http.NewRequest("POST", queryUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	httpClient := &http.Client{Timeout: time.Second * 30}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get 普通get请求
func (client *SimpleHTTPClient) Get(queryURL string, headers map[string]string, paramsMap map[string]string) ([]byte, error) {
	params := url.Values{}
	urlObj, err := url.Parse(queryURL)
	if err != nil {
		return nil, err
	}

	for k, v := range paramsMap {
		params.Set(k, v)
	}

	urlObj.RawQuery = params.Encode()
	urlPath := urlObj.String()
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	httpClient := &http.Client{Timeout: time.Second * 30}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Upload 上传文件
func (client *SimpleHTTPClient) Upload(queryUrl string, filePath string) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filePath)
	if err != nil {
		return nil, err
	}

	//打开文件句柄操作
	fh, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(queryUrl, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// GetSimpleHTTPClient 获取http客户端
func GetSimpleHTTPClient() *SimpleHTTPClient {
	if simpleHTTPClient == nil {
		simpleHTTPClientOnce.Do(func() {
			simpleHTTPClient = &SimpleHTTPClient{}
		})
	}
	return simpleHTTPClient
}
