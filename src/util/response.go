package util

type Response struct {
	body interface{}
	statusCode int
	apiErrorCode int
}

func NewResponse() (r *Response) {
	return &Response{}
}

func (r *Response) WithApiError(body interface{}, statusCode int, apiErrorCode int) *Response {
	r.body = body
	r.statusCode = statusCode
	r.apiErrorCode = apiErrorCode

	return r
}

func (r *Response) WithJson(body interface{}, statusCode int) *Response {
	r.body = body
	r.statusCode = statusCode

	return r
}

func (r *Response) WithStatusCode(statusCode int) *Response {
	r.statusCode = statusCode

	return r
}

func (r *Response) GetBody() interface{} {
	return r.body
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) GetApiErrorCode() int {
	return r.apiErrorCode
}
