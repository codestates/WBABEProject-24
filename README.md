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
- #### 전체 API 요약
    ![전체 API 요약](https://user-images.githubusercontent.com/115597002/209684052-a4b78738-4fb6-4f39-974a-0ceef820987a.png)

- ### 메뉴 생성
  - **요청**
  <details open>
    <summary>요청 데이터</summary>
    <b>body</b> : {"name":"치즈김밥","price":7000,"hotGrade":1,"isAvailable":true}
  </details>

    ![메뉴 생성 요청](https://user-images.githubusercontent.com/115597002/209683063-3c7fc7a2-e47d-4ea4-8841-218ea17eaa2a.png)
  - **응답**
    ![메뉴 생성 응답](https://user-images.githubusercontent.com/115597002/209683091-8c3c1026-8015-4707-9873-19551851a423.png)

- ### 메뉴 리스트 출력
  - **요청**
   <details open>
    <summary>요청 데이터</summary>
    <b>query</b> : sort=[recommend(추천순)|score(평점순)|most(주문순)|new(최신순)]
  </details>

    ![메뉴 리스트 요청](https://user-images.githubusercontent.com/115597002/209682679-253965c2-014a-4737-8ae2-39c9a5257df6.png)
  - **응답**
    ![메뉴 리스트 응답](https://user-images.githubusercontent.com/115597002/209682710-a6f9aa21-2328-4544-a90e-efddf4ff4729.png)

- ### 메뉴 수정
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : name=메뉴_이름<br>
        <b>body</b> : {"price":18000,"hotGrade":4,"isAvailable":true}
    </details>

    ![메뉴 수정 요청](https://user-images.githubusercontent.com/115597002/209682890-7141310e-9b35-4d4c-8616-9439adce538c.png)
  - **응답**
    ![메뉴 수정 응답](https://user-images.githubusercontent.com/115597002/209682946-be534416-5e0e-4856-85fd-fb2852d988df.png)

- ### 메뉴 삭제
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : name=메뉴_이름
    </details>

    ![메뉴 삭제 요청](https://user-images.githubusercontent.com/115597002/209682998-e05052ca-92f0-48bf-9bdb-e48f739833d5.png)
  - **응답**
    ![메뉴 삭제 응답](https://user-images.githubusercontent.com/115597002/209683031-9cd550ee-6502-46da-8f7c-196b339c111e.png)

- ### 주문 생성
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>body</b> : {"menuList":["낙곱새", "치즈김밥"],"address":"인천","phone":"010-1111-2222"}
    </details>

    ![주문 생성 요청](https://user-images.githubusercontent.com/115597002/209683387-4d1e8619-ae25-439c-9cd2-6d30fa50bbd2.png)
  - **응답**
    ![주문 생성 응답](https://user-images.githubusercontent.com/115597002/209683416-7e303fe0-37d2-4b2a-9618-f53fade97a32.png)

- ### 주문 리스트 출력
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>query</b> : status=[active|deactive|complete|all]<br>
        - active : 대기, 주문, 조리, 배달 상태<br>
        - deactive : active에 해당하지 않는 상태 <br>
        - complete : 완료 상태<br>
        - all : 모든 상태
    </details>

    ![주문 리스트 요청](https://user-images.githubusercontent.com/115597002/209683131-52d5b8a9-b61e-4a04-8bfb-933b26b432a6.png)
  - **응답**
    ![주문 리스트 응답](https://user-images.githubusercontent.com/115597002/209683196-00c7d6b3-7f55-42ff-af9e-db90ddd45abd.png)

- ### 주문 메뉴 변경
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : seq=order_seq<br>
        <b>path</b> : type=[add|change]<br>
        <b>body</b> : {"menuList": ["낙곱새"]}
    </details>

    ![주문 메뉴 변경 요청-1](https://user-images.githubusercontent.com/115597002/209683226-0b30afec-9ce3-4e41-9e71-73c6d534c3c5.png)
    ![주문 메뉴 변경 요청-2](https://user-images.githubusercontent.com/115597002/209683267-f72d1170-ca61-4e34-b264-d04c4e7bc342.png)
  - **응답**
    ![주문 메뉴 변경 응답](https://user-images.githubusercontent.com/115597002/209683298-3389d89e-15eb-411f-ac91-3c4f4e0aa9ce.png)

- ### 주문 상태 변경
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : seq=order_seq<br>
        <b>path</b> : status=[대기|주문|조리|배달|완료]
    </details>

    ![주문 상태 변경 요청](https://user-images.githubusercontent.com/115597002/209683321-d2280e83-b65a-4840-8ad4-dc5d29721dce.png)
  - **응답**
    ![주문 상태 변경 응답](https://user-images.githubusercontent.com/115597002/209683361-443e51e9-b071-4259-9703-aac10ec4d1bc.png)

- ### 리뷰 생성
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>body</b> : {"comment": "너무 맛있어요!","menuName": "치즈김밥","orderSeq": "1672149611-0000000002","score": 4}
    </details>

    ![리뷰 생성 요청](https://user-images.githubusercontent.com/115597002/209682618-07328b94-26e4-4f23-9a34-1488845d47de.png)
  - **응답**
    ![리뷰 생성 응답](https://user-images.githubusercontent.com/115597002/209682647-b4d02dbe-3760-4f71-b03c-fbc596178a67.png)

- ### 리뷰 리스트 출력
  - **요청**
    <details open>
        <summary>요청 데이터</summary>
        <b>path</b> : menu=메뉴_이름
    </details>

    ![리뷰 리스트 요청](https://user-images.githubusercontent.com/115597002/209682527-74c0bfdf-c4c7-4bf8-92de-0f87b52e70b6.png)
  - **응답**
    ![리뷰 리스트 응답](https://user-images.githubusercontent.com/115597002/209682575-7620ead6-ad32-432a-9737-75bb6f645c3e.png)
