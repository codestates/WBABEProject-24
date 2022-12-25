# Online Ordering System
> Wemade Blockchain Academy 2022 Back-End Final Project

## 프로젝트 소개
주문자(고객)와 피주문자(사업자)를 위한 온라인 주문 시스템 구현

## 기술 스택
- Go
- Gin-Gonic
- MongoDB

## 개발 환경 구성
- Go 설치
- MongoDB 설치

## 설정
- conf/config.toml

## 실행 방법
```
git clone https://github.com/codestates/WBABEProject-24.git
cd WBABEProject-24
go mod tidy
go run main.go
```

## DB 구성
![erd](https://user-images.githubusercontent.com/115597002/209469922-4c4f85fe-3065-4417-8754-569d51c6742d.PNG)

## 기능 소개
- 전체 API
![전체 API](https://user-images.githubusercontent.com/115597002/209470202-08b07bdc-b65b-43fb-8008-205ab6dd530a.PNG)
- 메뉴 생성
![메뉴 생성](https://user-images.githubusercontent.com/115597002/209469949-bb2c2fc6-81b6-4d12-9784-05e62d4f21cc.PNG)
- 메뉴 리스트 조회
![메뉴 리스트](https://user-images.githubusercontent.com/115597002/209470019-91ef9724-0060-4553-90c1-adf6f6223227.PNG)
- 메뉴 변경
![메뉴 변경](https://user-images.githubusercontent.com/115597002/209470038-27c47543-a92f-41f9-b051-622f59329fed.PNG)
- 메뉴 삭제 (실제 데이터를 삭제하지 않고, delete 플래그 설정)
![메뉴 삭제](https://user-images.githubusercontent.com/115597002/209469991-cfcb3d80-c0bf-48f7-8539-7082ade8dd8a.PNG)
- 주문 리스트 조회
![주문 리스트](https://user-images.githubusercontent.com/115597002/209470000-62d4f3d1-5b3e-42f4-b0dd-ca8fd4a09f0b.PNG)
- 주문 생성
![주문 생성](https://user-images.githubusercontent.com/115597002/209469957-ac3222e9-878e-4ed1-9569-ded243792687.PNG)
- 주문 메뉴 변경
![주문 변경](https://user-images.githubusercontent.com/115597002/209470030-fae3f173-f153-4b3c-87c4-f22fbdeb4f02.PNG)
- 주문 상태 변경
![주문 상태 변경](https://user-images.githubusercontent.com/115597002/209470052-9636149e-b9c7-4fdc-a6a1-8a01e68099b8.PNG)
- 리뷰 생성
![리뷰 생성](https://user-images.githubusercontent.com/115597002/209469968-792108a8-0386-4c68-8450-d33101e62329.PNG)
- 리뷰 리스트 조회
![리뷰 리스트](https://user-images.githubusercontent.com/115597002/209470007-db854059-ca9e-4aec-b0bc-4a388b589a28.PNG)