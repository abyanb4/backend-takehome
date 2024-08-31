package comment

import "github.com/go-playground/validator/v10"

func ValidateCreatePostRequest(req *CreateCommentRequest) error {
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
