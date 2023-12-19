package request

import (
	"fmt"
	"log"
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func TransactionCreateRequestToTransactionDomain(request web.TransactionCreateRequest, create web.TransactionCreate, detail web.TransactionDetailCreate) *domain.Transaction {
	transaction := &domain.Transaction{
		CashierID: request.CashierID,
		MembershipID:       request.MembershipID,
		ConvertPointID :    request.ConvertPointID,
		Discount: 			create.Discount,
		TotalPrice: 		create.TotalPrice,
		Tax: 				create.Tax,
		TotalPayment: 		create.TotalPayment,
		Details:   			make([]domain.TransactionDetail, 0),
		TransactionPayment: request.TransactionPayment,
	}
	for _, detailRequest:= range request.Details {
		productDetailId := detailRequest.ProductDetailID
		productPrice, priceExists := detail.ProductPrice[productDetailId]
		subTotal, subTotalExists := detail.SubTotal[productDetailId]

		if !priceExists || !subTotalExists {
			log.Printf("ProductDetailID %d not found in ProductPrice or SubTotal", productDetailId)
			return nil
		}

		transactionDetail := &domain.TransactionDetail{
			ProductDetailID: detailRequest.ProductDetailID,
			Price: productPrice,
			SubTotal: subTotal,
			Quantity:  detailRequest.Quantity,
			Notes: detailRequest.Notes,
		}
		transaction.Details = append(transaction.Details, *transactionDetail)
	}
	return transaction
}


func TransactionDomainToTransactionSchema(request domain.Transaction) *schema.Transaction {
	transaction := &schema.Transaction{
		CashierID: request.CashierID,
		MembershipID:       request.MembershipID,
		ConvertPointID :    request.ConvertPointID,
		Discount: request.Discount,
		TotalPrice: request.TotalPrice,
		Tax: request.Tax,
		TotalPayment: request.TotalPayment,
	}

	for _, detailRequest:= range request.Details {
		transactionDetail := schema.TransactionDetail{
			ProductDetailID: detailRequest.ProductDetailID,
			Quantity:  detailRequest.Quantity,
			Notes: detailRequest.Notes,
		}
		transaction.Details = append(transaction.Details, transactionDetail)
	}
	return transaction
}

func CreateTransactionPaymentRequestToMidtransChargeRequest(transaction *domain.Transaction, membership *domain.Membership, paymentMethod *domain.PaymentMethod, discount, tax float64) *coreapi.ChargeReq{
	var midtransItemDetails []midtrans.ItemDetails
	for _, detailRequest:= range transaction.Details {
		item := midtrans.ItemDetails{
			ID: fmt.Sprintf("PRODUCT-%d", detailRequest.ProductDetailID),
			Name: "Product Qbills",
			Price: int64(detailRequest.Price),
			Qty: int32(detailRequest.Quantity),
			Brand: "QBILLS",
			MerchantName: "QBILLS",
		}
		midtransItemDetails = append(midtransItemDetails, item)

	}
	midtransItemDetails = append(midtransItemDetails, midtrans.ItemDetails{
		ID: "Tax",
		Name: "Tax",
		Price: int64(tax),
		Qty: 1,
		Brand: "QBILLS",
		MerchantName: "QBILLS",
	})

	midtransItemDetails = append(midtransItemDetails, midtrans.ItemDetails{
		ID: "Discount",
		Name: "Discount",
		Price: int64(-discount),
		Qty: 1,
		Brand: "QBILLS",
		MerchantName: "QBILLS",
	})

	chargeRequest := coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.TransactionPayment.Invoice,
			GrossAmt: int64(transaction.TotalPayment), // Ubah tipe data ke int64
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: membership.Name,
			Phone: membership.PhoneNumber,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(paymentMethod.Name),

		},
		Items: &midtransItemDetails, // Items bukan pointer
	}
	return &chargeRequest
}

func ChargeResponseToTransactionPayment(response *coreapi.ChargeResponse, transaction *domain.Transaction) *domain.Transaction {
	parseTransactionTime, _ := time.Parse("2006-01-02 15:04:05", response.TransactionTime)
	var vaNumber string
	if transaction.TransactionPayment.PaymentMethodID == 3 {
		vaNumber = response.PermataVaNumber
	} else {
		vaNumber = response.VaNumbers[0].VANumber
	}
	fmt.Println(response, "ini midtrans")
	return &domain.Transaction{
		UpdatedAt: parseTransactionTime,
		CashierID: transaction.CashierID,
		MembershipID:       transaction.MembershipID,
		ConvertPointID :    transaction.ConvertPointID,
		Discount: 			transaction.Discount,
		TotalPrice: 		transaction.TotalPrice,
		Tax: 				transaction.Tax,
		TotalPayment: 		transaction.TotalPayment,
		Details:   			transaction.Details,
		TransactionPayment: domain.TransactionPayment{
			ID: transaction.TransactionPayment.ID,
			TransactionID: transaction.TransactionPayment.TransactionID,
			CreatedAt: parseTransactionTime,
			UpdatedAt: parseTransactionTime,
			PaymentMethodID: transaction.TransactionPayment.PaymentMethodID,
			PaymentMethod: transaction.TransactionPayment.PaymentMethod,
			Invoice: transaction.TransactionPayment.Invoice,
			PaymentStatus: response.TransactionStatus,
			VANumber: vaNumber,
		},
	}
}