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
      "iam.googleapis.com",
      "apigateway.googleapis.com",
      "servicecontrol.googleapis.com",
      "secretmanager.googleapis.com",
      "certificatemanager.googleapis.com",
      "cloudbuild.googleapis.com",
      "artifactregistry.googleapis.com",
      "compute.googleapis.com",
      "cloudresourcemanager.googleapis.com",
      "run.googleapis.com",
      "containerscanning.googleapis.com",
      #下↓部分はFirebase用
      "maps-android-backend.googleapis.com",
      "maps-ios-backend.googleapis.com",
      "cloudbilling.googleapis.com",
      "firebase.googleapis.com",
      "serviceusage.googleapis.com",
      "apikeys.googleapis.com",
      "firebaserules.googleapis.com",
      "firebasestorage.googleapis.com",
      "storage.googleapis.com",
      "cloudscheduler.googleapis.com"
    ]
  )

  service            = each.key
  project            = var.project_id
  disable_on_destroy = false
}

