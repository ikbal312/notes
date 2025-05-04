resource "aws_vpc" "main_vpc" {
  cidr_block = "10.0.0.0/16"
}
resource "aws_subnet" "main_subnet" {
  vpc_id     = aws_vpc.main_vpc.id
  cidr_block = "10.0.0.0/24"
}

