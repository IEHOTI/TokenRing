package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Генерация нового случайного сообщения
func (n *Node) generateNewToken() {
	// Ожидание
	time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)

	// Берется любой получатель (кроме самого отправителя)
	receiver := n.ID
	for receiver == n.ID {
		receiver = rand.Intn(n.NodesCount)
	}

	// Вычисление хэша получателя
	receiverHash := calculateHash(receiver)

	// Генерация случайных данных
	messages := []string{
		"Привет на 100 лет",
		"Тестовое сообщение",
		"Данные для обработки",
		"Простой мусор",
		"Вирус PETYA",
		"Троянский конь",
	}
	data := fmt.Sprintf("%s %d", messages[rand.Intn(len(messages))], n.ID)

	token := Token{
		Data:     data,
		Receiver: receiverHash,
		TTL:      n.NodesCount * 2, // Время жизни пропорциональное по отношению к размеру кольца
		Sender:   n.ID,
	}

	fmt.Printf("Узел %d: Отправляю новое сообщение узлу %d: '%s'\n",
		n.ID, receiver, token.Data)

	n.Outgoing <- token
}

// Отправка первого сообщения в кольцо из главного потока
func sendInitialToken(nodes []*Node) {
	if len(nodes) < 2 {
		log.Fatal("Для работы требуется как минимум 2 узла")
	}

	// Случайный получатель
	receiver := 1
	if len(nodes) > 2 {
		receiver = rand.Intn(len(nodes)-1) + 1
	}

	// Вычисление хэша получателя
	receiverHash := calculateHash(receiver)

	token := Token{
		Data:     "Первое сообщение от главного потока",
		Receiver: receiverHash,
		TTL:      len(nodes) * 2,
		Sender:   -1, // т.к. из главного потока
	}

	// Отправка сообщения первому узлу
	nodes[0].Outgoing <- token

	fmt.Printf("\nГлавный поток: Отправляю первое сообщение узлу %d через узел 0\n", receiver)
	fmt.Printf("Всего узлов в кольце: %d\n\n", len(nodes))
}
