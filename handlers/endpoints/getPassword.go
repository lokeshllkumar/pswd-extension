package endpoints

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func GetPasswordHandler(c *gin.Context) {
	service := c.Param("service")
	if service == "" {
		c.JSON(400, gin.H{"error": "Service parameter is required"})
		return
	}

	var cmd *exec.Cmd

	username := c.Param("username")

	if username == "" {
		cmd = exec.Command("./pswd-cli", "get", "--service", service)
	} else {
		cmd = exec.Command("./pswd-cli", "get", "--service", service, "--username", username)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute get subcommand"})
		return
	}

	var records map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &records); err != nil {
		c.JSON(500, gin.H{"error": "Error while parsing output"})
		return
	}

	c.JSON(200, records)
}