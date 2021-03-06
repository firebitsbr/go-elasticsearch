// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// FlushSynced - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-synced-flush.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) FlushSynced(options ...*Option) (*FlushSyncedResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["FlushSynced"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &FlushSyncedResponse{resp}, err
}

// FlushSyncedResponse is the response for FlushSynced
type FlushSyncedResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
