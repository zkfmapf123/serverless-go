provider "aws" {
  profile = lookup(var.provider_config, "profile")
  region  = lookup(var.provider_config, "region")
}
