package request

type CategoryUpdateReq struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1"`
}
