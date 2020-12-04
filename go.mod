module t-blog-back

go 1.13

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190318221613-d196dffd7c2b
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190523182746-aaccbc9213b0
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190318195719-6c81ef8f67ca
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190529010454-aa71c3f32488
	google.golang.org/appengine => github.com/golang/appengine v1.6.1-0.20190515044707-311d3c5cf937
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190522204451-c2c4e71fbf69
	google.golang.org/grpc => github.com/grpc/grpc-go v1.30.0
)

require (
	github.com/appleboy/gin-jwt v2.5.0+incompatible
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/astaxie/beego v1.12.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.52.0
	github.com/golang/protobuf v1.4.2
	github.com/json-iterator/go v1.1.9
	github.com/modern-go/reflect2 v1.0.1
	github.com/rs/xid v1.2.1
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4 // indirect
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200709005830-7a2ca40e9dc3 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/dgrijalva/jwt-go.v3 v3.2.0 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)
