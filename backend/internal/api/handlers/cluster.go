package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/freedyc/devin-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func init() {
	clusterNodes = []models.ClusterNode{
		{
			ID:        1,
			Name:      "primary-node",
			IPAddress: "127.0.0.1",
			Role:      "primary",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func GetClusterNodes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nodes": clusterNodes,
		"total": len(clusterNodes),
	})
}

func CreateClusterNode(c *gin.Context) {
	var req struct {
		Name      string `json:"name" binding:"required"`
		IPAddress string `json:"ip_address" binding:"required"`
		Role      string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node := models.ClusterNode{
		ID:        len(clusterNodes) + 1,
		Name:      req.Name,
		IPAddress: req.IPAddress,
		Role:      req.Role,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	clusterNodes = append(clusterNodes, node)
	c.JSON(http.StatusCreated, node)
}

func DeleteClusterNode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}

	for i, node := range clusterNodes {
		if node.ID == id {
			clusterNodes = append(clusterNodes[:i], clusterNodes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
}

func SyncClusterConfiguration(c *gin.Context) {
	var syncResults []map[string]interface{}

	for _, node := range clusterNodes {
		if node.Role == "secondary" && node.Status == "active" {
			result := map[string]interface{}{
				"node_id":   node.ID,
				"node_name": node.Name,
				"status":    "synced",
				"timestamp": time.Now(),
			}
			syncResults = append(syncResults, result)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cluster configuration synchronized",
		"results": syncResults,
	})
}
