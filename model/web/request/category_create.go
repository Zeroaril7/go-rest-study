package request

type CategoryCreateReq struct {
	Name string `validate:"required,min=1"`
}
