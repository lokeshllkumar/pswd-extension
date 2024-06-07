package endpoints

import (
	_ "bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/utils"
)

func AddPasswordHandler(c *gin.Context) {
	var req utils.PasswordEntry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
	}
	cmd := exec.Command("pswd-cli", "add", "--service", req.Service, "--username", req.Username, "--password", req.Password)
	// var out bytes.Buffer
	// cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute add subcommand"})
		return
	}

	c.JSON(200, gin.H{"message": "Password pushed to database successfully"})
}
