package ipfs

import (
	"io"
	"net/http"
	"net/url"

	ipfsApi "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
)

// IPFSClient is a struct that connects to an IPFS node.
type IPFSClient struct {
	client *ipfsApi.Shell
}

// NewIPFSClient creates a new instance of the IPFSClient struct.
func NewIPFSClient(ipfsUrl string) *IPFSClient {
	u, err := url.Parse(ipfsUrl)
	if err != nil {
		log.WithError(err).Fatal("Could not parse IPFS URL")
	}

	ipfsHost := u.Scheme + "://" + u.Host
	ipfsUsername := u.User.Username()
	ipfsPassword, _ := u.User.Password()

	log.WithField("host", u.Redacted()).Info("Starting IPFS client")

	httpClient := &http.Client{
		Transport: authTransport{
			RoundTripper: http.DefaultTransport,
			Username:     ipfsUsername,
			Password:     ipfsPassword,
		},
	}

	shell := ipfsApi.NewShellWithClient(ipfsHost, httpClient)
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
	Username string
	Password string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.Username, t.Password)
	return t.RoundTripper.RoundTrip(r)
}
