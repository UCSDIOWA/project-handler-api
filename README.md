# project-handler-api

## Overview ##
This repository contains the necessary files to host restful API's using Protocol Buffers (a.k.a protobuf) under golang to run a database. Information on protocol buffers
can be found on [protobufs Google Developers site](https://developers.google.com/protocol-buffers/docs/proto3).
All of the endpoints are hosted using [Heroku](https://www.heroku.com). The database was implemented using [MongoDB](https://mongodb.com)
with the help of the public MongoDB driver [mgo](https://github.com/globalsign/mgo) and is being hosted using [mLab](https://mlab.com).
This repository handles requests from multiple pages of our website dealing with user profile information.

## Program Execution ##
Make sure [mgo](https://github.com/globalsign/mgo), [glog](https://github.com/golang/glog), [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), 
[cors](https://github.com/rs/cors), and [grpc](https://godoc.org/google.golang.org/grpc) are installed in your golang environemnt. To execute the program 
run the server.go file as follows,

	go run server.go

This will execute the server file.

## Endpoints ##
Each endpoint expects to receive specific fields to process a request. The following are the expectations for each endpoint and the response. The API link is https://tea-project-management-api.herokuapp.com/. Note that each endpoint only accepts POST requests, so the link will only display a "Not Found" message since GET requests aren't supported.

| Endpoint | Request | Response |
|:--------:|---------|----------|
| createproject | string email = 1;<br>string xid = 2;<br>string title = 3;<br>string projectleader = 4;<br>int32 percentdone = 5;<br>int32 groupsize = 6;<br>bool isprivate = 7;<br>repeated string tags = 8;<br>string deadline = 9;<br>string calendarid = 10;<br>string description = 11;<br>bool done = 12;<br>repeated string joinrequests = 13;<br>repeated string memberslist = 14;<br>repeated string milestones = 15;<br>repeated string pinnedannouncements = 16;<br>repeated string unpinnedannouncements = 17; | bool success = 1; |
| updateproject | string xid = 1;<br>string title = 2;<br>string projectleader = 3;<br>int32 percentdone = 4;<br>int32 groupsize = 5;<br>bool isprivate = 6;<br>repeated string tags = 7;<br>string deadline = 8;<br>string calendarid = 9;<br>string description = 10;<br>bool done = 11;<br>repeated string joinrequests = 12;<br>repeated string memberslist = 13;<br>repeated string milestones = 14;<br>repeated string pinnedannouncements = 15;<br>repeated string unpinnedannouncements = 16; | bool success = 1;<br>string xid = 2;<br>string title = 3;<br>string projectleader = 4;<br>int32 percentdone = 5;<br>int32 groupsize = 6;<br>bool isprivate = 7;<br>repeated string tags = 8;<br>string deadline = 9;<br>string calendarid = 10;<br>string description = 11;<br>bool done = 12;<br>repeated string joinrequests = 13;<br>repeated string memberslist = 14;<br>repeated string milestones = 15;<br>repeated string pinnedannouncements = 16;<br>repeated string unpinnedannouncements = 17;|
| getallprojects | string email = 1; | bool success = 1;<br>repeated Projects projects = 2; |
| getprojects | repeated string xid = 1; | bool success = 1;<br>repeated Projects projects = 2; |
| getprojects | string xid = 1;<br>string email = 2; | bool success = 1; |



## Types ##
These are the outlines of some of the custom types which are returned.

| Type | Fields |
|:--------:|---------|
| Projects | string xid = 1;<br>string title = 2;<br>string projectleader = 3;<br>int32 percentdone = 4;<br>int32 groupsize = 5;<br>bool isprivate = 6;<br>repeated string tags = 7;<br>string deadline = 8;<br>string calendarid = 9;<br>string description = 10;<br>bool done = 11;<br>repeated string joinrequests = 12;<br>repeated string memberslist = 13;<br>repeated string milestones = 14;<br>repeated string pinnedannouncements = 15;<br>repeated string unpinnedannouncements = 16; |
