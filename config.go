package truequeslib

type Config struct {
	Endpoints Endpoints `yaml:"endpoints"`
}

type Endpoints struct {
	Adverts string `yaml:"adverts"`
}
