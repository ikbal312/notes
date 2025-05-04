provider "aws" {
  region = "ap-southeast-1"
}

resource "aws_vpc" "main_vpc" {
  cidr_block         = "10.0.0.0/16"
  enable_dns_support = true
  tags = {
    Name = "main_vpc"
  }


}
