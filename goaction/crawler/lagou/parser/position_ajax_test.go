package parser

import (
	"testing"
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
)


func TestParseJobJson(t *testing.T)  {
	contents, err := ioutil.ReadFile("positionAjax.json")
	if err != nil {
		panic(err)
	}

	job := Position{}
	err = json.Unmarshal(contents, &job)
	if err != nil {
		log.Fatal(err)
	}

	for i, r := range job.Content.PositionResult.Result {
		fmt.Printf("#%3d 职位：%s 描述：%s/%s/%s/%s/%s 公司：%s 融资：%s 公司规模: %s 职位标签：%s 职位诱惑：%s Url: https://www.lagou.com/jobs/%d.html\n", i, r.PositionName,
			r.Salary, r.City, r.WorkYear, r.Education, r.JobNature, r.CompanyShortName, r.FinanceStage, r.CompanySize, r.PositionLables, r.PositionAdvantage, r.PositionId)
	}
}