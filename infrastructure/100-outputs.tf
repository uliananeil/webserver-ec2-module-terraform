# Display dns information

output "http_ip" {
  value = {
    for instance in aws_instance.http :
    instance.id => instance.private_ip
  }
}

output "db_ip" {
  value = {
    for instance in aws_instance.db :
    instance.id => instance.private_ip
  }
}

output "vpc_cidr_block" {
  value = aws_vpc.terraform.cidr_block
}

output "http_subnet_cidr_block" {
  value = aws_subnet.http.cidr_block
}

output "db_subnet_cidr_block" {
  value = aws_subnet.db.cidr_block
}