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
