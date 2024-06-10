package endpoints

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/utils"
)

func GetPasswordsHandler(c *gin.Context) {
	cmd := exec.Command("/usr/local/bin/pswd-cli/pswd-cli", "list")

	stdout, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute list subcommand"})
		return
	}

	fmt.Printf("%s\n", string(stdout))

	records, err := utils.ParseCLIOutput(string(stdout))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while parsing output"})
		return
	}

	c.JSON(200, records)
}
