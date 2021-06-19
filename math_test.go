package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(1, 1)

	if total != 30 {
		t.Errorf("Resultado inválido: %d, esperado %d", total, 30)
	}
}
