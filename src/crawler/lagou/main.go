package main

import (
	"github.com/ericchiang/css"
	"golang.org/x/net/html"
	"strings"
	"fmt"
	"feilin.com/gocourse/goaction/crawler/fetcher"
)

/*
	拉勾招聘
	- 爬取职位信息
	- 爬取公司信息
*/

const url = "https://www.lagou.com/zhaopin/"
const allCityRequstUrl = "https://www.lagou.com/gongsi/allCity.html?option=0-0-0"
// 根据公司找职位

func main() {

	resp, err := fetcher.Fetch(allCityRequstUrl)
	if err != nil {
		panic(err)
	}


	sel, err := css.Compile(".city_list>li>a")
	if err != nil {
		panic(err)
	}

	node, err := html.Parse(strings.NewReader(string(resp)))
	if err != nil {
		panic(err)
	}

	for i, ele := range sel.Select(node) {
		fmt.Printf("#%3d, data: %s\n", i, ele.Attr)
		//html.Render(os.Stdout, ele)
	}

	//fmt.Println()
}
