package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func TransactionDomainToTransactionResponse(transaction *domain.Transaction) *web.TransactionResponse {
	response := &web.TransactionResponse{
		ID:              transaction.ID,
		CreatedAt: transaction.CreatedAt,
		Cashier: web.CashierTransactionResponse{
			ID: transaction.Cashier.ID,
			Username: transaction.Cashier.Username,
		},
		Membership: web.MembershipTransactionResponse{
			ID: transaction.Membership.ID,
			Name: transaction.Membership.Name,
			CodeMember: transaction.Membership.CodeMember,
			Point: transaction.ConvertPoint.Point,
			PhoneNumber: transaction.Membership.PhoneNumber,
		},
		ConvertPointID: transaction.ConvertPointID,
		Status: transaction.Status,
		Discount: transaction.Discount,
		TotalPrice: transaction.TotalPrice,
		Tax: transaction.Tax,
		TotalPayment: transaction.TotalPayment,
	}


	for _ , detail := range transaction.Details { 
		response.Details = append(response.Details, web.TransactionDetailResponse{
			ID: detail.ID,
			TransactionID: detail.TransactionID,
			ProductDetailID: detail.ProductDetailID,
			Price: detail.Price,
			Quantity: detail.Quantity,
			SubTotal: detail.SubTotal,
			Notes: detail.Notes,
			ProductDetail: web.ProductDetailTransactionResponse{
				ID: detail.ProductDetail.ID,
				ProductID: detail.ProductDetail.ProductID,
				Product: web.ProductTransactionResponse{
					ID: detail.ProductDetail.Product.ID,
					Name: detail.ProductDetail.Product.Name,
					Ingredients: detail.ProductDetail.Product.Ingredients,
					Image: detail.ProductDetail.Product.Image,
				},
				Price: detail.ProductDetail.Price,
				TotalStock: detail.ProductDetail.TotalStock,
				Size: detail.ProductDetail.Size,
			},
		})
	}
	return response 

}

func TransactionSchemaToTransactionDomain(transaction *schema.Transaction) *domain.Transaction {
	return &domain.Transaction{
		ID:              transaction.ID,
		CashierID:        transaction.CashierID,
		MembershipID: transaction.MembershipID,
		ConvertPointID: transaction.ConvertPointID,
		Status: transaction.Status,
		Discount: transaction.Discount,
		TotalPrice: transaction.TotalPrice,
		Tax: transaction.Tax,
		TotalPayment: transaction.TotalPayment,
		Details: []domain.TransactionDetail{},
	}
}


func ConvertTransactionResponse(transactions []domain.Transaction) []web.TransactionResponse {
	var results []web.TransactionResponse

	for _, transaction := range transactions {
		TransactionResponse := web.TransactionResponse{
			ID:              transaction.ID,
			ConvertPointID: transaction.ConvertPointID,
			Status: transaction.Status,
			Discount: transaction.Discount,
			TotalPrice: transaction.TotalPrice,
			Tax: transaction.Tax,
			TotalPayment: transaction.TotalPayment,
			// Details: transaction.Details,
		}
		results = append(results, TransactionResponse)
	}
	return results
}