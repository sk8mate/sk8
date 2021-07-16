terraform { 
    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = "~> 3.27"
        }
    }
    required_version = "1.0.1"
}

provider "aws" {
    region  = "eu-west-1"
}

provider "aws" {
    alias = "us_east_1"
    region = "us-east-1"
}

# Backend
resource "aws_s3_bucket" "terraform_state" {
  bucket = "sk8-terraform-state"

  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }
}

resource "aws_dynamodb_table" "terraform_locks" {
  name         = "sk8_terraform_locks"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }
}

terraform {
  backend "s3" {
    bucket         = "sk8-terraform-state"
    key            = "terraform.tfstate"
    region         = "eu-west-1"
    dynamodb_table = "sk8_terraform_locks"
    encrypt        = true
  }
}
