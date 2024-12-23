// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: subs-payment.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tier          string  `protobuf:"bytes,2,opt,name=tier,proto3" json:"tier,omitempty"`
	PricePerMonth float32 `protobuf:"fixed32,3,opt,name=price_per_month,json=pricePerMonth,proto3" json:"price_per_month,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subscription) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *Subscription) GetPricePerMonth() float32 {
	if x != nil {
		return x.PricePerMonth
	}
	return 0
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentGateway  string                 `protobuf:"bytes,3,opt,name=payment_gateway,json=paymentGateway,proto3" json:"payment_gateway,omitempty"`
	Amount          float32                `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency        string                 `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	TransactionDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"` // can be null
	Status          string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Url             string                 `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Payment) GetPaymentGateway() string {
	if x != nil {
		return x.PaymentGateway
	}
	return ""
}

func (x *Payment) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Payment) GetTransactionDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionDate
	}
	return nil
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Payment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UserSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Subscription *Subscription          `protobuf:"bytes,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Duration     int64                  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"` // can be null
	Payment      *Payment               `protobuf:"bytes,6,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *UserSubscription) Reset() {
	*x = UserSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSubscription) ProtoMessage() {}

func (x *UserSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSubscription.ProtoReflect.Descriptor instead.
func (*UserSubscription) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{2}
}

func (x *UserSubscription) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserSubscription) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserSubscription) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

func (x *UserSubscription) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UserSubscription) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *UserSubscription) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type CreateUserSubcriptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Tier     string `protobuf:"bytes,2,opt,name=tier,proto3" json:"tier,omitempty"`
	Duration int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *CreateUserSubcriptionReq) Reset() {
	*x = CreateUserSubcriptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserSubcriptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserSubcriptionReq) ProtoMessage() {}

func (x *CreateUserSubcriptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserSubcriptionReq.ProtoReflect.Descriptor instead.
func (*CreateUserSubcriptionReq) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserSubcriptionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserSubcriptionReq) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *CreateUserSubcriptionReq) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type CreateUserSubcriptionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Subscription *Subscription          `protobuf:"bytes,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Duration     int64                  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Payment      *Payment               `protobuf:"bytes,6,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *CreateUserSubcriptionResp) Reset() {
	*x = CreateUserSubcriptionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserSubcriptionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserSubcriptionResp) ProtoMessage() {}

func (x *CreateUserSubcriptionResp) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserSubcriptionResp.ProtoReflect.Descriptor instead.
func (*CreateUserSubcriptionResp) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserSubcriptionResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateUserSubcriptionResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserSubcriptionResp) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

func (x *CreateUserSubcriptionResp) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CreateUserSubcriptionResp) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *CreateUserSubcriptionResp) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type GetUserSubcriptionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserSubcriptionsReq) Reset() {
	*x = GetUserSubcriptionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSubcriptionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSubcriptionsReq) ProtoMessage() {}

func (x *GetUserSubcriptionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSubcriptionsReq.ProtoReflect.Descriptor instead.
func (*GetUserSubcriptionsReq) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserSubcriptionsReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserSubcriptionsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSubscriptions []*UserSubscription `protobuf:"bytes,1,rep,name=user_subscriptions,json=userSubscriptions,proto3" json:"user_subscriptions,omitempty"`
}

func (x *GetUserSubcriptionsResp) Reset() {
	*x = GetUserSubcriptionsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSubcriptionsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSubcriptionsResp) ProtoMessage() {}

func (x *GetUserSubcriptionsResp) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSubcriptionsResp.ProtoReflect.Descriptor instead.
func (*GetUserSubcriptionsResp) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserSubcriptionsResp) GetUserSubscriptions() []*UserSubscription {
	if x != nil {
		return x.UserSubscriptions
	}
	return nil
}

type CompletePaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallbackToken string  `protobuf:"bytes,1,opt,name=callback_token,json=callbackToken,proto3" json:"callback_token,omitempty"`
	Id            string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId    string  `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	PaymentMethod string  `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	PaidAmount    float32 `protobuf:"fixed32,5,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	Status        string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	PaidAt        string  `protobuf:"bytes,7,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
}

func (x *CompletePaymentReq) Reset() {
	*x = CompletePaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePaymentReq) ProtoMessage() {}

func (x *CompletePaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePaymentReq.ProtoReflect.Descriptor instead.
func (*CompletePaymentReq) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{7}
}

func (x *CompletePaymentReq) GetCallbackToken() string {
	if x != nil {
		return x.CallbackToken
	}
	return ""
}

func (x *CompletePaymentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompletePaymentReq) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *CompletePaymentReq) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *CompletePaymentReq) GetPaidAmount() float32 {
	if x != nil {
		return x.PaidAmount
	}
	return 0
}

func (x *CompletePaymentReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CompletePaymentReq) GetPaidAt() string {
	if x != nil {
		return x.PaidAt
	}
	return ""
}

type CompletePaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *CompletePaymentResp) Reset() {
	*x = CompletePaymentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePaymentResp) ProtoMessage() {}

func (x *CompletePaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePaymentResp.ProtoReflect.Descriptor instead.
func (*CompletePaymentResp) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{8}
}

