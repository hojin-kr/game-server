# indie-game-server-architecture
초기 비용 관리 및 용이한 확장성을 고려한 서버 아키텍처
서비스 운영의 불확실성을 고려해서 초기에는 비용이 발생하지 않지만 사용량에 따라 스케일 조정이 용이한 아키텍처

# Concept
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


# Setup
## Instance Deploy to AppEngine
Setup GCP service account key & product ID
setting up github secrets
```
service_account_key: ${{ secrets.SA_KEY }}
project_id: ${{ secrets.PROJECT_ID }}
```
## GCP AppEngine Enable   
![image](https://user-images.githubusercontent.com/22079767/143732125-7f4e9d47-859a-4df7-8d4a-284e90ab5e5c.png)
