package cli

import (
	"flag"
)

func FlagListen() *string {
	return flag.String("listen", ":8080", "Listen address and port [ADDR]:PORT (default: ':8080')")
}
func FlagMongoDB() *string {
	return flag.String("mongodb", "127.0.0.1:27017/rok4", "MongoDB url. Format: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options] (Default: 127.0.0.1:27017)")
}