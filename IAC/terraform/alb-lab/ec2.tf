resource "aws_security_group" "ec2_sg" {
  name        = "EC2SecurityGroup"
  description = "security group for ec2 instances"
  vpc_id      = aws_vpc.my_vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
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
}



resource "aws_key_pair" "web_key" {
  key_name   = "web_key"
  public_key = file("~/.ssh/web_key.pub")
}

resource "aws_instance" "web_server" {
  count                       = 2
  ami                         = "ami-060e277c0d4cce553"
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.public_subnets[count.index].id
  vpc_security_group_ids      = [aws_security_group.ec2_sg.id]
  key_name                    = aws_key_pair.web_key.key_name
  associate_public_ip_address = true

  user_data = <<-EOF
            #!/bin/bash
            apt-get update
            apt-get install apache2 -y 
            HOSTNAME=$(ostname)
            IP_ADDRESS=$(hostname -I | cut -d' ' -f1)
            echo "<html><head><title>Server Details</title></head><body><h1>Server Details</h1><p><strong>Hostname:</strong> $HOSTNAME</p><p><strong>IP Address:</strong> $IP_ADDRESS</p></body></html>" | tee /var/www/html/index.html
            systemctl restart apache2    
            EOF

  tags = {
    Name = "web-server-${count.index + 1}"
  }
}


