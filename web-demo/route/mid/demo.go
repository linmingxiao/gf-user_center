package mid

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"net/http"
)

func MyFitDemo(w *fst.GFResponse, r *http.Request) {
	logx.Info("Handler my fit demo.")
}
