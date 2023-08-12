#!/bin/bash

read -p "Enter your MySQL username: " sqlUsername
read -s -p "Enter your MySQL password: " dbPassword
echo

DB_NAME="librago"

mysql -u "$sqlUsername" -p"$dbPassword" -e "CREATE DATABASE IF NOT EXISTS $DB_NAME;"

migrate -path ../database/migration/ -database "mysql://${sqlUsername}:${dbPassword}@tcp(localhost:3306)/${DB_NAME}" -verbose up

read -p "Enter admin password: " adminPassword

cat <<EOF > ../config.yaml
DB_USERNAME: "$sqlUsername"
DB_PASSWORD: "$dbPassword"
DB_HOST: "127.0.0.1:3306"
DB_NAME: "librago"
JWT_SECRET: "$adminPassword"
EOF