module "lambda" {
  source = "../"

  provider_config = {
    profile = "default"
    region  = "ap-northeast-2"
  }

  lambda_source_code = base64encode("./function.zip")

  lambda_config = {
    iam_name = "Basic-Lambda-Role"
    name     = "test-lambda"
  }

  lambda_hardware = {
    timeout     = "10"
    memory_size = "128"
    runtime     = "nodejs18.x"
  }

  env_vars = {
    "a" : 10,
    "b" : 20,
    "c" : 30
  }

  tags = {
  }
}
