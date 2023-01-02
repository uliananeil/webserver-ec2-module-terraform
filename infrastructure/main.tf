provider "aws" {
  profile = "default"
  region = "eu-central-1"
}

terraform {
    backend "s3" {
    bucket         = "terraform-state-bucket-12645"
    key            = "terraform.tfstate"
    region  = "eu-central-1"
  }  
}
