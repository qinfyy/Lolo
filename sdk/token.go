package sdk

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gucooing/lolo/db"
	"gucooing/lolo/protocol/quick"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"

	"gucooing/lolo/config"
)

type Token struct {
	ID   uint32 `json:"id"`
	Time int64  `json:"time"`
	Key  string `json:"key"`
}

func (s *Server) GenToken(id uint32) string {
	t := &Token{
		ID:   id,
		Time: time.Now().Add(10 * time.Minute).Unix(),
		Key:  config.GetGucooingApiKey(),
	}
	tokenBin, err := sonic.Marshal(t)
	if err != nil {
		return ""
	}
	if s.rsa != nil {
		tokenBin, err = s.rsa.Encode(tokenBin)
		if err != nil {
			return ""
		}
	}
	return 逆天转换(tokenBin)
}

func (s *Server) ToToken(token string) (*Token, error) {
	bin := 逆天转回(token)
	if s.rsa != nil {
		tokenBin, err := s.rsa.Decrypt(bin)
		if err != nil {
			return nil, err
		}
		bin = tokenBin
	}
	t := &Token{}
	if err := sonic.Unmarshal(bin, t); err != nil {
		return nil, err
	}
	if t.Key != config.GetGucooingApiKey() {
		return nil, errors.New("token nil")
	}
	return t, nil
}

const (
	seg = "☃️"
)

func 逆天转换(bin []byte) string {
	str := strings.Builder{}
	for _, b := range bin {
		str.WriteString(seg)
		str.WriteString(strconv.Itoa(int(b) + 100))
	}
	return str.String()
}

func 逆天转回(str string) []byte {
	bin := make([]byte, 0)
	for _, s := range strings.Split(str, seg) {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		bin = append(bin, byte(i-100))
	}
	return bin
}

func (s *Server) checkSdkToken(c *gin.Context) {
	var req quick.CheckSdkTokenRequest
	rsp := &quick.CheckSdkTokenResponse{}
	defer c.JSON(http.StatusOK, rsp)
	if err := c.ShouldBind(&req); err != nil {
		rsp.Code = -1
		return
	}
	_, err := s.ToToken(req.Token)
	if err != nil {
		rsp.Code = 1
		return
	}
	userCheck, err := db.OrCreateOFQuickCheck(req.UID)
	if err != nil {
		rsp.Code = 1
		return
	}
	if req.Token != userCheck.GateToken {
		rsp.Code = 2
		return
	}
	rsp.Uid = req.UID
}
