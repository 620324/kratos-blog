package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/redis/go-redis/v9"
	user2 "kratos-blog/api/user"
	"kratos-blog/pkg/auth"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
	h "net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// PathList 当对文章进行修改时,需要清除redis缓存的接口
	PathList = []string{
		"/api/addBlog",
		"/api/updateBlog",
		"/api/updateIndividual",
		"/api/deleteBlog",
		"/api/updateOnly",
		"/api/cacheBlog",
	}
)

type FilterInterface interface {
	DeleteCache() http.FilterFunc
	FilterPermission(whiteList, blackList []string) http.FilterFunc
}

type FilterRepo struct {
	data *Data
	log  *log.Helper
	mu   sync.Mutex
	wg   sync.WaitGroup
}

func NewFilterRepo(data *Data, l log.Logger) *FilterRepo {
	lg := log.NewHelper(log.With(l, "module", "filter"))
	return &FilterRepo{data: data, log: lg}
}

// DeleteCache :dev 清除redis缓存
func (f *FilterRepo) DeleteCache() http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(writer h.ResponseWriter, request *h.Request) {
			method := request.Method
			if f.data.role.IsPathInList(request.URL.Path, PathList) {
				if method == h.MethodPost || method == h.MethodDelete || method == h.MethodPut {
					f.log.Info(f.data.rdb.Del(CTX, server.AdminNotes).Result())
					f.log.Info(f.data.rdb.Del(CTX, server.Notes).Result())
				}
			}
			handler.ServeHTTP(writer, request)
		})
	}
}

func (f *FilterRepo) FilterPermission(whiteList []string) http.FilterFunc {
	return f.data.role.FilterPermission(whiteList)
}

func WritePermissionError(w http.ResponseWriter) {
	w.WriteHeader(401)
	rep := map[string]interface{}{
		"code":   401,
		"result": vo.PERMISSION_ERROR,
	}
	byteRep, _ := json.Marshal(&rep)
	w.Write(byteRep)
}

func (f *FilterRepo) CommentLimiter() http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var pathList = []string{
				"/api/addComment",
				"/api/addReward",
			}
			for _, wlPath := range pathList {
				if strings.HasPrefix(req.URL.Path, wlPath) {
					token := req.Header.Get(Token)
					if len(token) == 0 {
						WritePermissionError(w)
						return
					}

					username := auth.GetLoginName(token)
					if len(username) == 0 {
						WritePermissionError(w)
						return
					}
					if !f.data.rdb.HExists(context.Background(), server.TokenMangerKey, username).Val() {
						WritePermissionError(w)
						return
					}
					res, err := f.data.uc.GetUser(context.Background(), &user2.GetUserRequest{
						Name: username,
					})
					if err != nil {
						panic(err)
					}
					if r, err := strconv.ParseBool(res.Data[5]); r || err != nil {
						WritePermissionError(w)
						return
					}
					prefixUsernameKey := fmt.Sprintf("comment_%s", username)
					data, e := f.data.rdb.Get(context.Background(), prefixUsernameKey).Result()
					if e != nil && !errors.Is(e, redis.Nil) {
						// 处理错误
						f.log.Log(log.LevelError, e)
						return
					}
					if len(data) > 0 {
						if i, err := strconv.Atoi(data); err == nil {
							if i >= 12 {
								_, err = f.data.uc.SetBlack(context.Background(), &user2.SetBlackRequest{
									Name: username,
								})
								if err != nil {
									f.log.Log(log.LevelError, err)
									return
								}
								f.data.rdb.Del(context.Background(), prefixUsernameKey)
								WritePermissionError(w)
								return
							}
							i += 1
							f.data.rdb.IncrBy(context.Background(), prefixUsernameKey, int64(i))
						}
					} else {
						f.data.rdb.Set(context.Background(), prefixUsernameKey, 1, time.Second*60)
					}
				}
			}
			handler.ServeHTTP(w, req)
			return
		})
	}
}
