package wca

type IAudioEndpointVolumeEvents struct {
	VTable *IAudioEndpointVolumeEventsVtbl
}

type IAudioEndpointVolumeEventsVtbl struct {
	OnNotify uintptr
}
