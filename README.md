## Contents
- [Brief](#brief)
- [Run](#run)
- [Usage](#usage)
- [Gotchas](#gotchas)
- [References](#references)

### [Brief](#brief)
Trac server using gRPC gateway [1]. 

### [Setup](#setup)

Edit trac.proto

Generate trac.pb.go
```
cd ./pb
./protoc.sh
```

Generate trac.pb.gw.go
```
cd ./pb
./gw.sh 

```

### [Build](#build)

Build using:
```
cd cmd-server
go build
cd cmd-gateway
go build
```

### [Run](#run)
Run using:
```
./run.sh --OR--
./cmd-server/cmd-server -stderrthreshold=INFO -dbdb=m0v -dbuser=sridhar -dbpw=rcsp8 &
./cmd-gateway/cmd-gateway -stderrthreshold=INFO &
```
Starts gRPC server on 9090 and http gateway on 8080. 

### [Usage](#usage)

Post request using 
```
curl -d '{"user":12, "coord":{"altitude":32, "point":{"latitude":32, "longitude":33}}, "timestamp_value": "2012-02-03T14:14:14Z"}' localhost:8080/v1/trac
```

### [Gotchas](#gotchas)

Must provide URL containing timestamps in double quotes to curl otherwise parsing error:
```
curl "localhost:8080/v1/trac/list/15?period.start=2012-03-02T12:00:00Z&period.end=2012-04-02T12:00:00Z"
curl localhost:8080/v1/trac/list/15?&track=12sr // no double quotes reqd.
```

### [References](#references)
+ [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
+ [Funnel](https://github.com/agnivade/funnel)
+ [Annotations](https://github.com/google/go-genproto/blob/master/googleapis/api/annotations/http.pb.go)

