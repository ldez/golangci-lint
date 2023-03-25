//golangcitest:args -Edepguard
//golangcitest:config_path testdata/configs/depguard.yml
package testdata

import (
	"compress/gzip" // want "import 'compress/gzip' is not allowed from list 'Main': in the denylist"
	"log"           // want "import 'log' is not allowed from list 'Main': don't use log"
)

func SpewDebugInfo() {
	log.Println(gzip.BestCompression)
}
