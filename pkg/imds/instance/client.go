package instance

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const instanceEndpoint = "http://169.254.169.254/metadata/instance"

type ClientConfig struct {
	Logger micrologger.Logger

	HttpClient *http.Client
}

type Client struct {
	logger micrologger.Logger

	httpClient *http.Client
}

func NewClient(config ClientConfig) (*Client, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.HttpClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.HttpClient must not be empty", config)
	}

	client := &Client{
		logger: config.Logger,
		httpClient: config.HttpClient,
	}

	return client, nil
}

func (c *Client) GetMetadata(ctx context.Context) (metadata *Metadata, err error) {
	c.logger.LogCtx(ctx, "level", "debug", "message", "fetching instance metadata")

	req, err := http.NewRequest("GET", instanceEndpoint, nil)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	req.Header.Add("Metadata", "True")

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2019-03-11")
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.LogCtx(ctx, "level", "error", "message", "error when sending request to the IMDS")
		return nil, microerror.Mask(err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			c.logger.LogCtx(ctx, "level", "error", "message", "error when closing IMDS response body")
		}
	}()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.LogCtx(ctx, "level", "error", "message", "error when reading IMDS response body")
		return nil, microerror.Mask(err)
	}

	var instanceMetadata Metadata
	err = json.Unmarshal(responseBody, &instanceMetadata)
	if err != nil {
		c.logger.LogCtx(ctx, "level", "error", "message", "error when parsing IMDS JSON response")
		return nil, microerror.Mask(err)
	}

	c.logger.LogCtx(ctx, "level", "debug", "message", "successfully fetched instance metadata")

	return &instanceMetadata, nil
}
