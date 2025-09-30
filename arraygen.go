package goutil

import (
	"math/rand"
	"time"
)

// ArrayGenerate retourne un tableau d'entiers aléatoires entre 1 et 10000.
// La taille du tableau est spécifiée par l'argument n.
func ArrayGenerate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, n)
	for i := range result {
		result[i] = rand.Intn(10000) + 1
	}
	return result
}
