module backend

go 1.16

require github.com/astaxie/beego v1.12.1

require (
	github.com/gorilla/websocket v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mailru/easyjson v0.7.7
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/smartystreets/goconvey v1.6.4
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect

)

require (
	github.com/AndreasBriese/bbloom v0.0.0-20190825152654-46b345b51c96
	github.com/GoAdminGroup/go-admin v1.2.23
	github.com/PuerkitoBio/goquery v1.7.1
	github.com/cayleygraph/quad v1.2.4
	github.com/cjoudrey/gluaurl v0.0.0-20161028222611-31cbb9bef199
	github.com/cloudflare/cloudflare-go v0.26.0
	github.com/dghubble/go-twitter v0.0.0-20211002212826-ad02880e616b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/color v1.13.0
	github.com/fofapro/fofa-go v0.0.0-20200317042037-c0caee09013d
	github.com/geziyor/geziyor v0.0.0-20211021191925-369b42cbc6c5
	github.com/go-ini/ini v1.63.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/kataras/golog v0.1.7
	github.com/miekg/dns v1.1.43
	github.com/rakyll/statik v0.1.7
	github.com/schollz/progressbar/v3 v3.8.3
	github.com/yl2chen/cidranger v1.0.2
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9
	golang.org/x/net v0.0.0-20211020060615-d418f374d309
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
	layeh.com/gopher-json v0.0.0-20201124131017-552bb3c4c3bf
)

require (
	github.com/Ice3man543/nvd v1.0.8
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/alecthomas/jsonschema v0.0.0-20210818095345-1014919a589c
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/andygrunwald/go-jira v1.14.0
	github.com/antchfx/htmlquery v1.2.3
	github.com/apex/log v1.9.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/bluele/gcache v0.0.2
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/caffix/eventbus v0.0.0-20210920204849-80ad577dc223
	github.com/caffix/netmap v0.0.0-20210920205002-55c52c6d2030
	github.com/caffix/pipeline v0.0.0-20211016020340-7e6f803c0762
	github.com/caffix/queue v0.0.0-20210927150606-04a67d1f0607
	github.com/caffix/resolve v0.0.0-20211019003042-25065a619cd5
	github.com/caffix/service v0.0.0-20210920205156-38bde8eb0503
	github.com/caffix/stringset v0.0.0-20210920202210-bde5591d523d
	github.com/chromedp/cdproto v0.0.0-20211025030258-2570df970243 // indirect
	github.com/corpix/uarand v0.1.1
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-rod/rod v0.101.7
	github.com/google/go-github v17.0.0+incompatible
	github.com/itchyny/gojq v0.12.4
	github.com/json-iterator/go v1.1.12
	github.com/julienschmidt/httprouter v1.3.0
	github.com/karlseguin/ccache v2.0.3+incompatible
	github.com/karrick/godirwalk v1.16.1
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/owenrumney/go-sarif v1.0.11
	github.com/pkg/errors v0.9.1
	github.com/projectdiscovery/clistats v0.0.8
	github.com/projectdiscovery/fastdialer v0.0.13-0.20210917073912-cad93d88e69e
	github.com/projectdiscovery/goflags v0.0.7
	github.com/projectdiscovery/gologger v1.1.4
	github.com/projectdiscovery/hmap v0.0.2-0.20210917080408-0fd7bd286bfa
	github.com/projectdiscovery/interactsh v0.0.4
	github.com/projectdiscovery/nuclei-updatecheck-api v0.0.0-20210914222811-0a072d262f77
	github.com/projectdiscovery/rawhttp v0.0.7
	github.com/projectdiscovery/retryabledns v1.0.13-0.20210916165024-76c5b76fd59a
	github.com/projectdiscovery/retryablehttp-go v1.0.2
	github.com/projectdiscovery/stringsutil v0.0.0-20210830151154-f567170afdd9
	github.com/projectdiscovery/yamldoc-go v1.0.2
	github.com/prometheus/common v0.32.1 // indirect
	github.com/remeh/sizedwaitgroup v1.0.0
	github.com/rs/xid v1.3.0
	github.com/segmentio/ksuid v1.0.4
	github.com/shirou/gopsutil/v3 v3.21.7
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.4.1
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/temoto/robotstxt v1.1.2 // indirect
	github.com/tj/go-update v2.2.5-0.20200519121640-62b4b798fd68+incompatible
	github.com/valyala/fasttemplate v1.2.1
	github.com/xanzy/go-gitlab v0.50.3
	go.uber.org/atomic v1.9.0
	go.uber.org/multierr v1.7.0
	go.uber.org/ratelimit v0.2.0
	golang.org/x/sys v0.0.0-20211025112917-711f33c9992c // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7
	golang.org/x/tools v0.1.7 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.2
)
