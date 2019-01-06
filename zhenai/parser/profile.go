package parser

import (
	"sky.com/case/crawler-go/engine"
	"regexp"
	"sky.com/case/crawler-go/model"
	"strconv"
)

var (
	// 名字、性别、年龄、身高、体重、收入、婚姻、学历、职业、户口、星座
	AgeRgx      = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
	HeightRgx   = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
	WeightRgx   = regexp.MustCompile(`cm</div><div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
	IncomeRgx   = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
	MarriageRgx = regexp.MustCompile(`<div class="purple-btns"[^>]*><div class="m-btn purple"[^>]*>([^<]+)</div>`)
	HukouRgx    = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)
	XinzuoRgx   = regexp.MustCompile(`岁</div><div class="m-btn purple"[^>]*>(.+)座[^<]+</div>`)
)

func ParseProfile(content []byte, name string) engine.ParserRes {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(content, AgeRgx))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(content, HeightRgx))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(content, WeightRgx))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(content, IncomeRgx)
	profile.Marriage = extractString(content, MarriageRgx)
	profile.Hukou = extractString(content, HukouRgx)
	profile.Xinzuo = extractString(content, XinzuoRgx)

	return engine.ParserRes{Items: []interface{}{profile}}
}

func extractString(content []byte, rgx *regexp.Regexp) string {
	matchs := rgx.FindSubmatch(content)
	if len(matchs) < 2 {
		return ""
	}

	return string(matchs[1])
}
