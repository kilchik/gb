package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type someStruct struct {
	m sync.RWMutex
}

func (s *someStruct) A() {
	s.m.RLock()
	s.B()
	s.m.RUnlock()
}

func (s *someStruct) B() {
	s.m.RLock() // Это рекурсивный «лок», который будет подсвечен линтером
	s.m.RUnlock()
}

func getContents(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("get: %w", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body: %w", err)
	}

	return string(body), nil
}


type foo struct {
	A int `json:"a"`
	B int `json:"a"`
}

// 42: dupl tag


func main()  {

	f := &foo{}
	fmt.Println(json.Unmarshal([]byte(`{"a":42,"b":43}`), f))
	fmt.Println(f)




	//m1 := map[int]string{}
	//m2 := map[int]string{}
	//m3 := make(map[int]string)
	//
	//m1[42] = "Bob"
	//m2[42] = "Bob"
	//m3[42] = "Bob"
	//
	//var m map[string]interface{}
	//json.Unmarshal([]byte(`{"foo":"bar"}`), &m)

	options := []string{"foo", "bar"}

	var features []string
	for _, o := range options {
		features = append(features, o)
	}
}
