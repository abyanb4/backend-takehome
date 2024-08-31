package post

import "github.com/go-playground/validator/v10"

func ValidateCreatePostRequest(req *CreatePostRequest) error {
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdatePostRequest(req *UpdatePostRequest) error {
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
