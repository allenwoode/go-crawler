package parser

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"regexp"
	"strconv"
	"feilin.com/gocourse/goaction/crawler/model"
	"log"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
//<td><span class="label">年龄：</span>28岁</td>
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		log.Printf("convert error %v", err)
	} else {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		log.Printf("convert error %v", err)
	} else {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		log.Printf("convert error %v", err)
	} else {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)


	results := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	//log.Printf("parser profile %v", results.Items)

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//results.Items = append(results.Items, name)
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	return results
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return "0"
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}