output "public_ip" {
  value = aws_instance.ec2.public_ip
}

output "vpc_id" {
  value = module.vpc.vpc_id
}

output "security_group_id" {
  value = aws_security_group.ec2_sg.id
}
