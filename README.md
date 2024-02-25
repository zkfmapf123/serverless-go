# Serverless-go-deploy-agent

## Desc
- Lambda 함수 생성 / 삭제 / 배포 Agent
- Only Golang...

## 실행방법

```sh
  make build

  ## help
  ./agent

  ## command
  ./agent [command] -p [profile] -r [region] -f [function path]

  ## 혹시 나는 명령어로써 사용하고 싶다.!!
  mv agent /usr/local/bin
  agent -f <function-path> command
```

## Default Parameter

|     | 설명               | 기본값         |
| --- | ------------------ | -------------- |
| -p  | profile            | default        |
| -r  | region             | ap-northeast-2 |
| -f  | function root path | functions      |

## 사용방법

```sh
  project-root
    |- functions
      |- add
        |- config.yml ## add Function Setting
        |- ...
      |- min
        |- config.yml ## min Function Setting 
        |- ... 
    |- agent.yml      ## Global Setting
    |- agent  
```

1. Global Setting (agent.yaml)

- project-root/agent.yaml
- Global Setting 값은 agent.yml을 활용합니다.

```yml
configs:
  state_bucket: dk-s3-lambda-state ## Required
  region: ap-northeast-2 ## Required
  profile: default ## Required
```

2. 함수 Setting (config.yaml)

- path/function-path/config.yaml 
- 람다함수 생성 / 배포에 필요한 구성을 정의합니다.
- Required Options은 모두 기재하여야 합니다.
- Tag는 Pascal Case로 치환됩니다.
- 환경변수는 Uppercase로 치환됩니다.

```yml
## Required
configs:
  state_s3_bucket: lambda-dk-s3-bucket      ## Lambda의 S3 Version에 필요한 버켓
  role_arn: Basic-Lambda-Role               ## Lambda에 사용될 IAM 권한
  function_name: add_function               ## Lambda 함수 이름

## Required
handler_config:
  timeout: 60                               
  memory_size: 128

## Required (Translate Pascal Case)
tags:
  Env: test

## Required (Translate Upper Case)
envs:
  var_A: 10
  var_B: 20
```

## Dependency

- <a href="https://registry.terraform.io/modules/zkfmapf123/lambda/lee/latest">Terraform Lambda Module</a>
- <a href="https://docs.aws.amazon.com/ko_kr/code-library/latest/ug/go_2_iam_code_examples.html"> aws-sdk-v2-golang lib </a>