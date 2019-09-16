package helpers

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpWrapperGet(address string, body map[string]string, headers map[string]string) (responseBody []byte, err error) {
	responseBody = []byte{}

	urlData := url.Values{}
	for k, v := range body {
		urlData.Set(k, v)
	}

	client := getClient(address)

	addressUrl := fmt.Sprintf("%v?%v", address, urlData.Encode())
	request, err := http.NewRequest("GET", addressUrl, strings.NewReader(urlData.Encode()))
	if err != nil {
		return
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getClient(addressUrl string) *http.Client {
	if strings.HasPrefix(addressUrl, "https") {
		transCfg := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}
		return &http.Client{Transport: transCfg}
	} else {
		return &http.Client{}
	}
}