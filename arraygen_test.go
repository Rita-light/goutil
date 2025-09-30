package goutil

import "testing"

func TestArrayGenerate(t *testing.T) {
	size := 10
	arr := ArrayGenerate(size)

	if len(arr) != size {
		t.Errorf("ArrayGenerate(%d) a retourné %d éléments", size, len(arr))
	}

	for _, val := range arr {
		if val < 1 || val > 10000 {
			t.Errorf("Valeur %d hors de l'intervalle [1, 10000]", val)
		}
	}
}
