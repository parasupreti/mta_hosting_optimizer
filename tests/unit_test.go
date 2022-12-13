package tests

import (
	"mta_hosting_optimiser/services"
	"testing"
)

func Test1(t *testing.T) {
	ips := services.GetIpConfigService()
	if ips == nil {
		t.Errorf("returned empty ipconfig service object")
		return
	}
}
