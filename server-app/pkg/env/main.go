package env

import "os"

/* DATABASE MYSQL */
var MYSQL_HOST string = os.Getenv("MYSQL_HOST")
var MYSQL_PORT string = os.Getenv("MYSQL_PORT")
var MYSQL_USER string = os.Getenv("MYSQL_USER")
var MYSQL_PASSWORD string = os.Getenv("MYSQL_PASSWORD")
var MYSQL_DATABASE string = os.Getenv("MYSQL_DATABASE")
