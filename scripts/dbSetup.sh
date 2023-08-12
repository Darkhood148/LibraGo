#!/bin/bash

read -p "Enter your MySQL username: " sqlUsername
read -s -p "Enter your MySQL password: " dbPassword
echo

DB_NAME="librago"

migrate -path database/migration/ -database "mysql://${sqlUsername}:${dbPassword}@tcp(localhost:3306)/${DB_NAME}" -verbose up
