package midtranspkg

import (
	midtransrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/midtrans_request"
	"github.com/alifdwt/haiwan-go/pkg/dotenv"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type SnapClient struct {
	client snap.Client
}

func setupGlobalMidtransConfig() {
	config, err := dotenv.LoadConfig("../..")
	if err != nil {
		log.Error("cannot load .env file", err)
	}

	midtrans.ClientKey = config.MIDTRANS_CLIENT_KEY
	midtrans.ServerKey = config.MIDTRANS_SERVER_KEY
	midtrans.Environment = midtrans.Sandbox
}

func NewSnapClient() *SnapClient {
	setupGlobalMidtransConfig()

	var s snap.Client

	s.New(midtrans.ServerKey, midtrans.Sandbox)

	return &SnapClient{s}
}

func (s *SnapClient) CreateTransaction(request midtransrequest.CreateMidtransRequest) (*snap.Response, error) {
	params := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "order-csb-" + uuid.New().String(),
			GrossAmt: int64(request.GrossAmount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: false,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.FirstName,
			LName: request.LastName,
			Email: request.Email,
			Phone: request.Phone,
		},
		Callbacks: &snap.Callbacks{
			Finish: "http://localhost:300/user/order",
		},
	}

	res, err := snap.CreateTransaction(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
