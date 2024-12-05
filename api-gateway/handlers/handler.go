package handlers

import pb "api-gateway/pb/generated"

type Handlers struct {
	DateClient    pb.SwipeServiceClient
	PaymentClient pb.SubPaymentClient
	ProfileClient pb.ProfileServiceClient
	UserClient    pb.UserServiceClient
}
