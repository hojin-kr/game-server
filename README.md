# 개요
Go 게임 서버 프레임워크 개발

# GCP Architecture Diagram
![image](https://user-images.githubusercontent.com/22079767/154786972-c1e06d6a-15db-4eec-8c35-cb1e29434541.png)

# 특징
- 확장성 (비용 0의 개발환경에서 편한 프로덕션 전환)
- GCP 매니지드 서비스 서버
- gRPC 도입
- GitOps

# 스펙
- Go lang
- gRPC
- go-gin (rest)
- GCP Datastore
- GCP AppEngine (rest)
- GCP CloudRun

---

# 배포
## Deploy REST API Server
REST API 서버 배포는 GCP AppEngine을 사용합니다.
REST API 서버는 `rest` 브랜치에 작업되어있습니다.
```
git checkout rest
```
1.  Setup GCP service account key & Project ID
GCP 서비스 어카운트및 프로젝트 아이디를 github repository Secrets 에 정의합니다.  

```
service_account_key: ${{ secrets.SA_KEY }}
project_id: ${{ secrets.PROJECT_ID }}
```
![image](https://user-images.githubusercontent.com/22079767/144077080-504aeb7c-ae48-4d99-b36c-e6d99216a9ad.png)

2. Enable Require GCP APIs AppEngine, Datastore
GCP의 서비스를 처음 사용한다면 GCP APIs에서 API 사용을 설정합니다.

https://console.cloud.google.com/appengine
https://console.cloud.google.com/datastore

3. Run Git Action workflow
배포 정보를 deployments에서 배포 타겟 별로 관리합니다.

- app name : /deployments app_anme
- --no-stop-previous-version : bool
![image](https://user-images.githubusercontent.com/22079767/144077357-0c05438e-87e0-46c0-8ad3-5e1a21380cc3.png)

### Deploy gRPC SERVER
gRPC 서버 배포는 GCP CloudRun을 사용합니다.
gRPC 서버는 `main` 브랜치에 작업되어있습니다.   

```
# 빌드 및 푸시
gcloud builds submit . --tag=gcr.io/${PROJECT_ID}/haru
# test/cloudrun.yaml의 정의에 맞춰 클라우드런에 배포
gcloud run services replace deployments/test/cloudrun.yaml --region=asia-northeast3
```


# gRPC에 대해서 
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
