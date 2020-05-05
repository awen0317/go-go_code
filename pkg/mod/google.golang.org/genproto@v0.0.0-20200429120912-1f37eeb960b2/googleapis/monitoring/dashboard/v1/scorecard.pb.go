// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/dashboard/v1/scorecard.proto

package dashboard

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// A widget showing the latest value of a metric, and how this value relates to
// one or more thresholds.
type Scorecard struct {
	// Fields for querying time series data from the
	// Stackdriver metrics API.
	TimeSeriesQuery *TimeSeriesQuery `protobuf:"bytes,1,opt,name=time_series_query,json=timeSeriesQuery,proto3" json:"time_series_query,omitempty"`
	// Defines the optional additional chart shown on the scorecard. If
	// neither is included - then a default scorecard is shown.
	//
	// Types that are valid to be assigned to DataView:
	//	*Scorecard_GaugeView_
	//	*Scorecard_SparkChartView_
	DataView isScorecard_DataView `protobuf_oneof:"data_view"`
	// The thresholds used to determine the state of the scorecard given the
	// time series' current value. For an actual value x, the scorecard is in a
	// danger state if x is less than or equal to a danger threshold that triggers
	// below, or greater than or equal to a danger threshold that triggers above.
	// Similarly, if x is above/below a warning threshold that triggers
	// above/below, then the scorecard is in a warning state - unless x also puts
	// it in a danger state. (Danger trumps warning.)
	//
	// As an example, consider a scorecard with the following four thresholds:
	// {
	//   value: 90,
	//   category: 'DANGER',
	//   trigger: 'ABOVE',
	// },
	// {
	//   value: 70,
	//   category: 'WARNING',
	//   trigger: 'ABOVE',
	// },
	// {
	//   value: 10,
	//   category: 'DANGER',
	//   trigger: 'BELOW',
	// },
	// {
	//   value: 20,
	//   category: 'WARNING',
	//   trigger: 'BELOW',
	// }
	//
	// Then: values less than or equal to 10 would put the scorecard in a DANGER
	// state, values greater than 10 but less than or equal to 20 a WARNING state,
	// values strictly between 20 and 70 an OK state, values greater than or equal
	// to 70 but less than 90 a WARNING state, and values greater than or equal to
	// 90 a DANGER state.
	Thresholds           []*Threshold `protobuf:"bytes,6,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Scorecard) Reset()         { *m = Scorecard{} }
func (m *Scorecard) String() string { return proto.CompactTextString(m) }
func (*Scorecard) ProtoMessage()    {}
func (*Scorecard) Descriptor() ([]byte, []int) {
	return fileDescriptor_0914d2b6c0dbb42b, []int{0}
}

func (m *Scorecard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scorecard.Unmarshal(m, b)
}
func (m *Scorecard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scorecard.Marshal(b, m, deterministic)
}
func (m *Scorecard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scorecard.Merge(m, src)
}
func (m *Scorecard) XXX_Size() int {
	return xxx_messageInfo_Scorecard.Size(m)
}
func (m *Scorecard) XXX_DiscardUnknown() {
	xxx_messageInfo_Scorecard.DiscardUnknown(m)
}

var xxx_messageInfo_Scorecard proto.InternalMessageInfo

func (m *Scorecard) GetTimeSeriesQuery() *TimeSeriesQuery {
	if m != nil {
		return m.TimeSeriesQuery
	}
	return nil
}

type isScorecard_DataView interface {
	isScorecard_DataView()
}

type Scorecard_GaugeView_ struct {
	GaugeView *Scorecard_GaugeView `protobuf:"bytes,4,opt,name=gauge_view,json=gaugeView,proto3,oneof"`
}

type Scorecard_SparkChartView_ struct {
	SparkChartView *Scorecard_SparkChartView `protobuf:"bytes,5,opt,name=spark_chart_view,json=sparkChartView,proto3,oneof"`
}

func (*Scorecard_GaugeView_) isScorecard_DataView() {}

func (*Scorecard_SparkChartView_) isScorecard_DataView() {}

func (m *Scorecard) GetDataView() isScorecard_DataView {
	if m != nil {
		return m.DataView
	}
	return nil
}

func (m *Scorecard) GetGaugeView() *Scorecard_GaugeView {
	if x, ok := m.GetDataView().(*Scorecard_GaugeView_); ok {
		return x.GaugeView
	}
	return nil
}

func (m *Scorecard) GetSparkChartView() *Scorecard_SparkChartView {
	if x, ok := m.GetDataView().(*Scorecard_SparkChartView_); ok {
		return x.SparkChartView
	}
	return nil
}

func (m *Scorecard) GetThresholds() []*Threshold {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Scorecard) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Scorecard_GaugeView_)(nil),
		(*Scorecard_SparkChartView_)(nil),
	}
}

// A gauge chart shows where the current value sits within a pre-defined
// range. The upper and lower bounds should define the possible range of
// values for the scorecard's query (inclusive).
type Scorecard_GaugeView struct {
	// The lower bound for this gauge chart. The value of the chart should
	// always be greater than or equal to this.
	LowerBound float64 `protobuf:"fixed64,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	// The upper bound for this gauge chart. The value of the chart should
	// always be less than or equal to this.
	UpperBound           float64  `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scorecard_GaugeView) Reset()         { *m = Scorecard_GaugeView{} }
func (m *Scorecard_GaugeView) String() string { return proto.CompactTextString(m) }
func (*Scorecard_GaugeView) ProtoMessage()    {}
func (*Scorecard_GaugeView) Descriptor() ([]byte, []int) {
	return fileDescriptor_0914d2b6c0dbb42b, []int{0, 0}
}

func (m *Scorecard_GaugeView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scorecard_GaugeView.Unmarshal(m, b)
}
func (m *Scorecard_GaugeView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scorecard_GaugeView.Marshal(b, m, deterministic)
}
func (m *Scorecard_GaugeView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scorecard_GaugeView.Merge(m, src)
}
func (m *Scorecard_GaugeView) XXX_Size() int {
	return xxx_messageInfo_Scorecard_GaugeView.Size(m)
}
func (m *Scorecard_GaugeView) XXX_DiscardUnknown() {
	xxx_messageInfo_Scorecard_GaugeView.DiscardUnknown(m)
}

var xxx_messageInfo_Scorecard_GaugeView proto.InternalMessageInfo

func (m *Scorecard_GaugeView) GetLowerBound() float64 {
	if m != nil {
		return m.LowerBound
	}
	return 0
}

func (m *Scorecard_GaugeView) GetUpperBound() float64 {
	if m != nil {
		return m.UpperBound
	}
	return 0
}

// A sparkChart is a small chart suitable for inclusion in a table-cell or
// inline in text. This message contains the configuration for a sparkChart
// to show up on a Scorecard, showing recent trends of the scorecard's
// timeseries.
type Scorecard_SparkChartView struct {
	// The type of sparkchart to show in this chartView.
	SparkChartType SparkChartType `protobuf:"varint,1,opt,name=spark_chart_type,json=sparkChartType,proto3,enum=google.monitoring.dashboard.v1.SparkChartType" json:"spark_chart_type,omitempty"`
	// The lower bound on data point frequency in the chart implemented by
	// specifying the minimum alignment period to use in a time series query.
	// For example, if the data is published once every 10 minutes it would not
	// make sense to fetch and align data at one minute intervals. This field is
	// optional and exists only as a hint.
	MinAlignmentPeriod   *duration.Duration `protobuf:"bytes,2,opt,name=min_alignment_period,json=minAlignmentPeriod,proto3" json:"min_alignment_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Scorecard_SparkChartView) Reset()         { *m = Scorecard_SparkChartView{} }
func (m *Scorecard_SparkChartView) String() string { return proto.CompactTextString(m) }
func (*Scorecard_SparkChartView) ProtoMessage()    {}
func (*Scorecard_SparkChartView) Descriptor() ([]byte, []int) {
	return fileDescriptor_0914d2b6c0dbb42b, []int{0, 1}
}

func (m *Scorecard_SparkChartView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scorecard_SparkChartView.Unmarshal(m, b)
}
func (m *Scorecard_SparkChartView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scorecard_SparkChartView.Marshal(b, m, deterministic)
}
func (m *Scorecard_SparkChartView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scorecard_SparkChartView.Merge(m, src)
}
func (m *Scorecard_SparkChartView) XXX_Size() int {
	return xxx_messageInfo_Scorecard_SparkChartView.Size(m)
}
func (m *Scorecard_SparkChartView) XXX_DiscardUnknown() {
	xxx_messageInfo_Scorecard_SparkChartView.DiscardUnknown(m)
}

var xxx_messageInfo_Scorecard_SparkChartView proto.InternalMessageInfo

func (m *Scorecard_SparkChartView) GetSparkChartType() SparkChartType {
	if m != nil {
		return m.SparkChartType
	}
	return SparkChartType_SPARK_CHART_TYPE_UNSPECIFIED
}

func (m *Scorecard_SparkChartView) GetMinAlignmentPeriod() *duration.Duration {
	if m != nil {
		return m.MinAlignmentPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Scorecard)(nil), "google.monitoring.dashboard.v1.Scorecard")
	proto.RegisterType((*Scorecard_GaugeView)(nil), "google.monitoring.dashboard.v1.Scorecard.GaugeView")
	proto.RegisterType((*Scorecard_SparkChartView)(nil), "google.monitoring.dashboard.v1.Scorecard.SparkChartView")
}

