locals {
  owners      = var.organization
  environment = var.environment
  name        = "${var.organization}-${var.environment}"
  common_tags = {
    Owner       = local.owners
    Environment = local.environment
  }
}


