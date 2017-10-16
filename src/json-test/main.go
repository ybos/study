package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

var s string = `{
  "data": [
    {
      "token": "9vVbnw8qI213fQUTjFX5lGHxtc8AAAAArQQAAABFverOdyZkqOBAfZwM_YEmzRqn-keh_Q0NdKWTGVjqoEPMDczvUEiKRSs46dCRcA",
      "userid": "1234567890",
      "address": "%E6%99%93%E5%B3%B0%E5%9B%AD",
      "f": "android",
      "version": "5.11.3",
      "wm_latitude": "22954393",
      "wm_logintoken": "9vVbnw8qI213fQUTjFX5lGHxtc8AAAAArQQAAABFverOdyZkqOBAfZwM_YEmzRqn-keh_Q0NdKWTGVjqoEPMDczvUEiKRSs46dCRcA",
      "wm_mac": "14%3Ad1%3A1f%3A8c%3A59%3A55",
      "request_id": "656A40AD-8897-496D-86BC-8579B1C25901",
      "uuid": "xxxx",
      "wm_actual_longitude": "113321605",
      "wm_actual_latitude": "22954631",
      "wm_ctype": "android",
      "app": "4",
      "wm_visitid": "f6de1958-25a0-40d5-9981-0d95585e3bb7",
      "wm_did": "A0000059B8DDD9",
      "platform": "4",
      "seq_id": "80502",
      "wm_dversion": "22_5.1.1",
      "wm_longitude": "113322115",
      "wm_channel": "1009",
      "wm_uuid": "xxxx",
      "wm_dtype": "HUAWEI%20TIT-CL00",
      "push_token": "dpsh9a98b44ca669a13deef1f850de4988f5atpu",
      "partner": "4",
      "wm_appversion": "5.11.3",
      "utm_medium": "android",
      "utm_content": "A0000059B8DDD9",
      "utm_term": "51103",
      "utm_source": "1009",
      "ci": "20",
      "utm_campaign": "AwaimaiBwaimai",
      "lat": "22.954631",
      "lng": "113.321605"
    },
    {
      "trace_tag": "%7B%22extra%22:%7B%22matched_longitude%22:%22123458872%22,%22auto_longitude%22:%22123460416%22,%22auto_latitude%22:%2241801396%22,%22matched_ch_address%22:%22%E4%B8%AD%E8%A1%97%E4%B8%8A%E4%B9%98%22,%22matched_latitude%22:%2241802940%22,%22is_match%22:%221%22%7D,%22tgt_page%22:%22p_activity%22,%22src_page%22:%22p_homepage%22,%22req_time%22:%221506675652736%22%7D",
      "userid": "1234567891",
      "token": "7XNJpaQQT4KJywf_ZCP3h9xD-tsAAAAAhwQAAOzi_ZHQzx9DtWLXZcVZzV6BvCZmvfOs5UKB5jCjzJ0gHcHUh56yBDJGzyF6c2AL4g",
      "lat": "41.802940",
      "lng": "123.458872",
      "f": "iphone",
      "address": "%E4%B8%AD%E8%A1%97%E4%B8%8A%E4%B9%98",
      "wm_ctype": "iphone",
      "wm_dversion": "11.0",
      "utm_term": "5.11.0",
      "wm_dtype": "iPhone%207%20Plus%20(Global)",
      "ci": "66",
      "wm_did": "6430E434-FE85-40CD-85FE-67B600161ACB",
      "uuid": "xxxx",
      "wm_latitude": "41802940",
      "wm_channel": "2000",
      "wm_visitid": "8FD9E8D3-575E-467F-8E8B-7652BFA852D62017-09-29-17-00613",
      "utm_content": "xxxx",
      "utm_source": "2000",
      "utm_medium": "iphone",
      "partner": "4",
      "wm_logintoken": "7XNJpaQQT4KJywf_ZCP3h9xD-tsAAAAAhwQAAOzi_ZHQzx9DtWLXZcVZzV6BvCZmvfOs5UKB5jCjzJ0gHcHUh56yBDJGzyF6c2AL4g",
      "version": "5.11.0",
      "wm_uuid": "xxxx",
      "platform": "5",
      "request_id": "AF61BF96-B773-450E-BC68-2215FEFB62C6",
      "seq_id": "69726",
      "wm_actual_longitude": "123460416",
      "wm_actual_latitude": "41801396",
      "wm_appversion": "5.11.0",
      "wm_longitude": "123458872",
      "utm_campaign": "AwaimaiBpushGhomepageH211889",
      "app": "4"
    }
]
}`

type (
	Info struct {
		Token               string `json:"token"`
		Userid              string `json:"userid"`
		Address             string `json:"address"`
		F                   string `json:"f"`
		Version             string `json:"version"`
		Wm_latitude         string `json:"wm_latitude"`
		Wm_logintoken       string `json:"wm_logintoken"`
		Wm_mac              string `json:"wm_mac"`
		Request_id          string `json:"request_id"`
		Uuid                string `json:"uuid"`
		Wm_actual_longitude string `json:"wm_actual_longitude"`
		Wm_actual_latitude  string `json:"wm_actual_latitude"`
		Wm_ctype            string `json:"wm_ctype"`
		App                 string `json:"app"`
		Wm_visitid          string `json:"wm_visitid"`
		Wm_did              string `json:"wm_did"`
		Platform            string `json:"platform"`
		Seq_id              string `json:"seq_id"`
		Wm_dversion         string `json:"wm_dversion"`
		Wm_longitude        string `json:"wm_longitude"`
		Wm_channel          string `json:"wm_channel"`
		Wm_uuid             string `json:"wm_uuid"`
		Wm_dtype            string `json:"wm_dtype"`
		Push_token          string `json:"push_token"`
		Partner             string `json:"partner"`
		Wm_appversion       string `json:"wm_appversion"`
		Utm_medium          string `json:"utm_medium"`
		Utm_content         string `json:"utm_content"`
		Utm_term            string `json:"utm_term"`
		Utm_source          string `json:"utm_source"`
		Ci                  string `json:"ci"`
		Utm_campaign        string `json:"utm_campaign"`
		Lat                 string `json:"lat"`
		Lng                 string `json:"lng"`
	}

	DataStruct struct {
		Data []Info `json:"data"`
	}
)

func main() {
	var data DataStruct
	//	err := json.NewDecoder(s).Decode(&data)
	err := json.Unmarshal([]byte(s), &data)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, d := range data.Data {
		elem := reflect.ValueOf(&d).Elem()
		type_ := elem.Type()

		var keys = make([]string, 0)

		map_ := map[string]interface{}{}
		for i := 0; i < type_.NumField(); i++ {
			keys = append(keys, type_.Field(i).Name)

			map_[type_.Field(i).Name] = elem.Field(i).Interface()
		}

		sort.Sort(sort.StringSlice(keys))

		for _, k := range keys {
			fmt.Println(k, "\t\t\t", map_[k])
		}

		fmt.Println("------------------------------------------------\n")
	}

}
