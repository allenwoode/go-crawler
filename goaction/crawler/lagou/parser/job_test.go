package parser

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"github.com/satori/go.uuid"
	"encoding/json"
)

type Position struct {
	Content struct{
		PageNo int
		PageSize int
		HrInfoMap map[int]struct{
			UserId int
			Phone string
			PositionName string
			ReceiveEmail string
			RealName string
			Portraint string
			UserLeavel string
			CanTalk bool
		}
		PositionResult struct{
			HiTags string
			HotLabels string
			LocationInfo struct{
				BusinessZone string
				City string
				District string
				IsAllhotBusinessZone bool
				LocationCode string
				QueryByGisCode bool
			}
			QueryAnalysisInfo interface{}
			ResultSize int
			TotalCount int
			StrategyProperty interface{}
			Result []struct{
				//AdWord int
				//AppShow int
				//Approve int
				BusinessZones []string
				City string
				CompanyFullName string
				//CompanyId int
				CompanyLabelList []string
				//CompanyLogo string
				CompanyShortName string
				CompanySize string
				CreateTime string
				//Deliver int
				District string
				Education string //学历要求
				//Explain string
				FinanceStage string //融资等级
				FirstType string
				//FormatCreateTime string
				//GradeDescription string
				Hitags []string
				//ImState string
				//IndustryField string
				JobNature string //全职|兼职
				LastLogin int
				Latitude string
				Longitude string
				Linestaion string //路线
				//PcShow int
				//Plus string
				PositionAdvantage string //诱惑
				PositionId int
				PositionLables []string //职位标签
				PositionName string //职位名称
				//PromotionScoreExplain string
				PublisherId int
				ResumeProcessDay int
				ResumeProcessRate int
				Salary string //工资
				//Score int
				SecondType string //职位分类
				StationName string //地铁站名
				SubwayLine string //地铁线路
				WorkYear string //工作年限
			}
		}
	}
	Code int
	Success bool
	ResubmitToken string
	RequestId string
	Msg string
} 

func getUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}

func TestParseJob(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.lagou.com/jobs/positionAjax.json?px=default&city=北京", nil)
	format := "JSESSIONID=%s; _ga=GA1.2.1875835759.1524646095; _gid=GA1.2.1147638266.1524646095; user_trace_token=%s; LGUID=%s; index_location_city=%s; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1524646102; LGSID=%s; TG-TRACK-CODE=search_code; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1524714822; LGRID=%s; SEARCH_ID=%s"
	cookie := fmt.Sprintf(format, getUuid(), getUuid(), getUuid(), "%E5%85%A8%E5%9B%BD", getUuid(), getUuid(), getUuid())

	//fmt.Println(cookie)
	//form := url.Values{}
	//form.Add("pn", "2000")
	//form.Add("first", "true")
	//form.Add("kd", "java")
	//req.PostForm = form

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Host", "www.lagou.com")
	req.Header.Add("Origin", "https://www.lagou.com")
	req.Header.Add("Referer", "https://www.lagou.com/jobs/list_?&px=new&city=%E5%8C%97%E4%BA%AC")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.117 Safari/537.36")
	req.Header.Add("X-Anit-Forge-Code", "0")
	req.Header.Add("X-Anit-Forge-Token", "None")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	//req.Form = url.Values{}
	//req.Form.Add("pn", "")
	//req.Form.Add("first", "true")
	//req.Form.Add("kd", "java")

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error: %v", err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request error: %s", http.StatusText(resp.StatusCode))
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(contents))

	job := Position{}
	err = json.Unmarshal(contents, &job)
	if err != nil {
		log.Fatal(err)
	}

	for i, r := range job.Content.PositionResult.Result {
		fmt.Printf("#%3d 职位：%s 描述：%s/%s/%s/%s/%s 公司：%s 融资：%s 公司规模: %s 职位标签：%s 职位诱惑：%s Url: https://www.lagou.com/jobs/%d.html\n",
			i, r.PositionName, r.Salary, r.City, r.WorkYear, r.Education, r.JobNature, r.CompanyShortName, r.FinanceStage, r.CompanySize, r.PositionLables, r.PositionAdvantage, r.PositionId)
	}
	//fmt.Println(job.Content.PositionResult.Result)
}
