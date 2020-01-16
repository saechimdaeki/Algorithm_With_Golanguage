
## https://github.com/Jeiwan/blockchain_go/tree/part_1  (branch별로 바꿔가며 따라해보면 잘 작동 )

##### 코드설명 한글로 해주는 블로그 https://mingrammer.com/building-blockchain-in-go-part-1/

### 블록체인은 특정한 구조를 지닌 데이터베이스로 순서가 지정된 링크드 리스트이다
- 이런점을 Go언어에서는 배열과 맵을 활용하여 구현할수있음 (배열은 해시를 유지 맵은 해시블록쌍 유지 )


### 해시캐시(Hashcash)
- 1.공개적으로 알려진 데이터를 가져온다 (이메일의 경우 수신자의 이메일 주소, 비트코인의 경우 블록의 헤더를 들 수 있다)
- 2.여기에 카운터를 더한다. 카운터는 0부터 시작한다.
- 3.data + counter의 해시를 구한다.
- 4.해시가 특정 요구사항을 충족하는지 확인한다.
- 1)만족한다면 알고리즘을 끝낸다.
- 2)그렇지 않다면 카운터를 증가시켜 3번과 4번 스텝을 반복한다.