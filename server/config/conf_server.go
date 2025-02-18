package config

import "fmt"

type Server struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Env    string `yaml:"env" `
	JwtKey string `yaml:"jwt_key" `
}

func (s *Server) AddHP() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
