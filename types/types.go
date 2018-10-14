package types


type Server struct {
	Host string
	Port int
}

type Login struct{
	Username string
	Password string
}

type Database struct{
	Server
	Login
	Schema string
}