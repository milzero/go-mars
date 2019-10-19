package util

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func HttpPost(urlStr string, header map[string]string, body []byte) ([]byte, int, error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), 0, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return nil, resp.StatusCode, nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}

func HttpGet(urlStr string, header map[string]string) ([]byte, int, error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*30)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 30))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 30,
		},
	}
	req, _ := http.NewRequest("GET", urlStr, nil)

	req.Header.Set("Content-Type", "text/html;charset=utf-8")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return nil, resp.StatusCode, nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}

func HttpDelete(url string, header map[string]string) ([]byte, int, error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	req, _ := http.NewRequest("DELETE", url, nil)

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return nil, resp.StatusCode, nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}
