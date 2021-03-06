// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chargehive/chtype/payment_method_schema.proto

package chtype

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PaymentMethodSchema int32

const (
	PAYMENT_METHOD_SCHEMA_INVALID        PaymentMethodSchema = 0
	PAYMENT_METHOD_SCHEMA_ENCRYPTED_CARD PaymentMethodSchema = 1
	PAYMENT_METHOD_SCHEMA_CARD           PaymentMethodSchema = 2
	PAYMENT_METHOD_SCHEMA_APPLE_PAY      PaymentMethodSchema = 3
	PAYMENT_METHOD_SCHEMA_GOOGLE_PAY     PaymentMethodSchema = 4
	PAYMENT_METHOD_SCHEMA_PAYPAL         PaymentMethodSchema = 5
	PAYMENT_METHOD_SCHEMA_DIRECT_DEBIT   PaymentMethodSchema = 6
)

var PaymentMethodSchema_name = map[int32]string{
	0: "PAYMENT_METHOD_SCHEMA_INVALID",
	1: "PAYMENT_METHOD_SCHEMA_ENCRYPTED_CARD",
	2: "PAYMENT_METHOD_SCHEMA_CARD",
	3: "PAYMENT_METHOD_SCHEMA_APPLE_PAY",
	4: "PAYMENT_METHOD_SCHEMA_GOOGLE_PAY",
	5: "PAYMENT_METHOD_SCHEMA_PAYPAL",
	6: "PAYMENT_METHOD_SCHEMA_DIRECT_DEBIT",
}

var PaymentMethodSchema_value = map[string]int32{
	"PAYMENT_METHOD_SCHEMA_INVALID":        0,
	"PAYMENT_METHOD_SCHEMA_ENCRYPTED_CARD": 1,
	"PAYMENT_METHOD_SCHEMA_CARD":           2,
	"PAYMENT_METHOD_SCHEMA_APPLE_PAY":      3,
	"PAYMENT_METHOD_SCHEMA_GOOGLE_PAY":     4,
	"PAYMENT_METHOD_SCHEMA_PAYPAL":         5,
	"PAYMENT_METHOD_SCHEMA_DIRECT_DEBIT":   6,
}

func (x PaymentMethodSchema) String() string {
	return proto.EnumName(PaymentMethodSchema_name, int32(x))
}

func (PaymentMethodSchema) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{0}
}

type PaymentMethodSchemaApplePay struct {
	SessionUrl           string   `protobuf:"bytes,1,opt,name=session_url,json=sessionUrl,proto3" json:"session_url,omitempty"`
	InitiativeContext    string   `protobuf:"bytes,2,opt,name=initiative_context,json=initiativeContext,proto3" json:"initiative_context,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Session              string   `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaApplePay) Reset()         { *m = PaymentMethodSchemaApplePay{} }
func (m *PaymentMethodSchemaApplePay) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaApplePay) ProtoMessage()    {}
func (*PaymentMethodSchemaApplePay) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{0}
}
func (m *PaymentMethodSchemaApplePay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaApplePay.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaApplePay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaApplePay.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaApplePay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaApplePay.Merge(m, src)
}
func (m *PaymentMethodSchemaApplePay) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaApplePay.Size(m)
}
func (m *PaymentMethodSchemaApplePay) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaApplePay.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaApplePay proto.InternalMessageInfo

func (m *PaymentMethodSchemaApplePay) GetSessionUrl() string {
	if m != nil {
		return m.SessionUrl
	}
	return ""
}

func (m *PaymentMethodSchemaApplePay) GetInitiativeContext() string {
	if m != nil {
		return m.InitiativeContext
	}
	return ""
}

func (m *PaymentMethodSchemaApplePay) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PaymentMethodSchemaApplePay) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type PaymentMethodSchemaGooglePay struct {
	LowValueToken        string   `protobuf:"bytes,1,opt,name=low_value_token,json=lowValueToken,proto3" json:"low_value_token,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CardDetails          string   `protobuf:"bytes,3,opt,name=card_details,json=cardDetails,proto3" json:"card_details,omitempty"`
	CardNetwork          string   `protobuf:"bytes,4,opt,name=card_network,json=cardNetwork,proto3" json:"card_network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaGooglePay) Reset()         { *m = PaymentMethodSchemaGooglePay{} }
func (m *PaymentMethodSchemaGooglePay) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaGooglePay) ProtoMessage()    {}
func (*PaymentMethodSchemaGooglePay) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{1}
}
func (m *PaymentMethodSchemaGooglePay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaGooglePay.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaGooglePay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaGooglePay.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaGooglePay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaGooglePay.Merge(m, src)
}
func (m *PaymentMethodSchemaGooglePay) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaGooglePay.Size(m)
}
func (m *PaymentMethodSchemaGooglePay) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaGooglePay.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaGooglePay proto.InternalMessageInfo

