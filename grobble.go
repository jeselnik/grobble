package grobble

type Track struct {
	Timestamp            int
	Artist, Album, Title string
}

type Service interface {
	Login() error
	Scrobble(Track) error
	/* returns slice of successful tracks, slice of failed tracks, error */
	BatchScrobble([]Track) ([]Track, []Track, error)
	CapabilityBatchScrobble() bool
	GetServiceName() string
}