func init() {
	proto.RegisterFile("google/monitoring/dashboard/v1/scorecard.proto", fileDescriptor_0914d2b6c0dbb42b)
}

var fileDescriptor_0914d2b6c0dbb42b = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xc9, 0xc0, 0x8c, 0x54, 0x47, 0x2a, 0x60, 0xb1, 0x28, 0x5d, 0xc0, 0x68, 0x56, 0x45,
	0x42, 0xb6, 0xda, 0xd9, 0xa0, 0xb0, 0xa2, 0x33, 0x52, 0x41, 0x68, 0xa4, 0x92, 0x56, 0x23, 0x04,
	0x8b, 0xc8, 0x8d, 0x1f, 0xae, 0x45, 0x62, 0x07, 0xdb, 0x69, 0xd5, 0x03, 0x70, 0x19, 0x2e, 0xc0,
	0x9d, 0x38, 0x05, 0x8a, 0xd3, 0x84, 0xc9, 0x02, 0xc2, 0xf2, 0x3d, 0xff, 0xff, 0xe7, 0x3f, 0xcf,
	0x2f, 0x88, 0x08, 0xad, 0x45, 0x06, 0x34, 0xd7, 0x4a, 0x3a, 0x6d, 0xa4, 0x12, 0x94, 0x33, 0xbb,
	0xdd, 0x68, 0x66, 0x38, 0xdd, 0x4d, 0xa9, 0x4d, 0xb5, 0x81, 0x94, 0x19, 0x4e, 0x0a, 0xa3, 0x9d,
	0xc6, 0xcf, 0x6a, 0x3d, 0xf9, 0xa3, 0x27, 0xad, 0x9e, 0xec, 0xa6, 0xe3, 0x97, 0x3d, 0xbc, 0x1c,
	0x9c, 0x91, 0xa9, 0xad, 0x69, 0xe3, 0x23, 0x8d, 0xfa, 0x6a, 0x53, 0x7e, 0xa1, 0xbc, 0x34, 0xcc,
	0x49, 0xad, 0xea, 0xf3, 0x8b, 0xef, 0xa7, 0x68, 0xb0, 0x6a, 0x12, 0xe0, 0xcf, 0xe8, 0xb1, 0x93,
	0x39, 0x24, 0x16, 0x8c, 0x04, 0x9b, 0x7c, 0x2b, 0xc1, 0x1c, 0x46, 0xc1, 0x79, 0x30, 0x09, 0x67,
	0x94, 0xfc, 0x3b, 0x17, 0x59, 0xcb, 0x1c, 0x56, 0xde, 0xf7, 0xa1, 0xb2, 0xc5, 0x0f, 0x5d, 0xb7,
	0x81, 0xd7, 0x08, 0x09, 0x56, 0x0a, 0x48, 0x76, 0x12, 0xf6, 0xa3, 0x07, 0x9e, 0x7a, 0xd9, 0x47,
	0x6d, 0xb3, 0x91, 0x45, 0xe5, 0xbd, 0x95, 0xb0, 0x7f, 0x7b, 0x2f, 0x1e, 0x88, 0xa6, 0xc0, 0x1c,
	0x3d, 0xb2, 0x05, 0x33, 0x5f, 0x93, 0x74, 0xcb, 0x8c, 0xab, 0xd9, 0xa7, 0x9e, 0xfd, 0xea, 0xff,
	0xd9, 0xab, 0x8a, 0x70, 0x55, 0x01, 0x8e, 0x17, 0x0c, 0x6d, 0xa7, 0x83, 0xdf, 0x21, 0xe4, 0xb6,
	0x06, 0xec, 0x56, 0x67, 0xdc, 0x8e, 0xce, 0xce, 0xef, 0x4f, 0xc2, 0xd9, 0x8b, 0xde, 0x89, 0x34,
	0x8e, 0xf8, 0x8e, 0x79, 0x7c, 0x83, 0x06, 0xed, 0xa7, 0xe0, 0xe7, 0x28, 0xcc, 0xf4, 0x1e, 0x4c,
	0xb2, 0xd1, 0xa5, 0xe2, 0x7e, 0xd4, 0x41, 0x8c, 0x7c, 0x6b, 0x5e, 0x75, 0x2a, 0x41, 0x59, 0x14,
	0xad, 0xe0, 0xa4, 0x16, 0xf8, 0x96, 0x17, 0x8c, 0x7f, 0x06, 0x68, 0xd8, 0x8d, 0x8f, 0x3f, 0x76,
	0x47, 0xe2, 0x0e, 0x05, 0x78, 0xf2, 0x70, 0x46, 0x7a, 0x47, 0xd2, 0x92, 0xd6, 0x87, 0x02, 0xee,
	0x8e, 0xa1, 0xaa, 0xf1, 0x7b, 0xf4, 0x24, 0x97, 0x2a, 0x61, 0x99, 0x14, 0x2a, 0x07, 0xe5, 0x92,
	0x02, 0x8c, 0xd4, 0x75, 0xac, 0x70, 0xf6, 0xb4, 0xa1, 0x37, 0xcb, 0x46, 0xae, 0x8f, 0xcb, 0x16,
	0xe3, 0x5c, 0xaa, 0x37, 0x8d, 0x6b, 0xe9, 0x4d, 0xf3, 0x10, 0x0d, 0x38, 0x73, 0xcc, 0x3f, 0xd9,
	0xfc, 0x47, 0x80, 0x2e, 0x52, 0x9d, 0xf7, 0xe4, 0x9b, 0x0f, 0xdb, 0x37, 0x5b, 0x56, 0x77, 0x2c,
	0x83, 0x4f, 0x8b, 0xa3, 0x43, 0xe8, 0x8c, 0x29, 0x41, 0xb4, 0x11, 0x54, 0x80, 0xf2, 0x09, 0x68,
	0x7d, 0xc4, 0x0a, 0x69, 0xff, 0xf6, 0xb7, 0xbc, 0x6e, 0x8b, 0x5f, 0x27, 0x93, 0x85, 0x97, 0x47,
	0xd1, 0x55, 0xa6, 0x4b, 0x1e, 0x45, 0x37, 0xad, 0x25, 0x8a, 0xae, 0x1b, 0x59, 0x14, 0xdd, 0x4e,
	0x37, 0x67, 0x1e, 0x7f, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x74, 0x34, 0x3c, 0x8d, 0xdb, 0x03,
	0x00, 0x00,
}
