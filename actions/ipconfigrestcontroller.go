package actions

import (
	"mta_hosting_optimiser/services"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func IpConfigHandler(c buffalo.Context) error {
	ipConfigServiceImpl := services.GetIpConfigService()
	response, err := ipConfigServiceImpl.GetInefficientHosts()
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"message": "Server issue. Please try again later!!"}))
	}
	return c.Render(http.StatusOK, r.JSON(response))
}
