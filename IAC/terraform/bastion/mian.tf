provider "aws" {
  region = "ap-southeast-1"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "4.0.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["ap-southeast-1a"]
  private_subnets = ["10.0.1.0/24"]
  public_subnets  = ["10.0.2.0/24"]

  enable_nat_gateway = true
  single_nat_gateway = true

  enable_dns_hostnames = true
  enable_dns_support   = true

  map_public_ip_on_launch = true

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

resource "aws_security_group" "public-sg" {

  vpc_id = module.vpc.vpc_id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "public-sg"
  }
}

resource "aws_security_group" "private-sg" {
  vpc_id = module.vpc.vpc_id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["10.0.2.0/24"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "private-sg"
  }
}

resource "aws_key_pair" "main" {
  key_name   = "main-key"
  public_key = file("~/.ssh/web_key.pub")
}

resource "aws_instance" "public" {
  ami           = "ami-060e277c0d4cce553"
  instance_type = "t2.micro"
  subnet_id     = module.vpc.public_subnets[0]
  key_name      = aws_key_pair.main.key_name
  tags = {
    Name = "public-instance"
  }
  security_groups = [aws_security_group.public-sg.id]
}


resource "aws_instance" "private" {
  ami           = "ami-060e277c0d4cce553"
  instance_type = "t2.micro"
  subnet_id     = module.vpc.private_subnets[0]
  key_name      = aws_key_pair.main.key_name
  tags = {
    Name = "private-instance"
  }
  security_groups = [aws_security_group.private-sg.id]
}


output "public_ins" {
  value = aws_instance.public.public_ip

}
