provider "google" {
    project = "lets-try-tech-ltt-sys-prod"
    region = "us-east1"
    zone = "us-east1-b"
}

resource "google_compute_address" "ip_vm_home_page" {
  name = "ip-home-page"
}

resource "google_compute_firewall" "default" {
    name    = "http-ports-permission"
    network = "default"
    description ="Allow HTTP/HTTPS from anywhere"
    
    allow {
        protocol = "tcp"
        ports    = ["80","443"]
    }
}

resource "google_compute_instance" "vm_home_page" {
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
        access_config {
            nat_ip = "${google_compute_address.ip_vm_home_page.address}"
        }
    }
}