package swagger

import (
	"log"
	"os/exec"
)

func GenerateSwaggerDocs() {
	cmd := exec.Command("swag", "init",
		"--parseDependency",
		"--parseInternal",
		"--parseDepth", "1",
		"--output", "./docs",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to generate Swagger docs: %v\nOutput: %s", err, output)
		return
	}

	log.Println("Swagger documentation generated successfully")
}
