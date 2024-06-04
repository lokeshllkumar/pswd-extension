package endpoints

import (
	"bytes"
	"os/exec"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetPasswordsHandler(c *gin.Context) {
	cmd := exec.Command("pswd-cli", "list")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute list subcommand"})
		return
	}

	var records []map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &records); err != nil {
		c.JSON(500, gin.H{"error": "Error while parsing output"})
		return
	}

	c.JSON(200, records)
}
