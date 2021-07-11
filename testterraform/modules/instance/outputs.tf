output "instance_name" {
  value = google_compute_instance.testing_instance.name
}

output "instance_cpu" {
  value = google_compute_instance.testing_instance.cpu_platform
}