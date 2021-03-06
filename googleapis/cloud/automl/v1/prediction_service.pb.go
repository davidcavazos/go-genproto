// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/prediction_service.proto

package automl

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
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

// Request message for [PredictionService.Predict][google.cloud.automl.v1.PredictionService.Predict].
type PredictRequest struct {
	// Name of the model requested to serve the prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Payload to perform a prediction on. The payload must match the
	// problem type that the model was trained to solve.
	Payload *ExamplePayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific parameters, any string must be up to 25000
	// characters long.
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60105ec759f54a4, []int{0}
}

func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetPayload() *ExamplePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Response message for [PredictionService.Predict][google.cloud.automl.v1.PredictionService.Predict].
type PredictResponse struct {
	// Prediction result.
	// Translation and Text Sentiment will return precisely one payload.
	Payload []*AnnotationPayload `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific prediction response metadata.
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a60105ec759f54a4, []int{1}
}

func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (m *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(m, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetPayload() []*AnnotationPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.automl.v1.PredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.PredictRequest.ParamsEntry")
	proto.RegisterType((*PredictResponse)(nil), "google.cloud.automl.v1.PredictResponse")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.PredictResponse.MetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/prediction_service.proto", fileDescriptor_a60105ec759f54a4)
}

var fileDescriptor_a60105ec759f54a4 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0x36, 0xda, 0xda, 0x09, 0xfe, 0x2d, 0x1a, 0xe3, 0x22, 0x18, 0x22, 0xb4, 0x31, 0xe0,
	0x0e, 0x89, 0x88, 0xba, 0xb5, 0x60, 0x5a, 0x8a, 0x28, 0x16, 0xd6, 0x08, 0xbd, 0x90, 0x40, 0x18,
	0x37, 0xe3, 0x76, 0x75, 0x76, 0xce, 0xb8, 0x3b, 0x9b, 0x18, 0xc4, 0x1b, 0x5f, 0xc1, 0x2b, 0x1f,
	0xc1, 0x57, 0xe9, 0xa5, 0xbe, 0x82, 0x78, 0x21, 0xf8, 0x02, 0x5e, 0xc9, 0xee, 0x4c, 0x93, 0xa6,
	0x75, 0xb1, 0xde, 0xcd, 0xee, 0x77, 0xbe, 0xef, 0x7c, 0x67, 0xe6, 0x7c, 0x98, 0x84, 0x00, 0x21,
	0x67, 0x24, 0xe0, 0x90, 0x8d, 0x08, 0xcd, 0x14, 0xc4, 0x9c, 0x8c, 0x3b, 0x44, 0x26, 0x6c, 0x14,
	0x05, 0x2a, 0x02, 0x31, 0x4c, 0x59, 0x32, 0x8e, 0x02, 0xe6, 0xca, 0x04, 0x14, 0xd8, 0x35, 0x4d,
	0x70, 0x0b, 0x82, 0xab, 0x09, 0xee, 0xb8, 0xe3, 0x5c, 0x33, 0x42, 0x54, 0x46, 0x84, 0x0a, 0x01,
	0x8a, 0xe6, 0xe4, 0x54, 0xb3, 0x9c, 0xb2, 0x36, 0xf3, 0xca, 0xa1, 0xa4, 0x53, 0x0e, 0x74, 0x64,
	0x08, 0x6b, 0x25, 0x84, 0x11, 0x55, 0x74, 0x18, 0x29, 0x16, 0x1f, 0x28, 0x5f, 0x2f, 0x29, 0x8c,
	0xc0, 0x14, 0xdc, 0x30, 0x05, 0x1c, 0x44, 0x98, 0x64, 0x42, 0x44, 0x22, 0x24, 0x20, 0x59, 0xb2,
	0xe0, 0xef, 0xca, 0x21, 0xf7, 0x01, 0x8f, 0x98, 0x50, 0x1a, 0x68, 0xfe, 0x42, 0xf8, 0x9c, 0xaf,
	0xef, 0xa2, 0xcf, 0xde, 0x66, 0x2c, 0x55, 0xb6, 0x8d, 0x4f, 0x09, 0x1a, 0xb3, 0x3a, 0x6a, 0xa0,
	0xd6, 0x4a, 0xbf, 0x38, 0xdb, 0x0f, 0xf1, 0xb2, 0xf1, 0x5f, 0xb7, 0x1a, 0xa8, 0x55, 0xed, 0xae,
	0xba, 0x7f, 0xbf, 0x27, 0x77, 0xfb, 0x1d, 0x8d, 0x25, 0x67, 0xbe, 0xae, 0xee, 0x1f, 0xd0, 0xec,
	0x27, 0x78, 0x49, 0xd2, 0x84, 0xc6, 0x69, 0xbd, 0xd2, 0xa8, 0xb4, 0xaa, 0xdd, 0x6e, 0x99, 0xc0,
	0xa2, 0x1b, 0xd7, 0x2f, 0x48, 0xdb, 0x42, 0x25, 0xd3, 0xbe, 0x51, 0x70, 0xee, 0xe3, 0xea, 0xa1,
	0xdf, 0xf6, 0x05, 0x5c, 0x79, 0xc3, 0xa6, 0xc6, 0x6f, 0x7e, 0xb4, 0x2f, 0xe1, 0xd3, 0x63, 0xca,
	0x33, 0x56, 0x98, 0x5d, 0xe9, 0xeb, 0x0f, 0xcf, 0xba, 0x87, 0x9a, 0x3f, 0x10, 0x3e, 0x3f, 0xeb,
	0x90, 0x4a, 0x10, 0x29, 0xb3, 0xb7, 0xe6, 0xc3, 0xa1, 0xc2, 0xdb, 0xcd, 0x32, 0x6f, 0xbd, 0xd9,
	0x73, 0x1e, 0x9b, 0xef, 0x19, 0x3e, 0x13, 0x33, 0x45, 0xf3, 0xf7, 0xab, 0x5b, 0x85, 0xca, 0x9d,
	0x7f, 0x4e, 0xa8, 0xfb, 0xbb, 0x3b, 0x86, 0xa7, 0x87, 0x9c, 0xc9, 0x38, 0xeb, 0xf8, 0xec, 0x02,
	0xf4, 0x3f, 0x83, 0x76, 0x7f, 0x23, 0x7c, 0xd1, 0x9f, 0x2d, 0xf9, 0x73, 0xbd, 0xe3, 0xf6, 0x67,
	0x84, 0x97, 0xcd, 0x5f, 0x7b, 0xf5, 0x64, 0x2f, 0xe0, 0xac, 0x9d, 0x70, 0x8e, 0xe6, 0xc6, 0xc7,
	0x6f, 0xdf, 0x3f, 0x59, 0x77, 0x9b, 0xdd, 0x7c, 0x3f, 0xdf, 0xe7, 0x7b, 0xb3, 0x21, 0x13, 0x78,
	0xcd, 0x02, 0x95, 0x92, 0x36, 0xe1, 0x10, 0xe8, 0x95, 0x24, 0x6d, 0x12, 0xc3, 0x88, 0xf1, 0x94,
	0xb4, 0x3f, 0x78, 0x26, 0x87, 0x1e, 0x6a, 0x3b, 0x8f, 0xf7, 0x7b, 0x97, 0x8d, 0xb8, 0x6e, 0x48,
	0x65, 0x94, 0xba, 0x01, 0xc4, 0x5f, 0x7b, 0xee, 0x9e, 0x52, 0x32, 0xf5, 0x08, 0x99, 0x4c, 0x26,
	0x47, 0xc0, 0x3c, 0x12, 0x7b, 0x3a, 0x1d, 0xb7, 0x24, 0xa7, 0xea, 0x15, 0x24, 0xf1, 0xe6, 0x17,
	0x84, 0x9d, 0x00, 0xe2, 0x12, 0xe3, 0x9b, 0xb5, 0x63, 0x17, 0xe3, 0xe7, 0x61, 0xf0, 0xd1, 0x8b,
	0x07, 0x86, 0x11, 0x02, 0xa7, 0x22, 0x74, 0x21, 0x09, 0x49, 0xc8, 0x44, 0x11, 0x15, 0x32, 0xef,
	0x7b, 0x34, 0x8c, 0xeb, 0xfa, 0xb4, 0x6f, 0xd5, 0x1e, 0x15, 0x35, 0x83, 0xad, 0x1c, 0x1f, 0xf4,
	0x32, 0x05, 0x3b, 0x7c, 0xb0, 0xdb, 0xf9, 0x69, 0x5d, 0xd5, 0x80, 0xe7, 0x15, 0x88, 0xe7, 0x15,
	0xd0, 0x53, 0xcf, 0xdb, 0xed, 0xbc, 0x5c, 0x2a, 0xd4, 0x6f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x10, 0xe1, 0x94, 0x89, 0xaa, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.automl.v1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/automl/v1/prediction_service.proto",
}
