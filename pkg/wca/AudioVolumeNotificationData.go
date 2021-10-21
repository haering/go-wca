package wca

import "github.com/go-ole/go-ole"

type AudioVolumeNotificationData struct {
	GuidEventContext ole.GUID
	Muted            bool
	MasterVolume     float32
	Channels         uint
	ChannelVolumes   [1]float32
}
