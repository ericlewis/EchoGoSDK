package examples

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ericlewis/EchoGoSDK/pkg/buttons"
	"github.com/ericlewis/EchoGoSDK/pkg/led"
	"github.com/ericlewis/EchoGoSDK/pkg/mic"
)

func main() {
	// Init Buttons
	if err := buttons.Init(); err != nil {
		panic(err)
	}

	buttons.ListenForVolumeButtons(func() {
		// Up
		led.SetColorAll(0, 255, 0)
	}, func() {
		// Down
		led.SetColorAll(255, 0, 0)
	})

	// Init LEDs
	if err := led.Init(); err != nil {
		panic(err)
	}

	// Clear all LED lights
	if err := led.Clear(); err != nil {
		panic(err)
	}

	// Prepare microphone
	if err := mic.Init(); err != nil {
		panic(err)
	}

	// Record microphone for 5 seconds
	micDevice := mic.GetDevice()
	audioStream := make(chan []byte)
	go func() {
		err := micDevice.GetAudioStream(micDevice.DeviceConfig, audioStream)
		if err != nil {
			panic(err)
		}
	}()

	recordingSeconds := 5
	fmt.Println("Recording for", recordingSeconds, "seconds...")
	start := time.Now()
	dataBuffer := new(bytes.Buffer)
	for {
		audioData := <-audioStream
		dataBuffer.Write(audioData)
		if time.Now().Sub(start).Seconds() >= float64(recordingSeconds) {
			fmt.Println("Stopping!")
			break
		}
	}
	close(audioStream)
	fmt.Println("Recorded", len(dataBuffer.Bytes()), "bytes")

	// Playing Audio through speaker
	// speaker.GetDevice().SendAudioStream(<YOUR WAV DATA>)

	// Run a fancy RGB light animation
	if err := led.Fun(); err != nil {
		panic(err)
	}
}
