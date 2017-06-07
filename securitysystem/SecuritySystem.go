package securitysystem

import (
	"fmt"
	
	"github.com/loranbriggs/go-camera"
	"github.com/rhionin/pi-security-golang/motionsensor"
)

type (
	SecuritySystem struct {
		camera *camera.Camera
	}
)

func New(cameraOutputDir string) SecuritySystem {
	fmt.Println("On motion detect, saving pictures to ", cameraOutputDir)
	cam := camera.New(cameraOutputDir)
	return SecuritySystem{
		camera: cam,
	}
}

func (ss *SecuritySystem) Start() {
	fmt.Println("Starting security app")
	motionsensor.StartSensor(ss.onMotionStart, ss.onMotionStop)
	fmt.Println("Done")
}

func (ss *SecuritySystem) onMotionStart(interface{}) {
	fmt.Println("Motion Detected")
	ss.camera.Capture()
	fmt.Println("Captured picture")
}

func (ss *SecuritySystem) onMotionStop(interface{}) {
	fmt.Println("Motion Stopped")
}
