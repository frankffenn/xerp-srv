module github.com/frankffenn/xerp-srv

go 1.15

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/frankffenn/xerp-srv/go-utils v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/google/uuid v1.1.3
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/frankffenn/xerp-srv/go-utils => ../go-utils
