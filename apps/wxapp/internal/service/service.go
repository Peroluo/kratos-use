package service

import (
	"context"
	"fmt"

	pb "app/apps/wxapp/api"
	"app/apps/wxapp/internal/dao"

	"github.com/bilibili/kratos/pkg/conf/paladin"

	common_service_v1 "app/apps/common/api"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, wire.Bind(new(pb.WxServiceServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// GetUser ...
func (s *Service) GetUser(ctx context.Context, req *common_service_v1.CommonReq) (resp *common_service_v1.CommonRes, err error) {
	resp = &common_service_v1.CommonRes{Data: "2323"}
	fmt.Println(resp)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
