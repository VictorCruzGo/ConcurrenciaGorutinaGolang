package main

import (
	"testing"
)

func TestMiSuma(t *testing.T) {
	//----------tablas de pruebas (table test)
	type prueba struct {
		datos     []int
		respuesta int
	}

	pruebas := []prueba{
		prueba{
			datos:     []int{2, 4, 6},
			respuesta: 12,
		},
		prueba{
			datos:     []int{1, 5, 2},
			respuesta: 8,
		},
		prueba{
			datos:     []int{0, -1, 1},
			respuesta: 0,
		},
		prueba{
			datos:     []int{-10, 4, 20},
			respuesta: 14,
		},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}
}
