package ipfs

import (
	"io"
	"net/http"

	ipfsApi "github.com/ipfs/go-ipfs-api"
)

// IPFSClient is a struct that connects to an IPFS node.
type IPFSClient struct {
	client *ipfsApi.Shell
}

// NewIPFSClient creates a new instance of the IPFSClient struct.
func NewIPFSClient(url, projectId, projectSecret string) *IPFSClient {
	httpClient := &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}

	shell := ipfsApi.NewShellWithClient(url, httpClient)
	return &IPFSClient{client: shell}
}

// Add uploads a file to IPFS.
func (client *IPFSClient) Add(r io.Reader) (string, error) {
	hash, err := client.client.Add(r)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}
