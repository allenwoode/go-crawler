package model

import "encoding/json"

type Profile struct {
	Name       string // 昵称
	Age        int    // 年龄
	Height     int    // 身高
	Weight     int    // 体重
	Gender     string // 性别
	Income     string // 收入
	Marriage   string // 婚姻
	Education  string // 教育
	Occupation string // 职业
	Hokou      string // 户口
	Xinzuo     string // 星座
	House      string // 购房
	Car        string // 购车
}

// json string to object
func FromJsonObj(obj interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(obj)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
