package static

import "net/http"

type MethodHttp string

const (
	MethodHttp_Post = http.MethodPost
	MethodHttp_Get  = http.MethodGet
)

func (mh MethodHttp) String() string {
	return string(mh)
}
