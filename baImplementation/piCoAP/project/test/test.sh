#!/bin/bash
mv clienttest.go client_test.go
go test 
mv client_test.go clienttest.go

