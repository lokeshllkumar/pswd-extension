package endpoints

import (
	"bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func DeletePasswordHandler(c *gin.Context) {
	service := c.Query("service")
	if service == "" {
		c.JSON(400, gin.H{"error": "Service parameter is required"})
		return
	}

	var cmd *exec.Cmd

	username := c.Query("username")

	if username == "" {
		cmd = exec.Command("pswd-cli", "update", "--service", "service")
	} else {
		cmd = exec.Command("pswd-cli", "update", "--service", service, "--username", username)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failes to execute delete subcommand"})
		return
	}

	c.JSON(200, gin.H{"message": "Password record deleted successfully"})
}