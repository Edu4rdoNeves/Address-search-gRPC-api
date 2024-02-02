package dto

import (
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/model"
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/pb"
)

func ViaCepAddresToPbAddresModel(viaCep *model.AddressByViaCep) *pb.AddressByViaCepResponse {
	return &pb.AddressByViaCepResponse{
		Cep:         viaCep.Cep,
		Logradouro:  viaCep.Logradouro,
		Complemento: viaCep.Complemento,
		Bairro:      viaCep.Bairro,
		Localidade:  viaCep.Localidade,
		Uf:          viaCep.Uf,
		Ibge:        viaCep.Ibge,
		Gia:         viaCep.Gia,
		Ddd:         viaCep.Ddd,
		Siafi:       viaCep.Siafi,
	}

}

func CorreiosAddresToPbAddresModel(correios *model.AddressByCorreios) *pb.AddressCorreiosResponse {
	return &pb.AddressCorreiosResponse{
		Cep:          correios.Cep,
		State:        correios.State,
		City:         correios.City,
		Neighborhood: correios.Neighborhood,
		Street:       correios.Street,
		Service:      correios.Service,
	}

}
