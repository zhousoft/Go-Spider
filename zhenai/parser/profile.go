package parser

import (
	"Go-Spider/engine"
	"Go-Spider/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

//ParseProfile 解析用户信息
func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	profile.Age = extractInt(contents, ageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Weight = extractInt(contents, weightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.HuKou = extractString(contents, hukouRe)
	profile.XingZuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	intStr := extractString(contents, re)
	intV, err := strconv.Atoi(intStr)
	if err != nil {
		return 0
	}
	return intV
}
