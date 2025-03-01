package grobble

type Track struct {
	Artist, Album, Title, Timestamp string
}

type Service interface {
	Auth() error
	Login() error
	Scrobble(Track) error
	BatchScrobble([]Track) (int, int)
	CapabilityBatchScrobble() bool
	GetServiceName() string
}
