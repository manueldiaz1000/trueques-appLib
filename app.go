package truequeslib

type Api struct {
	Version string
}

type BaseResponse struct {
	Version string `json:"version"`
	Error   string `json:"error"`
}

func InitApi(version string) Api {
	return Api{
		Version: version,
	}
}

func (api *Api) GetBaseReponse(err string) BaseResponse {
	return BaseResponse{
		Version: api.Version,
		Error:   err,
	}
}
