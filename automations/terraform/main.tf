provider "google" {
    project = "lets-try-tech-ltt-sys-prod"
    region = "us-east1"
    zone = "us-east1-b"
}

resource "google_compute_instance" "vm_home-page" {
    name = "home-page"
    machine_type = "f1-micro"
    zone = "us-east1-b"
    
    boot_disk {
        initialize_params {
            image = "debian-cloud/debian-9"
        }
    }
    network_interface {
        network = "default"
        access_config = {}
    }
}