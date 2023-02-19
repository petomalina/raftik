package raftik

type Server struct {
}

func NewServer(config *Config) (*Server, error) {
	return &Server{}, nil
}
