package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/azure-imds-client/pkg/imds/instance"
)

func main() {
	ctx := context.Background()
	client := &http.Client{}

	logger, err := micrologger.New(micrologger.Config{})
	if err != nil {
		panic(err)
	}

	instanceClientConfig := instance.ClientConfig{
		Logger: logger,
		HttpClient: client,
	}
	instanceClient, err := instance.NewClient(instanceClientConfig)
	if err != nil {
		panic(err)
	}

	instanceMetadata, err := instanceClient.GetMetadata(ctx)
	if err != nil {
		panic(err)
	}

	prettyJson, err := json.MarshalIndent(instanceMetadata, "", "    ")
	fmt.Print(string(prettyJson))
}
