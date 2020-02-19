package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSc int
	Error Err
}

var (
	ErrorResponseBodyParseFailed = ErrResponse{
		HttpSc: 400,
		Error: Err{
			Error: "Request Body Cannot be Parsed!",
			ErrorCode: "001",
		},
	}
	ErrorNotAuthUser = ErrResponse{
		HttpSc: 401,
		Error: Err{
			Error: "User Authentication Failed!",
			ErrorCode:"002",
		},
	}
	ErrorDBError = ErrResponse{
		500,
		Err {
			"DB ops Failed",
			"003",
		},
	}
	ErrorInternalFaults = ErrResponse{
		HttpSc: 500,
		Error: Err{
			Error:     "Internal Fault",
			ErrorCode: "004",
		},
	}
)
