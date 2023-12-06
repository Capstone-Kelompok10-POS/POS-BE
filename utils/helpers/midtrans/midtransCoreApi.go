package midtrans

import (
	"qbills/configs"
	"qbills/models/domain"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransCoreApi interface {
	ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	CheckTransactionStatus(invoice string) (string, *domain.PaymentTransactionStatus, error)
}

type MidtransCoreApiImpl struct {
	Client coreapi.Client
}

func NewMidtransCoreApi(config *configs.MidtransConfig) MidtransCoreApi {
	client := coreapi.Client{}
	client.New(config.ServerKey, midtrans.Sandbox)

	return &MidtransCoreApiImpl{
		Client: client,
	}
}
func (coreApi *MidtransCoreApiImpl) ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	response, err := coreApi.Client.ChargeTransaction(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (coreApi *MidtransCoreApiImpl) CheckTransactionStatus(orderId string) (string, *domain.PaymentTransactionStatus, error) {
	response, err := coreApi.Client.CheckTransaction(orderId)

	parseSettlementTime, _ := time.Parse("2006-01-02 15:04:05", response.SettlementTime)
	
	if err != nil {
		return "", nil, err
	} else {
		if response != nil {
			transactionStatus := domain.PaymentTransactionStatus{
				OrderID: response.OrderID,
				TransactionStatus: response.TransactionStatus,
				FraudStatus: response.FraudStatus,
				SettlementTime: parseSettlementTime,
			}

			if transactionStatus.TransactionStatus == "capture" {
				if transactionStatus.FraudStatus == "challenge" {
					return "challenge", &transactionStatus, nil
				} else if transactionStatus.FraudStatus == "accept" {
					return "success", &transactionStatus, nil
				}
			} else if transactionStatus.TransactionStatus == "settlement" {
				return "success", &transactionStatus, nil
			} else if transactionStatus.TransactionStatus == "deny" {
				return "deny", &transactionStatus, nil
			} else if transactionStatus.TransactionStatus == "cancel" || transactionStatus.TransactionStatus == "expire" {
				return "failure", &transactionStatus, nil
			} else if transactionStatus.TransactionStatus == "pending" {
				return "pending", &transactionStatus, nil
			}
		}
	}

	return "", nil, nil
}