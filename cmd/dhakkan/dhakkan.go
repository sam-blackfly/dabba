package main

import (
	"fmt"

	"github.com/sam-blackfly/dabba/internal/colors"
)

func main() {

	fmt.Printf("%s All checks %s\n", colors.Info("✓"), colors.Success("PASSED"))
}
