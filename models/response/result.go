package response

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Result) Fail() *Result {
	r.Code = 400
	r.Message = "服务内部异常"
	return r
}

func (r *Result) Success(data interface{}) *Result {
	r.Message = "success"
	r.Code = 200
	r.Data = data
	return r
}
