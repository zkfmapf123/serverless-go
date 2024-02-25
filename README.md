# Serverless-go-deploy-agent

## Desc
- Lambda 함수 생성 / 삭제 / 배포 Agent

## 실행방법

```sh
  make build

  ## help
  ./agent

  ## command
  ./agent [command] -p [profile] -r [region] -f [function path]
```

## Default Parameter

|     | 설명               | 기본값         |
| --- | ------------------ | -------------- |
| -p  | profile            | default        |
| -r  | region             | ap-northeast-2 |
| -f  | function root path | functions      |

## 완료된 기능

- [x] Lambda Dashboard
- [x] Create Lambda
- [x] Delete Lambda
- [ ] Deploy Lambda
  - [x] 간단한 배포
  - [ ] 세부사항 변경 하는 배포
  - [ ] Versioning S3 + 배포
- [ ] Rollback Lambda

## 사용방법

1. Global Setting (yaml)

- Global Setting 값은 agent.yml을 활용합니다.

```yml
configs:
  state_bucket: dk-s3-lambda-state ## Required
  region: ap-northeast-2 ## Required
  profile: default ## Required
```

2. 함수 Setting (yaml)
- 람다함수 생성 / 배포에 필요한 구성을 정의합니다.
- Tag는 Pascal Case로 치환됩니다.
- 환경변수는 Uppercase로 치환됩니다.

```yml
## Required
configs:
  state_s3_bucket: lambda-dk-s3-bucket
  role_arn: Basic-Lambda-Role
  filename: function.zip
  handler: handler.go
  function_name: add_function

## Required
handler_config:
  timeout: 60
  memory_size: 128

## Required (Only Uppercase)
tags:
  Env: test

## Required (Only Uppercase)
envs:
  var_A: 10
  var_B: 20
```

## Dependency

- <a href="https://registry.terraform.io/modules/zkfmapf123/lambda/lee/latest">Terraform Lambda Module</a>
- <a href="https://docs.aws.amazon.com/ko_kr/code-library/latest/ug/go_2_iam_code_examples.html"> aws-sdk-v2-golang lib </a>