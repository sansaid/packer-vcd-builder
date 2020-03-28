package vcdbuilder

import (
	"fmt"
	"github.com/google/uuid"
)

func generateRandomVmName() (string, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, fmt.Errorf("Error generating random VM name: %s", err)
	}
	
	vmName := fmt.Sprintf("packergen-%s", string(uuid))
	return vmName, err
}