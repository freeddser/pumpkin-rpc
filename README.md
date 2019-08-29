# Go GRPC Example

# env
* export DG_NODE_NO=1
* export DG_PSQL_DB_HOST=localhost
* export DG_PSQL_DB_NAME=imbi
* export DG_PSQL_DB_PASSWORD=123456
* export DG_PSQL_DB_PORT=5432
* export DG_PSQL_DB_USER=scpman


$ make build-server

Building server...

Done.

$ make build-client

Building client...

Done.

$ make run-server

Running server binary...

INFO[0000] here                                          fields.time="2019-08-29T04:28:55Z" hostname=gavin-V8 source="server.go:73"

localhost gavin imbi


$ make run-client

Running client binary...

INFO[0000] Hello gRPC.                                   fields.time="2019-08-29T04:30:06Z" hostname=gavin-V8 source="client.go:74"



    --------------
    customers:<id:101 name:"gavin" email:"gavin@aa.com" phone:"1010101" > customers:<id:102 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:103 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1245712384 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1619005440 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:-419426304 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:-679473152 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1635782656 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1182797824 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1058194552302604288 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1058194603582164992 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1058195264700944384 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1058195282296049664 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1058214119926140928 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1058347749390422016 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1058723537407184896 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059620889038950400 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059620892511834112 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059620896202821632 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059623778448838656 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059623782764777472 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059625617080717312 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059633067133505536 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059633069893357568 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059633072535769088 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059648770779648000 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059648773409476608 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059648777003995136 name:"bob" email:"1@x.com" phone:"1876527" > customers:<id:1059655029012566016 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059655060201410560 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059655809245384704 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059655823476658176 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059655836881653760 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059662219081420800 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059662590763864064 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059662595100774400 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059662598724653056 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059664289381486592 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1059664295912017920 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1075662565188571136 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1166924726283538432 name:"gavin1" email:"gavin@aa.com" phone:"1010101" > customers:<id:1001 name:"gavin1" email:"gavin@aa.com" phone:"1010101" >
    =========


* https://github.com/grpc-ecosystem/grpc-gateway

* curl -X GET -k https://localhost:50055/v1/customer/all
* curl -X POST -k https://localhost:50055/v1/customer -d '{"id": "1","name":"gavin","email":"x","one":"1876527"}'

Monitor

https://localhost:50055/v1/monitor

monitor client

cmd/monitor && go run getMonitor.go  -c ../../config.toml


--------------------------------------------------------------------
|avg_time:2.44ms  |max_time:3.15ms  |method:GET  |min_time:1.73ms  |request_url:/v1/customer/all  |times:2  |total_time:4.88ms|