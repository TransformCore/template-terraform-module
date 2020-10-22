terraform {
  required_version = ">= 0.13"
}

provider "aws" {
  region = var.region
}

variable "region" {
  type        = string
  description = "(required) region in aws"
}

resource "aws_instance" "complete" {
  ami                    = "ami-05c424d59413a2876"
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.complete_sg.id]

  user_data = <<EOF
#!/bin/bash
echo "Hello, ENGINE" > index.html
nohup busybox httpd -f -p 8080 &
EOF
}

resource "aws_security_group" "complete_sg" {
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "public_ip" {
  value = aws_instance.complete.public_ip
}
