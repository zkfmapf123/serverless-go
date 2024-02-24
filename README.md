# Serverless-go-deploy-agent

## 초기설정

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
  - [ ] Versioning S3
  - [ ] Deploy
- [ ] Rollback Lambda + S3 Versioning

## Dependency

- <a href="https://registry.terraform.io/modules/zkfmapf123/lambda/lee/latest">Terraform Lambda Module</a>
- <a href="https://docs.aws.amazon.com/ko_kr/code-library/latest/ug/go_2_iam_code_examples.html"> aws-sdk-v2-golang lib </a>