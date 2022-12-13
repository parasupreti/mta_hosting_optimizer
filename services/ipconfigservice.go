package services

type IpConfigService interface {
	GetInefficientHosts() ([]map[string]string, error)
}