func (x *CompletePaymentResp) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type GetPaymentByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId int64 `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *GetPaymentByIDReq) Reset() {
	*x = GetPaymentByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentByIDReq) ProtoMessage() {}

func (x *GetPaymentByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentByIDReq.ProtoReflect.Descriptor instead.
func (*GetPaymentByIDReq) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{9}
}

func (x *GetPaymentByIDReq) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

type GetPaymentByIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *GetPaymentByIDResp) Reset() {
	*x = GetPaymentByIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subs_payment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentByIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentByIDResp) ProtoMessage() {}

func (x *GetPaymentByIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_subs_payment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentByIDResp.ProtoReflect.Descriptor instead.
func (*GetPaymentByIDResp) Descriptor() ([]byte, []int) {
	return file_subs_payment_proto_rawDescGZIP(), []int{10}
}

func (x *GetPaymentByIDResp) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

var File_subs_payment_proto protoreflect.FileDescriptor

var file_subs_payment_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x2d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x80,
	0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x45, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0xfd, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x63, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x02, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x67, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a,
	0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x75, 0x62, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x12,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61,
	0x69, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x69,
	0x64, 0x41, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x75,
	0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x44,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x32, 0xff, 0x02, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x73,
	0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x60, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x20, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subs_payment_proto_rawDescOnce sync.Once
	file_subs_payment_proto_rawDescData = file_subs_payment_proto_rawDesc
)

func file_subs_payment_proto_rawDescGZIP() []byte {
	file_subs_payment_proto_rawDescOnce.Do(func() {
		file_subs_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_subs_payment_proto_rawDescData)
	})
	return file_subs_payment_proto_rawDescData
}

var file_subs_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_subs_payment_proto_goTypes = []any{
	(*Subscription)(nil),              // 0: sub_payment.Subscription
	(*Payment)(nil),                   // 1: sub_payment.Payment
	(*UserSubscription)(nil),          // 2: sub_payment.UserSubscription
	(*CreateUserSubcriptionReq)(nil),  // 3: sub_payment.CreateUserSubcriptionReq
	(*CreateUserSubcriptionResp)(nil), // 4: sub_payment.CreateUserSubcriptionResp
	(*GetUserSubcriptionsReq)(nil),    // 5: sub_payment.GetUserSubcriptionsReq
	(*GetUserSubcriptionsResp)(nil),   // 6: sub_payment.GetUserSubcriptionsResp
	(*CompletePaymentReq)(nil),        // 7: sub_payment.CompletePaymentReq
	(*CompletePaymentResp)(nil),       // 8: sub_payment.CompletePaymentResp
	(*GetPaymentByIDReq)(nil),         // 9: sub_payment.GetPaymentByIDReq
	(*GetPaymentByIDResp)(nil),        // 10: sub_payment.GetPaymentByIDResp
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
}
var file_subs_payment_proto_depIdxs = []int32{
	11, // 0: sub_payment.Payment.transaction_date:type_name -> google.protobuf.Timestamp
	0,  // 1: sub_payment.UserSubscription.subscription:type_name -> sub_payment.Subscription
	11, // 2: sub_payment.UserSubscription.end_date:type_name -> google.protobuf.Timestamp
	1,  // 3: sub_payment.UserSubscription.payment:type_name -> sub_payment.Payment
	0,  // 4: sub_payment.CreateUserSubcriptionResp.subscription:type_name -> sub_payment.Subscription
	11, // 5: sub_payment.CreateUserSubcriptionResp.end_date:type_name -> google.protobuf.Timestamp
	1,  // 6: sub_payment.CreateUserSubcriptionResp.payment:type_name -> sub_payment.Payment
	2,  // 7: sub_payment.GetUserSubcriptionsResp.user_subscriptions:type_name -> sub_payment.UserSubscription
	1,  // 8: sub_payment.CompletePaymentResp.payment:type_name -> sub_payment.Payment
	1,  // 9: sub_payment.GetPaymentByIDResp.payment:type_name -> sub_payment.Payment
	3,  // 10: sub_payment.SubPayment.CreateUserSubcription:input_type -> sub_payment.CreateUserSubcriptionReq
	5,  // 11: sub_payment.SubPayment.GetUserSubcriptions:input_type -> sub_payment.GetUserSubcriptionsReq
	7,  // 12: sub_payment.SubPayment.CompletePayment:input_type -> sub_payment.CompletePaymentReq
	9,  // 13: sub_payment.SubPayment.GetPaymentByID:input_type -> sub_payment.GetPaymentByIDReq
	4,  // 14: sub_payment.SubPayment.CreateUserSubcription:output_type -> sub_payment.CreateUserSubcriptionResp
	6,  // 15: sub_payment.SubPayment.GetUserSubcriptions:output_type -> sub_payment.GetUserSubcriptionsResp
	8,  // 16: sub_payment.SubPayment.CompletePayment:output_type -> sub_payment.CompletePaymentResp
	10, // 17: sub_payment.SubPayment.GetPaymentByID:output_type -> sub_payment.GetPaymentByIDResp
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_subs_payment_proto_init() }
func file_subs_payment_proto_init() {
	if File_subs_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subs_payment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Subscription); i {
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
		file_subs_payment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Payment); i {
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
		file_subs_payment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserSubscription); i {
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
		file_subs_payment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserSubcriptionReq); i {
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
		file_subs_payment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserSubcriptionResp); i {
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
		file_subs_payment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserSubcriptionsReq); i {
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
		file_subs_payment_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserSubcriptionsResp); i {
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
		file_subs_payment_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CompletePaymentReq); i {
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
		file_subs_payment_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CompletePaymentResp); i {
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
		file_subs_payment_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetPaymentByIDReq); i {
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
		file_subs_payment_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetPaymentByIDResp); i {
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
			RawDescriptor: file_subs_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subs_payment_proto_goTypes,
		DependencyIndexes: file_subs_payment_proto_depIdxs,
		MessageInfos:      file_subs_payment_proto_msgTypes,
	}.Build()
	File_subs_payment_proto = out.File
	file_subs_payment_proto_rawDesc = nil
	file_subs_payment_proto_goTypes = nil
	file_subs_payment_proto_depIdxs = nil
}
