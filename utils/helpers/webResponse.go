package helpers

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

type TSuccessResponseWithTotal struct {
	Meta    TResponseMeta `json:"meta"`
	Total   int           `json:"total"`
	Results interface{}   `json:"results"`
}

type TSuccessResponseWithMeta struct {
	Meta       TResponseMeta `json:"meta"`
	Pagination any           `json:"pagination,omitempty"`
	Results    interface{}   `json:"results"`
}

type TErrorResponse struct {
	Meta TResponseMeta `json:"meta"`
}

func SuccessResponse(message string, data interface{}) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
		}
	} else {
		return TSuccessResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
			Results: data,
		}
	}
}

func SuccessResponseWithMeta(message string, data interface{}, pagination any) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
		}
	} else {
		return TSuccessResponseWithMeta{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
			Pagination: pagination,
			Results:    data,
		}
	}
}

func SuccessResponseWithTotal(message string, data interface{}, total int) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
		}
	} else {
		return TSuccessResponseWithTotal{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
			Total:   total,
			Results: data,
		}
	}
}

func ErrorResponse(message string) interface{} {
	return TErrorResponse{
		Meta: TResponseMeta{
			Success: false,
			Message: message,
		},
	}
}
