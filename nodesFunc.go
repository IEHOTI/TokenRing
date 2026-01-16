package main

import (
	"fmt"
)

func newNode(id int, incoming <-chan Token, outgoing chan<- Token, nodesCount int) *Node {
	return &Node{
		ID:         id,
		Incoming:   incoming,
		Outgoing:   outgoing,
		NodesCount: nodesCount,
	}
}

// Запуск узла для обработки соббщений
func (n *Node) runNode() {
	for token := range n.Incoming {
		n.MessageCount++

		// Проверка времени жизни
		if token.TTL <= 0 {
			fmt.Printf("Узел %d: TTL истек для сообщения от узла %d\n", n.ID, token.Sender)
			continue
		}

		// Проверка, является ли узел получателем
		if n.isReceiver(token.Receiver) {
			fmt.Printf("[TTL=%d] Узел %d: Получено сообщение от узла %d: '%s'\n",
				token.TTL, n.ID, token.Sender, token.Data)

			// Генерация нового сообщения для продолжения работы
			n.generateNewToken()
			continue
		}

		// Отправка сообщения следующему узлу
		token.TTL--
		token.Sender = n.ID
		fmt.Printf("[TTL=%d] Узел %d: Отправка сообщения следующему узлу\n", token.TTL, n.ID)

		n.Outgoing <- token
	}
}

// Создание кольцевой сети узлов
func createRing(nodesCount int) []*Node {
	// Создание каналов связи между узлами
	channels := make([]chan Token, nodesCount)
	for i := range channels {
		channels[i] = make(chan Token, 10)
	}

	nodes := make([]*Node, nodesCount)

	// Связывка узлов в кольцо
	for i := 0; i < nodesCount; i++ {
		incoming := channels[i]
		outgoing := channels[(i+1)%nodesCount]
		nodes[i] = newNode(i, incoming, outgoing, nodesCount)
	}

	return nodes
}

// Запуск всех узлов
func startNodes(nodes []*Node) {
	for _, node := range nodes {
		go node.runNode()
	}
}
