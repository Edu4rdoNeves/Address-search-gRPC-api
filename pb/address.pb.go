// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: proto/address.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep string `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
}

func (x *AddressRequest) Reset() {
	*x = AddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequest) ProtoMessage() {}

func (x *AddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequest.ProtoReflect.Descriptor instead.
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_address_proto_rawDescGZIP(), []int{0}
}

func (x *AddressRequest) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

type SearchAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressCorreiosResponse *AddressCorreiosResponse `protobuf:"bytes,1,opt,name=addressCorreiosResponse,proto3" json:"addressCorreiosResponse,omitempty"`
	AddressByViaCepResponse *AddressByViaCepResponse `protobuf:"bytes,2,opt,name=addressByViaCepResponse,proto3" json:"addressByViaCepResponse,omitempty"`
}

func (x *SearchAddressResponse) Reset() {
	*x = SearchAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAddressResponse) ProtoMessage() {}

func (x *SearchAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAddressResponse.ProtoReflect.Descriptor instead.
func (*SearchAddressResponse) Descriptor() ([]byte, []int) {
	return file_proto_address_proto_rawDescGZIP(), []int{1}
}

func (x *SearchAddressResponse) GetAddressCorreiosResponse() *AddressCorreiosResponse {
	if x != nil {
		return x.AddressCorreiosResponse
	}
	return nil
}

func (x *SearchAddressResponse) GetAddressByViaCepResponse() *AddressByViaCepResponse {
	if x != nil {
		return x.AddressByViaCepResponse
	}
	return nil
}

type AddressCorreiosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep          string `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
	State        string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	City         string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Neighborhood string `protobuf:"bytes,4,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
	Street       string `protobuf:"bytes,5,opt,name=street,proto3" json:"street,omitempty"`
	Service      string `protobuf:"bytes,6,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *AddressCorreiosResponse) Reset() {
	*x = AddressCorreiosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressCorreiosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressCorreiosResponse) ProtoMessage() {}

func (x *AddressCorreiosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressCorreiosResponse.ProtoReflect.Descriptor instead.
func (*AddressCorreiosResponse) Descriptor() ([]byte, []int) {
	return file_proto_address_proto_rawDescGZIP(), []int{2}
}

func (x *AddressCorreiosResponse) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *AddressCorreiosResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AddressCorreiosResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressCorreiosResponse) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *AddressCorreiosResponse) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *AddressCorreiosResponse) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type AddressByViaCepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep         string `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
	Logradouro  string `protobuf:"bytes,2,opt,name=logradouro,proto3" json:"logradouro,omitempty"`
	Complemento string `protobuf:"bytes,3,opt,name=complemento,proto3" json:"complemento,omitempty"`
	Bairro      string `protobuf:"bytes,4,opt,name=bairro,proto3" json:"bairro,omitempty"`
	Localidade  string `protobuf:"bytes,5,opt,name=localidade,proto3" json:"localidade,omitempty"`
	Uf          string `protobuf:"bytes,6,opt,name=uf,proto3" json:"uf,omitempty"`
	Ibge        string `protobuf:"bytes,7,opt,name=ibge,proto3" json:"ibge,omitempty"`
	Gia         string `protobuf:"bytes,8,opt,name=gia,proto3" json:"gia,omitempty"`
	Ddd         string `protobuf:"bytes,9,opt,name=ddd,proto3" json:"ddd,omitempty"`
	Siafi       string `protobuf:"bytes,10,opt,name=siafi,proto3" json:"siafi,omitempty"`
}

func (x *AddressByViaCepResponse) Reset() {
	*x = AddressByViaCepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressByViaCepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressByViaCepResponse) ProtoMessage() {}

func (x *AddressByViaCepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressByViaCepResponse.ProtoReflect.Descriptor instead.
func (*AddressByViaCepResponse) Descriptor() ([]byte, []int) {
	return file_proto_address_proto_rawDescGZIP(), []int{3}
}

func (x *AddressByViaCepResponse) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *AddressByViaCepResponse) GetLogradouro() string {
	if x != nil {
		return x.Logradouro
	}
	return ""
}

