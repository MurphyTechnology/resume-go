package resume

import "resume/models/dao"

type Resume struct {
	User      dao.User        `json:"user"`
	Skill     []dao.Skill     `json:"skill"`
	Education []dao.Education `json:"education"`
	Company   []CompanyRes    `json:"company"`
	Works     []dao.Works     `json:"works"`
	Honor     []dao.Honor     `json:"honor"`
}
