package truequeslib

type Api struct {
	Name    string
	Version string
}

type BaseResponse struct {
	Version string `json:"version"`
	Error   string `json:"error"`
}

func InitApi(name string, version string) Api {
	return Api{
		Name:    name,
		Version: version,
	}
}

func (api *Api) GetBaseReponse(err string) BaseResponse {
	return BaseResponse{
		Version: api.Version,
		Error:   err,
	}
}
