package speaker

import (
	"github.com/ericlewis/GoTinyAlsa/pkg/pcm"
	"github.com/ericlewis/GoTinyAlsa/pkg/tinyalsa"
)

const CardNr = 0
const DeviceNr = 23

// GetDevice returns the pre-configured speaker alsa device
func GetDevice() tinyalsa.AlsaDevice {
	return tinyalsa.NewDevice(CardNr, DeviceNr, pcm.Config{
		Channels:    2,
		SampleRate:  48000,
		PeriodSize:  2048,
		PeriodCount: 4,
		Format:      tinyalsa.PCM_FORMAT_S16_LE,
	})
}
