package smartHomeSamsungTVCommon

// DeviceEndpointResponse holds the current request state
type DeviceEndpointResponse struct {
	TVSwitchedOn  bool
	TVSwitchedOff bool
	VolumeUp      int
	VolumeDown    int
	VolumeMute    bool
	Play          bool
	Pause         bool
	OK            bool
}
