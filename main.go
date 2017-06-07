package main

import (
	"os"

	"github.com/rhionin/pi-security-golang/securitysystem"
)

func main() {
	cameraOutputDir := os.Getenv("CAMERA_OUTPUT_DIR")
	
	ss := securitysystem.New(cameraOutputDir)
	ss.Start()
}


