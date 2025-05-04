output "vpc_id" {
  description = "The ID of the vpc"
  value       = module.vpc.vpc_id
}


output "public_subnet_ids" {
  description = "list of ids for the public subnets"
  value       = module.vpc.public_subnets
}

output "private_subnet_ids" {
  description = "list ids for the private subnets"
  value       = module.vpc.private_subnets
}

output "nat_gateway_public_ips" {
  description = "list of public elastic ips created for aws NAT Gateway"
  value       = module.vpc.nat_public_ips
}

output "vpc_cidr_block" {
  description = "the cidr block of the vpc"
  value       = module.vpc.vpc_cidr_block
}
