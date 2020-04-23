package response

import "resume/models/dao"

type Resume struct {
	User  dao.User    `json:"user"`
	Skill []dao.Skill `json:"skill"`
	Event []dao.Event `json:"event"`
}
