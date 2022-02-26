// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/proto/configuration/cloud/google/google.proto

package google

import (
	http "github.com/buildbarn/bb-storage/pkg/proto/configuration/http"
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

type GoogleClientConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Credentials:
	//	*GoogleClientConfiguration_CredentialsJson
	Credentials isGoogleClientConfiguration_Credentials `protobuf_oneof:"credentials"`
	HttpClient  *http.ClientConfiguration               `protobuf:"bytes,2,opt,name=http_client,json=httpClient,proto3" json:"http_client,omitempty"`
	Scopes      []string                                `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *GoogleClientConfiguration) Reset() {
	*x = GoogleClientConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cloud_google_google_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleClientConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleClientConfiguration) ProtoMessage() {}

func (x *GoogleClientConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cloud_google_google_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleClientConfiguration.ProtoReflect.Descriptor instead.
func (*GoogleClientConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cloud_google_google_proto_rawDescGZIP(), []int{0}
}

func (m *GoogleClientConfiguration) GetCredentials() isGoogleClientConfiguration_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (x *GoogleClientConfiguration) GetCredentialsJson() string {
	if x, ok := x.GetCredentials().(*GoogleClientConfiguration_CredentialsJson); ok {
		return x.CredentialsJson
	}
	return ""
}

func (x *GoogleClientConfiguration) GetHttpClient() *http.ClientConfiguration {
	if x != nil {
		return x.HttpClient
	}
	return nil
}

func (x *GoogleClientConfiguration) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type isGoogleClientConfiguration_Credentials interface {
	isGoogleClientConfiguration_Credentials()
}

type GoogleClientConfiguration_CredentialsJson struct {
	CredentialsJson string `protobuf:"bytes,1,opt,name=credentials_json,json=credentialsJson,proto3,oneof"`
}

func (*GoogleClientConfiguration_CredentialsJson) isGoogleClientConfiguration_Credentials() {}

var File_pkg_proto_configuration_cloud_google_google_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_cloud_google_google_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x52, 0x0a,
	0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e,
	0x2f, 0x62, 0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_cloud_google_google_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_cloud_google_google_proto_rawDescData = file_pkg_proto_configuration_cloud_google_google_proto_rawDesc
)

func file_pkg_proto_configuration_cloud_google_google_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_cloud_google_google_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_cloud_google_google_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_cloud_google_google_proto_rawDescData)
	})
	return file_pkg_proto_configuration_cloud_google_google_proto_rawDescData
}

var file_pkg_proto_configuration_cloud_google_google_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_configuration_cloud_google_google_proto_goTypes = []interface{}{
	(*GoogleClientConfiguration)(nil), // 0: buildbarn.configuration.cloud.google.GoogleClientConfiguration
	(*http.ClientConfiguration)(nil),  // 1: buildbarn.configuration.http.ClientConfiguration
}
var file_pkg_proto_configuration_cloud_google_google_proto_depIdxs = []int32{
	1, // 0: buildbarn.configuration.cloud.google.GoogleClientConfiguration.http_client:type_name -> buildbarn.configuration.http.ClientConfiguration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_cloud_google_google_proto_init() }
func file_pkg_proto_configuration_cloud_google_google_proto_init() {
	if File_pkg_proto_configuration_cloud_google_google_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_cloud_google_google_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleClientConfiguration); i {
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
	file_pkg_proto_configuration_cloud_google_google_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GoogleClientConfiguration_CredentialsJson)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_cloud_google_google_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_cloud_google_google_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_cloud_google_google_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_cloud_google_google_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_cloud_google_google_proto = out.File
	file_pkg_proto_configuration_cloud_google_google_proto_rawDesc = nil
	file_pkg_proto_configuration_cloud_google_google_proto_goTypes = nil
	file_pkg_proto_configuration_cloud_google_google_proto_depIdxs = nil
}
