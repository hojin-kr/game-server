# 개요
Google Cloud Platform CloudRun&AppEngine 에 최적화된 확장이 용이한 서버 프레임워크

# 특징
- Google Cloud Platform AppEngine 최적화
	- 자동화 배포
	- Memcache
- Google Cloud Platform Datastore 로Persistent 데이터 관리
	- Account
- stdlogging
- GitoOps Github Action CI/CD 자동화
- REST full API
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

### Check AppEngine Deploy & Ping Test
![image](https://user-images.githubusercontent.com/22079767/144080614-745250b4-4acd-421e-920d-52087308bd8d.png)

https://console.cloud.google.com/appengine/services

```
➜  indie-game-server git:(main) curl https://{{YOUR_APPENGINE_ADDRESS}}.appspot.com/ping
{"message":"pong"}
```


### build submit & run
```
gcloud builds submit -t gcr.io/{project_id}/haru:0.1.0 .
gcloud run services replace deployments/dev/cloudrun.yaml --region=asia-northeast3
```