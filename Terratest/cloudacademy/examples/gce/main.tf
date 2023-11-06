terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.75.0"
    }
  }
}

provider "google" {
  project = "nova-vx"
  region  = "us-central1"
}

module "webserver" {
  source       = "../../"
  name         = var.name
  machine_type = var.machine_type
  zone         = var.zone
  project_id   = "nova-vx"
  region       = "us-central1"
}
