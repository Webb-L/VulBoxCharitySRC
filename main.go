package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) >= 2 {
		token = os.Args[1]
	} else {
		log.Fatalln("缺少Token！")
		return
	}

	_ = os.Remove("data.csv")
	getManufacturer()
}

var (
	page  = 1
	token = ""
)

// 获取数据。
func getManufacturer() {
	client := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://user.vulbox.com/api/hacker/bugs/business?page=%d&per_page=20", page), nil)
	if err != nil {
		log.Fatalln("创建请求体失败！")
		return
	}
	request.Header.Add("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("发送请求失败！")
		return
	}

	log.Printf("读取第%d页的数据。\n", page)

	if response.StatusCode != http.StatusOK {
		log.Fatalf("状态码：%d\n\n", response.StatusCode)
		return
	}

	defer response.Body.Close()
	readAll, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("响应体的数据失败！")
		return
	}

	var business Business
	err = json.Unmarshal(readAll, &business)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if len(business.Data.Data) == 0 {
		log.Println("查询结束！")
		return
	}
	for _, data := range business.Data.Data {
		Manufacturer{
			Name: data.BusName,
			Url:  data.BusUrl,
		}.save()
	}
	page++
	getManufacturer()
}

type Business struct {
	Code int `json:"code"`
	Data struct {
		CurrentPage int `json:"current_page"`
		Data        []struct {
			BusName string `json:"bus_name"`
			BusUrl  string `json:"bus_url"`
			BusType int    `json:"bus_type"`
		} `json:"data"`
		LastPage int `json:"last_page"`
		PerPage  int `json:"per_page"`
		Total    int `json:"total"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type Manufacturer struct {
	Name string
	Url  string
}

// 保存厂商名字和网站。
func (mf Manufacturer) save() error {
	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalln("读取文件失败！")
		return err
	}

	_, err = file.Write([]byte(fmt.Sprintf("\"%s\",\"%s\"\n", mf.Name, mf.Url)))
	if err != nil {
		log.Fatalln("写入数据失败！")
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
