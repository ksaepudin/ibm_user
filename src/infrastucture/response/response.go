package response

import "ibm_users_accsess_management/src/entity"

// Format ...
func Format(code int, err error, data ...interface{}) (statusCode int, b *entity.Body) {
	var (
		msg string
		d   interface{}

		pg = entity.Pagination{}
	)

	if err != nil {
		msg = err.Error()
	}

	if len(data) >= 1 {
		d = data[0]
	}

	b = &entity.Body{
		Code:    code,
		Data:    d,
		Message: msg,
	}

	if len(data) > 1 {
		pg.TotalEntries = data[1]
		pg.CurrentPage = data[2]
		pg.Limit = data[3]

		b.Pagination = &pg
	}

	statusCode = code

	return
}
