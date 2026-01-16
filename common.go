package main

import "crypto/sha3"

// Вычисление SHA3-256 хэша идентификатора узла
func calculateHash(nodeID int) []byte {
	hash := sha3.New256()
	hash.Write([]byte{byte(nodeID)})
	return hash.Sum(nil)
}

// Проверка, что узел является получателем
func (n *Node) isReceiver(receiverHash []byte) bool {
	hash := calculateHash(n.ID)
	for i := 0; i < len(hash); i++ {
		if hash[i] != receiverHash[i] {
			return false
		}
	}
	return true
}
