terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 2.70"
    }
  }
}

provider "aws" {
  profile = "default"
  region = "us-east-1"
}

# Naming with underscore to avoid conflicts with terraform syntax
resource "aws_instance" "golang_app_2020" {
  ami           = "ami-0947d2ba12ee1ff75"
  instance_type = "t2.micro"
  key_name = "golang-app-2020"
  tags = {
    Name = "golang_app_2020"
  }
}
