package endpoints

import (
	_ "bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func UpdatePasswordHandler(c *gin.Context) {
	service := c.Query("service")
	username := c.Query("username")
	newPassword := c.Query("newPassword")

	if service == "" || username == "" || newPassword == "" {
		c.JSON(400, gin.H{"error": "Service, username and newPassword parameters are required"})
		return
	}

	cmd := exec.Command("pswd-cli", "update", "--service", service, "--username", username, "--newPassword", newPassword)

	// var out bytes.Buffer
	// cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute update subcommand"})
		return
	}

	c.JSON(200, gin.H{"message": "Password record updated successfully"})
}
