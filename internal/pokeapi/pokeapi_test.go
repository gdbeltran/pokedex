package pokeapi

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gdbeltran/pokedexcli/internal/pokecache"
)

// mockRoundTripper implements http.RoundTripper
type mockRoundTripper struct {
	responseBody string
	statusCode   int
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		StatusCode: m.statusCode,
		Body:       ioutil.NopCloser(strings.NewReader(m.responseBody)),
		Header:     make(http.Header),
	}
	return resp, nil
}

func TestListLocations_Mocked(t *testing.T) {
	mockResp := `{"results":[{"name":"kanto-route-1"}],"next":null,"previous":null}`
	client := &Client{
		httpClient: http.Client{
			Transport: &mockRoundTripper{
				responseBody: mockResp,
				statusCode:   200,
			},
		},
		cache: pokecache.NewCache(1 * time.Minute),
	}
	resp, err := client.ListLocations(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(resp.Results) != 1 || resp.Results[0].Name != "kanto-route-1" {
		t.Errorf("unexpected response: %+v", resp)
	}
}
