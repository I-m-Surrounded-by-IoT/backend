module github.com/I-m-Surrounded-by-IoT/backend

go 1.21

require (
	github.com/Boostport/mjml-go v0.14.6
	github.com/IBM/sarama v1.43.1
	github.com/caarlos0/env/v9 v9.0.0
	github.com/eclipse/paho.mqtt.golang v1.4.3
	github.com/emersion/go-sasl v0.0.0-20231106173351-e73c9f7bad43
	github.com/emersion/go-smtp v0.21.0
	github.com/gin-contrib/cors v1.7.1
	github.com/gin-gonic/gin v1.9.1
	github.com/go-kratos/aegis v0.2.0
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20240401040316-f715eb0bd808
	github.com/go-kratos/kratos/v2 v2.7.3
	github.com/go-redsync/redsync/v4 v4.12.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.6.0
	github.com/joho/godotenv v1.5.1
	github.com/json-iterator/go v1.1.12
	github.com/panjf2000/ants/v2 v2.9.1
	github.com/redis/go-redis/v9 v9.5.1
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/ulule/limiter/v3 v3.11.2
	github.com/zijiren233/go-colorable v0.0.0-20230930131441-997304c961cb
	github.com/zijiren233/logrus-kafka-hook v0.1.3
	github.com/zijiren233/stream v0.5.2
	github.com/zijiren233/timewheel-redis v0.3.5
	go.etcd.io/etcd/client/v3 v3.5.13
	go.uber.org/automaxprocs v1.5.3
	golang.org/x/crypto v0.22.0
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8
	golang.org/x/net v0.22.0
	golang.org/x/sync v0.6.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240401170217-c3f982113cda
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	gorm.io/driver/postgres v1.5.7
	gorm.io/gorm v1.25.9
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/bytedance/sonic v1.11.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.19.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tetratelabs/wazero v1.7.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	go.etcd.io/etcd/api/v3 v3.5.13 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.13 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
