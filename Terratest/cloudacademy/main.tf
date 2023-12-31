resource "google_compute_instance" "server" {
  name         = var.name
  machine_type = var.machine_type
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2204-lts"
    }
  }

  network_interface {
    subnetwork = "default"
    access_config {}
  }
}

resource "google_project_service" "enable_facilities" {
  for_each = toset(
    [
      "firebasestorage.googleapis.com",
      "storage.googleapis.com"
    ]
  )

  service            = each.key
  project            = var.project_id
  disable_on_destroy = false
}

