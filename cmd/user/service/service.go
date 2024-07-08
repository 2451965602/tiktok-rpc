package service

import (
	"context"
	"strconv"
	"tiktokrpc/cmd/user/dal/db"
	milvus "tiktokrpc/cmd/user/dal/miluvs"
	"tiktokrpc/cmd/user/pkg/ai/rpc"
	"tiktokrpc/cmd/user/pkg/errmsg"
	"tiktokrpc/kitex_gen/user"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) Register(req *user.RegisterRequest) (*db.User, error) {
	return db.CreateUser(s.ctx, req.Username, req.Password)
}

func (s *UserService) Login(req *user.LoginRequest) (*db.UserInfoDetail, error) {

	return db.LoginCheck(s.ctx, req)
}

func (s *UserService) GetInfo(req *user.InfoRequest) (*db.UserInfoDetail, error) {
	return db.GetInfo(s.ctx, req.UserId)
}

func (s *UserService) GetInfoByName(req *user.NameToInfoRequest) (*db.UserInfoDetail, error) {
	return db.GetInfoByName(s.ctx, req.UserName)
}

func (s *UserService) UploadAvatar(req *user.UploadRequest) (*db.User, error) {
	return db.UploadAvatar(s.ctx, strconv.FormatInt(req.UserId, 10), req.AvatarUrl)
}

func (s *UserService) MFAGet(req *user.MFAGetRequest) (*db.MFA, error) {
	return db.MFAGet(s.ctx, strconv.FormatInt(req.UserId, 10))
}

func (s *UserService) MFABind(req *user.MFABindRequest) error {
	return db.MFABind(s.ctx, strconv.FormatInt(req.UserId, 10), req.Secret, req.Code)
}

func (s *UserService) MFAStatus(req *user.MFAStatusRequest) error {

	if req.ActionType != "0" && req.ActionType != "1" {
		return errmsg.IllegalParamError
	}

	err := db.MFAStatus(s.ctx, strconv.FormatInt(req.UserId, 10), req.Code, req.ActionType)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) AiUploadImages(req *user.UploadImagesRequest) error {

	if err := milvus.CreateCollection(s.ctx, req); err != nil {
		return err
	}

	vector, err := rpc.GerVector(req.ImgPath)
	if err != nil {
		return err
	}
	if err := milvus.InsertData(s.ctx, vector, req.ImgUrl, req.CollectionName); err != nil {
		return err
	}

	return nil
}

func (s *UserService) AiSearchImages(req *user.SearchImagesRequest) (string, error) {

	vector, err := rpc.GerVector(req.ImgPath)
	if err != nil {
		return "", err
	}

	url, err := milvus.Search(vector, req)
	if err != nil {
		return "", err
	}

	_ = milvus.InsertData(s.ctx, vector, req.ImgUrl, req.CollectionName)

	return url, nil
}
