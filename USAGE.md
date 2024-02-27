<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"openapi"
	"github.com/himanshusahu0001/openapimodels/components"
	"github.com/himanshusahu0001/openapimodels/operations"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ClusterAPI.ClusterAPIListClusters(ctx, operations.ClusterAPIListClustersRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.ConfigClusterStatusList != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->