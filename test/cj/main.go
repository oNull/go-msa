package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// 可以根据需要添加其他字段
}

func main() {
	var ids string
	for i := 1; i <= 12; i++ {
		url := "https://cjdropshipping.com/zone/995e9a4a-c3c6-4c99-99ce-745682dbb6ec?pageNum=" + strconv.Itoa(i)
		fmt.Println(url)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Cookie", "_tt_enable_cookie=1; _ttp=3EB7Rd_vmO7d2fFMuH41M_uKjJo; _fbp=fb.1.1709709200343.1798877158; _scid=7d3179df-8871-4d4f-99be-7f4d50126cb8; __qca=P0-177973453-1709709201088; hasYindao=1; _ati=281817070140; lng=en; language=en|en; googtrans=/en/en; g_state={\"i_p\":1710913374819,\"i_l\":1}; _ga_SXFZSX9SSV=deleted; cookieyes-consent=consentid:d0JFblQ4VkU2MzBZQWNVdjZ5bGZPWXkySUZBeEV0MHE,consent:yes,action:no,necessary:yes,functional:yes,analytics:yes,performance:yes,advertisement:yes,other:yes; country=AF; INGRESSCOOKIE=1724375888.837.5802.6994|f85a63221c8fd9b049d9a82e34821ae3; csrfToken=squLO-wXFyO5NLAJUpCj7Awr; _gid=GA1.2.11193010.1724375897; _clck=o9x7jw%7C2%7Cfok%7C0%7C1526; currency=currency%253DUSD%253BNAME%253D%25E7%25BE%258E%25E5%2585%2583%253BID%253DUS%253Bsymbol%253D%2524%253Brate%253D1; temp-agent-flag=1; _ScCbts=%5B%5D; _sctr=1%7C1724342400000; _gcl_au=1.1.96130921.1717552187.691698810.1724375972.1724375974; SHOWMESSAGE=S; loginName=Tmljb2xlRFNN; token=Q0o6ZDRhY2Q5MDBhMmQ0NDU3MDk5MWMzOTYxNjk4M2UzMmEx; name=Tmljb2xlIERTTQ==; firstName=Tmljb2xl; lastName=RFNN; email=bmljb2xlQGRyb3BzaGlwbWFuLmNvbQ==; phone=ODYtMTM5NTU0MTcwMjM=; relateSalesman=S3Jpc2U=; salesmanId=OGMxNWU1NDUzNmIxNDg5NzlhMDk4ZDJkMTUxNTQ4OWM=; cjLoginName=NicoleDSM; accessToken=CJ:eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJDSjI1MzYxMzAsZmFsc2UiLCJpbmZvIjp7InVzZXJJZCI6IjE3NDkyNDc5NjM5NDQ5ODQ1NzYiLCJ0b0NqRXJwSWQiOm51bGwsInRvQ2pGbG93ZXJOYW1lIjpudWxsLCJpbnRlcm5hbEVtcGxveWVlSWQiOm51bGwsImludGVybmFsRW1wbG95ZWVOYW1lIjpudWxsLCJzdWJVc2VySWQiOm51bGwsImNvZFN0YXR1cyI6bnVsbCwic3lzdGVtIjoiQ0oiLCJwbGF0Zm9ybSI6Mn0sImlhdCI6MTcyNDM3NTk3NSwiZXhwIjoxNzU1NDc5OTc1fQ.YgrUyify6MIPa3hZmsSaoD7JgFcocxax88YFSeWAEeE; cjLoginToken=CJ:d4acd900a2d44570991c39616983e32a1; vip=1; userId=MTc0OTI0Nzk2Mzk0NDk4NDU3Ng==; avatar=aHR0cHM6Ly9jYy13ZXN0LXVzYS5vc3MtdXMtd2VzdC0xLmFsaXl1bmNzLmNvbS9iYTYyNzIzOC04N2Q3LTQxYTktODZmMS05NjBiNGNjMDdlYmYucG5n; emailVerifyStatus=1; isFirstLogin=0; loginTime=1724375975807; memberLevelTagType=Free; plusExpirationTime=1721231999000; num=CJ2536130; moneyLevelType=4; specifyVisible=1; anonymousToken=fdc28748b4bb0398f3e69de9490d8b722463bad48d550cffb6fcc6bb2cc431d9; _scid_r=7d3179df-8871-4d4f-99be-7f4d50126cb8; _ga_2FHGJ78Y24=GS1.1.1724375897.93.1.1724376164.0.0.0; _ga_S2Q09M5ZGM=GS1.1.1724375897.93.1.1724376164.0.0.0; _rdt_uuid=1709709201336.988c17d8-3148-4b33-a5fd-ac4c60c489d9; _clsk=tzephj%7C1724376165751%7C4%7C1%7Ct.clarity.ms%2Fcollect; _uetsid=999cae5060ed11ef9c932fbe1d59f364; _uetvid=fd43c7d0db8811eeb352391be5356700; _ga=GA1.2.2124388390.1709709201; _gat_gtag_UA_88409817_1=1; _ga_SXFZSX9SSV=GS1.1.1724375896.112.1.1724376243.60.0.1310846664")
		req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		bodyStr := string(body)
		//fmt.Println(bodyStr)
		// 提取window.viewMoreList的内容
		start := strings.Index(bodyStr, "window.viewMoreList=") + len("window.viewMoreList=")
		end := strings.Index(bodyStr[start:], "</script>")
		if end == -1 {
			// 如果没有找到分号，则尝试找到闭合的数组括号
			end = strings.Index(bodyStr[start:], "]")
		}
		if end == -1 {
			fmt.Println("Failed to find the end of the array")
			return
		}
		jsonStr := bodyStr[start : start+end]

		fixedStr := strings.ReplaceAll(jsonStr, "undefined", "null")
		// 解析JSON
		var items []Item
		err = json.Unmarshal([]byte(fixedStr), &items)
		if err != nil {
			fmt.Println("Failed to unmarshal JSON:", err)
			return
		}
		// 输出结果
		// 创建一个字符串来保存id，并用逗号隔开

		num := 0
		for _, item := range items {
			ids += "'" + item.ID + "',"
			num++
		}

		fmt.Println("数量：", num)
		//for _, item := range items {
		//	fmt.Printf("ID: %s, Name: %s\n", item.ID, item.Name)
		//	// ... 可以输出其他字段
		//}

	}
	// 将ids字符串写入到文本文件中
	err := ioutil.WriteFile("ids.txt", []byte(ids), 0644) // 0644是文件权限
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}
	fmt.Println("IDs written to file successfully.")
}
