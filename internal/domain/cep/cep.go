package cep

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
		return ErrInvalidCep
	}

	if len(c.CEP) != 8 {
		return ErrInvalidCep
	}

	return nil
}
