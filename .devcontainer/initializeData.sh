#!/bin/bash
echo "initialize.sh - 1607"

cd datamodel
git status
# due to docker_compose configuration, the mysql container can be accessed by the service name, which is "db"
export MYSQL="mysql -h db -P 3306 -u root -ppassword "
export ROOT=/tmp
make mysql-clean language-mysql

sleep infinity