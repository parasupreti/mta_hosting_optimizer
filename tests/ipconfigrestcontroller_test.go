package tests

import (
	"mta_hosting_optimiser/services"
	"testing"
)

func TestIpConfigController(t *testing.T) {
	// Single qoute enclosed and comma separted awbs list
	ips := services.GetIpConfigService()
	resp, err := ips.GetInefficientHosts()
	if err != nil {
		t.Errorf("Error (%v)", err)
		return
	}
	mockOutput := []map[string]string{{"hostfqdn": "mta-prod-3"}, {"hostfqdn": "mta-prod-1"}}
	for key := range resp {
		if mockOutput[key]["hostfqdn"] != resp[key]["hostfqdn"] {
			t.Errorf("expected (%v), got (%v)", mockOutput[key]["hostfqdn"], resp[key]["hostfqdn"])
			return
		}
	}
}
