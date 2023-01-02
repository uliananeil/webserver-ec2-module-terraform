provider "aws" {
  profile = "default"
  region = "eu-central-1"
}

terraform {
    backend "s3" {
    bucket         = "state-bucket-134681"
    key            = "terraform.tfstate"
    region  = "eu-central-1"
  }  
}
