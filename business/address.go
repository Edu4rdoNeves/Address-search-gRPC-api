package business

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Edu4rdoNeves/Address-search-gRPC-api/model"
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/utils"
)

type IAddressBusiness interface {
	SearchAddressByViaCep(cep string) (*model.AddressByViaCep, error)
	SearchAddressByCorreios(cep string) (*model.AddressByCorreios, error)
}

type AddressBusiness struct {
}

func NewAddressBusiness() IAddressBusiness {
	return &AddressBusiness{}
}

func (a *AddressBusiness) SearchAddressByViaCep(cep string) (*model.AddressByViaCep, error) {
	responseData, err := utils.Get(fmt.Sprintf(os.Getenv("VIACEP_BASE_URL"), cep))
	if err != nil {
		return nil, err
	}

	address := &model.AddressByViaCep{}
	err = json.Unmarshal(responseData, address)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. Error: %v", err)
	}

	return address, nil
}

func (a *AddressBusiness) SearchAddressByCorreios(cep string) (*model.AddressByCorreios, error) {
	responseData, err := utils.Get(fmt.Sprintf(os.Getenv("CORREIOS_BASE_URL"), cep))
	if err != nil {
		return nil, err
	}

	address := &model.AddressByCorreios{}
	err = json.Unmarshal(responseData, address)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. Error: %v", err)
	}

	return address, nil
}
