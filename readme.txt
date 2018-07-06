https://thebook.io/006806/ch08/


// to run
go run *.go

// test with "curl"
curl http://localhost:8080/users/1
curl http://localhost:8080/users/1/addr/2
curl http://localhost:8080/users
curl http://localhost:8080/users -d ''
curl http://localhost:8080/users/1/addr -d ''



//////  Web Server in docker container /////
//////  in Windows 10  ///////////////////////
// docker container 안에서
(1) docker container를 launch할때
     docker run --rm -it -p 8080:8080 my18 /bin/bash
     (-p 8080:8080 가 추가됨에 주의)
     (8080은 server가 지켜보는 port 번호임)
// windows 10 cmd 창 안에서
(1) ipconfig /all (windows command 창에서)
(2) Hyper-V Virtual Ethernet Adapter의 IP address를 찾음
     예, 172.20.172.65(기본 설정), 10.0.75.1(기본 설정)
(3) Web browser에서 아래와 같이 하면 됨
     http://172.20.172.65:8080/



8.3 인가 8.4 할 차례 (2018/06/01)
