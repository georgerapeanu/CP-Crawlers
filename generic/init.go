package generic

func Init() error {
	var err error
	err = httpClientInit()

	if err != nil {
		return err
	}

	return nil
}
