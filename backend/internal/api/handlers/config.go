package handlers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReloadBind9(c *gin.Context) {
	cmd := exec.Command("rndc", "reload")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Failed to reload BIND9: %v, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to reload BIND9",
			"output": string(output),
		})
		return
	}

	logrus.Info("BIND9 reloaded successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "BIND9 reloaded successfully",
		"output":  string(output),
	})
}

func RestartDHCP(c *gin.Context) {
	cmd := exec.Command("systemctl", "restart", "isc-dhcp-server")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Failed to restart DHCP: %v, output: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to restart DHCP",
			"output": string(output),
		})
		return
	}

	logrus.Info("DHCP restarted successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "DHCP restarted successfully",
		"output":  string(output),
	})
}

func ValidateBind9Config(c *gin.Context) {
	cmd := exec.Command("named-checkconf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("BIND9 configuration validation failed: %v, output: %s", err, string(output))
		c.JSON(http.StatusBadRequest, gin.H{
			"valid":  false,
			"error":  "Configuration validation failed",
			"output": string(output),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"message": "Configuration is valid",
		"output":  string(output),
	})
}

func ValidateDHCPConfig(c *gin.Context) {
	cmd := exec.Command("dhcpd", "-t", "-cf", "/etc/dhcp/dhcpd.conf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("DHCP configuration validation failed: %v, output: %s", err, string(output))
		c.JSON(http.StatusBadRequest, gin.H{
			"valid":  false,
			"error":  "Configuration validation failed",
			"output": string(output),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"message": "Configuration is valid",
		"output":  string(output),
	})
}
