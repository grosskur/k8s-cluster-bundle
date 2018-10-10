// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/bundle/v1alpha1/cluster_bundle.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The ClusterBundle is a packaging format for Kubernetes Components.
type ClusterBundle struct {
	// Required. Kubernetes API Version for the Bundle.
	// Must be gke.io/k8s-cluster-bundle/v1alpha1.
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Required. The Kubernetes `kind` for this object. Must be 'ClusterBundle'.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Required. Kubernetes ObjectMeta proto. The Metadata.name field must be
	// filled out for each Bundle.
	Metadata *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Required. Version-string for this bundle. The version should be a SemVer
	// string (see https://semver.org/) of the form X.Y.Z (Major.Minor.Patch).
	// Generally speaking, a major-version (changes should indicate breaking
	// changes, minor-versions should indicate backwards compatible features, and
	// patch changes should indicate backwords compatible. If there are any
	// changes to the bundle, then the version string must be incremented.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Spec for the ClusterBundle, which specifies the intended Kubernetes cluster
	// configuration.
	Spec                 *ClusterBundleSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClusterBundle) Reset()         { *m = ClusterBundle{} }
func (m *ClusterBundle) String() string { return proto.CompactTextString(m) }
func (*ClusterBundle) ProtoMessage()    {}
func (*ClusterBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b30f666a9d42204, []int{0}
}

