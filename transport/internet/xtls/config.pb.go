// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: transport/internet/xtls/config.proto

package xtls

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Certificate_Usage int32

const (
	Certificate_ENCIPHERMENT     Certificate_Usage = 0
	Certificate_AUTHORITY_VERIFY Certificate_Usage = 1
	Certificate_AUTHORITY_ISSUE  Certificate_Usage = 2
)

// Enum value maps for Certificate_Usage.
var (
	Certificate_Usage_name = map[int32]string{
		0: "ENCIPHERMENT",
		1: "AUTHORITY_VERIFY",
		2: "AUTHORITY_ISSUE",
	}
	Certificate_Usage_value = map[string]int32{
		"ENCIPHERMENT":     0,
		"AUTHORITY_VERIFY": 1,
		"AUTHORITY_ISSUE":  2,
	}
)

func (x Certificate_Usage) Enum() *Certificate_Usage {
	p := new(Certificate_Usage)
	*p = x
	return p
}

func (x Certificate_Usage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Certificate_Usage) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_xtls_config_proto_enumTypes[0].Descriptor()
}

func (Certificate_Usage) Type() protoreflect.EnumType {
	return &file_transport_internet_xtls_config_proto_enumTypes[0]
}

func (x Certificate_Usage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Certificate_Usage.Descriptor instead.
func (Certificate_Usage) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_xtls_config_proto_rawDescGZIP(), []int{0, 0}
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TLS certificate in x509 format.
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// TLS key in x509 format.
	Key          []byte            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Usage        Certificate_Usage `protobuf:"varint,3,opt,name=usage,proto3,enum=xray.transport.internet.xtls.Certificate_Usage" json:"usage,omitempty"`
	OcspStapling int64             `protobuf:"varint,4,opt,name=ocsp_stapling,json=ocspStapling,proto3" json:"ocsp_stapling,omitempty"`
	// TLS certificate path
	CertificatePath string `protobuf:"bytes,5,opt,name=certificate_path,json=certificatePath,proto3" json:"certificate_path,omitempty"`
	// TLS Key path
	KeyPath string `protobuf:"bytes,6,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_xtls_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_xtls_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_transport_internet_xtls_config_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Certificate) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Certificate) GetUsage() Certificate_Usage {
	if x != nil {
		return x.Usage
	}
	return Certificate_ENCIPHERMENT
}

func (x *Certificate) GetOcspStapling() int64 {
	if x != nil {
		return x.OcspStapling
	}
	return 0
}

func (x *Certificate) GetCertificatePath() string {
	if x != nil {
		return x.CertificatePath
	}
	return ""
}

func (x *Certificate) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not to allow self-signed certificates.
	AllowInsecure bool `protobuf:"varint,1,opt,name=allow_insecure,json=allowInsecure,proto3" json:"allow_insecure,omitempty"`
	// List of certificates to be served on server.
	Certificate []*Certificate `protobuf:"bytes,2,rep,name=certificate,proto3" json:"certificate,omitempty"`
	// Override server name.
	ServerName string `protobuf:"bytes,3,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// Lists of string as ALPN values.
	NextProtocol []string `protobuf:"bytes,4,rep,name=next_protocol,json=nextProtocol,proto3" json:"next_protocol,omitempty"`
	// Whether or not to enable session (ticket) resumption.
	EnableSessionResumption bool `protobuf:"varint,5,opt,name=enable_session_resumption,json=enableSessionResumption,proto3" json:"enable_session_resumption,omitempty"`
	// If true, root certificates on the system will not be loaded for
	// verification.
	DisableSystemRoot bool `protobuf:"varint,6,opt,name=disable_system_root,json=disableSystemRoot,proto3" json:"disable_system_root,omitempty"`
	// The minimum TLS version.
	MinVersion string `protobuf:"bytes,7,opt,name=min_version,json=minVersion,proto3" json:"min_version,omitempty"`
	// The maximum TLS version.
	MaxVersion string `protobuf:"bytes,8,opt,name=max_version,json=maxVersion,proto3" json:"max_version,omitempty"`
	// Specify cipher suites, except for TLS 1.3.
	CipherSuites string `protobuf:"bytes,9,opt,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
	// Whether the server selects its most preferred ciphersuite.
	PreferServerCipherSuites bool `protobuf:"varint,10,opt,name=prefer_server_cipher_suites,json=preferServerCipherSuites,proto3" json:"prefer_server_cipher_suites,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_xtls_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_xtls_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_internet_xtls_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetAllowInsecure() bool {
	if x != nil {
		return x.AllowInsecure
	}
	return false
}

