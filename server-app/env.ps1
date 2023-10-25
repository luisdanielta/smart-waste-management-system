$ENV = @{
    "JWT_KEY" = "88U)3F!!SlirpLZP(a^v-l0pj(ivN+f+~&J9m-6t~TZ51@l2F2"

    # database - qfsd_dashboard_postgresql
    "MYSQL_HOST" = "localhost"
    "MYSQL_PORT" = "3306"
    "MYSQL_DATABASE" = "swms"
    "MYSQL_USER" = "admin-swms"
    "MYSQL_PASSWORD" = "swms-db-pwd"

    # server
    "PORT_SERVER" = "9095"
}

foreach ($key in $ENV.Keys)
{
    $value = $ENV[$key]
    [Environment]::SetEnvironmentVariable($key, $value)
}
Write-Host "All Environment variables loaded successfully."