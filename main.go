package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// Парсинг аргументов командной строки и задание количества узлов по умолчанию
	var nodesCount int
	flag.IntVar(&nodesCount, "nodes", 3, "Количество узлов в кольцевой сети")
	flag.Parse()

	if nodesCount < 2 {
		log.Fatal("Количество узлов должно быть не менее 2")
	}

	fmt.Printf("Создание кольцевой сети из %d узлов...\n", nodesCount)

	// Создание сети
	nodes := createRing(nodesCount)

	// Запуск
	startNodes(nodes)

	// Отправка первого сообщения для начала работы
	sendInitialToken(nodes)

	// Время жизни программы (30 секунд)
	fmt.Printf("\nПрограмма запущена\n")
	time.Sleep(30 * time.Second)

	// Статистика после завершения
	fmt.Printf("\nСтатистика:\n")
	totalMessages := 0
	for _, node := range nodes {
		fmt.Printf("Узел %d: обработал %d сообщений\n", node.ID, node.MessageCount)
		totalMessages += node.MessageCount
	}
	fmt.Printf("Всего обработано сообщений: %d\n", totalMessages)
}
