package service

import (
	midtransrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/midtrans_request"
	midtranspkg "github.com/alifdwt/haiwan-go/pkg/midtrans_pkg"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransService struct {
	snapClient *midtranspkg.SnapClient
}

func NewMidtransService(snapClient *midtranspkg.SnapClient) *midtransService {
	return &midtransService{
		snapClient: snapClient,
	}
}

func (s *midtransService) CreateTransaction(request *midtransrequest.CreateMidtransRequest) (*snap.Response, error) {
	res, err := s.snapClient.CreateTransaction(*request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
