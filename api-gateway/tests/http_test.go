package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUserInfo(t *testing.T) {
	data := make(map[string]interface{})
	data["username"] = "root"
	data["password"] = "123456"
	bytesData, err := json.Marshal(data)
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", "http://localhost:8080/base/login", reader)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("uid", "1")
	//request.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjEsIlJvbGUiOlsxLDIsM10sIkV4cGlyZVRpbWUiOjE2MDA4NDA2MDh9.8AgQSnYCbzIzyWPsy8CVD09d4FczXdsM2QLDeL2q98s")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
