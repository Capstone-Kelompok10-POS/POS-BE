package request

import (
	"log"
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func TransactionCreateRequestToTransactionDomain(request web.TransactionCreateRequest, create web.TransactionCreate, detail web.TransactionDetailCreate) *domain.Transaction {
	transaction := &domain.Transaction{
		CashierID: request.CashierID,
		MembershipID:       request.MembershipID,
		ConvertPointID :    request.ConvertPointID,
		Status:             create.Status,
		Discount: 			create.Discount,
		TotalPrice: 		create.TotalPrice,
		Tax: 				create.Tax,
		TotalPayment: 		create.TotalPayment,
		Details:   			make([]domain.TransactionDetail, 0),
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
		Status : request.Status,
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
