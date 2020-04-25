package resume

import "resume/models/dao"

type CompanyRes struct {
	Info  dao.Company `json:"info"`
	Event []dao.Event `json:"event"`
}
