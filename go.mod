module gogs.io/gogs

go 1.14

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/cloudflare/cfssl v1.5.0 // indirect
	github.com/editorconfig/editorconfig-core-go/v2 v2.3.7
	github.com/fatih/color v1.9.0 // indirect
	github.com/fxamacker/cbor/v2 v2.2.0 // indirect
	github.com/go-macaron/binding v1.1.1
	github.com/go-macaron/cache v0.0.0-20190810181446-10f7c57e2196
	github.com/go-macaron/captcha v0.2.0
	github.com/go-macaron/csrf v0.0.0-20190812063352-946f6d303a4c
	github.com/go-macaron/gzip v0.0.0-20160222043647-cad1c6580a07
	github.com/go-macaron/i18n v0.5.0
	github.com/go-macaron/session v0.0.0-20190805070824-1a3cdc6f5659
	github.com/go-macaron/toolbox v0.0.0-20190813233741-94defb8383c6
	github.com/go-task/task v2.2.0+incompatible // indirect
	github.com/gogs/chardet v0.0.0-20150115103509-2404f7772561
	github.com/gogs/cron v0.0.0-20171120032916-9f6c956d3e14
	github.com/gogs/git-module v1.1.4
	github.com/gogs/go-gogs-client v0.0.0-20200128182646-c69cb7680fd4
	github.com/gogs/go-libravatar v0.0.0-20191106065024-33a75213d0a0
	github.com/gogs/minwinsvc v0.0.0-20170301035411-95be6356811a
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/issue9/identicon v1.0.1
	github.com/jaytaylor/html2text v0.0.0-20190408195923-01ec452cbe43
	github.com/json-iterator/go v1.1.10
	github.com/kevinburke/go-bindata v3.22.0+incompatible // indirect
	github.com/klauspost/compress v1.8.6 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mattn/go-zglob v0.0.3 // indirect
	github.com/mcuadros/go-version v0.0.0-20190830083331-035f6764e8d2 // indirect
	github.com/microcosm-cc/bluemonday v1.0.4
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/msteinert/pam v0.0.0-20190215180659-f29b9f28d6f9
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/niklasfasching/go-org v0.1.9
	github.com/olekukonko/tablewriter v0.0.4
	github.com/pkg/errors v0.9.1
	github.com/pquerna/otp v1.2.0
	github.com/prometheus/client_golang v1.8.0
	github.com/radovskyb/watcher v1.0.7 // indirect
	github.com/russross/blackfriday v1.6.0
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sergi/go-diff v1.1.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/stretchr/testify v1.6.1
	github.com/unknwon/cae v1.0.2
	github.com/unknwon/com v1.0.1
	github.com/unknwon/i18n v0.0.0-20190805065654-5c6446a380b6
	github.com/unknwon/paginater v0.0.0-20170405233947-45e5d631308e
	github.com/urfave/cli v1.22.5
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	golang.org/x/text v0.3.4
	gopkg.in/DATA-DOG/go-sqlmock.v2 v2.0.0-20180914054222-c19298f520d0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/ldap.v2 v2.5.1
	gopkg.in/macaron.v1 v1.4.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.5
	gorm.io/gorm v1.20.7
	mvdan.cc/sh v2.6.4+incompatible // indirect
	unknwon.dev/clog/v2 v2.1.2
	webauthn v0.0.0-00010101000000-000000000000
	xorm.io/builder v0.3.6
	xorm.io/core v0.7.2
	xorm.io/xorm v0.8.0
)

// +heroku goVersion go1.15
// +heroku install ./

// For local `webauthn` import
replace webauthn => ./webauthn
replace github.com/gogs/go-gogs-client => ./go-gogs-client