package listenbrainz

const (
	baseURL = "https://api.listenbrainz.org/1/"
)

type ListenBrainz struct {
	Token, BaseURL string
}

func New(l ListenBrainz) *ListenBrainz {
	baseURL := baseURL
	if l.BaseURL != "" {
		baseURL = l.BaseURL
	}
	return &ListenBrainz{Token: l.Token, BaseURL: baseURL}
}

func (s *ListenBrainz) CapabilityBatchScrobble() bool {
	return false
}

func (s *ListenBrainz) GetServiceName() string {
	return "ListenBrainz"
}
