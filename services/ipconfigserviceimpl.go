package services

import (
	"mta_hosting_optimiser/helpers"
	"mta_hosting_optimiser/models"

	"github.com/gobuffalo/envy"
)

type IpConfigServiceImpl struct {
}

var ipConfigService IpConfigService

func GetIpConfigService() IpConfigService {
	if ipConfigService == nil {
		ipConfigService = IpConfigServiceImpl{}
	}
	return ipConfigService
}

func (impl IpConfigServiceImpl) GetInefficientHosts() ([]map[string]string, error) {
	ipconfig := []models.Ipconfig{}
	threshold := envy.Get("IP_CONFIG_THRESHOLD", "1")
	err := models.DB.RawQuery("select hostfqdn from ipconfigs group by hostfqdn having count(active) <= 1 union select hostfqdn from ipconfigs where active=? group by hostfqdn having count(active) <= ?", threshold, threshold).All(&ipconfig)
	if err != nil {
		helpers.ErrorLog(err)
		return nil, err
	}
	response := []map[string]string{}
	for _, ips := range ipconfig {
		response = append(response, map[string]string{"hostfqdn": ips.Hostfqdn})
	}
	return response, nil
}
