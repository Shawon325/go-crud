package requests

type UserRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email,max=255"`
}

func (r *UserRequest) Validate() map[string]string {
	if err := validate.Struct(r); err != nil {
		return formatValidationErrors(err)
	}

	return nil
}
