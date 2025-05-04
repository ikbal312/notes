variable "vpc_name" {
  description = "Name of the vpc"
  type        = string
  default     = "myvpc"
}

variable "vpc_cidr_block" {
  description = "CIDR Block for the vpc"
  type        = string
  default     = "10.0.0.0/16"
}

variable "vpc_availability_zones" {
  description = "list of availability zones"
  type        = list(string)
  default     = ["ap-souteast-1a", "ap-southeast-1b"]
}

variable "vpc_public_subnets" {
  description = "CIDR Block for public subnet"
  type        = list(string)
  default     = ["10.0.101.0/24", "10.0.102.0/24"]
}

variable "vpc_private_subnets" {
  description = "CIDR Block for private subnet"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24"]
}

variable "vpc_enable_nat_gateway" {
  description = "Enable NAT Gateways for Private Subnet"
  type        = bool
  default     = true
}


variable "vpc_single_nat_gateway" {
  description = "Use a single NAT Gateway to reduce costs"
  type        = bool
  default     = true
}

