#!/bin/bash
curl -X GET "localhost:8080/v1/account?ID=5190813503979520"
curl -X POST localhost:8080/v1/account
curl -X POST localhost:8080/v1/profile -d '{"ID":5190813503979520, "Nickname":"test"}' -H "Content-Type: application/json"
curl -X GET "localhost:8080/v1/profile?ID=5190813503979520"
curl -X POST localhost:8080/v1/event/boss/attack -d '{"ID":1, "Point":100}' -H "Content-Type: application/json"
curl -X GET "localhost:8080/v1/event/boss/attack?ID=1"