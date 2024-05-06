terraform {
  required_providers {
    kind = {
      source = "tehcyx/kind"
      version = "0.4.0"
    }
  }
}

provider "kind" {
  # Configuration options
}

# Create a cluster with kind of the name "test-cluster" with kubernetes version v1.27.1
resource "kind_cluster" "default" {
    name = "test-cluster"
    node_image = "kindest/node:v1.27.1"
}