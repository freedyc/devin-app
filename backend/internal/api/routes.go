package api

import (
	"github.com/freedyc/devin-app/backend/internal/api/handlers"
	"github.com/freedyc/devin-app/backend/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	api := router.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)

		dns := api.Group("/dns")
		{
			dns.GET("/zones", handlers.GetDNSZones)
			dns.POST("/zones", handlers.CreateDNSZone)
			dns.GET("/zones/:id", handlers.GetDNSZone)
			dns.PUT("/zones/:id", handlers.UpdateDNSZone)
			dns.DELETE("/zones/:id", handlers.DeleteDNSZone)

			dns.GET("/zones/:id/records", handlers.GetDNSRecords)
			dns.POST("/zones/:id/records", handlers.CreateDNSRecord)
			dns.GET("/records/:id", handlers.GetDNSRecord)
			dns.PUT("/records/:id", handlers.UpdateDNSRecord)
			dns.DELETE("/records/:id", handlers.DeleteDNSRecord)
		}

		dhcp := api.Group("/dhcp")
		{
			dhcp.GET("/scopes", handlers.GetDHCPScopes)
			dhcp.POST("/scopes", handlers.CreateDHCPScope)
			dhcp.GET("/scopes/:id", handlers.GetDHCPScope)
			dhcp.PUT("/scopes/:id", handlers.UpdateDHCPScope)
			dhcp.DELETE("/scopes/:id", handlers.DeleteDHCPScope)

			dhcp.GET("/scopes/:id/reservations", handlers.GetDHCPReservations)
			dhcp.POST("/scopes/:id/reservations", handlers.CreateDHCPReservation)
			dhcp.GET("/reservations/:id", handlers.GetDHCPReservation)
			dhcp.PUT("/reservations/:id", handlers.UpdateDHCPReservation)
			dhcp.DELETE("/reservations/:id", handlers.DeleteDHCPReservation)

			dhcp.GET("/leases", handlers.GetDHCPLeases)
			dhcp.GET("/scopes/:id/leases", handlers.GetDHCPLeasesByScope)
		}

		gslb := api.Group("/gslb")
		{
			gslb.GET("/policies", handlers.GetGSLBPolicies)
			gslb.POST("/policies", handlers.CreateGSLBPolicy)
			gslb.GET("/policies/:id", handlers.GetGSLBPolicy)
			gslb.PUT("/policies/:id", handlers.UpdateGSLBPolicy)
			gslb.DELETE("/policies/:id", handlers.DeleteGSLBPolicy)

			gslb.GET("/policies/:id/pools", handlers.GetGSLBPools)
			gslb.POST("/policies/:id/pools", handlers.CreateGSLBPool)
			gslb.GET("/pools/:id", handlers.GetGSLBPool)
			gslb.PUT("/pools/:id", handlers.UpdateGSLBPool)
			gslb.DELETE("/pools/:id", handlers.DeleteGSLBPool)

			gslb.GET("/pools/:id/servers", handlers.GetGSLBServers)
			gslb.POST("/pools/:id/servers", handlers.CreateGSLBServer)
			gslb.GET("/servers/:id", handlers.GetGSLBServer)
			gslb.PUT("/servers/:id", handlers.UpdateGSLBServer)
			gslb.DELETE("/servers/:id", handlers.DeleteGSLBServer)
		}

		config := api.Group("/config")
		{
			config.POST("/bind9/reload", handlers.ReloadBind9)
			config.POST("/dhcp/restart", handlers.RestartDHCP)
			config.GET("/bind9/validate", handlers.ValidateBind9Config)
			config.GET("/dhcp/validate", handlers.ValidateDHCPConfig)
		}
	}
}