func (m *PaymentMethodSchemaGooglePay) GetLowValueToken() string {
	if m != nil {
		return m.LowValueToken
	}
	return ""
}

func (m *PaymentMethodSchemaGooglePay) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PaymentMethodSchemaGooglePay) GetCardDetails() string {
	if m != nil {
		return m.CardDetails
	}
	return ""
}

func (m *PaymentMethodSchemaGooglePay) GetCardNetwork() string {
	if m != nil {
		return m.CardNetwork
	}
	return ""
}

type PaymentMethodSchemaPayPal struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaPayPal) Reset()         { *m = PaymentMethodSchemaPayPal{} }
func (m *PaymentMethodSchemaPayPal) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaPayPal) ProtoMessage()    {}
func (*PaymentMethodSchemaPayPal) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{2}
}
func (m *PaymentMethodSchemaPayPal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaPayPal.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaPayPal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaPayPal.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaPayPal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaPayPal.Merge(m, src)
}
func (m *PaymentMethodSchemaPayPal) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaPayPal.Size(m)
}
func (m *PaymentMethodSchemaPayPal) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaPayPal.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaPayPal proto.InternalMessageInfo

type PaymentMethodSchemaEncryptedCard struct {
	NameOnCard           []byte   `protobuf:"bytes,1,opt,name=name_on_card,json=nameOnCard,proto3" json:"name_on_card,omitempty"`
	Number               []byte   `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	ValidFromMonth       []byte   `protobuf:"bytes,3,opt,name=valid_from_month,json=validFromMonth,proto3" json:"valid_from_month,omitempty"`
	ValidFromYear        []byte   `protobuf:"bytes,4,opt,name=valid_from_year,json=validFromYear,proto3" json:"valid_from_year,omitempty"`
	ExpiryMonth          []byte   `protobuf:"bytes,5,opt,name=expiry_month,json=expiryMonth,proto3" json:"expiry_month,omitempty"`
	ExpiryYear           []byte   `protobuf:"bytes,6,opt,name=expiry_year,json=expiryYear,proto3" json:"expiry_year,omitempty"`
	IssueNumber          []byte   `protobuf:"bytes,7,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	TransportKeyId       string   `protobuf:"bytes,8,opt,name=transport_key_id,json=transportKeyId,proto3" json:"transport_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaEncryptedCard) Reset()         { *m = PaymentMethodSchemaEncryptedCard{} }
func (m *PaymentMethodSchemaEncryptedCard) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaEncryptedCard) ProtoMessage()    {}
func (*PaymentMethodSchemaEncryptedCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{3}
}
func (m *PaymentMethodSchemaEncryptedCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaEncryptedCard.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaEncryptedCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaEncryptedCard.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaEncryptedCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaEncryptedCard.Merge(m, src)
}
func (m *PaymentMethodSchemaEncryptedCard) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaEncryptedCard.Size(m)
}
func (m *PaymentMethodSchemaEncryptedCard) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaEncryptedCard.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaEncryptedCard proto.InternalMessageInfo

