package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func TransactionDomainToTransactionResponse(transaction *domain.Transaction) *web.TransactionResponse {
	createdAt := transaction.CreatedAt
	response := &web.TransactionResponse{
		ID:        transaction.ID,
		CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
		Cashier: web.CashierTransactionResponse{
			ID:       transaction.Cashier.ID,
			Fullname: transaction.Cashier.Fullname,
			Username: transaction.Cashier.Username,
		},
		Membership: web.MembershipTransactionResponse{
			ID:          transaction.Membership.ID,
			Name:        transaction.Membership.Name,
			CodeMember:  transaction.Membership.CodeMember,
			TotalPoint:  transaction.Membership.TotalPoint,
			PhoneNumber: transaction.Membership.PhoneNumber,
		},
		ConvertPointID: transaction.ConvertPointID,
		Discount:       transaction.Discount,
		TotalPrice:     transaction.TotalPrice,
		Tax:            transaction.Tax,
		TotalPayment:   transaction.TotalPayment,
		TransactionPayment: web.TransactionPaymentResponse{
			ID:            transaction.TransactionPayment.ID,
			TransactionID: transaction.ID,
			CreatedAt:     transaction.TransactionPayment.CreatedAt,
			UpdateAt:      transaction.TransactionPayment.UpdatedAt,
			PaymentMethod: web.PaymentMethodResponse{
				ID:            transaction.TransactionPayment.PaymentMethod.ID,
				PaymentTypeID: transaction.TransactionPayment.PaymentMethod.PaymentTypeID,
				PaymentType: web.PaymentTypeResponse{
					ID:       transaction.TransactionPayment.PaymentMethod.PaymentType.ID,
					TypeName: transaction.TransactionPayment.PaymentMethod.PaymentType.TypeName,
				},
				Name: transaction.TransactionPayment.PaymentMethod.Name,
			},
			Invoice:       transaction.TransactionPayment.Invoice,
			VANumber:      transaction.TransactionPayment.VANumber,
			PaymentStatus: transaction.TransactionPayment.PaymentStatus,
		},
	}

	for _, detail := range transaction.Details {
		response.Details = append(response.Details, web.TransactionDetailResponse{
			ID:            detail.ID,
			TransactionID: detail.TransactionID,
			Price:         detail.Price,
			Quantity:      detail.Quantity,
			SubTotal:      detail.SubTotal,
			Notes:         detail.Notes,
			ProductDetail: web.ProductDetailTransactionResponse{
				ID:        detail.ProductDetail.ID,
				ProductID: detail.ProductDetail.ProductID,
				Product: web.ProductTransactionResponse{
					ID:          detail.ProductDetail.Product.ID,
					Name:        detail.ProductDetail.Product.Name,
					Ingredients: detail.ProductDetail.Product.Ingredients,
					Image:       detail.ProductDetail.Product.Image,
				},
				Price:      detail.ProductDetail.Price,
				TotalStock: detail.ProductDetail.TotalStock,
				Size:       detail.ProductDetail.Size,
			},
		})
	}
	return response

}
func TransactionMonthlyRevenueDomainToTransactionMonthlyRevenueResponse(transactionMonthly *domain.TransactionMonthlyRevenue) *web.TransactionMonthlyRevenueResponse{
	return &web.TransactionMonthlyRevenueResponse{
		Year: transactionMonthly.Year,
		Month: transactionMonthly.Month,
		Revenue: transactionMonthly.Revenue,
	}
}

func TransactionYearlyRevenueDomainToTransactionYearlyRevenueResponse(transactionMonthly *domain.TransactionYearlyRevenue) *web.TransactionYearlyRevenueResponse{
	return &web.TransactionYearlyRevenueResponse{
		Year: transactionMonthly.Year,
		Revenue: transactionMonthly.Revenue,
	}
}

func TransactionDailyDomainToTransactionDailyResponse(transactionDaily *domain.TransactionDailyRevenue) *web.TransactionDailyRevenueResponse{
	return &web.TransactionDailyRevenueResponse{
		Day: transactionDaily.Day.Format("2006-01-02"),
		Success: transactionDaily.Success,
		Pending: transactionDaily.Pending,
		Cancelled: transactionDaily.Cancelled,
		Revenue: transactionDaily.Revenue,
	}
}

