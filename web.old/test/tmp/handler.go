package tmp

import ()


type tmp struct {
	Handler handler
}

func NewHandler() *tmp {
	return &tmp{
		Handler: &file{
			num: 10,
		},
	}
}
