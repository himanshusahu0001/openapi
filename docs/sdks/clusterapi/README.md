# ClusterAPI
(*ClusterAPI*)

### Available Operations

* [ClusterAPIListClusters](#clusterapilistclusters) - List Clusters

## ClusterAPIListClusters

List Clusters

### Example Usage

```go
package main

import(
	"github.com/himanshusahu0001/openapi/models/components"
	"openapi"
	"context"
	"github.com/himanshusahu0001/openapi/models/operations"
	"log"
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ClusterAPIListClustersRequest](../../models/operations/clusterapilistclustersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.ClusterAPIListClustersResponse](../../models/operations/clusterapilistclustersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