func TransactionDomainToTransactionResponseNoMembership(transaction *domain.Transaction) *web.TransactionResponse {
	createdAt := transaction.CreatedAt
	response := &web.TransactionResponse{
		ID:        transaction.ID,
		CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
		Cashier: web.CashierTransactionResponse{
			ID:       transaction.Cashier.ID,
			Fullname: transaction.Cashier.Fullname,
			Username: transaction.Cashier.Username,
		},
		Membership: web.MembershipTransactionResponse{
			Name: "Anonymous",
		},
		ConvertPointID: transaction.ConvertPointID,
		Discount:       transaction.Discount,
		TotalPrice:     transaction.TotalPrice,
		Tax:            transaction.Tax,
		TotalPayment:   transaction.TotalPayment,
		TransactionPayment: web.TransactionPaymentResponse{
			ID:            transaction.TransactionPayment.ID,
			TransactionID: transaction.ID,
			PaymentMethod: web.PaymentMethodResponse{
				ID:            transaction.TransactionPayment.PaymentMethod.ID,
				PaymentTypeID: transaction.TransactionPayment.PaymentMethod.PaymentTypeID,
				PaymentType: web.PaymentTypeResponse{
					ID:       transaction.TransactionPayment.PaymentMethod.PaymentType.ID,
					TypeName: transaction.TransactionPayment.PaymentMethod.PaymentType.TypeName,
				},
				Name: transaction.TransactionPayment.PaymentMethod.Name,
			},
			Invoice:       transaction.TransactionPayment.Invoice,
			VANumber:      transaction.TransactionPayment.VANumber,
			PaymentStatus: transaction.TransactionPayment.PaymentStatus,
		},
	}

	for _, detail := range transaction.Details {
		response.Details = append(response.Details, web.TransactionDetailResponse{
			ID:            detail.ID,
			TransactionID: detail.TransactionID,
			Price:         detail.Price,
			Quantity:      detail.Quantity,
			SubTotal:      detail.SubTotal,
			Notes:         detail.Notes,
			ProductDetail: web.ProductDetailTransactionResponse{
				ID:        detail.ProductDetail.ID,
				ProductID: detail.ProductDetail.ProductID,
				Product: web.ProductTransactionResponse{
					ID:          detail.ProductDetail.Product.ID,
					Name:        detail.ProductDetail.Product.Name,
					Ingredients: detail.ProductDetail.Product.Ingredients,
					Image:       detail.ProductDetail.Product.Image,
				},
				Price:      detail.ProductDetail.Price,
				TotalStock: detail.ProductDetail.TotalStock,
				Size:       detail.ProductDetail.Size,
			},
		})
	}
	return response

}

func TransactionSchemaToTransactionDomain(transaction *schema.Transaction) *domain.Transaction {
	return &domain.Transaction{
		ID:             transaction.ID,
		CashierID:      transaction.CashierID,
		MembershipID:   transaction.MembershipID,
		ConvertPointID: transaction.ConvertPointID,
		Discount:       transaction.Discount,
		TotalPrice:     transaction.TotalPrice,
		Tax:            transaction.Tax,
		TotalPayment:   transaction.TotalPayment,
		Details:        []domain.TransactionDetail{},
	}
}

func ConvertTransactionResponse(transactions []domain.Transaction) []web.TransactionResponse {
	var results []web.TransactionResponse
	for _, transaction := range transactions {
		results = append(results, *TransactionDomainToTransactionResponse(&transaction))
	}
	return results
}

func ConvertTransactionMonthlyRevenueResponse(transactions []domain.TransactionMonthlyRevenue) []web.TransactionMonthlyRevenueResponse {
	var results []web.TransactionMonthlyRevenueResponse
	for _, transaction  := range transactions {
		results = append(results, *TransactionMonthlyRevenueDomainToTransactionMonthlyRevenueResponse(&transaction))
	}
	return results
}

