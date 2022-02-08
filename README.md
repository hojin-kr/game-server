# Issue
main, grpc branch 배포시 AppEngine으로 배포 불가
-> gRPC서버에 대해 GCP AppEngine에서 지원하지 않음
-> gRPC서버 대응을 위해 Cloud Run 배포 추가 예정

# 개요
Google Cloud Platform AppEngine 에 최적화된 확장이 용이한 게임 서버 프레임워크  
gRPC 구현

# 특징
- Google Cloud Platform AppEngine 최적화
	- 자동화 배포
	- Memcache
- Google Cloud Platform Datastore 로Persistent 데이터 관리
	- Account
- stdlogging
- GitoOps Github Action CI/CD 자동화
- gRPC 구현
- TDD, Unit Test 
- 적은 서버 초기 비용
- golang web socket
---

# Using
## Instance Deploy to AppEngine

### Setup GCP service account key & product ID
setting up secrets
```
service_account_key: ${{ secrets.SA_KEY }}
project_id: ${{ secrets.PROJECT_ID }}
```
![image](https://user-images.githubusercontent.com/22079767/144077080-504aeb7c-ae48-4d99-b36c-e6d99216a9ad.png)

### Enable Require GCP APIs AppEngine, Datastore
https://console.cloud.google.com/appengine
https://console.cloud.google.com/datastore

### Run Git Action workflow
- app name : /deployments app_anme
- --no-stop-previous-version : bool
![image](https://user-images.githubusercontent.com/22079767/144077357-0c05438e-87e0-46c0-8ad3-5e1a21380cc3.png)

### gRPC
gRPC를 사용합니다. 
protobuffer 정의 파일 변경후 언어에 맞춰 빌드를 수행
- Server Go
- Client Unity (C#) 

> 참고
> https://hojin-kr.github.io/2022-02-02-Go-서버와-Unity-클라이언트-gRPC-도입/

```
#!/bin/bash
protoc --go_out=./cmd --go_opt=paths=source_relative \
    --go-grpc_out=./cmd --go-grpc_opt=paths=source_relative \
    proto/haru.proto
protoc --csharp_out=./proto/csharp \ 
    --plugin=protoc-gen-csharp_grpc=/Users/hojin/Work/hojin/haru/proto/plugins/grpc_csharp_plugin \
    --csharp_grpc_out=./proto/csharp \
    proto/haru.proto
```

