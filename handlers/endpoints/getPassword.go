package endpoints

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/utils"
)

func GetPasswordHandler(c *gin.Context) {
	service := c.Query("service")
	if service == "" {
		c.JSON(400, gin.H{"error": "Service parameter is required"})
		return
	}

	var cmd *exec.Cmd

	username := c.Query("username")

	if username == "" {
		cmd = exec.Command("/usr/local/bin/pswd-cli/pswd-cli", "get", "--service", service)
	} else {
		cmd = exec.Command("/usr/local/bin/pswd-cli/pswd-cli", "get", "--service", service, "--username", username)
	}

	stdout, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute get subcommand"})
		return
	}

	records, err := utils.ParseCLIOutput(string(stdout))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while parsing output"})
		return
	}

	c.JSON(200, records)
}
