provider "aws" {
  region = "eu-central-1"
}

terraform {
    backend "s3" {
    bucket         = "terraform-state-bucket-12645"
    key            = "test/terraform.tfstate"
    region  = "eu-central-1"
  }  
}
