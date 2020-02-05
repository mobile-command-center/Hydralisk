package goods

type Validator interface {
	Validate() error
}

func Validate(v Validator) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}
