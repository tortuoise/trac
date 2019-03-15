Title: trac
Category: readme
Tags: software, go, grpc, postgres 
Date: 21st March 2018
Format: markdown

## [Contents](#contents){#contents}
    - [Brief](#brief)
    - [Components](#components)
        - [Hardware](#hardware)
        - [Software](#software)
    - [Theory](#theory)
    - [Gotchas](#gotchas)
    - [References](#references)

### [Brief](#brief){#brief}

### [Components](#components){#components}

### [Usage](#usage){#usage} 

Post request using 
```
curl -d '{"user":12, "coord":{"altitude":32, "point":{"latitude":32, "longitude":33}}, "timestamp_value": "2012-02-03T14:14:14Z"}' localhost:8080/v1/trac
```



### [Gotchas](#gotchas){#gotchas}

Must provide URL containing timestamps in double quotes to curl otherwise parsing error:
```
curl "localhost:8080/v1/trac/list/15?period.start=2012-03-02T12:00:00Z&period.end=2012-04-02T12:00:00Z"
curl localhost:8080/v1/trac/list/15?&track=12sr // no double quotes reqd.
```

### [References](#references){#references}
+ [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
+ [Funnel](https://github.com/agnivade/funnel)
+ [Annotations](https://github.com/google/go-genproto/blob/master/googleapis/api/annotations/http.pb.go)

