package api

type SystemAPI interface {
	ReportStatus(data interface{})
	FetchCfg(serviceID int)
}

type UserAPI interface {
	Authorize(userID int64, userToken string) (string, int)
	ModifyProperty(userID int64, token string) bool
}
