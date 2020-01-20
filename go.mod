module markman-server

go 1.13

require (
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/validator/v10 v10.0.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.5.0
	github.com/ugorji/go v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/sys v0.0.0-20191112214154-59a1497f0cea // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

// replace (
// 	./tools/config => ./tools/config
// 	github.com/xemxx/markman-server/config => ./config
// 	github.com/xemxx/markman-server/middleware => ./middleware
// 	github.com/xemxx/markman-server/models => ./models
// 	github.com/xemxx/markman-server/routers => ./routers
// )
