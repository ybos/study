package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// 详细的城市数据结构
	Data struct {
		ID   int    `json:"id"`
		PID  int    `json:"pid"`
		NAME string `json:"name"`
	}

	// 省
	Area struct {
		P  []Data   `json:"p"`
		C  []Data   `json:"c"`
		D  []Data   `json:"d"`
		DA struct{} `json:"data"`
	}

//	// gResult  映射到从搜索拿到的结构文档
//	gResult struct {
//		GsearchResultClass string `json:"GsearchResultClass"`
//		UnescapedURL       string `json:"unescapedUrl"`
//		URL                string `json:"url"`
//		VisibleURL         string `json:"visibleUrl"`
//		CacheURL           string `json:"cacheUrl"`
//		Title              string `json:"title"`
//		TitleNoFormatting  string `json:"titleNoFormatting"`
//		Content            string `json:"content"`
//	}

//	// gResponse 包含顶级的文档
//	gResponse struct {
//		ResponseData struct {
//			Results []gResult `json:"results"`
//		} `json:"responseData"`
//	}
)

func main() {
	uri := "http://shop-api.ecovacs.cn/shopApi/commonApi/region"

	// 获取内容
	resp, err := http.Get(uri)

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var area Area
	err = json.NewDecoder(resp.Body).Decode(&area)

	if err != nil {
		log.Println("Error2: ", err)
		return
	}

	fmt.Println(area)
	fmt.Println("----------------\n")

	fmt.Println("p:")
	fmt.Println(area.P)
	fmt.Println("----------------\n")

	fmt.Println("c:")
	fmt.Println(area.C)
	fmt.Println("----------------\n")

	fmt.Println("d:")
	fmt.Println(area.D)
	fmt.Println("----------------\n")
}