func (x *Config) GetCertificate() []*Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Config) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *Config) GetNextProtocol() []string {
	if x != nil {
		return x.NextProtocol
	}
	return nil
}

func (x *Config) GetEnableSessionResumption() bool {
	if x != nil {
		return x.EnableSessionResumption
	}
	return false
}

func (x *Config) GetDisableSystemRoot() bool {
	if x != nil {
		return x.DisableSystemRoot
	}
	return false
}

func (x *Config) GetMinVersion() string {
	if x != nil {
		return x.MinVersion
	}
	return ""
}

func (x *Config) GetMaxVersion() string {
	if x != nil {
		return x.MaxVersion
	}
	return ""
}

func (x *Config) GetCipherSuites() string {
	if x != nil {
		return x.CipherSuites
	}
	return ""
}

func (x *Config) GetPreferServerCipherSuites() bool {
	if x != nil {
		return x.PreferServerCipherSuites
	}
	return false
}

var File_transport_internet_xtls_config_proto protoreflect.FileDescriptor

var file_transport_internet_xtls_config_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e,
	0x78, 0x74, 0x6c, 0x73, 0x22, 0xb9, 0x02, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x2e, 0x78, 0x74, 0x6c, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6f, 0x63, 0x73, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6f, 0x63, 0x73, 0x70, 0x53, 0x74, 0x61, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x22, 0x44, 0x0a, 0x05, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4e, 0x43, 0x49, 0x50, 0x48, 0x45, 0x52, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x10, 0x02,
	0x22, 0xd4, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x2e, 0x78, 0x74, 0x6c, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x3a, 0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75,
	0x69, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x42, 0x76, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x78, 0x74, 0x6c, 0x73, 0x50, 0x01, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78,
	0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x78, 0x74, 0x6c, 0x73,
	0xaa, 0x02, 0x1c, 0x58, 0x72, 0x61, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x58, 0x74, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_xtls_config_proto_rawDescOnce sync.Once
	file_transport_internet_xtls_config_proto_rawDescData = file_transport_internet_xtls_config_proto_rawDesc
)

func file_transport_internet_xtls_config_proto_rawDescGZIP() []byte {
	file_transport_internet_xtls_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_xtls_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_xtls_config_proto_rawDescData)
	})
	return file_transport_internet_xtls_config_proto_rawDescData
}

var file_transport_internet_xtls_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transport_internet_xtls_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_internet_xtls_config_proto_goTypes = []interface{}{
	(Certificate_Usage)(0), // 0: xray.transport.internet.xtls.Certificate.Usage
	(*Certificate)(nil),    // 1: xray.transport.internet.xtls.Certificate
	(*Config)(nil),         // 2: xray.transport.internet.xtls.Config
}
var file_transport_internet_xtls_config_proto_depIdxs = []int32{
	0, // 0: xray.transport.internet.xtls.Certificate.usage:type_name -> xray.transport.internet.xtls.Certificate.Usage
	1, // 1: xray.transport.internet.xtls.Config.certificate:type_name -> xray.transport.internet.xtls.Certificate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transport_internet_xtls_config_proto_init() }
func file_transport_internet_xtls_config_proto_init() {
	if File_transport_internet_xtls_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_xtls_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_transport_internet_xtls_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_transport_internet_xtls_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_xtls_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_xtls_config_proto_depIdxs,
		EnumInfos:         file_transport_internet_xtls_config_proto_enumTypes,
		MessageInfos:      file_transport_internet_xtls_config_proto_msgTypes,
	}.Build()
	File_transport_internet_xtls_config_proto = out.File
	file_transport_internet_xtls_config_proto_rawDesc = nil
	file_transport_internet_xtls_config_proto_goTypes = nil
	file_transport_internet_xtls_config_proto_depIdxs = nil
}
