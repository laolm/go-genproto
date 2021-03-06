// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/dataset.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A workspace for solving a single, particular machine learning (ML) problem.
// A workspace contains examples that may be annotated.
type Dataset struct {
	// Required.
	// The dataset metadata that is specific to the problem type.
	//
	// Types that are valid to be assigned to DatasetMetadata:
	//	*Dataset_TranslationDatasetMetadata
	//	*Dataset_ImageClassificationDatasetMetadata
	//	*Dataset_TextClassificationDatasetMetadata
	//	*Dataset_ImageObjectDetectionDatasetMetadata
	//	*Dataset_VideoClassificationDatasetMetadata
	//	*Dataset_VideoObjectTrackingDatasetMetadata
	//	*Dataset_TextExtractionDatasetMetadata
	//	*Dataset_TextSentimentDatasetMetadata
	//	*Dataset_TablesDatasetMetadata
	DatasetMetadata isDataset_DatasetMetadata `protobuf_oneof:"dataset_metadata"`
	// Output only. The resource name of the dataset.
	// Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the dataset to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User-provided description of the dataset. The description can be up to
	// 25000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The number of examples in the dataset.
	ExampleCount int32 `protobuf:"varint,21,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Output only. Timestamp when this dataset was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag                 string   `protobuf:"bytes,17,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f2b1dc66a1e92da, []int{0}
}

