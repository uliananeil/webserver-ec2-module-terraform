provider "aws" {
  profile = "default"
  region = "eu-central-1"
}

terraform {
    backend "s3" {
    bucket         = "${STATE_BUCKET}"
    key            = "terraform.tfstate"
    region  = "eu-central-1"
  }  
}
