package domain

type ServiceMessage struct {
	Response    string
	RequestText string
	MaxToken    int
	Model       string
	ChatId      int
}

//type Meta struct {
//	MaxToken int
//}
