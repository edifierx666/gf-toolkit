package gf_toolkit

import (
  "testing"
  "time"

  "github.com/edifierx666/gf-toolkit/gtjwt"
  "github.com/gogf/gf/v2/container/gvar"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/golang-jwt/jwt/v5"
)

type Some struct {
  A string
}
type CustomClaims struct {
  Data interface{} `json:"data"`
  *jwt.RegisteredClaims
}

func TestA(t *testing.T) {

  gtjwtIns := gtjwt.New()
  token, err := gtjwtIns.Token(
    CustomClaims{
      Data:             Some{A: "dddddddddd"},
      RegisteredClaims: &jwt.RegisteredClaims{},
    },
  )
  parse, errr := gtjwtIns.Parse(token)
  g.Dump(token, err, parse, errr)

  gtjwtIns1 := gtjwt.New()
  token1, err1 := gtjwtIns1.Token(
    g.Map{"aaaaa": "???????????"},
  )
  parse1, errr1 := gtjwtIns1.Parse(token1)

  g.Dump(token1, err1, parse1, errr1)

}

func TestA1(t *testing.T) {
  gtjwtIns2 := gtjwt.New()
  token2, err2 := gtjwtIns2.Token(
    Some{A: "哈哈哈哈"},
  )
  parse2, errr2 := gtjwtIns2.Parse(token2)

  g.Dump(token2, err2, parse2, errr2)

  m := ""

  err2 = gtjwtIns2.ParseWithDst(token2, m)
  g.Dump(err2, m)
}

type CTime struct {
  time.Time
}
type N struct {
  ExpiresAt *jwt.NumericDate
}

func TestA2(t *testing.T) {
  gvar.New(
    &jwt.RegisteredClaims{},
  ).Map(
    gvar.MapOption{
      OmitEmpty: true,
    },
  )
}
