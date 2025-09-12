// Package filesend: Holds file send service
package filesend

type Config struct {
	Enabled bool
}

type Service struct {
	config Config
	ready  bool
}

func (s *Service) Enabled() bool {
	return s.config.Enabled
}

func (s *Service) Configure(c Config) {
	s.config = c
}

func CreateService(conf Config) Service {
	return Service{
		ready:  true,
		config: conf,
	}
}
