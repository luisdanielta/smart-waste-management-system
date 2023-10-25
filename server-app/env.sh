#!/bin/bash
declare -A ENV

ENV["MYSQL_HOST"]="localhost"
ENV["MYSQL_PORT"]="3306"
ENV["MYSQL_USER"]="admin-swms"
ENV["MYSQL_PASSWORD"]="swms-db-pwd"
ENV["MYSQL_DATABASE"]="swms"

ENV["PORT_SERVER"]="9095"


# set env to system
for key in ${!ENV[@]}; do
    export $key=${ENV[$key]}
done

echo "env set success"