package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Traitement(url string) string {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Cookie", "__cf_bm=qfyI8Hsb8DV1Rrw16Ty.iVFlQ7JrbQqUii65A0_VxEY-1647267811-0-AbAU3Pao1XQoNIodtQn0iGOIgOKpPyWOTUQJagpQ8CABKjIqKiZXeMl4bJryS/TWZrDuZmNi5ayzefcpMhQeHjI=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
