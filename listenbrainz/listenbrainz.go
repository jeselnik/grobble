package listenbrainz

const (
	baseURL = "https://api.listenbrainz.org/1/"
)

type Params struct {
	Token, BaseURL string
}

type ListenBrainz struct {
	token, baseURL string
}

func New(p Params) *ListenBrainz {
	baseURL := baseURL
	if p.BaseURL != "" {
		baseURL = p.BaseURL
	}
	return &ListenBrainz{token: p.Token, baseURL: baseURL}
}

func (s *ListenBrainz) GetServiceName() string {
	return "ListenBrainz"
}
