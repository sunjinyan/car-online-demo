package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Mongo *dao.Mongo
	Logger *zap.Logger
	rentalpb.UnimplementedProfileServiceServer
}


func (s *Service) GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (s *Service)  SubmitProfile(ctx context.Context,iden *rentalpb.Identity) (*rentalpb.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProfile not implemented")
}
func (s *Service)  ClearProfile(ctx context.Context,cpr *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearProfile not implemented")
}
func (s *Service)  GetProfilePhoto(ctx context.Context,gpr *rentalpb.GetProfilePhotoRequest) (*rentalpb.GetProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfilePhoto not implemented")
}
func (s *Service)  CreateProfilePhoto(ctx context.Context,cpr *rentalpb.CreateProfilePhotoRequest) (*rentalpb.CreateProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfilePhoto not implemented")
}
func (s *Service)  CompleteProfilePhoto(ctx context.Context,cpr *rentalpb.CompleteProfilePhotoRequest) (*rentalpb.Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteProfilePhoto not implemented")
}
func (s *Service)  ClearProfilePhoto(ctx context.Context,cpr *rentalpb.ClearProfilePhotoRequest) (*rentalpb.ClearProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearProfilePhoto not implemented")
}
func (s *Service)  UploadFilePhoto(ctx context.Context,ufr *rentalpb.UploadFilePhotoRequest) (*rentalpb.UploadFilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFilePhoto not implemented")
}