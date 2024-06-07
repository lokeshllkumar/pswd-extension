package endpoints

import (
	"bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/utils"
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

	records, err := utils.ParseCLIOutput(out.String())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while parsing output"})
		return
	}

	c.JSON(200, records)
}
