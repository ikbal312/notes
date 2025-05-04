# input variable

variable "aws_region" {
  description = "Region where aws resource will be created"
  type        = string
  default     = "ap-southeast-1"
}

variable "environment" {
  description = "Environment Variable used as  a prefix (e.g.: dev, prod)"
  type        = string
  default     = "dev"
}



variable "organization" {
  description = "Organization for organizational purpose"
  type        = string
  default     = "poridhi"
}


