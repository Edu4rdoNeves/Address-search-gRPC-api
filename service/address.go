package service

import (
	"context"
	"net/http"

	"github.com/Edu4rdoNeves/Address-search-gRPC-api/business"
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/dto"
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/pb"
	"google.golang.org/grpc/status"
)

type AddressService struct {
	pb.UnimplementedAddressServicesServer
	business business.IAddressBusiness
}

func NewAddressService(Business business.IAddressBusiness) *AddressService {
	return &AddressService{business: Business}
}

func (a *AddressService) SearchAddress(ctx context.Context, in *pb.AddressRequest) (*pb.SearchAddressResponse, error) {
	cep := in.GetCep()
	viaCepChannel := make(chan *pb.AddressByViaCepResponse)
	correiosChannel := make(chan *pb.AddressCorreiosResponse)

	go func() {
		responseData, err := a.business.SearchAddressByViaCep(cep)
		if err != nil {
			status.Errorf(http.StatusBadRequest, "Search address error by via cep api. Error: %v", err)
			return
		}

		viaCepResponse := dto.ViaCepAddresToPbAddresModel(responseData)

		viaCepChannel <- viaCepResponse
	}()

	go func() {
		responseData, err := a.business.SearchAddressByCorreios(cep)
		if err != nil {
			status.Errorf(http.StatusBadRequest, "Search address error by correios api. Error: %v", err)
			return
		}

		correiosResponse := dto.CorreiosAddresToPbAddresModel(responseData)
		correiosChannel <- correiosResponse
	}()

	select {
	case addressViaCep := <-viaCepChannel:
		return &pb.SearchAddressResponse{
			AddressByViaCepResponse: addressViaCep,
		}, nil

	case addressCorreios := <-correiosChannel:
		return &pb.SearchAddressResponse{
			AddressCorreiosResponse: addressCorreios,
		}, nil

	default:
		return nil, status.Errorf(http.StatusRequestTimeout, "expired waiting time.")
	}
}
