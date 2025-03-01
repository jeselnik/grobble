package grobble

type Track struct {
	Timestamp            int
	Artist, Album, Title string
}

type Service interface {
	Auth() error
	Login() error
	Scrobble(Track) error
	BatchScrobble([]Track) (int, int)
	CapabilityBatchScrobble() bool
	GetServiceName() string
}