func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (m *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(m, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

type isDataset_DatasetMetadata interface {
	isDataset_DatasetMetadata()
}

type Dataset_TranslationDatasetMetadata struct {
	TranslationDatasetMetadata *TranslationDatasetMetadata `protobuf:"bytes,23,opt,name=translation_dataset_metadata,json=translationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageClassificationDatasetMetadata struct {
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `protobuf:"bytes,24,opt,name=image_classification_dataset_metadata,json=imageClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextClassificationDatasetMetadata struct {
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `protobuf:"bytes,25,opt,name=text_classification_dataset_metadata,json=textClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageObjectDetectionDatasetMetadata struct {
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `protobuf:"bytes,26,opt,name=image_object_detection_dataset_metadata,json=imageObjectDetectionDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoClassificationDatasetMetadata struct {
	VideoClassificationDatasetMetadata *VideoClassificationDatasetMetadata `protobuf:"bytes,31,opt,name=video_classification_dataset_metadata,json=videoClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoObjectTrackingDatasetMetadata struct {
	VideoObjectTrackingDatasetMetadata *VideoObjectTrackingDatasetMetadata `protobuf:"bytes,29,opt,name=video_object_tracking_dataset_metadata,json=videoObjectTrackingDatasetMetadata,proto3,oneof"`
}

type Dataset_TextExtractionDatasetMetadata struct {
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `protobuf:"bytes,28,opt,name=text_extraction_dataset_metadata,json=textExtractionDatasetMetadata,proto3,oneof"`
}

type Dataset_TextSentimentDatasetMetadata struct {
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `protobuf:"bytes,30,opt,name=text_sentiment_dataset_metadata,json=textSentimentDatasetMetadata,proto3,oneof"`
}

type Dataset_TablesDatasetMetadata struct {
	TablesDatasetMetadata *TablesDatasetMetadata `protobuf:"bytes,33,opt,name=tables_dataset_metadata,json=tablesDatasetMetadata,proto3,oneof"`
}

func (*Dataset_TranslationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageObjectDetectionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoObjectTrackingDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextExtractionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextSentimentDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TablesDatasetMetadata) isDataset_DatasetMetadata() {}

func (m *Dataset) GetDatasetMetadata() isDataset_DatasetMetadata {
	if m != nil {
		return m.DatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTranslationDatasetMetadata() *TranslationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TranslationDatasetMetadata); ok {
		return x.TranslationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageClassificationDatasetMetadata() *ImageClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageClassificationDatasetMetadata); ok {
		return x.ImageClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextClassificationDatasetMetadata() *TextClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextClassificationDatasetMetadata); ok {
		return x.TextClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageObjectDetectionDatasetMetadata() *ImageObjectDetectionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageObjectDetectionDatasetMetadata); ok {
		return x.ImageObjectDetectionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetVideoClassificationDatasetMetadata() *VideoClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_VideoClassificationDatasetMetadata); ok {
		return x.VideoClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetVideoObjectTrackingDatasetMetadata() *VideoObjectTrackingDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_VideoObjectTrackingDatasetMetadata); ok {
		return x.VideoObjectTrackingDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextExtractionDatasetMetadata() *TextExtractionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextExtractionDatasetMetadata); ok {
		return x.TextExtractionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextSentimentDatasetMetadata() *TextSentimentDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextSentimentDatasetMetadata); ok {
		return x.TextSentimentDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTablesDatasetMetadata() *TablesDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TablesDatasetMetadata); ok {
		return x.TablesDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dataset) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Dataset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Dataset) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func (m *Dataset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Dataset) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Dataset) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Dataset_TranslationDatasetMetadata)(nil),
		(*Dataset_ImageClassificationDatasetMetadata)(nil),
		(*Dataset_TextClassificationDatasetMetadata)(nil),
		(*Dataset_ImageObjectDetectionDatasetMetadata)(nil),
		(*Dataset_VideoClassificationDatasetMetadata)(nil),
		(*Dataset_VideoObjectTrackingDatasetMetadata)(nil),
		(*Dataset_TextExtractionDatasetMetadata)(nil),
		(*Dataset_TextSentimentDatasetMetadata)(nil),
		(*Dataset_TablesDatasetMetadata)(nil),
	}
}

func init() {
	proto.RegisterType((*Dataset)(nil), "google.cloud.automl.v1beta1.Dataset")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/dataset.proto", fileDescriptor_1f2b1dc66a1e92da)
}

var fileDescriptor_1f2b1dc66a1e92da = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6a, 0xd4, 0x4e,
	0x14, 0xc7, 0x7f, 0xe9, 0xcf, 0x3f, 0x38, 0x5b, 0x45, 0x07, 0x4a, 0xe3, 0x76, 0xdb, 0xdd, 0xb6,
	0xda, 0xae, 0x17, 0x26, 0xb4, 0x5e, 0x88, 0x29, 0xa8, 0xfd, 0x23, 0x2a, 0x58, 0x95, 0xb5, 0xf4,
	0x42, 0x0a, 0x61, 0x36, 0x3b, 0x0d, 0xa3, 0x49, 0x26, 0x24, 0x67, 0xcb, 0x8a, 0x77, 0x82, 0x0f,
	0x20, 0x28, 0xf8, 0x02, 0xbe, 0x8c, 0x6f, 0xe0, 0x2b, 0xf8, 0x14, 0x32, 0x67, 0x66, 0xad, 0x9a,
	0xdd, 0x49, 0xef, 0xb2, 0x73, 0x3e, 0x39, 0xe7, 0x93, 0xef, 0x0c, 0xb3, 0xe4, 0x56, 0x2c, 0x65,
	0x9c, 0x70, 0x3f, 0x4a, 0xe4, 0x70, 0xe0, 0xb3, 0x21, 0xc8, 0x34, 0xf1, 0x4f, 0x36, 0xfa, 0x1c,
	0xd8, 0x86, 0x3f, 0x60, 0xc0, 0x4a, 0x0e, 0x5e, 0x5e, 0x48, 0x90, 0x74, 0x41, 0xa3, 0x1e, 0xa2,
	0x9e, 0x46, 0x3d, 0x83, 0x36, 0xd7, 0x6d, 0x7d, 0x44, 0xca, 0x62, 0xae, 0xbb, 0x34, 0xbb, 0x36,
	0x10, 0x58, 0x3f, 0xe1, 0xa5, 0x21, 0xd7, 0xac, 0x24, 0x1f, 0x19, 0xaf, 0xe6, 0x6d, 0x2b, 0x57,
	0xb0, 0xac, 0x4c, 0x18, 0x08, 0x99, 0x19, 0xdc, 0x6a, 0x7a, 0x22, 0x06, 0x5c, 0x1a, 0xb0, 0x6d,
	0x40, 0xfc, 0xd5, 0x1f, 0x1e, 0xfb, 0x20, 0x52, 0x5e, 0x02, 0x4b, 0x73, 0x03, 0xb4, 0x0c, 0xc0,
	0x72, 0xe1, 0xb3, 0x2c, 0x93, 0x80, 0x63, 0x8c, 0xfe, 0xca, 0x8f, 0x06, 0xb9, 0xb8, 0xa7, 0x03,
	0xa4, 0xef, 0x49, 0xeb, 0x0f, 0x91, 0xd0, 0xe4, 0x1a, 0xa6, 0x1c, 0x98, 0x7a, 0x76, 0xe7, 0x3b,
	0x4e, 0xb7, 0xb1, 0x79, 0xd7, 0xb3, 0x24, 0xec, 0x1d, 0x9c, 0x36, 0x30, 0x6d, 0xf7, 0xcd, 0xeb,
	0x4f, 0xfe, 0xeb, 0x35, 0x61, 0x6a, 0x95, 0x7e, 0x76, 0xc8, 0x4d, 0xdc, 0x81, 0x30, 0x4a, 0x58,
	0x59, 0x8a, 0x63, 0x11, 0x4d, 0xd1, 0x70, 0x51, 0xe3, 0x81, 0x55, 0xe3, 0xa9, 0xea, 0xb4, 0xfb,
	0x57, 0xa3, 0xaa, 0xce, 0x8a, 0xa8, 0xa5, 0xe8, 0x27, 0x87, 0xdc, 0x50, 0xbb, 0x58, 0x6b, 0x75,
	0x1d, 0xad, 0xee, 0xdb, 0xc3, 0xe1, 0x23, 0xa8, 0x93, 0x5a, 0x86, 0x3a, 0x88, 0x7e, 0x75, 0xc8,
	0xba, 0x8e, 0x4a, 0xf6, 0xdf, 0xf0, 0x08, 0xc2, 0x01, 0x07, 0x1e, 0x4d, 0xd6, 0x6a, 0xa2, 0xd6,
	0xc3, 0xfa, 0xb0, 0x5e, 0x60, 0xab, 0xbd, 0x71, 0xa7, 0xaa, 0xd8, 0xaa, 0xa8, 0xc7, 0x70, 0x17,
	0xf1, 0x74, 0xd6, 0xe6, 0xd5, 0x3e, 0xc3, 0x2e, 0x1e, 0xaa, 0x4e, 0xb5, 0xbb, 0x78, 0x52, 0x4b,
	0xd1, 0x2f, 0x0e, 0x59, 0xd3, 0x5a, 0x26, 0x31, 0x28, 0x58, 0xf4, 0x56, 0x64, 0x71, 0xd5, 0x6b,
	0xf1, 0xac, 0x5e, 0x3a, 0x89, 0x03, 0xd3, 0x68, 0x9a, 0x97, 0x95, 0xa2, 0x1f, 0x1d, 0xd2, 0xc1,
	0xd3, 0xc5, 0x47, 0xca, 0x68, 0x72, 0x52, 0x2d, 0x34, 0x0a, 0x6a, 0x4f, 0xd6, 0xa3, 0xdf, 0x3d,
	0xaa, 0x32, 0x8b, 0x60, 0x03, 0xe8, 0x07, 0x87, 0xb4, 0xd1, 0xa3, 0xe4, 0x99, 0xba, 0x3f, 0x32,
	0xa8, 0x6a, 0x2c, 0xa1, 0xc6, 0xbd, 0x5a, 0x8d, 0x57, 0xe3, 0x16, 0x55, 0x8b, 0x16, 0x58, 0xea,
	0x34, 0x21, 0xf3, 0xfa, 0x66, 0xad, 0xce, 0x5e, 0xc6, 0xd9, 0x9b, 0xf6, 0xd9, 0xf8, 0x6e, 0x75,
	0xe8, 0x1c, 0x4c, 0x2a, 0x50, 0x4a, 0xce, 0x65, 0x2c, 0xe5, 0xae, 0xd3, 0x71, 0xba, 0x97, 0x7a,
	0xf8, 0x4c, 0x97, 0xc9, 0xec, 0x40, 0x94, 0x79, 0xc2, 0xde, 0x85, 0x58, 0x9b, 0xc1, 0x5a, 0xc3,
	0xac, 0x3d, 0x57, 0x48, 0x87, 0x34, 0x06, 0xbc, 0x8c, 0x0a, 0x91, 0xab, 0x1c, 0xdd, 0xff, 0x0d,
	0x71, 0xba, 0x44, 0x57, 0xc9, 0x65, 0x3e, 0x62, 0x69, 0x9e, 0xf0, 0x30, 0x92, 0xc3, 0x0c, 0xdc,
	0xb9, 0x8e, 0xd3, 0x3d, 0xdf, 0x9b, 0x35, 0x8b, 0xbb, 0x6a, 0x8d, 0x6e, 0x91, 0x46, 0x54, 0x70,
	0x06, 0x3c, 0x54, 0x59, 0xb8, 0x57, 0xf0, 0xfb, 0x9a, 0xe3, 0xef, 0x1b, 0xdf, 0xe5, 0xde, 0xc1,
	0xf8, 0x2e, 0xef, 0x11, 0x8d, 0xab, 0x05, 0xa5, 0xce, 0x81, 0xc5, 0xee, 0x35, 0xad, 0xae, 0x9e,
	0x77, 0x28, 0xb9, 0xfa, 0x6f, 0x6a, 0x3b, 0xdf, 0x1c, 0xd2, 0x8e, 0x64, 0x6a, 0x4b, 0xed, 0xa5,
	0xf3, 0x7a, 0xdb, 0x94, 0x63, 0x99, 0xb0, 0x2c, 0xf6, 0x64, 0x11, 0xfb, 0x31, 0xcf, 0x50, 0xc1,
	0xd7, 0x25, 0x96, 0x8b, 0x72, 0xe2, 0x1f, 0xd1, 0x96, 0xfe, 0xf9, 0x7d, 0x66, 0xe1, 0x31, 0x82,
	0x47, 0xbb, 0x0a, 0x3a, 0xda, 0x1e, 0x82, 0xdc, 0x4f, 0x8e, 0x0e, 0x35, 0xf4, 0x73, 0x66, 0x49,
	0x57, 0x83, 0x00, 0xcb, 0x41, 0x80, 0xf5, 0x67, 0x41, 0x60, 0x80, 0xfe, 0x05, 0x1c, 0x76, 0xe7,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xfd, 0x61, 0xd6, 0xe6, 0x07, 0x00, 0x00,
}