func (m *ClusterBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterBundle.Unmarshal(m, b)
}
func (m *ClusterBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterBundle.Marshal(b, m, deterministic)
}
func (m *ClusterBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterBundle.Merge(m, src)
}
func (m *ClusterBundle) XXX_Size() int {
	return xxx_messageInfo_ClusterBundle.Size(m)
}
func (m *ClusterBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterBundle.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterBundle proto.InternalMessageInfo

func (m *ClusterBundle) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *ClusterBundle) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ClusterBundle) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ClusterBundle) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClusterBundle) GetSpec() *ClusterBundleSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// A spec object that wraps the Cluster Bundle.
type ClusterBundleSpec struct {
	// Configuration for the nodes.
	NodeConfigs []*NodeConfig `protobuf:"bytes,1,rep,name=node_configs,json=nodeConfigs,proto3" json:"node_configs,omitempty"`
	// Configuration for the nodes, specified as external files. The process of
	// inlining reads node config files into node configs, and so after
	// inlining, this list will be empty.
	NodeConfigFiles []*File `protobuf:"bytes,2,rep,name=node_config_files,json=nodeConfigFiles,proto3" json:"node_config_files,omitempty"`
	// Kubernetes objects grouped into cluster components and versioned together.
	// These could be applications or they could be some sort of supporting
	// object collection.
	Components []*ClusterComponent `protobuf:"bytes,3,rep,name=components,proto3" json:"components,omitempty"`
	// Cluster components that are specified externally as Files. The process of inlining
	// for a bundle reads component files into components, and so after
	// inlining, this list will be empty.
	ComponentFiles       []*File  `protobuf:"bytes,4,rep,name=component_files,json=componentFiles,proto3" json:"component_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterBundleSpec) Reset()         { *m = ClusterBundleSpec{} }
func (m *ClusterBundleSpec) String() string { return proto.CompactTextString(m) }
func (*ClusterBundleSpec) ProtoMessage()    {}
func (*ClusterBundleSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b30f666a9d42204, []int{1}
}

func (m *ClusterBundleSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterBundleSpec.Unmarshal(m, b)
}
func (m *ClusterBundleSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterBundleSpec.Marshal(b, m, deterministic)
}
func (m *ClusterBundleSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterBundleSpec.Merge(m, src)
}
func (m *ClusterBundleSpec) XXX_Size() int {
	return xxx_messageInfo_ClusterBundleSpec.Size(m)
}
func (m *ClusterBundleSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterBundleSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterBundleSpec proto.InternalMessageInfo

func (m *ClusterBundleSpec) GetNodeConfigs() []*NodeConfig {
	if m != nil {
		return m.NodeConfigs
	}
	return nil
}

func (m *ClusterBundleSpec) GetNodeConfigFiles() []*File {
	if m != nil {
		return m.NodeConfigFiles
	}
	return nil
}

func (m *ClusterBundleSpec) GetComponents() []*ClusterComponent {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *ClusterBundleSpec) GetComponentFiles() []*File {
	if m != nil {
		return m.ComponentFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterBundle)(nil), "pkg.apis.bundle.v1alpha1.ClusterBundle")
	proto.RegisterType((*ClusterBundleSpec)(nil), "pkg.apis.bundle.v1alpha1.ClusterBundleSpec")
}

func init() {
	proto.RegisterFile("pkg/apis/bundle/v1alpha1/cluster_bundle.proto", fileDescriptor_6b30f666a9d42204)
}

var fileDescriptor_6b30f666a9d42204 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbb, 0x6e, 0xdb, 0x30,
	0x14, 0x86, 0x21, 0x5b, 0xbd, 0xd1, 0x6d, 0x0d, 0x73, 0x22, 0x3c, 0xb4, 0x86, 0xd1, 0xc1, 0x70,
	0x6b, 0xa9, 0x6e, 0x97, 0x6e, 0x0d, 0x2c, 0x20, 0x06, 0x8c, 0xdc, 0xa0, 0x00, 0x19, 0xb2, 0x08,
	0x94, 0x44, 0xcb, 0x8c, 0x28, 0x92, 0x10, 0x29, 0xbf, 0x5d, 0x1e, 0x26, 0x6f, 0x12, 0x88, 0xba,
	0xc4, 0x41, 0xa0, 0x38, 0x9b, 0xce, 0xe1, 0xff, 0x7f, 0xe7, 0x27, 0x75, 0xc0, 0x42, 0xa6, 0x89,
	0x8b, 0x25, 0x55, 0x6e, 0x58, 0xf0, 0x98, 0x11, 0x77, 0xbf, 0xc4, 0x4c, 0xee, 0xf0, 0xd2, 0x8d,
	0x58, 0xa1, 0x34, 0xc9, 0x83, 0xaa, 0xef, 0xc8, 0x5c, 0x68, 0x01, 0x91, 0x4c, 0x13, 0xa7, 0x94,
	0x3b, 0x75, 0xbb, 0x91, 0x8f, 0x7f, 0x75, 0x82, 0xaa, 0x3a, 0x88, 0x44, 0x96, 0x09, 0x5e, 0x71,
	0xc6, 0xbf, 0x8f, 0x8e, 0x8d, 0x44, 0x26, 0x05, 0x27, 0x5c, 0xd7, 0x8e, 0x79, 0xa7, 0x83, 0x8b,
	0xb8, 0xa4, 0xf3, 0x2d, 0x4d, 0x8e, 0x6a, 0x45, 0x78, 0x47, 0x22, 0x1d, 0x64, 0x44, 0xe3, 0x4a,
	0x3b, 0x7d, 0xb0, 0xc0, 0x17, 0xaf, 0x9a, 0xb9, 0x32, 0x62, 0xf8, 0x1d, 0x0c, 0xb0, 0xa4, 0xc1,
	0x9e, 0xe4, 0x8a, 0x0a, 0x8e, 0xac, 0x89, 0x35, 0xfb, 0xe4, 0x03, 0x2c, 0xe9, 0x4d, 0xd5, 0x81,
	0x10, 0xd8, 0x29, 0xe5, 0x31, 0xea, 0x99, 0x13, 0xf3, 0x0d, 0x4f, 0xc0, 0xc7, 0x12, 0x1a, 0x63,
	0x8d, 0x51, 0x7f, 0x62, 0xcd, 0x06, 0x7f, 0x7e, 0x38, 0x5d, 0x6f, 0xe5, 0x5c, 0x9a, 0x14, 0xe7,
	0x44, 0x63, 0xbf, 0x75, 0x41, 0x04, 0x3e, 0x34, 0x23, 0x6d, 0x03, 0x6e, 0x4a, 0xf8, 0x1f, 0xd8,
	0x4a, 0x92, 0x08, 0xbd, 0x33, 0xdc, 0x9f, 0xdd, 0xdc, 0x67, 0xf7, 0xb8, 0x96, 0x24, 0xf2, 0x8d,
	0x71, 0x7a, 0xdf, 0x03, 0xa3, 0x17, 0x67, 0x70, 0x0d, 0x3e, 0x1f, 0x3c, 0x9d, 0x42, 0xd6, 0xa4,
	0xff, 0x7a, 0xec, 0x0b, 0x11, 0x13, 0xcf, 0x88, 0xfd, 0x01, 0x6f, 0xbf, 0x15, 0xdc, 0x80, 0xd1,
	0x01, 0x28, 0xd8, 0x52, 0x46, 0x14, 0xea, 0x19, 0xda, 0xb7, 0x6e, 0xda, 0x29, 0x65, 0xc4, 0x1f,
	0x3e, 0x71, 0xca, 0xba, 0x64, 0x81, 0xf6, 0xcf, 0x2b, 0xd4, 0x37, 0x90, 0xf9, 0xd1, 0x1b, 0x7b,
	0x8d, 0xc5, 0x3f, 0x70, 0xc3, 0x35, 0x18, 0xb6, 0x55, 0x9d, 0xca, 0x7e, 0x53, 0xaa, 0xaf, 0xad,
	0xcd, 0x84, 0x5a, 0x9d, 0xdd, 0x6e, 0x12, 0xaa, 0x77, 0x45, 0xe8, 0x44, 0x22, 0x73, 0xd7, 0x42,
	0x24, 0x8c, 0x78, 0x4c, 0x14, 0xf1, 0x15, 0xc3, 0x7a, 0x2b, 0xf2, 0xcc, 0x4d, 0xff, 0xa9, 0x45,
	0xbd, 0xb9, 0x8b, 0x7a, 0xe7, 0xba, 0x76, 0x30, 0x7c, 0x6f, 0x16, 0xef, 0xef, 0x63, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0xe8, 0xf4, 0xde, 0x7b, 0x03, 0x00, 0x00,
}