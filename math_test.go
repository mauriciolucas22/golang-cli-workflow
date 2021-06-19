package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(20, 10)

	if total != 30 {
		t.Errorf("Resultado inválido: %d, esperado %d", total, 30)
	}
}
