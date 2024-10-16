package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8889")

	work := func() {
		drone.TakeOff()
		gobot.After(5*time.Second, func() {
			drone.Backward(10)
		})

		gobot.After(5*time.Second, func() {
			drone.Hover()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
