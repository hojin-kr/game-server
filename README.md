# indie-game-server-architecture
초기 비용 관리 및 용이한 확장성을 고려한 서버 아키텍처
서비스 운영의 불확실성을 고려해서 초기에는 비용이 발생하지 않지만 사용량에 따라 스케일 조정이 용이한 아키텍처

# Architecture Concept
## 인스턴스
### gcp app engine
- 쉬운 배포
- 자유로운 스케일 조정
- 멤캐시 지원 (개런티는 없지만 제한없이 가능)
- 무료범위:(f1티어 1인스턴스730h)
---
## 동적 저장소
### gcp datastore
- 초기 비용 발생하지 않음
- 완전 관리형 서비스
- 무료범위:(1기가 저장, 한달 5만 읽기, 2만 쓰기)
유저 정보 및 공용보스등 경쟁요소 저장 및 기타 서버에 동적 데이터 저장 필요한 것들
---
## 정적 저장소
### github repository
- 버전 관리
게임 코드와 분리된 프라이빗 레포지토리로 세팅
raw 포맷으로 제공하거나 packages 서비스 사용
도메인과 부가정보를 정적 데이터에 넣어서 점검 상태 등등 컨트롤 가능하도록
---
## 서버 컨테이너 저장소
### github packages
- github CI/CD 통합
서버 Container image 저장소 필요
github packages쓰고 용량을 관리
---
## CI/CD
### github action
- github CI/CD 통합
build해서 packages에 올리고 원하면 배포까지
---
## 도메인
정적 데이터로 도메인 관리하고 앱 실행시 받아가서 사용하도록 해서 도메인 변경 가능하도록

---
## 로깅
stdout

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
# Models
```

// Account account infomation
type Account struct {
	RegisterTimestamp int64
	DeviceID          string
	GoogleID          string
	AppleID           string
	LineID            string
	KakaoID           string
	ID                int64 `datastore:"-"`
}

// Profile profile inforamtion
type Profile struct {
	Nickname string
	ID       int64 `datastore:"-"`
}

// Wallet wallet
type Wallet struct {
	Coin string
	ID   int64 `datastore:"-"`
}
```

# APIs
```
[GIN-debug] POST   /api/v1/accounts/get      --> main.getAccount (3 handlers)
[GIN-debug] POST   /api/v1/profiles/get      --> main.getProfile (3 handlers)
[GIN-debug] POST   /api/v1/profiles/set      --> main.setProfile (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
```
