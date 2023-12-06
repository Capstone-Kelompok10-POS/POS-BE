package helpers

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
type TSuccessResponseWithTotal struct {
	Meta    TResponseMeta `json:"meta"`
	Total   int           `json:"total"`
	Results interface{}   `json:"results"`
}

>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
type TSuccessResponseWithMeta struct {
	Meta       TResponseMeta `json:"meta"`
	Pagination any           `json:"pagination,omitempty"`
	Results    interface{}   `json:"results"`
}

<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
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

>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
func ErrorResponse(message string) interface{} {
	return TErrorResponse{
		Meta: TResponseMeta{
			Success: false,
			Message: message,
		},
	}
}
