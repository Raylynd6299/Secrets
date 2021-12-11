// Package user provides ...
package userHandler

import "net/http"

func UserInfo(writerResponse http.ResponseWriter, requestInfo *http.Request) {
	http.ServeFile(writerResponse, requestInfo, "./assets/index.html")
}
