#!/bin/bash
cd datamodel
export MYSQL="mysql -h localhost -P 3306 -u root -ppassword --local_infile=1"
make external-mysql language-mysql

