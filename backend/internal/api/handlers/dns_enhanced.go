package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/freedyc/devin-app/backend/internal/models"
	"github.com/freedyc/devin-app/backend/pkg/bind9"
	"github.com/gin-gonic/gin"
)

var (
	dnsViews        []models.DNSView
	forwardingZones []models.ForwardingZone
	stubZones       []models.StubZone
	recursionConfig models.RecursionConfig
)

func init() {
	recursionConfig = models.RecursionConfig{
		ID:               1,
		Enabled:          true,
		AllowRecursion:   []string{"any"},
		Forwarders:       []string{"8.8.8.8", "8.8.4.4"},
		ForwardFirst:     true,
		DNSSECValidation: "auto",
	}
}

func GetDNSViews(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"views": dnsViews,
		"total": len(dnsViews),
	})
}

func CreateDNSView(c *gin.Context) {
	var req models.CreateViewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	view := models.DNSView{
		ID:           len(dnsViews) + 1,
		Name:         req.Name,
		MatchClients: req.MatchClients,
		Recursion:    req.Recursion,
		Enabled:      true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	dnsViews = append(dnsViews, view)
	c.JSON(http.StatusCreated, view)
}

func DeleteDNSView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid view ID"})
		return
	}

	for i, view := range dnsViews {
		if view.ID == id {
			dnsViews = append(dnsViews[:i], dnsViews[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "View deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "View not found"})
}

func GetForwardingZones(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"zones": forwardingZones,
		"total": len(forwardingZones),
	})
}

func CreateForwardingZone(c *gin.Context) {
	var req models.CreateForwardingZoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zone := models.ForwardingZone{
		ID:         len(forwardingZones) + 1,
		Name:       req.Name,
		Forwarders: req.Forwarders,
		Forward:    req.Forward,
		Enabled:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	forwardingZones = append(forwardingZones, zone)
	c.JSON(http.StatusCreated, zone)
}

func DeleteForwardingZone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	for i, zone := range forwardingZones {
		if zone.ID == id {
			forwardingZones = append(forwardingZones[:i], forwardingZones[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Forwarding zone deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Forwarding zone not found"})
}

func GetStubZones(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"zones": stubZones,
		"total": len(stubZones),
	})
}

func CreateStubZone(c *gin.Context) {
	var req models.CreateStubZoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zone := models.StubZone{
		ID:        len(stubZones) + 1,
		Name:      req.Name,
		Masters:   req.Masters,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	stubZones = append(stubZones, zone)
	c.JSON(http.StatusCreated, zone)
}

func DeleteStubZone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	for i, zone := range stubZones {
		if zone.ID == id {
			stubZones = append(stubZones[:i], stubZones[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Stub zone deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Stub zone not found"})
}

func GetRecursionConfig(c *gin.Context) {
	c.JSON(http.StatusOK, recursionConfig)
}

func UpdateRecursionConfig(c *gin.Context) {
	var req models.RecursionConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recursionConfig = req
	c.JSON(http.StatusOK, recursionConfig)
}

func DeployDNSConfiguration(c *gin.Context) {
	config := bind9.EnhancedNamedConfig{
		Recursion:       recursionConfig,
		Views:           dnsViews,
		ForwardingZones: forwardingZones,
		StubZones:       stubZones,
		DefaultZones:    convertToEnhancedZones(dnsZones),
	}

	configPath := "/etc/bind/named.conf"
	
	content, err := bind9.GenerateEnhancedNamedConf(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate configuration",
			"details": err.Error(),
		})
		return
	}

	if err := bind9.WriteEnhancedNamedConf(configPath, config); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Configuration generated successfully",
			"note": "Manual deployment required - run with sudo privileges",
			"config_path": configPath,
			"config_content": content,
			"deployment_command": "sudo systemctl reload bind9",
		})
		return
	}

	if err := reloadBind9Service(); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{
			"message": "Configuration written but service reload failed",
			"config_path": configPath,
			"reload_error": err.Error(),
			"manual_reload": "sudo systemctl reload bind9",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "DNS configuration deployed and service reloaded successfully",
		"config_path": configPath,
	})
}

func convertToEnhancedZones(zones []models.DNSZone) []models.DNSZoneEnhanced {
	enhanced := make([]models.DNSZoneEnhanced, len(zones))
	for i, zone := range zones {
		enhanced[i] = models.DNSZoneEnhanced{
			DNSZone:  zone,
			ZoneType: "master",
		}
	}
	return enhanced
}
