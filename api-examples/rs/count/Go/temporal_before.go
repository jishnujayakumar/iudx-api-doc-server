package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://pune.iudx.org.in/resource-server/pscdcl/v1/count"

	payload := strings.NewReader("{\n\t\"id\": \"rbccps.org/aa9d66a000d94a78895de8d4c0b3a67f3450e531/pscdcl/aqm-bosch-climo/ABC Farm House Junction_4\",\n\t\"time\": \"2019-08-19T00:00:00.000Z\",\n\t\"TRelation\": \"before\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}