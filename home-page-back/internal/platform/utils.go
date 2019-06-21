package platform

func SetupAll(setupFuncs ...func() error) error {
	for _, s := range setupFuncs {
		if err := s(); err != nil {
			return err
		}
	}
	return nil
}
