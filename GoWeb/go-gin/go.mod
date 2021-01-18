module github.com/EDDYCJY/go-gin-example

go 1.15

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/robfig/cron v1.2.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210113000019-eaf3bda374d2 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./conf
	github.com/EDDYCJY/go-gin-example/docs => ./docs
	github.com/EDDYCJY/go-gin-example/middleware => ./middleware
	github.com/EDDYCJY/go-gin-example/models => ./models
	github.com/EDDYCJY/go-gin-example/pkg/e => ./pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./pkg/setting
	//github.com/EDDYCJY/go-gin-example/pkg/file => ./file
	github.com/EDDYCJY/go-gin-example/pkg/util => ./pkg/util
	github.com/EDDYCJY/go-gin-example/routers => ./routers
)
