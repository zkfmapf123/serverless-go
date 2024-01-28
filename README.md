# Serverless-go-deploy-agent

## 초기설정

```sh
  ## help
  ./main

  ## command
  ## profile, region 미기재시 ~/.aws/credential의 Default로 동작합니다.
  ./main [command] -p [profile] -r [region]
```

## 완료된 기능

- [x] Lambda Dashboard
- [ ] Deploy Lambda
- [ ] Rollback Lambda
- [ ] Delete Lambda

## 추후 추가되야 할 작업

- Lambda를 만들어주는건 아직안됨 -> 테라폼으로 관리할수있도록 해야 함

  - 테라폼으로 만들어주되, output 할 수 있도록
  - state bucket도 yml에서 추출하자...

- 현재는 S3 + zip 형태로 하지만, 추후에는 Dockerfile로도 배포할수있도록 해야함
