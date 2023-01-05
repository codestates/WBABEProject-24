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
- **web, swagger, db 관련 부분을 알맞게 설정 필요**

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
- ### 전체 API 요약
    ![전체 API](https://user-images.githubusercontent.com/115597002/210808682-f442d425-a38c-4d7f-bb0e-b85733353e3c.png)

- ### 메뉴 생성
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>body</b> : {"name":"딸기빙수","price":12000,"hotGrade":1,"isAvailable":true}
    </details>

    ![메뉴 생성 요청](https://user-images.githubusercontent.com/115597002/210808958-e8c0c4bb-c12a-4f4d-b3c9-363a7f5b3343.png)
  - **응답**
    ![메뉴 생성 응답](https://user-images.githubusercontent.com/115597002/210808972-e4584eea-2b97-40cb-8158-90c156348464.png)

- ### 메뉴 리스트 출력
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>query</b> : sort=[recommend(추천순)|score(평점순)|most(주문순)|new(최신순)]
    </details>

    ![메뉴 리스트 요청](https://user-images.githubusercontent.com/115597002/210808851-d5551720-96d6-4f32-aa99-95e1091448c4.png)
  - **응답**
    ![메뉴 리스트 응답](https://user-images.githubusercontent.com/115597002/210808875-6d6eef1a-aece-47c4-b6f7-73faaeaf7f82.png)

- ### 메뉴 수정
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : name=메뉴_이름<br>
        <b>body</b> : {"price":13000,"hotGrade":1,"isAvailable":true}
    </details>

    ![메뉴 수정 요청](https://user-images.githubusercontent.com/115597002/210809016-e052b09c-d0fc-4979-8778-10919d98738f.png)
  - **응답**
    ![메뉴 수정 응답](https://user-images.githubusercontent.com/115597002/210809035-7021ecad-9695-44f0-8b2b-767f6a1bb643.png)

- ### 메뉴 삭제
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : name=메뉴_이름
    </details>

    ![메뉴 삭제 요청](https://user-images.githubusercontent.com/115597002/210809079-3f23f968-7429-462d-9077-7150813ae906.png)
  - **응답**
    ![메뉴 삭제 응답](https://user-images.githubusercontent.com/115597002/210809100-e6a391dd-34f5-45db-ade9-94d003a7724c.png)

- ### 주문 생성
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>body</b> : {"menuList":["낙곱새", "치즈김밥"],"address":"부산","phone":"010-1111-2222"}
    </details>

    ![주문 생성 요청](https://user-images.githubusercontent.com/115597002/210809185-5b2fa979-6a0e-4559-9db4-270372e9ef71.png)
  - **응답**
    ![주문 생성 응답](https://user-images.githubusercontent.com/115597002/210809201-ef82c819-34a3-45d0-b591-aeb4fccba2b7.png)

- ### 주문 리스트 출력
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>query</b> : status=[active|complete|all]<br>
        - active : 대기, 주문, 조리, 배달 상태<br>
        - complete : 완료 상태<br>
        - all : 모든 상태
    </details>

    ![주문 리스트 요청](https://user-images.githubusercontent.com/115597002/210809147-1a4b4ca8-b600-4472-b0b3-0fe98a322a1a.png)
  - **응답**
    ![주문 리스트 응답](https://user-images.githubusercontent.com/115597002/210809158-fd9a8ae2-128d-4375-a2ad-6a52420bd4d1.png)

- ### 주문 메뉴 변경
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : seq=order_seq<br>
        <b>path</b> : type=[add|change]<br>
        <b>body</b> : {"menuList": ["낙곱새"]}
    </details>

    ![주문 메뉴 변경 요청-1](https://user-images.githubusercontent.com/115597002/210809238-be291a78-0f7a-4dc1-9dcf-6fc049eba768.png)
    ![주문 메뉴 변경 요청-2](https://user-images.githubusercontent.com/115597002/210809250-f4ad1421-fe4a-4005-bbc1-97a86f2d9812.png)
  - **응답**
    ![주문 메뉴 변경 응답](https://user-images.githubusercontent.com/115597002/210809258-5c274ca0-13c0-469d-8a20-a75d93d23f4a.png)

- ### 주문 상태 변경
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : seq=order_seq<br>
        <b>path</b> : status=[대기|주문|조리|배달|완료]
    </details>

    ![주문 상태 변경 요청](https://user-images.githubusercontent.com/115597002/210809266-f43b647e-0c85-41d0-a6b4-8d5cb5291c58.png)
  - **응답**
    ![주문 상태 변경 응답](https://user-images.githubusercontent.com/115597002/210809285-d9b484a6-efe2-4ce1-86c7-6f7584ca1c91.png)

- ### 리뷰 생성
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>body</b> : {"comment": "너무 맛있어요!","menuName": "낙곱새","orderSeq": "1672929029-0000000001","score": 5}
    </details>

    ![리뷰 생성 요청](https://user-images.githubusercontent.com/115597002/210809375-994ad4b5-c659-4a12-8188-f1f2a7cba65e.png)
  - **응답**
    ![리뷰 생성 응답](https://user-images.githubusercontent.com/115597002/210809399-7d79e312-e759-464d-88ef-03999315fdea.png)

- ### 리뷰 리스트 출력
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : menu=메뉴_이름
    </details>

    ![리뷰 리스트 요청](https://user-images.githubusercontent.com/115597002/210809348-c12e91de-fce0-427a-8be2-8f19d0afa072.png)
  - **응답**
    ![리뷰 리스트 응답](https://user-images.githubusercontent.com/115597002/210809358-570db165-93f9-41a9-8e25-0a2a61fb7c0d.png)
