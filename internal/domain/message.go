package domain

type ServiceMessage struct {
	Response string
	Request  string
	Meta
}

type Meta struct {
	MaxToken int
}
