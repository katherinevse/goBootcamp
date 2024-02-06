package api

// точка входа в программу
type Api struct {
}

// API constructor: build base API instance
func New() *Api {
	return &Api{}
}

// Start http,configure, router,database connection
func (api *Api) Start() error {
	return nil
}
