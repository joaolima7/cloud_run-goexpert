package cep

import (
	"testing"
)

func TestNewCep_Success(t *testing.T) {
	c, err := NewCep("15085240")
	if err != nil {
		t.Errorf("esperava erro nil, mas recebeu: %v", err)
	}
	if c == nil {
		t.Fatal("esperava instância de Cep, mas recebeu nil")
	}
	if c.CEP != "15085240" {
		t.Errorf("esperava CEP '15085240', mas recebeu %s", c.CEP)
	}
}

func TestNewCep_EmptyCep(t *testing.T) {
	_, err := NewCep("")
	if err == nil {
		t.Error("esperava erro para CEP vazio, mas recebeu nil")
	}
}

func TestNewCep_InvalidLength(t *testing.T) {
	_, err := NewCep("1234567")
	if err == nil {
		t.Error("esperava erro para CEP com menos de 8 dígitos, mas recebeu nil")
	}

	_, err = NewCep("123456789")
	if err == nil {
		t.Error("esperava erro para CEP com mais de 8 dígitos, mas recebeu nil")
	}
}
