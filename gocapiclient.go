package gocapiclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/guardian/gocapiclient/queries"
	"net/http"
)

const ClientVersion = "0.1"

type GuardianContentClient struct {
	ApiKey             string
	UserAgent          string
	HttpClient         *http.Client
	TargetUrl          string
	ThriftDeserializer *thrift.TDeserializer
}

func (contentClient GuardianContentClient) makeCapiRequest(q queries.Query) (*http.Response, error) {
	url := q.GetUrl(contentClient.TargetUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Currently this client only supports Thrift
	req.Header.Set("Accept", "application/x-thrift")
	req.Header.Set("User-Agent", "content-api-go-client/v"+string(ClientVersion))

	values := req.URL.Query()
	values.Add("api-key", contentClient.ApiKey)
	values.Add("format", "thrift")

	req.URL.RawQuery = values.Encode()

	return contentClient.HttpClient.Do(req)
}

func (contentClient GuardianContentClient) GetResponse(q queries.Query) error {
	response, err := contentClient.makeCapiRequest(q)
	if err != nil {
		return err
	}

	return q.Deserialize(contentClient.ThriftDeserializer, response.Body)
}

func NewGuardianContentClient(targetUrl string, apiKey string) *GuardianContentClient {
	transport := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTCompactProtocolFactory().GetProtocol(transport)

	deser := &thrift.TDeserializer{
		transport,
		protocol}

	userAgent := "content-api-go-client/v" + string(ClientVersion)
	client := GuardianContentClient{
		ApiKey:             apiKey,
		UserAgent:          userAgent,
		HttpClient:         &http.Client{},
		TargetUrl:          targetUrl,
		ThriftDeserializer: deser,
	}

	return &client
}
