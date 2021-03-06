// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

package message

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Header_HmacHashFunction int32

const (
	Header_MD5  Header_HmacHashFunction = 0
	Header_SHA1 Header_HmacHashFunction = 1
)

var Header_HmacHashFunction_name = map[int32]string{
	0: "MD5",
	1: "SHA1",
}
var Header_HmacHashFunction_value = map[string]int32{
	"MD5":  0,
	"SHA1": 1,
}

func (x Header_HmacHashFunction) Enum() *Header_HmacHashFunction {
	p := new(Header_HmacHashFunction)
	*p = x
	return p
}
func (x Header_HmacHashFunction) String() string {
	return proto.EnumName(Header_HmacHashFunction_name, int32(x))
}
func (x *Header_HmacHashFunction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Header_HmacHashFunction_value, data, "Header_HmacHashFunction")
	if err != nil {
		return err
	}
	*x = Header_HmacHashFunction(value)
	return nil
}

type Field_ValueType int32

const (
	Field_STRING  Field_ValueType = 0
	Field_BYTES   Field_ValueType = 1
	Field_INTEGER Field_ValueType = 2
	Field_DOUBLE  Field_ValueType = 3
	Field_BOOL    Field_ValueType = 4
)

var Field_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BYTES",
	2: "INTEGER",
	3: "DOUBLE",
	4: "BOOL",
}
var Field_ValueType_value = map[string]int32{
	"STRING":  0,
	"BYTES":   1,
	"INTEGER": 2,
	"DOUBLE":  3,
	"BOOL":    4,
}

func (x Field_ValueType) Enum() *Field_ValueType {
	p := new(Field_ValueType)
	*p = x
	return p
}
func (x Field_ValueType) String() string {
	return proto.EnumName(Field_ValueType_name, int32(x))
}
func (x *Field_ValueType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Field_ValueType_value, data, "Field_ValueType")
	if err != nil {
		return err
	}
	*x = Field_ValueType(value)
	return nil
}

type Header struct {
	MessageLength    *uint32                  `protobuf:"varint,1,req,name=message_length" json:"message_length,omitempty"`
	HmacHashFunction *Header_HmacHashFunction `protobuf:"varint,3,opt,name=hmac_hash_function,enum=message.Header_HmacHashFunction,def=0" json:"hmac_hash_function,omitempty"`
	HmacSigner       *string                  `protobuf:"bytes,4,opt,name=hmac_signer" json:"hmac_signer,omitempty"`
	HmacKeyVersion   *uint32                  `protobuf:"varint,5,opt,name=hmac_key_version" json:"hmac_key_version,omitempty"`
	Hmac             []byte                   `protobuf:"bytes,6,opt,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}

const Default_Header_HmacHashFunction Header_HmacHashFunction = Header_MD5

func (m *Header) GetMessageLength() uint32 {
	if m != nil && m.MessageLength != nil {
		return *m.MessageLength
	}
	return 0
}

func (m *Header) GetHmacHashFunction() Header_HmacHashFunction {
	if m != nil && m.HmacHashFunction != nil {
		return *m.HmacHashFunction
	}
	return Default_Header_HmacHashFunction
}

func (m *Header) GetHmacSigner() string {
	if m != nil && m.HmacSigner != nil {
		return *m.HmacSigner
	}
	return ""
}

func (m *Header) GetHmacKeyVersion() uint32 {
	if m != nil && m.HmacKeyVersion != nil {
		return *m.HmacKeyVersion
	}
	return 0
}

func (m *Header) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type Field struct {
	Name             *string          `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	ValueType        *Field_ValueType `protobuf:"varint,2,opt,name=value_type,enum=message.Field_ValueType,def=0" json:"value_type,omitempty"`
	Representation   *string          `protobuf:"bytes,3,opt,name=representation" json:"representation,omitempty"`
	ValueString      []string         `protobuf:"bytes,4,rep,name=value_string" json:"value_string,omitempty"`
	ValueBytes       [][]byte         `protobuf:"bytes,5,rep,name=value_bytes" json:"value_bytes,omitempty"`
	ValueInteger     []int64          `protobuf:"varint,6,rep,packed,name=value_integer" json:"value_integer,omitempty"`
	ValueDouble      []float64        `protobuf:"fixed64,7,rep,packed,name=value_double" json:"value_double,omitempty"`
	ValueBool        []bool           `protobuf:"varint,8,rep,packed,name=value_bool" json:"value_bool,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}

const Default_Field_ValueType Field_ValueType = Field_STRING

func (m *Field) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Field) GetValueType() Field_ValueType {
	if m != nil && m.ValueType != nil {
		return *m.ValueType
	}
	return Default_Field_ValueType
}

func (m *Field) GetRepresentation() string {
	if m != nil && m.Representation != nil {
		return *m.Representation
	}
	return ""
}

func (m *Field) GetValueString() []string {
	if m != nil {
		return m.ValueString
	}
	return nil
}

func (m *Field) GetValueBytes() [][]byte {
	if m != nil {
		return m.ValueBytes
	}
	return nil
}

func (m *Field) GetValueInteger() []int64 {
	if m != nil {
		return m.ValueInteger
	}
	return nil
}

func (m *Field) GetValueDouble() []float64 {
	if m != nil {
		return m.ValueDouble
	}
	return nil
}

func (m *Field) GetValueBool() []bool {
	if m != nil {
		return m.ValueBool
	}
	return nil
}

type Message struct {
	Uuid             []byte   `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Timestamp        *int64   `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	Type             *string  `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Logger           *string  `protobuf:"bytes,4,opt,name=logger" json:"logger,omitempty"`
	Severity         *int32   `protobuf:"varint,5,opt,name=severity,def=7" json:"severity,omitempty"`
	Payload          *string  `protobuf:"bytes,6,opt,name=payload" json:"payload,omitempty"`
	EnvVersion       *string  `protobuf:"bytes,7,opt,name=env_version" json:"env_version,omitempty"`
	Pid              *int32   `protobuf:"varint,8,opt,name=pid" json:"pid,omitempty"`
	Hostname         *string  `protobuf:"bytes,9,opt,name=hostname" json:"hostname,omitempty"`
	Fields           []*Field `protobuf:"bytes,10,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

const Default_Message_Severity int32 = 7

func (m *Message) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Message) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Message) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Message) GetLogger() string {
	if m != nil && m.Logger != nil {
		return *m.Logger
	}
	return ""
}

func (m *Message) GetSeverity() int32 {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return Default_Message_Severity
}

func (m *Message) GetPayload() string {
	if m != nil && m.Payload != nil {
		return *m.Payload
	}
	return ""
}

func (m *Message) GetEnvVersion() string {
	if m != nil && m.EnvVersion != nil {
		return *m.EnvVersion
	}
	return ""
}

func (m *Message) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *Message) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Message) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.Header_HmacHashFunction", Header_HmacHashFunction_name, Header_HmacHashFunction_value)
	proto.RegisterEnum("message.Field_ValueType", Field_ValueType_name, Field_ValueType_value)
}
