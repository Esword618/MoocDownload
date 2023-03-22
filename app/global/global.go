package global

import (
	"context"
	"sync"
)

var GlobalCtx context.Context

var GlobalDownloadInfo sync.Map
