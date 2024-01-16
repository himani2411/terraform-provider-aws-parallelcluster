---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pcluster_compute_fleet_status Resource - pcluster"
subcategory: ""
description: |-
  Update the status of the cluster compute fleet.
---

# pcluster_compute_fleet_status (Resource)

Update the status of the cluster compute fleet.

## Example Usage

```terraform
resource "pcluster_compute_fleet_status" "example_fleet" {
  cluster_name   = var.cluster_name
  status_request = "START_REQUESTED"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) The name of the cluster.
- `status_request` (String) The compute fleet status.

### Optional

- `region` (String) The AWS Region that the cluster is in.

### Read-Only

- `id` (String) ComputeFleetStatus identifier
- `last_status_update_time` (String) The timestamp that represents the last status update time.
- `status` (String) The compute fleet status.