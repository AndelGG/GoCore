package domain

type JsonParser interface {
	Parse() (ServiceMessage, error)
}
