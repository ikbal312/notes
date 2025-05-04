module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.2.0"

  # Basic details
  name            = "${local.name}-${var.vpc_name}"
  cidr            = var.vpc_cidr_block
  azs             = var.vpc_availability_zones
  public_subnets  = var.vpc_public_subnets
  private_subnets = var.vpc_private_subnets

  # NAT Gateway - outbount communication
  enable_nat_gateway = var.vpc_enable_nat_gateway
  single_nat_gateway = var.vpc_single_nat_gateway

  # VPC DNS parameter
  tags     = local.common_tags
  vpc_tags = local.common_tags

  # Additional Tags for subnet
  public_subnet_tags = {
    Type = "Public Subnets"
  }
  private_subnet_tags = {
    Type = "Private Subnet"
  }
}
