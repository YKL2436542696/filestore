package handler

import (
	"filestore/util"
	"fmt"
	"net/http"
	"os"
	"time"

	dblayer "filestore/db"
)

const (
	pwd_salt = "#890"
)

// SignUpHandler 处理用户注册请求
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := os.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	encPasswd := util.Sha1([]byte(password + pwd_salt))
	suc := dblayer.UserSignup(username, encPasswd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}

}

// SignInHandler 登录接口
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// 返回上传html页面
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPasswd := util.Sha1([]byte(password + pwd_salt))
	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignIn(username, encPasswd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}
	// 2. 生成访问凭证（token）
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}
	// 3. 登录成功后重定向到首页
	//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

// UserInfoHandler: 用户信息查询接口
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析参数
	r.ParseForm()
	username := r.Form.Get("username")
	//token := r.Form.Get("token")

	//// 2.验证token是否有效
	//if !isTokenValid(token) {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}

	// 3. 查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 4. 组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

func isTokenValid(token string) bool {
	// TODO: 判断token时效性，是否过期
	// TODO： 从数据库表中查询username对应的token信息
	// TODO：对比两个token
	return len(token) >= 40
}

func GenToken(username string) string {
	// 40位字符 md5(username + timestamp + token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
