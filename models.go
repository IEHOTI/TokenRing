package main

// Token = структура сообщения в кольцевой сети ~ сообщение
type Token struct {
	Data     string
	Receiver []byte // SHA3-256 хэш идентификатора получателя
	TTL      int    // Время жизни
	Sender   int    // Отправитель для логирования
}

// Node = узел в кольцевой сети
type Node struct {
	ID           int
	Incoming     <-chan Token
	Outgoing     chan<- Token
	NodesCount   int
	MessageCount int
}
