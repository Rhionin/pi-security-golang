package motionsensor

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func StartSensor(onMotionStart, onMotionStop func(interface{})) {
	r := raspi.NewAdaptor()

	pir := gpio.NewPIRMotionDriver(r, "7")

	work := func() {
		fmt.Println("Registering motion detect events")
		pir.On(gpio.MotionDetected, onMotionStart)
		pir.On(gpio.MotionStopped, onMotionStop)
	}

	robot := gobot.NewRobot("motionBot",
		[]gobot.Connection{r},
		[]gobot.Device{pir},
		work,
	)

	fmt.Println("Starting worker")
	robot.Start()
}
