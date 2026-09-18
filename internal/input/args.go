package input

type ArgsSource struct {
	args []string
}

func NewArgsSource(args []string) *ArgsSource {
	return &ArgsSource{
		args: args,
	}
}

func (s *ArgsSource) Read(fn func(string) error) error {
	for _, arg := range s.args {
		if arg == "" {
			continue
		}

		if err := fn(arg); err != nil {
			return err
		}
	}

	return nil
}
