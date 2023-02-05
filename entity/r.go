package entity

type R struct {
	code    int
	data    any
	message string
}

func InitR(code int, data any, message string) R {
	return R{code: code, data: data, message: message}
}

func (r R) GetCode() int {
	return r.code
}
func (r R) SetCode(code int) {
	r.code = code
}

func (r R) GetData() any {
	return r.data
}
func (r R) SetData(data any) {
	r.data = data
}

func (r R) GetMessage() string {
	return r.message
}
func (r R) SetMessage(message string) {
	r.message = message
}