func (m *PaymentMethodSchemaEncryptedCard) GetNameOnCard() []byte {
	if m != nil {
		return m.NameOnCard
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetNumber() []byte {
	if m != nil {
		return m.Number
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetValidFromMonth() []byte {
	if m != nil {
		return m.ValidFromMonth
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetValidFromYear() []byte {
	if m != nil {
		return m.ValidFromYear
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetExpiryMonth() []byte {
	if m != nil {
		return m.ExpiryMonth
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetExpiryYear() []byte {
	if m != nil {
		return m.ExpiryYear
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetIssueNumber() []byte {
	if m != nil {
		return m.IssueNumber
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedCard) GetTransportKeyId() string {
	if m != nil {
		return m.TransportKeyId
	}
	return ""
}

type PaymentMethodSchemaCard struct {
	NameOnCard           string   `protobuf:"bytes,1,opt,name=name_on_card,json=nameOnCard,proto3" json:"name_on_card,omitempty"`
	ValidFromMonth       int32    `protobuf:"varint,2,opt,name=valid_from_month,json=validFromMonth,proto3" json:"valid_from_month,omitempty"`
	ValidFromYear        int32    `protobuf:"varint,3,opt,name=valid_from_year,json=validFromYear,proto3" json:"valid_from_year,omitempty"`
	ExpiryMonth          int32    `protobuf:"varint,4,opt,name=expiry_month,json=expiryMonth,proto3" json:"expiry_month,omitempty"`
	ExpiryYear           int32    `protobuf:"varint,5,opt,name=expiry_year,json=expiryYear,proto3" json:"expiry_year,omitempty"`
	IssueNumber          int32    `protobuf:"varint,6,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	Number               string   `protobuf:"bytes,7,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaCard) Reset()         { *m = PaymentMethodSchemaCard{} }
func (m *PaymentMethodSchemaCard) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaCard) ProtoMessage()    {}
func (*PaymentMethodSchemaCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{4}
}
func (m *PaymentMethodSchemaCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaCard.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaCard.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaCard.Merge(m, src)
}
func (m *PaymentMethodSchemaCard) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaCard.Size(m)
}
func (m *PaymentMethodSchemaCard) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaCard.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaCard proto.InternalMessageInfo

func (m *PaymentMethodSchemaCard) GetNameOnCard() string {
	if m != nil {
		return m.NameOnCard
	}
	return ""
}

func (m *PaymentMethodSchemaCard) GetValidFromMonth() int32 {
	if m != nil {
		return m.ValidFromMonth
	}
	return 0
}

func (m *PaymentMethodSchemaCard) GetValidFromYear() int32 {
	if m != nil {
		return m.ValidFromYear
	}
	return 0
}

func (m *PaymentMethodSchemaCard) GetExpiryMonth() int32 {
	if m != nil {
		return m.ExpiryMonth
	}
	return 0
}

func (m *PaymentMethodSchemaCard) GetExpiryYear() int32 {
	if m != nil {
		return m.ExpiryYear
	}
	return 0
}

func (m *PaymentMethodSchemaCard) GetIssueNumber() int32 {
	if m != nil {
		return m.IssueNumber
	}
	return 0
}

func (m *PaymentMethodSchemaCard) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type PaymentMethodSchemaEncryptedDirectDebit struct {
	AccountNumber        []byte   `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	SortCode             []byte   `protobuf:"bytes,2,opt,name=sort_code,json=sortCode,proto3" json:"sort_code,omitempty"`
	AccountHoldersName   []byte   `protobuf:"bytes,3,opt,name=account_holders_name,json=accountHoldersName,proto3" json:"account_holders_name,omitempty"`
	PayerReference       []byte   `protobuf:"bytes,4,opt,name=payer_reference,json=payerReference,proto3" json:"payer_reference,omitempty"`
	TransportKeyId       string   `protobuf:"bytes,5,opt,name=transport_key_id,json=transportKeyId,proto3" json:"transport_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaEncryptedDirectDebit) Reset() {
	*m = PaymentMethodSchemaEncryptedDirectDebit{}
}
func (m *PaymentMethodSchemaEncryptedDirectDebit) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaEncryptedDirectDebit) ProtoMessage()    {}
func (*PaymentMethodSchemaEncryptedDirectDebit) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{5}
}
func (m *PaymentMethodSchemaEncryptedDirectDebit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaEncryptedDirectDebit.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaEncryptedDirectDebit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaEncryptedDirectDebit.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaEncryptedDirectDebit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaEncryptedDirectDebit.Merge(m, src)
}
func (m *PaymentMethodSchemaEncryptedDirectDebit) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaEncryptedDirectDebit.Size(m)
}
func (m *PaymentMethodSchemaEncryptedDirectDebit) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaEncryptedDirectDebit.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaEncryptedDirectDebit proto.InternalMessageInfo

func (m *PaymentMethodSchemaEncryptedDirectDebit) GetAccountNumber() []byte {
	if m != nil {
		return m.AccountNumber
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedDirectDebit) GetSortCode() []byte {
	if m != nil {
		return m.SortCode
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedDirectDebit) GetAccountHoldersName() []byte {
	if m != nil {
		return m.AccountHoldersName
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedDirectDebit) GetPayerReference() []byte {
	if m != nil {
		return m.PayerReference
	}
	return nil
}

func (m *PaymentMethodSchemaEncryptedDirectDebit) GetTransportKeyId() string {
	if m != nil {
		return m.TransportKeyId
	}
	return ""
}

type PaymentMethodSchemaDirectDebit struct {
	AccountNumber        string   `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	SortCode             string   `protobuf:"bytes,2,opt,name=sort_code,json=sortCode,proto3" json:"sort_code,omitempty"`
	AccountHoldersName   string   `protobuf:"bytes,3,opt,name=account_holders_name,json=accountHoldersName,proto3" json:"account_holders_name,omitempty"`
	PayerReference       string   `protobuf:"bytes,4,opt,name=payer_reference,json=payerReference,proto3" json:"payer_reference,omitempty"`
	TransportKeyId       string   `protobuf:"bytes,5,opt,name=transport_key_id,json=transportKeyId,proto3" json:"transport_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethodSchemaDirectDebit) Reset()         { *m = PaymentMethodSchemaDirectDebit{} }
func (m *PaymentMethodSchemaDirectDebit) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodSchemaDirectDebit) ProtoMessage()    {}
func (*PaymentMethodSchemaDirectDebit) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf1a4ab3aa467b5, []int{6}
}
func (m *PaymentMethodSchemaDirectDebit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodSchemaDirectDebit.Unmarshal(m, b)
}
func (m *PaymentMethodSchemaDirectDebit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodSchemaDirectDebit.Marshal(b, m, deterministic)
}
func (m *PaymentMethodSchemaDirectDebit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodSchemaDirectDebit.Merge(m, src)
}
func (m *PaymentMethodSchemaDirectDebit) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodSchemaDirectDebit.Size(m)
}
func (m *PaymentMethodSchemaDirectDebit) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodSchemaDirectDebit.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodSchemaDirectDebit proto.InternalMessageInfo

func (m *PaymentMethodSchemaDirectDebit) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

func (m *PaymentMethodSchemaDirectDebit) GetSortCode() string {
	if m != nil {
		return m.SortCode
	}
	return ""
}

func (m *PaymentMethodSchemaDirectDebit) GetAccountHoldersName() string {
	if m != nil {
		return m.AccountHoldersName
	}
	return ""
}

func (m *PaymentMethodSchemaDirectDebit) GetPayerReference() string {
	if m != nil {
		return m.PayerReference
	}
	return ""
}

func (m *PaymentMethodSchemaDirectDebit) GetTransportKeyId() string {
	if m != nil {
		return m.TransportKeyId
	}
	return ""
}

func init() {
	proto.RegisterEnum("chargehive.chtype.PaymentMethodSchema", PaymentMethodSchema_name, PaymentMethodSchema_value)
	golang_proto.RegisterEnum("chargehive.chtype.PaymentMethodSchema", PaymentMethodSchema_name, PaymentMethodSchema_value)
	proto.RegisterType((*PaymentMethodSchemaApplePay)(nil), "chargehive.chtype.PaymentMethodSchemaApplePay")
	golang_proto.RegisterType((*PaymentMethodSchemaApplePay)(nil), "chargehive.chtype.PaymentMethodSchemaApplePay")
	proto.RegisterType((*PaymentMethodSchemaGooglePay)(nil), "chargehive.chtype.PaymentMethodSchemaGooglePay")
	golang_proto.RegisterType((*PaymentMethodSchemaGooglePay)(nil), "chargehive.chtype.PaymentMethodSchemaGooglePay")
	proto.RegisterType((*PaymentMethodSchemaPayPal)(nil), "chargehive.chtype.PaymentMethodSchemaPayPal")
	golang_proto.RegisterType((*PaymentMethodSchemaPayPal)(nil), "chargehive.chtype.PaymentMethodSchemaPayPal")
	proto.RegisterType((*PaymentMethodSchemaEncryptedCard)(nil), "chargehive.chtype.PaymentMethodSchemaEncryptedCard")
	golang_proto.RegisterType((*PaymentMethodSchemaEncryptedCard)(nil), "chargehive.chtype.PaymentMethodSchemaEncryptedCard")
	proto.RegisterType((*PaymentMethodSchemaCard)(nil), "chargehive.chtype.PaymentMethodSchemaCard")
	golang_proto.RegisterType((*PaymentMethodSchemaCard)(nil), "chargehive.chtype.PaymentMethodSchemaCard")
	proto.RegisterType((*PaymentMethodSchemaEncryptedDirectDebit)(nil), "chargehive.chtype.PaymentMethodSchemaEncryptedDirectDebit")
	golang_proto.RegisterType((*PaymentMethodSchemaEncryptedDirectDebit)(nil), "chargehive.chtype.PaymentMethodSchemaEncryptedDirectDebit")
	proto.RegisterType((*PaymentMethodSchemaDirectDebit)(nil), "chargehive.chtype.PaymentMethodSchemaDirectDebit")
	golang_proto.RegisterType((*PaymentMethodSchemaDirectDebit)(nil), "chargehive.chtype.PaymentMethodSchemaDirectDebit")
}

func init() {
	proto.RegisterFile("chargehive/chtype/payment_method_schema.proto", fileDescriptor_acf1a4ab3aa467b5)
}
func init() {
	golang_proto.RegisterFile("chargehive/chtype/payment_method_schema.proto", fileDescriptor_acf1a4ab3aa467b5)
}

var fileDescriptor_acf1a4ab3aa467b5 = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0xdd, 0xa4, 0x4d, 0xba, 0xb9, 0x4d, 0xbb, 0xe9, 0xb0, 0x82, 0xd0, 0x2e, 0x6d, 0x36, 0x2c,
	0xbb, 0x15, 0x52, 0x1b, 0x24, 0xf8, 0x01, 0xaf, 0x6d, 0xda, 0x88, 0x26, 0xb5, 0xbc, 0x61, 0xa5,
	0xa0, 0x95, 0x46, 0x13, 0xfb, 0x36, 0xb1, 0x6a, 0x7b, 0xac, 0xf1, 0x24, 0x5d, 0xff, 0x03, 0xef,
	0x3c, 0xf1, 0x0b, 0xfc, 0x02, 0x3c, 0xf3, 0xc4, 0x2f, 0x00, 0x4f, 0xbc, 0xf2, 0x05, 0xc8, 0x33,
	0x4e, 0x1b, 0xb4, 0xa6, 0x2d, 0xfb, 0xe6, 0x7b, 0xee, 0xd1, 0xb9, 0xf7, 0x1c, 0x5f, 0x69, 0xe0,
	0xc8, 0x9b, 0x31, 0x31, 0xc5, 0x59, 0xb0, 0xc0, 0x9e, 0x37, 0x93, 0x59, 0x82, 0xbd, 0x84, 0x65,
	0x11, 0xc6, 0x92, 0x46, 0x28, 0x67, 0xdc, 0xa7, 0xa9, 0x37, 0xc3, 0x88, 0x1d, 0x27, 0x82, 0x4b,
	0x4e, 0x76, 0x6e, 0xe8, 0xc7, 0x9a, 0xbe, 0x7b, 0x34, 0x0d, 0xe4, 0x6c, 0x3e, 0x39, 0xf6, 0x78,
	0xd4, 0x9b, 0xf2, 0x29, 0xef, 0x29, 0xe6, 0x64, 0x7e, 0xa1, 0x2a, 0x55, 0xa8, 0x2f, 0xad, 0xd0,
	0xfd, 0xb1, 0x02, 0x7b, 0x8e, 0x9e, 0x30, 0x50, 0x03, 0x5e, 0x29, 0x7d, 0x23, 0x49, 0x42, 0x74,
	0x58, 0x46, 0x0e, 0x60, 0x33, 0xc5, 0x34, 0x0d, 0x78, 0x4c, 0xe7, 0x22, 0x6c, 0x57, 0x3a, 0x95,
	0xc3, 0x86, 0x0b, 0x05, 0xf4, 0xad, 0x08, 0xc9, 0x11, 0x90, 0x20, 0x0e, 0x64, 0xc0, 0x64, 0xb0,
	0x40, 0xea, 0xf1, 0x58, 0xe2, 0x5b, 0xd9, 0xae, 0x2a, 0xde, 0xce, 0x4d, 0xc7, 0xd4, 0x0d, 0xf2,
	0x18, 0x6a, 0x92, 0x5f, 0x62, 0xdc, 0x5e, 0x53, 0x0c, 0x5d, 0x90, 0x36, 0x6c, 0x14, 0x92, 0xed,
	0x75, 0x85, 0x2f, 0xcb, 0xee, 0x4f, 0x15, 0x78, 0x52, 0xb2, 0xdf, 0x09, 0xe7, 0x53, 0xbd, 0xe0,
	0x73, 0x78, 0x14, 0xf2, 0x2b, 0xba, 0x60, 0xe1, 0x1c, 0xa9, 0x96, 0xd6, 0x4b, 0x6e, 0x85, 0xfc,
	0xea, 0x75, 0x8e, 0x8e, 0xd4, 0x88, 0x0e, 0x6c, 0xfa, 0x98, 0x7a, 0x22, 0x48, 0x64, 0x3e, 0x46,
	0x2f, 0xb8, 0x0a, 0x91, 0xa7, 0xd0, 0xf4, 0x98, 0xf0, 0xa9, 0x8f, 0x92, 0x05, 0x61, 0x5a, 0x6c,
	0xb8, 0x99, 0x63, 0x96, 0x86, 0xae, 0x29, 0x31, 0xca, 0x2b, 0x2e, 0x2e, 0x8b, 0x65, 0x15, 0x65,
	0xa8, 0xa1, 0xee, 0x1e, 0x7c, 0x5c, 0xb2, 0xaf, 0xc3, 0x32, 0x87, 0x85, 0xdd, 0x9f, 0xab, 0xd0,
	0x29, 0xe9, 0xda, 0xb1, 0x27, 0xb2, 0x44, 0xa2, 0x6f, 0x32, 0xe1, 0x93, 0x0e, 0x34, 0x63, 0x16,
	0x21, 0xe5, 0x31, 0xcd, 0x85, 0x95, 0x9d, 0xa6, 0x0b, 0x39, 0x76, 0x1e, 0x2b, 0xc6, 0x87, 0x50,
	0x8f, 0xe7, 0xd1, 0x04, 0x85, 0xb2, 0xd1, 0x74, 0x8b, 0x8a, 0x1c, 0x42, 0x6b, 0xc1, 0xc2, 0xc0,
	0xa7, 0x17, 0x82, 0x47, 0x34, 0xe2, 0xb1, 0x9c, 0x29, 0x17, 0x4d, 0x77, 0x5b, 0xe1, 0x5f, 0x0b,
	0x1e, 0x0d, 0x72, 0x34, 0x4f, 0x6d, 0x85, 0x99, 0x21, 0x13, 0xca, 0x4b, 0xd3, 0xdd, 0xba, 0x26,
	0x8e, 0x91, 0x89, 0xdc, 0x30, 0xbe, 0x4d, 0x02, 0x91, 0x15, 0x6a, 0x35, 0x45, 0xda, 0xd4, 0x98,
	0x96, 0x3a, 0x80, 0xa2, 0xd4, 0x32, 0x75, 0xbd, 0xad, 0x86, 0x96, 0x1a, 0x41, 0x9a, 0xce, 0x91,
	0x16, 0x3b, 0x6f, 0x68, 0x0d, 0x85, 0x0d, 0xaf, 0x17, 0x97, 0x82, 0xc5, 0x69, 0xc2, 0x85, 0xa4,
	0x97, 0x98, 0xd1, 0xc0, 0x6f, 0x3f, 0x54, 0xd9, 0x6e, 0x5f, 0xe3, 0xdf, 0x60, 0xd6, 0xf7, 0xbb,
	0xdf, 0x57, 0xe1, 0xa3, 0x92, 0x04, 0xff, 0x33, 0xb8, 0xc6, 0xbf, 0x82, 0x2b, 0x0b, 0x28, 0x8f,
	0xb0, 0x76, 0x9f, 0x80, 0xd6, 0x14, 0xf1, 0x8e, 0x80, 0xd6, 0x15, 0xe9, 0xb6, 0x80, 0x6a, 0x8a,
	0x71, 0x5b, 0x40, 0x75, 0xad, 0xb1, 0x1a, 0xd0, 0xcd, 0x1f, 0xdf, 0x50, 0xa6, 0x8a, 0xaa, 0xfb,
	0x77, 0x05, 0x5e, 0xdc, 0x76, 0x50, 0x56, 0x20, 0xd0, 0x93, 0x16, 0x4e, 0x02, 0x49, 0x3e, 0x83,
	0x6d, 0xe6, 0x79, 0x7c, 0x1e, 0xcb, 0xe5, 0x20, 0x7d, 0x59, 0x5b, 0x05, 0x5a, 0x8c, 0xda, 0x83,
	0x46, 0x9a, 0xff, 0x06, 0x8f, 0xfb, 0x58, 0xdc, 0xd7, 0xc3, 0x1c, 0x30, 0xb9, 0x8f, 0xe4, 0x0b,
	0x78, 0xbc, 0xd4, 0x98, 0xf1, 0xd0, 0x47, 0x91, 0xd2, 0x3c, 0xde, 0xe2, 0xca, 0x48, 0xd1, 0x3b,
	0xd5, 0xad, 0x21, 0x8b, 0x90, 0xbc, 0x80, 0x47, 0x09, 0xcb, 0x50, 0x50, 0x81, 0x17, 0x28, 0x30,
	0xf6, 0xb0, 0xb8, 0xb4, 0x6d, 0x05, 0xbb, 0x4b, 0xb4, 0xf4, 0x06, 0x6a, 0xa5, 0x37, 0xf0, 0x57,
	0x05, 0xf6, 0x4b, 0x4c, 0xdf, 0xed, 0xb5, 0x71, 0xa7, 0xd7, 0xc6, 0x3d, 0xbd, 0x36, 0xfe, 0x8f,
	0xd7, 0xc6, 0xfb, 0x7b, 0xfd, 0xfc, 0x87, 0x2a, 0x7c, 0x50, 0xe2, 0x95, 0x3c, 0x85, 0x4f, 0x1c,
	0x63, 0x3c, 0xb0, 0x87, 0x23, 0x3a, 0xb0, 0x47, 0xa7, 0xe7, 0x16, 0x7d, 0x65, 0x9e, 0xda, 0x03,
	0x83, 0xf6, 0x87, 0xaf, 0x8d, 0xb3, 0xbe, 0xd5, 0x7a, 0x40, 0x0e, 0xe1, 0x59, 0x39, 0xc5, 0x1e,
	0x9a, 0xee, 0xd8, 0x19, 0xd9, 0x16, 0x35, 0x0d, 0xd7, 0x6a, 0x55, 0xc8, 0x3e, 0xec, 0x96, 0x33,
	0x55, 0xbf, 0x4a, 0x3e, 0x85, 0x83, 0xf2, 0xbe, 0xe1, 0x38, 0x67, 0x36, 0x75, 0x8c, 0x71, 0x6b,
	0x8d, 0x3c, 0x83, 0x4e, 0x39, 0xe9, 0xe4, 0xfc, 0xfc, 0xa4, 0x60, 0xad, 0x93, 0x0e, 0x3c, 0x29,
	0x67, 0x39, 0xc6, 0xd8, 0x31, 0xce, 0x5a, 0x35, 0xf2, 0x1c, 0xba, 0xe5, 0x0c, 0xab, 0xef, 0xda,
	0xe6, 0x88, 0x5a, 0xf6, 0xcb, 0xfe, 0xa8, 0x55, 0x7f, 0xc9, 0x7f, 0xfb, 0x7d, 0xff, 0xc1, 0x2f,
	0x7f, 0xee, 0x57, 0xbe, 0xfb, 0x6a, 0xe5, 0xc9, 0x5b, 0x79, 0x3f, 0x97, 0x6f, 0x5d, 0xc8, 0xe2,
	0x69, 0xef, 0x9d, 0x77, 0xf5, 0xd7, 0xea, 0x8e, 0xa9, 0xb0, 0xd3, 0x60, 0x81, 0x6f, 0x4c, 0x85,
	0xfd, 0x51, 0xdd, 0x7d, 0x07, 0x7b, 0x33, 0x40, 0xc9, 0x7c, 0x26, 0xd9, 0xa4, 0xae, 0x04, 0xbf,
	0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa0, 0x68, 0xc0, 0xa4, 0x07, 0x00, 0x00,
}
