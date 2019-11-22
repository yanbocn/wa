package wa

import (
	"time"
)

// Sleep time.Sleep t second
func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}
