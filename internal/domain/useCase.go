package domain

type ResponderUseCase interface {
	SendMessage(message *ServiceMessage) (*ServiceMessage, error)
}
