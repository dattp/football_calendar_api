package common

type responseCustom struct {
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
	Options interface{} `json:"options"`
}

func Response(data interface{}, error interface{}, options interface{}) *responseCustom {
	return &responseCustom{
		Error:   error,
		Data:    data,
		Options: options,
	}
}

func ResponseSuccessLite(data interface{}) *responseCustom {
	return Response(data, nil, nil)
}

func ResponseError(error interface{}) *responseCustom {
	return Response(nil, error, nil)
}
