module github.com/lopezator/cockroach-gen

require (
	cloud.google.com/go v0.34.0
	contrib.go.opencensus.io/exporter/stackdriver v0.8.0 // indirect
	github.com/Azure/azure-storage-blob-go v0.0.0-20181022225951-5152f14ace1c
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/MichaelTJones/walk v0.0.0-20161122175330-4748e29d5718
	github.com/Microsoft/go-winio v0.4.11 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/Shopify/sarama v1.13.0
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/VividCortex/ewma v1.1.1
	github.com/abourget/teamcity v0.0.0-20170428031548-e241104394f9
	github.com/andy-kimball/arenaskl v0.0.0-20171206050650-224761e552af
	github.com/apache/thrift v0.0.0-20181211084444-2b7365c54f82 // indirect
	github.com/armon/circbuf v0.0.0-20150827004946-bbbad097214e
	github.com/aws/aws-sdk-go v1.15.90
	github.com/axiomhq/hyperloglog v0.0.0-20181115174702-e8c19f174915
	github.com/backtrace-labs/go-bcd v0.0.0-20171031183808-5d8e01b2f043
	github.com/benesch/cgosymbolizer v0.0.0-20180702220239-70e1ee2b39d3
	github.com/biogo/store v0.0.0-20160505134755-913427a1d5e8
	github.com/cenkalti/backoff v2.1.0+incompatible
	github.com/certifi/gocertifi v0.0.0-20180118203423-deb3ae2ef261 // indirect
	github.com/cockroachdb/apd v0.0.0-20181017181144-bced77f817b4
	github.com/cockroachdb/circuitbreaker v0.0.0-20181019191228-4f5b16865f3c
	github.com/cockroachdb/cmux v0.0.0-20170110192607-30d10be49292
	github.com/cockroachdb/cockroach-go v0.0.0-20181001143604-e0a95dfd547c
	github.com/cockroachdb/returncheck v0.0.0-20170227172625-e91bb28baf9d
	github.com/cockroachdb/ttycolor v0.0.0-20180623121402-5bed2b5a875c
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd
	github.com/cpuguy83/go-md2man v1.0.8 // indirect
	github.com/dgryski/go-metro v0.0.0-20180109044635-280f6062b5bc // indirect
	github.com/docker/distribution v2.7.0+incompatible
	github.com/docker/docker v0.0.0-20181207101903-a4a816b6bbed
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.3.3 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/gosigar v0.9.0
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a
	github.com/getsentry/raven-go v0.2.0
	github.com/ghemawat/stream v0.0.0-20171120220530-696b145b53b9
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-sql-driver/mysql v0.0.0-20181202171036-60d456a40278
	github.com/gogo/protobuf v1.0.0
	github.com/golang-commonmark/html v0.0.0-20180910111043-7d7c804e1d46 // indirect
	github.com/golang-commonmark/linkify v0.0.0-20180910111149-f05efb453a0e // indirect
	github.com/golang-commonmark/markdown v0.0.0-20180910011815-a8f139058164
	github.com/golang-commonmark/mdurl v0.0.0-20180910110917-8d018c6567d6 // indirect
	github.com/golang-commonmark/puny v0.0.0-20180910110745-050be392d8b8 // indirect
	github.com/golang/leveldb v0.0.0-20170107010102-259d9253d719
	github.com/golang/protobuf v1.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/google/btree v0.0.0-20180813153112-4030bb1f1f0c
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/martian v2.1.0+incompatible // indirect
	github.com/google/pprof v0.0.0-20181206194817-3ea8567a2e57
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.5.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/ianlancetaylor/cgosymbolizer v0.0.0-20170921033129-f5072df9c550 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20181102032728-5e5cf60278f6 // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.3.0+incompatible
	github.com/kisielk/gotool v1.0.0
	github.com/knz/go-libedit v0.0.0-20181205205044-f49778aa742b
	github.com/knz/strtime v0.0.0-20181018220328-af2256ee352c
	github.com/kr/pretty v0.1.0
	github.com/kr/text v0.1.0
	github.com/lib/pq v1.0.0
	github.com/lightstep/lightstep-tracer-go v0.15.6
	github.com/linkedin/goavro v0.0.0-20181018120728-1beee2a74088
	github.com/lufia/iostat v0.0.0-20170605150913-9f7362b77ad3
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20180109053238-3c758769bfa6 // indirect
	github.com/marusama/semaphore v0.0.0-20181027083059-edd5217b5829
	github.com/mattn/go-isatty v0.0.4
	github.com/mattn/go-runewidth v0.0.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.0
	github.com/montanaflynn/stats v0.0.0-20180911141734-db72e6cae808
	github.com/nlopes/slack v0.4.0
	github.com/olekukonko/tablewriter v0.0.1
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opennota/wd v0.0.0-20180911144301-b446539ab1e7 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.0.2
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5
	github.com/peterbourgon/g2s v0.0.0-20170223122336-d4e7ad98afea // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5
	github.com/pierrec/lz4 v0.0.0-20181005164709-635575b42742 // indirect
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v0.9.2
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910
	github.com/prometheus/common v0.0.0-20181126121408-4724e9255275
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/russross/blackfriday v1.5.2 // indirect
	github.com/sasha-s/go-deadlock v0.2.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shirou/gopsutil v2.18.11+incompatible
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24 // indirect
	github.com/sirupsen/logrus v1.2.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.2.2
	github.com/youtube/vitess v2.1.1+incompatible // indirect
	go.etcd.io/etcd v0.0.0-20181206185822-1900a8e26f2c
	go.opencensus.io v0.18.0 // indirect
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/net v0.0.0-20181201002055-351d144fa1fc
	golang.org/x/oauth2 v0.0.0-20181203162652-d668ce993890
	golang.org/x/perf v0.0.0-20180704124530-6e6d33e29852
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys v0.0.0-20181206074257-70b957f3b65e
	golang.org/x/text v0.3.0
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools v0.0.0-20181206194817-bcd4e47d0288
	google.golang.org/api v0.0.0-20180910000450-7ca32eb868bf
	google.golang.org/genproto v0.0.0-20181221010529-a1fde7408246 // indirect
	google.golang.org/grpc v1.16.0
	gopkg.in/linkedin/goavro.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.2.1
	gotest.tools v2.2.0+incompatible // indirect
	honnef.co/go/tools v0.0.0-20180728063816-88497007e858
	vitess.io/vitess v2.1.1+incompatible
)
