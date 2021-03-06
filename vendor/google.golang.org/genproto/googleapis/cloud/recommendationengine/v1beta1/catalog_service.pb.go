// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommendationengine/v1beta1/catalog_service.proto

package recommendationengine

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for CreateCatalogItem method.
type CreateCatalogItemRequest struct {
	// Required. The parent catalog resource name, such as
	// "projects/*/locations/global/catalogs/default_catalog".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The catalog item to create.
	CatalogItem          *CatalogItem `protobuf:"bytes,2,opt,name=catalog_item,json=catalogItem,proto3" json:"catalog_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateCatalogItemRequest) Reset()         { *m = CreateCatalogItemRequest{} }
func (m *CreateCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCatalogItemRequest) ProtoMessage()    {}
func (*CreateCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{0}
}

func (m *CreateCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCatalogItemRequest.Unmarshal(m, b)
}
func (m *CreateCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *CreateCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCatalogItemRequest.Merge(m, src)
}
func (m *CreateCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCatalogItemRequest.Size(m)
}
func (m *CreateCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCatalogItemRequest proto.InternalMessageInfo

func (m *CreateCatalogItemRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateCatalogItemRequest) GetCatalogItem() *CatalogItem {
	if m != nil {
		return m.CatalogItem
	}
	return nil
}

// Request message for GetCatalogItem method.
type GetCatalogItemRequest struct {
	// Required. Full resource name of catalog item, such as
	// "projects/*/locations/global/catalogs/default_catalog/catalogitems/some_catalog_item_id".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCatalogItemRequest) Reset()         { *m = GetCatalogItemRequest{} }
func (m *GetCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetCatalogItemRequest) ProtoMessage()    {}
func (*GetCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{1}
}

func (m *GetCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCatalogItemRequest.Unmarshal(m, b)
}
func (m *GetCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *GetCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCatalogItemRequest.Merge(m, src)
}
func (m *GetCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetCatalogItemRequest.Size(m)
}
func (m *GetCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCatalogItemRequest proto.InternalMessageInfo

func (m *GetCatalogItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for ListCatalogItems method.
type ListCatalogItemsRequest struct {
	// Required. The parent catalog resource name, such as
	// "projects/*/locations/global/catalogs/default_catalog".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Maximum number of results to return per page. If zero, the
	// service will choose a reasonable default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The previous ListCatalogItemsResponse.next_page_token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. A filter to apply on the list results.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCatalogItemsRequest) Reset()         { *m = ListCatalogItemsRequest{} }
func (m *ListCatalogItemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCatalogItemsRequest) ProtoMessage()    {}
func (*ListCatalogItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{2}
}

func (m *ListCatalogItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCatalogItemsRequest.Unmarshal(m, b)
}
func (m *ListCatalogItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCatalogItemsRequest.Marshal(b, m, deterministic)
}
func (m *ListCatalogItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCatalogItemsRequest.Merge(m, src)
}
func (m *ListCatalogItemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCatalogItemsRequest.Size(m)
}
func (m *ListCatalogItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCatalogItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCatalogItemsRequest proto.InternalMessageInfo

func (m *ListCatalogItemsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListCatalogItemsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCatalogItemsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCatalogItemsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Response message for ListCatalogItems method.
type ListCatalogItemsResponse struct {
	// The catalog items.
	CatalogItems []*CatalogItem `protobuf:"bytes,1,rep,name=catalog_items,json=catalogItems,proto3" json:"catalog_items,omitempty"`
	// If empty, the list is complete. If nonempty, the token to pass to the next
	// request's ListCatalogItemRequest.page_token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCatalogItemsResponse) Reset()         { *m = ListCatalogItemsResponse{} }
func (m *ListCatalogItemsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCatalogItemsResponse) ProtoMessage()    {}
func (*ListCatalogItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{3}
}

func (m *ListCatalogItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCatalogItemsResponse.Unmarshal(m, b)
}
func (m *ListCatalogItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCatalogItemsResponse.Marshal(b, m, deterministic)
}
func (m *ListCatalogItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCatalogItemsResponse.Merge(m, src)
}
func (m *ListCatalogItemsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCatalogItemsResponse.Size(m)
}
func (m *ListCatalogItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCatalogItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCatalogItemsResponse proto.InternalMessageInfo

func (m *ListCatalogItemsResponse) GetCatalogItems() []*CatalogItem {
	if m != nil {
		return m.CatalogItems
	}
	return nil
}

func (m *ListCatalogItemsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for UpdateCatalogItem method.
type UpdateCatalogItemRequest struct {
	// Required. Full resource name of catalog item, such as
	// "projects/*/locations/global/catalogs/default_catalog/catalogItems/some_catalog_item_id".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The catalog item to update/create. The 'catalog_item_id' field
	// has to match that in the 'name'.
	CatalogItem *CatalogItem `protobuf:"bytes,2,opt,name=catalog_item,json=catalogItem,proto3" json:"catalog_item,omitempty"`
	// Indicates which fields in the provided 'item' to update. If not
	// set, will by default update all fields. Optional.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCatalogItemRequest) Reset()         { *m = UpdateCatalogItemRequest{} }
func (m *UpdateCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCatalogItemRequest) ProtoMessage()    {}
func (*UpdateCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{4}
}

func (m *UpdateCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCatalogItemRequest.Unmarshal(m, b)
}
func (m *UpdateCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCatalogItemRequest.Merge(m, src)
}
func (m *UpdateCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCatalogItemRequest.Size(m)
}
func (m *UpdateCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCatalogItemRequest proto.InternalMessageInfo

func (m *UpdateCatalogItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCatalogItemRequest) GetCatalogItem() *CatalogItem {
	if m != nil {
		return m.CatalogItem
	}
	return nil
}

func (m *UpdateCatalogItemRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request message for DeleteCatalogItem method.
type DeleteCatalogItemRequest struct {
	// Required. Full resource name of catalog item, such as
	// "projects/*/locations/global/catalogs/default_catalog/catalogItems/some_catalog_item_id".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCatalogItemRequest) Reset()         { *m = DeleteCatalogItemRequest{} }
func (m *DeleteCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCatalogItemRequest) ProtoMessage()    {}
func (*DeleteCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{5}
}

func (m *DeleteCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCatalogItemRequest.Unmarshal(m, b)
}
func (m *DeleteCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCatalogItemRequest.Merge(m, src)
}
func (m *DeleteCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCatalogItemRequest.Size(m)
}
func (m *DeleteCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCatalogItemRequest proto.InternalMessageInfo

func (m *DeleteCatalogItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.CreateCatalogItemRequest")
	proto.RegisterType((*GetCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.GetCatalogItemRequest")
	proto.RegisterType((*ListCatalogItemsRequest)(nil), "google.cloud.recommendationengine.v1beta1.ListCatalogItemsRequest")
	proto.RegisterType((*ListCatalogItemsResponse)(nil), "google.cloud.recommendationengine.v1beta1.ListCatalogItemsResponse")
	proto.RegisterType((*UpdateCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.UpdateCatalogItemRequest")
	proto.RegisterType((*DeleteCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.DeleteCatalogItemRequest")
}

func init() {
	proto.RegisterFile("google/cloud/recommendationengine/v1beta1/catalog_service.proto", fileDescriptor_5d69b5a822db872f)
}

var fileDescriptor_5d69b5a822db872f = []byte{
	// 824 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x6b, 0x2b, 0x45,
	0x14, 0x67, 0x72, 0xff, 0xe0, 0x9d, 0xdc, 0x5b, 0xed, 0x80, 0xed, 0x92, 0x5a, 0x08, 0x5b, 0x90,
	0x36, 0xda, 0x1d, 0x1a, 0xa1, 0x42, 0x8b, 0xd8, 0x24, 0xad, 0xb5, 0x50, 0x31, 0xa6, 0x5a, 0x51,
	0x91, 0x30, 0xdd, 0x9c, 0x6c, 0xd7, 0xee, 0xce, 0xac, 0x3b, 0x93, 0x56, 0x2b, 0x7d, 0x11, 0x7d,
	0xf4, 0xc9, 0x0f, 0x20, 0xf8, 0xe8, 0x17, 0xf0, 0x5d, 0xc4, 0x07, 0xc1, 0x17, 0x7d, 0x2b, 0x3e,
	0x2a, 0xf8, 0x31, 0x64, 0x67, 0x76, 0xed, 0xe6, 0x1f, 0x4d, 0x1a, 0xef, 0xdb, 0xce, 0x9c, 0x33,
	0x67, 0x7e, 0xbf, 0x73, 0x7e, 0xe7, 0xec, 0xe0, 0xd7, 0x3d, 0x21, 0xbc, 0x00, 0xa8, 0x1b, 0x88,
	0x5e, 0x87, 0xc6, 0xe0, 0x8a, 0x30, 0x04, 0xde, 0x61, 0xca, 0x17, 0x1c, 0xb8, 0xe7, 0x73, 0xa0,
	0xe7, 0x1b, 0x27, 0xa0, 0xd8, 0x06, 0x75, 0x99, 0x62, 0x81, 0xf0, 0xda, 0x12, 0xe2, 0x73, 0xdf,
	0x05, 0x27, 0x8a, 0x85, 0x12, 0x64, 0xcd, 0x04, 0x70, 0x74, 0x00, 0x67, 0x54, 0x00, 0x27, 0x0d,
	0x50, 0x7a, 0x21, 0xbd, 0x8b, 0x45, 0x3e, 0x65, 0x9c, 0x0b, 0xa5, 0x9d, 0xa4, 0x09, 0x54, 0x5a,
	0xcc, 0x59, 0xdd, 0xc0, 0x07, 0xae, 0x52, 0xc3, 0xab, 0x53, 0x43, 0x4c, 0x0f, 0x6e, 0x4e, 0x7e,
	0xd0, 0x0f, 0x23, 0x11, 0x67, 0x17, 0xae, 0xa4, 0xe7, 0x02, 0xc1, 0xbd, 0xb8, 0xc7, 0xb9, 0xcf,
	0x3d, 0x2a, 0x22, 0x88, 0xfb, 0xe0, 0x2e, 0xa5, 0x4e, 0x7a, 0x75, 0xd2, 0xeb, 0x52, 0x08, 0x23,
	0xf5, 0x79, 0x6a, 0x2c, 0x0f, 0x1a, 0xbb, 0x3e, 0x04, 0x9d, 0x76, 0xc8, 0xe4, 0x99, 0xf1, 0xb0,
	0xbf, 0x41, 0xd8, 0x6a, 0xc4, 0xc0, 0x14, 0x34, 0x0c, 0xe6, 0x03, 0x05, 0x61, 0x0b, 0x3e, 0xed,
	0x81, 0x54, 0x64, 0x01, 0x3f, 0x8c, 0x58, 0x0c, 0x5c, 0x59, 0xa8, 0x8c, 0x56, 0x1f, 0xb5, 0xd2,
	0x15, 0xf9, 0x00, 0x3f, 0xce, 0x8a, 0xe0, 0x2b, 0x08, 0xad, 0x42, 0x19, 0xad, 0x16, 0xab, 0x9b,
	0xce, 0xc4, 0x25, 0x70, 0xf2, 0x97, 0x15, 0xdd, 0x9b, 0x85, 0xfd, 0x12, 0x7e, 0x7e, 0x1f, 0xd4,
	0x08, 0x2c, 0x04, 0xdf, 0xe7, 0x2c, 0x84, 0x14, 0x89, 0xfe, 0xb6, 0xbf, 0x46, 0x78, 0xf1, 0xd0,
	0x97, 0x79, 0x77, 0x79, 0x1b, 0xf6, 0x25, 0xfc, 0x28, 0x62, 0x1e, 0xb4, 0xa5, 0x7f, 0x09, 0x1a,
	0xf8, 0x83, 0xd6, 0x33, 0xc9, 0xc6, 0x91, 0x7f, 0x09, 0x64, 0x19, 0x63, 0x6d, 0x54, 0xe2, 0x0c,
	0xb8, 0x75, 0x4f, 0x1f, 0xd4, 0xee, 0xef, 0x26, 0x1b, 0x49, 0xcc, 0xae, 0x1f, 0x28, 0x88, 0xad,
	0xfb, 0x26, 0xa6, 0x59, 0xd9, 0xdf, 0x21, 0x6c, 0x0d, 0xe3, 0x90, 0x91, 0xe0, 0x12, 0xc8, 0x47,
	0xf8, 0x49, 0x3e, 0x59, 0xd2, 0x42, 0xe5, 0x7b, 0x33, 0x64, 0xeb, 0x71, 0x2e, 0x5b, 0x92, 0xbc,
	0x88, 0x9f, 0xe5, 0xf0, 0x99, 0x6a, 0xe7, 0x50, 0x17, 0x34, 0xb4, 0x27, 0xc9, 0x76, 0x33, 0x43,
	0x6e, 0xff, 0x82, 0xb0, 0xf5, 0x5e, 0xd4, 0x19, 0x5d, 0xe6, 0x11, 0xa9, 0x7d, 0x8a, 0x25, 0x26,
	0xdb, 0xb8, 0xd8, 0xd3, 0x50, 0xb4, 0x0e, 0x75, 0x96, 0x8b, 0xd5, 0x52, 0x16, 0x39, 0x93, 0xaa,
	0xf3, 0x46, 0x22, 0xd5, 0xb7, 0x98, 0x3c, 0x6b, 0x61, 0xe3, 0x9e, 0x7c, 0xdb, 0x0e, 0xb6, 0x76,
	0x21, 0x80, 0x49, 0x79, 0x54, 0x7f, 0x2c, 0xe2, 0xb9, 0xd4, 0xf5, 0xc8, 0xcc, 0x0b, 0xf2, 0x0f,
	0xc2, 0xf3, 0x43, 0x92, 0x27, 0x8d, 0x69, 0xa8, 0x8d, 0x69, 0x98, 0xd2, 0x1d, 0xf3, 0x63, 0x1f,
	0x7f, 0xf9, 0xc7, 0x5f, 0xdf, 0x16, 0x9a, 0xf6, 0xce, 0x7f, 0x83, 0xe0, 0x0b, 0x23, 0xd7, 0xd7,
	0xa2, 0x58, 0x7c, 0x02, 0xae, 0x92, 0xb4, 0x42, 0x03, 0xe1, 0x9a, 0xce, 0xa7, 0x95, 0x6c, 0xc0,
	0x48, 0x5a, 0xb9, 0xa2, 0x79, 0x41, 0x6c, 0xf5, 0x55, 0x8d, 0xfc, 0x89, 0xf0, 0x5c, 0x7f, 0x3b,
	0x91, 0x9d, 0x29, 0x20, 0x8e, 0xec, 0xc4, 0x3b, 0x93, 0x6c, 0x5e, 0xd7, 0x74, 0x4d, 0x34, 0xd7,
	0x06, 0xa9, 0xdd, 0x70, 0x4d, 0x76, 0x6f, 0x67, 0xda, 0x47, 0x94, 0x56, 0x2a, 0x57, 0xe4, 0x6f,
	0x84, 0x9f, 0x1b, 0xec, 0x3b, 0x52, 0x9f, 0x02, 0xde, 0x98, 0xe1, 0x51, 0x6a, 0xcc, 0x14, 0xc3,
	0x34, 0xbe, 0xfd, 0xa6, 0x26, 0x5a, 0x27, 0x33, 0x17, 0x95, 0x7c, 0x55, 0xc0, 0xf3, 0x43, 0xdd,
	0x3b, 0x95, 0x62, 0xc7, 0xf5, 0xfe, 0x9d, 0x8b, 0x19, 0x5f, 0xd7, 0xac, 0xbc, 0xd4, 0x5e, 0xce,
	0xb5, 0xb4, 0xe6, 0xfd, 0x4e, 0x75, 0xf6, 0x02, 0x0f, 0xa8, 0xf9, 0x27, 0x84, 0xe7, 0x87, 0x9a,
	0x7f, 0xaa, 0x34, 0x8c, 0x1b, 0x1d, 0xa5, 0x85, 0xa1, 0xf1, 0xb3, 0x97, 0xfc, 0x46, 0x07, 0x34,
	0x5b, 0xf9, 0x1f, 0x34, 0xfb, 0x1b, 0xc2, 0xe4, 0x40, 0xff, 0xe5, 0xfb, 0x54, 0xbb, 0x3b, 0x05,
	0x8b, 0xe1, 0xe3, 0x19, 0x8d, 0xe5, 0x2c, 0x4a, 0xee, 0xc9, 0xe0, 0xbc, 0x9d, 0x3d, 0x19, 0xec,
	0x96, 0xa6, 0x71, 0x68, 0xef, 0xcf, 0x3c, 0x66, 0xcc, 0x3b, 0x65, 0x0b, 0x55, 0x4a, 0xef, 0xff,
	0x5a, 0x5b, 0x19, 0x89, 0xd6, 0x20, 0x61, 0x91, 0x2f, 0x1d, 0x57, 0x84, 0xbf, 0xd7, 0x9c, 0x53,
	0xa5, 0x22, 0xb9, 0x45, 0xe9, 0xc5, 0xc5, 0xc5, 0x80, 0x91, 0xb2, 0x9e, 0x3a, 0x35, 0xcf, 0xa3,
	0xf5, 0x28, 0x60, 0xaa, 0x2b, 0xe2, 0xb0, 0xfe, 0x33, 0xc2, 0xeb, 0xae, 0x08, 0x27, 0xcf, 0x4b,
	0x13, 0x7d, 0xf8, 0x71, 0xea, 0xec, 0x89, 0x80, 0x71, 0xcf, 0x11, 0xb1, 0x47, 0x3d, 0xe0, 0xba,
	0xa6, 0xf4, 0xe6, 0xca, 0x09, 0x1e, 0x62, 0xdb, 0xa3, 0x8c, 0xdf, 0x17, 0x1e, 0xb4, 0xf6, 0x1a,
	0xb5, 0x83, 0x1f, 0x0a, 0x6b, 0xfb, 0xe6, 0x9e, 0x86, 0x06, 0xd5, 0xea, 0xf3, 0xdd, 0x33, 0xa0,
	0x8e, 0x37, 0xea, 0x49, 0xa0, 0x93, 0x87, 0xfa, 0xf6, 0x57, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x18, 0x1e, 0x90, 0x31, 0xe1, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	// Creates a catalog item.
	CreateCatalogItem(ctx context.Context, in *CreateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error)
	// Gets a specific catalog item.
	GetCatalogItem(ctx context.Context, in *GetCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error)
	// Gets a list of catalog items.
	ListCatalogItems(ctx context.Context, in *ListCatalogItemsRequest, opts ...grpc.CallOption) (*ListCatalogItemsResponse, error)
	// Updates a catalog item. Partial updating is supported. Non-existing
	// items will be created.
	UpdateCatalogItem(ctx context.Context, in *UpdateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error)
	// Deletes a catalog item.
	DeleteCatalogItem(ctx context.Context, in *DeleteCatalogItemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Bulk import of multiple catalog items. Request processing may be
	// synchronous. No partial updating supported. Non-existing items will be
	// created.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully updated.
	ImportCatalogItems(ctx context.Context, in *ImportCatalogItemsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) CreateCatalogItem(ctx context.Context, in *CreateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error) {
	out := new(CatalogItem)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/CreateCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCatalogItem(ctx context.Context, in *GetCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error) {
	out := new(CatalogItem)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/GetCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListCatalogItems(ctx context.Context, in *ListCatalogItemsRequest, opts ...grpc.CallOption) (*ListCatalogItemsResponse, error) {
	out := new(ListCatalogItemsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/ListCatalogItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateCatalogItem(ctx context.Context, in *UpdateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error) {
	out := new(CatalogItem)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/UpdateCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteCatalogItem(ctx context.Context, in *DeleteCatalogItemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/DeleteCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ImportCatalogItems(ctx context.Context, in *ImportCatalogItemsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/ImportCatalogItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	// Creates a catalog item.
	CreateCatalogItem(context.Context, *CreateCatalogItemRequest) (*CatalogItem, error)
	// Gets a specific catalog item.
	GetCatalogItem(context.Context, *GetCatalogItemRequest) (*CatalogItem, error)
	// Gets a list of catalog items.
	ListCatalogItems(context.Context, *ListCatalogItemsRequest) (*ListCatalogItemsResponse, error)
	// Updates a catalog item. Partial updating is supported. Non-existing
	// items will be created.
	UpdateCatalogItem(context.Context, *UpdateCatalogItemRequest) (*CatalogItem, error)
	// Deletes a catalog item.
	DeleteCatalogItem(context.Context, *DeleteCatalogItemRequest) (*empty.Empty, error)
	// Bulk import of multiple catalog items. Request processing may be
	// synchronous. No partial updating supported. Non-existing items will be
	// created.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully updated.
	ImportCatalogItems(context.Context, *ImportCatalogItemsRequest) (*longrunning.Operation, error)
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) CreateCatalogItem(ctx context.Context, req *CreateCatalogItemRequest) (*CatalogItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) GetCatalogItem(ctx context.Context, req *GetCatalogItemRequest) (*CatalogItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) ListCatalogItems(ctx context.Context, req *ListCatalogItemsRequest) (*ListCatalogItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalogItems not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateCatalogItem(ctx context.Context, req *UpdateCatalogItemRequest) (*CatalogItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteCatalogItem(ctx context.Context, req *DeleteCatalogItemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) ImportCatalogItems(ctx context.Context, req *ImportCatalogItemsRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCatalogItems not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_CreateCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/CreateCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCatalogItem(ctx, req.(*CreateCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/GetCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCatalogItem(ctx, req.(*GetCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListCatalogItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListCatalogItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/ListCatalogItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListCatalogItems(ctx, req.(*ListCatalogItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/UpdateCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateCatalogItem(ctx, req.(*UpdateCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/DeleteCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteCatalogItem(ctx, req.(*DeleteCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ImportCatalogItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportCatalogItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ImportCatalogItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/ImportCatalogItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ImportCatalogItems(ctx, req.(*ImportCatalogItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recommendationengine.v1beta1.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCatalogItem",
			Handler:    _CatalogService_CreateCatalogItem_Handler,
		},
		{
			MethodName: "GetCatalogItem",
			Handler:    _CatalogService_GetCatalogItem_Handler,
		},
		{
			MethodName: "ListCatalogItems",
			Handler:    _CatalogService_ListCatalogItems_Handler,
		},
		{
			MethodName: "UpdateCatalogItem",
			Handler:    _CatalogService_UpdateCatalogItem_Handler,
		},
		{
			MethodName: "DeleteCatalogItem",
			Handler:    _CatalogService_DeleteCatalogItem_Handler,
		},
		{
			MethodName: "ImportCatalogItems",
			Handler:    _CatalogService_ImportCatalogItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recommendationengine/v1beta1/catalog_service.proto",
}
