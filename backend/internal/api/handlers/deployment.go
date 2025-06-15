package handlers

import (
	"net/http"
	"os/exec"

	"github.com/freedyc/devin-app/backend/internal/models"
	"github.com/freedyc/devin-app/backend/pkg/bind9"
	"github.com/freedyc/devin-app/backend/pkg/dhcp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DeploySystemConfiguration(c *gin.Context) {
	var errors []string

	if err := deployDNSConfig(); err != nil {
		logrus.Errorf("Failed to deploy DNS config: %v", err)
		errors = append(errors, "DNS deployment failed: "+err.Error())
	}

	if err := deployDHCPConfig(); err != nil {
		logrus.Errorf("Failed to deploy DHCP config: %v", err)
		errors = append(errors, "DHCP deployment failed: "+err.Error())
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Deployment partially failed",
			"errors": errors,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "System configuration deployed successfully",
		"services": []string{"bind9", "dhcp"},
	})
}

func deployDNSConfig() error {
	for _, zone := range dnsZones {
		if zone.Enabled {
			records := getRecordsForZone(zone.ID)
			if err := bind9.WriteZoneFile("/etc/bind/zones", zone, records); err != nil {
				return err
			}
		}
	}

	if err := bind9.WriteNamedConf("/etc/bind/named.conf", dnsZones); err != nil {
		return err
	}

	return reloadBind9Service()
}

func deployDHCPConfig() error {
	if err := dhcp.WriteDHCPConf("/etc/dhcp/dhcpd.conf", dhcpScopes, dhcpReservations); err != nil {
		return err
	}

	return restartDHCPService()
}

func reloadBind9Service() error {
	cmd := exec.Command("rndc", "reload")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Failed to reload BIND9: %v, output: %s", err, string(output))
		return err
	}
	logrus.Info("BIND9 reloaded successfully")
	return nil
}

func restartDHCPService() error {
	cmd := exec.Command("systemctl", "restart", "isc-dhcp-server")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Failed to restart DHCP: %v, output: %s", err, string(output))
		return err
	}
	logrus.Info("DHCP restarted successfully")
	return nil
}

func getRecordsForZone(zoneID int) []models.DNSRecord {
	var records []models.DNSRecord
	for _, record := range dnsRecords {
		if record.ZoneID == zoneID {
			records = append(records, record)
		}
	}
	return records
}

func DeployDHCPConfiguration(c *gin.Context) {
	if err := deployDHCPConfig(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to deploy DHCP configuration",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "DHCP configuration deployed successfully",
	})
}
