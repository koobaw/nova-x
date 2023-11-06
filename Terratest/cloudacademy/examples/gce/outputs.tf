output "ip" {
    description = "IP address of the instance"
    value = module.webserver.ip
}

output "instance_id" {
    description = "ID of GCE instance"
    value = module.webserver.instance_id
}

output "storage" {
  value = module.webserver.storage.id
}