func (x *AddressByViaCepResponse) GetComplemento() string {
	if x != nil {
		return x.Complemento
	}
	return ""
}

func (x *AddressByViaCepResponse) GetBairro() string {
	if x != nil {
		return x.Bairro
	}
	return ""
}

func (x *AddressByViaCepResponse) GetLocalidade() string {
	if x != nil {
		return x.Localidade
	}
	return ""
}

func (x *AddressByViaCepResponse) GetUf() string {
	if x != nil {
		return x.Uf
	}
	return ""
}

func (x *AddressByViaCepResponse) GetIbge() string {
	if x != nil {
		return x.Ibge
	}
	return ""
}

func (x *AddressByViaCepResponse) GetGia() string {
	if x != nil {
		return x.Gia
	}
	return ""
}

func (x *AddressByViaCepResponse) GetDdd() string {
	if x != nil {
		return x.Ddd
	}
	return ""
}

func (x *AddressByViaCepResponse) GetSiafi() string {
	if x != nil {
		return x.Siafi
	}
	return ""
}

var File_proto_address_proto protoreflect.FileDescriptor

var file_proto_address_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x22, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x65, 0x70, 0x22, 0xc5, 0x01,
	0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x17, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x17, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x17, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x56, 0x69, 0x61, 0x43, 0x65,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x56, 0x69,
	0x61, 0x43, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x17, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x56, 0x69, 0x61, 0x43, 0x65, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x79, 0x56, 0x69, 0x61, 0x43, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x65,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x72, 0x61, 0x64, 0x6f, 0x75, 0x72, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x72, 0x61, 0x64, 0x6f, 0x75, 0x72,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x69, 0x72, 0x72, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x69, 0x72, 0x72, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x75,
	0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x62, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x69, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x64, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x64, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x61, 0x66, 0x69, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x69, 0x61, 0x66, 0x69, 0x32, 0x53, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_address_proto_rawDescOnce sync.Once
	file_proto_address_proto_rawDescData = file_proto_address_proto_rawDesc
)

func file_proto_address_proto_rawDescGZIP() []byte {
	file_proto_address_proto_rawDescOnce.Do(func() {
		file_proto_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_address_proto_rawDescData)
	})
	return file_proto_address_proto_rawDescData
}

var file_proto_address_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_address_proto_goTypes = []interface{}{
	(*AddressRequest)(nil),          // 0: pb.AddressRequest
	(*SearchAddressResponse)(nil),   // 1: pb.SearchAddressResponse
	(*AddressCorreiosResponse)(nil), // 2: pb.AddressCorreiosResponse
	(*AddressByViaCepResponse)(nil), // 3: pb.AddressByViaCepResponse
}
var file_proto_address_proto_depIdxs = []int32{
	2, // 0: pb.SearchAddressResponse.addressCorreiosResponse:type_name -> pb.AddressCorreiosResponse
	3, // 1: pb.SearchAddressResponse.addressByViaCepResponse:type_name -> pb.AddressByViaCepResponse
	0, // 2: pb.AddressServices.SearchAddress:input_type -> pb.AddressRequest
	1, // 3: pb.AddressServices.SearchAddress:output_type -> pb.SearchAddressResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_address_proto_init() }
func file_proto_address_proto_init() {
	if File_proto_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchAddressResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressCorreiosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressByViaCepResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_address_proto_goTypes,
		DependencyIndexes: file_proto_address_proto_depIdxs,
		MessageInfos:      file_proto_address_proto_msgTypes,
	}.Build()
	File_proto_address_proto = out.File
	file_proto_address_proto_rawDesc = nil
	file_proto_address_proto_goTypes = nil
	file_proto_address_proto_depIdxs = nil
}
