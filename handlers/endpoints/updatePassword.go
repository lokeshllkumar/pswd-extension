package endpoints

import (
	_ "bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/utils"
)

func UpdatePasswordHandler(c *gin.Context) {
	var updateEntry utils.PasswordEntry

	if err := c.BindJSON(&updateEntry); err != nil {
		c.JSON(500, gin.H{"error": "Failed to bind JSON data"})
		return
	}

	if updateEntry.Service == "" || updateEntry.Username == "" || updateEntry.Password == "" {
		c.JSON(400, gin.H{"error": "Service, username and newPassword parameters are required"})
		return
	}

	cmd := exec.Command("/usr/local/bin/pswd-cli/pswd-cli", "update", "--service", updateEntry.Service, "--username", updateEntry.Username, "--newPassword", updateEntry.Password)
	_, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute update subcommand"})
		return
	}

	c.JSON(200, gin.H{"message": "Password record updated successfully"})
}
