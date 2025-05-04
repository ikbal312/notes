
resource "aws_internet_gateway" "main_igw" {
  vpc_id = aws_vpc.main_vpc.id

}
resource "aws_route_table" "main_public" {
  vpc_id = aws_vpc.main_vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main_igw.id
  }
  tags = {
    Name = "rt_main_public"
  }
}
resource "aws_vpc" "main_vpc" {
  cidr_block = "10.0.0.0/16"
}
resource "aws_subnet" "main_subnet" {
  vpc_id     = aws_vpc.main_vpc.id
  cidr_block = "10.0.0.0/24"
}
resource "aws_route_table_association" "main_public_and_main_subnet_association" {
  subnet_id      = aws_subnet.main_subnet.id
  route_table_id = aws_route_table.main_public.id
}

resource "aws_key_pair" "web_key" {
  key_name   = "web_key"
  public_key = file("~/.ssh/web_key.pub") # need ssh key pair
}

resource "aws_security_group" "ssh_access" {
  name        = "sg_ssh"
  description = "allow ssh connection"
  vpc_id      = aws_vpc.main_vpc.id
  tags = {
    Name = "ssh_access"
  }
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]

  }
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]

  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# resource "aws_vpc_security_group_ingress_rule" "allow_ssh" {
#   security_group_id = aws_security_group.ssh_access.id
#   ip_protocol       = "tcp"
#   cidr_ipv4         = "0.0.0.0/0"
#   from_port         = "22"
#   to_port           = "22"
# }
# resource "aws_vpc_security_group_egress_rule" "allow_all" {
#   security_group_id = aws_vpc_security_group_ingress_rule.allow_ssh.id
#   cidr_ipv4         = "0.0.0.0/0"
#   ip_protocol       = "-1"

# }
resource "aws_instance" "web" {
  ami                         = "ami-060e277c0d4cce553" # Ubuntu 20.04 AMI ID
  instance_type               = "t2.micro"
  key_name                    = aws_key_pair.web_key.key_name
  subnet_id                   = aws_subnet.main_subnet.id
  vpc_security_group_ids      = [aws_security_group.ssh_access.id]
  associate_public_ip_address = true
  tags = {
    name        = "web_server"
    description = "An nginx web server on ubuntu"
  }
  # user_data = <<-EOF
  #   #!/bin/bash
  #   apt-get update
  #   apt-get install -y nginx
  # EOF
  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = file("~/.ssh/web_key")
      host        = self.public_ip
    }
    inline = [
      "sudo apt update -y",
      "sudo apt install -y nginx",
      "sudo systemctl enable nginx",
      "sudo systemctl start nginx"
    ]
  }
  provisioner "local-exec" {
    command = "echo instance${self.public_ip} crated! > instance_ip.txt"
  }



  # provisioner "remote-exec" {
  #   when = destroy
  #   connection {
  #     type        = "ssh"
  #     user        = "ubuntu"
  #     private_key = file("~/.ssh/web_key")
  #     host        = self.public_ip
  #   }
  #   inline = [
  #     "sudo systemctl stop nginx",
  #     "sudo apt remove -y nginx"
  #   ]
  # }
  # provisioner "local-exec" {
  #   when    = destroy
  #   command = "echo instance${self.public_ip} destroyed! > instance_ip.txt"
  # }


}

