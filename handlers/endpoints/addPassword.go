package endpoints

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/utils"
)

func AddPasswordHandler(c *gin.Context) {
	var newEntry utils.PasswordEntry

	if err := c.BindJSON(&newEntry); err != nil {
		c.JSON(500, gin.H{"error": "Failed to bind JSON data"})
		return
	}

	if newEntry.Service == "" || newEntry.Username == "" || newEntry.Password == "" {
		c.JSON(400, gin.H{"error": "service, username, and password are required fields"})
		return
	}

	cmd := exec.Command("/usr/local/bin/pswd-cli/pswd-cli", "add", "--service", newEntry.Service, "--username", newEntry.Username, "--password", newEntry.Password)
	_, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute add subcommand"})
		return
	}

	c.JSON(200, gin.H{"message": "Password pushed to database successfully"})
}
