package grobble

const (
	BATCH_SCROBBLE_MAX = 50
)

type Track struct {
	Timestamp            int
	Artist, Album, Title string
}

type Service interface {
	Login() error
	Scrobble(Track) error
	/* returns (slice of successful tracks, slice of failed tracks, error)
	has a limit of 50 tracks a request (aligning with last.fm API limitations) */
	BatchScrobble([]Track) ([]Track, []Track, error)
	CapabilityBatchScrobble() bool
	GetServiceName() string
}
