package cep

import "errors"

type Cep struct {
	CEP string
}

func NewCep(cep string) (*Cep, error) {
	newCep := &Cep{
		CEP: cep,
	}

	err := newCep.ValidateCep()
	if err != nil {
		return nil, err
	}

	return newCep, nil
}

func (c *Cep) ValidateCep() error {
	if c.CEP == "" {
		return errors.New("CEP não pode ser vazio")
	}

	if len(c.CEP) != 8 {
		return errors.New("CEP deve conter 8 dígitos")
	}

	return nil
}
