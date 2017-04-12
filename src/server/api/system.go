package api

type SystemAPI interface {
	ReportStatus(data interface{})
	FetchCfg(serviceID int)
}
