resource "google_compute_network" "main" {
  name                    = "main-vpc"
  auto_create_subnetworks = false
}
