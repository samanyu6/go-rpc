package server

import (
	"context"
	"go-rpc/db"
	"go-rpc/userpb"

	"google.golang.org/grpc"
)

var GrpcIdService *grpc.Server

type IdService struct {
	userpb.UnimplementedIdServiceServer
}

func init() {
	GrpcIdService = grpc.NewServer()
	userpb.RegisterIdServiceServer(GrpcIdService, &IdService{})
}

func (s *IdService) GetUser(ctx context.Context, request *userpb.GetId) (*userpb.IdResponse, error) {
	id := int(request.Id)
	userModelFromId := db.GetId(id)
	responseData := userpb.UserModel{
		Id:      int32(userModelFromId.Id),
		City:    userModelFromId.City,
		Fname:   userModelFromId.Fname,
		Phone:   int64(userModelFromId.Phone),
		Height:  float32(userModelFromId.Height),
		Married: userModelFromId.Married,
	}

	response := &userpb.IdResponse{
		UserModel: &responseData,
	}

	return response, nil
}

func (s *IdService) GetUserList(ctx context.Context, request *userpb.GetIdList) (*userpb.IdListResponse, error) {
	var idList []int
	var responseIdList []*userpb.UserModel

	for _, data := range request.Ids {
		idList = append(idList, int(data))
	}

	userModelForIds := db.GetList(idList)

	for _, data := range userModelForIds {
		resp := userpb.UserModel{
			Id:      int32(data.Id),
			Fname:   data.Fname,
			City:    data.City,
			Phone:   int64(data.Phone),
			Height:  float32(data.Height),
			Married: data.Married,
		}

		responseIdList = append(responseIdList, &resp)
	}

	response := &userpb.IdListResponse{
		UserModelArr: responseIdList,
	}

	return response, nil
}

func GetIdServiceInstance() *grpc.Server {
	return GrpcIdService
}
