# ClusterAPIListClustersRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Offset`                                              | **int*                                                | :heavy_minus_sign:                                    | offset from where to start the list of clusters.      |
| `Limit`                                               | **int*                                                | :heavy_minus_sign:                                    | limit the response to specified count.                |
| `State`                                               | [*operations.State](../../models/operations/state.md) | :heavy_minus_sign:                                    | filter list based on state of cluster.                |
| `GetMetrics`                                          | **bool*                                               | :heavy_minus_sign:                                    | get cluster metrics information.                      |
| `Search`                                              | **string*                                             | :heavy_minus_sign:                                    | text search for cluster.                              |