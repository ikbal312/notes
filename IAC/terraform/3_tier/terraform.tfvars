aws_region   = "ap-southeast-1"
environment  = "dev"
organization = "poridhi"

vpc_name       = "myvpc"
vpc_cidr_block = "10.0.0.0/16"

vpc_availability_zones = ["ap-souteat-1a", "ap-souteat-1b"]
vpc_public_subnets     = ["10.0.101.0/24", "10.0.102.0/24"]
vpc_private_subnets    = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24"]

vpc_enable_nat_gateway = true
vpc_single_nat_gateway = true

