package myhttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGetParser(url string, v interface{}) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic("http get error" + err.Error())
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, v)
}

func HttpPut(url, data string) {
	payload := strings.NewReader(data)
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
