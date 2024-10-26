package mail

var (
	Config  *EmailStationConfig
	Station *EmailStation
)

func Load() {
	if Station != nil {
		Station.Close()
	}

	Station = NewEmailStation(*Config)
	Station.Run()
}
