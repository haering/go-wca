package wca

type IAudioEndpointVolumeEvents struct {
	VTable *IAudioEndpointVolumeEventsVtbl
}

type IAudioEndpointVolumeEventsVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	OnNotify       uintptr
}
