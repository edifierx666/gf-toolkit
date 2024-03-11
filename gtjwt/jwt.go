package gtjwt

import (
  "fmt"

  "github.com/gogf/gf/v2/container/gvar"
  "github.com/gogf/gf/v2/errors/gerror"
  "github.com/gogf/gf/v2/util/gconv"
  "github.com/golang-jwt/jwt/v5"
)

type CustomClaims interface{}

type Jwt struct {
  ops *Options
}

type Options struct {
  SecretKey     string `json:"secretKey"`
  SigningMethod jwt.SigningMethod
}

type Option func(*Options)

func WithSecretKey(sk string) Option {
  return func(options *Options) {
    options.SecretKey = sk
  }
}

// WithSigningMethod func WithExpires(exp int64) Option {
//   return func(options *Options) {
//     options.Expires = exp
//   }
// }
func WithSigningMethod(method jwt.SigningMethod) Option {
  return func(options *Options) {
    options.SigningMethod = method
  }
}

func New(ops ...Option) (gtjwtIns *Jwt) {
  options := &Options{
    SecretKey:     "GTJWT_",
    SigningMethod: jwt.SigningMethodHS256,
  }
  for _, op := range ops {
    op(options)
  }
  return &Jwt{
    ops: options,
  }
}

func NewWithOptions(ops *Options) *Jwt {
  return &Jwt{ops: ops}
}

func (j *Jwt) Token(data CustomClaims) (string, error) {
  m := gvar.New(data).Map(
    gvar.MapOption{
      OmitEmpty: true,
    },
  )
  claims := jwt.MapClaims(m)
  token := jwt.NewWithClaims(j.ops.SigningMethod, claims)
  return token.SignedString([]byte(j.ops.SecretKey))
}

func (j *Jwt) Parse(token string) (*jwt.Token, error) {
  claims, err := jwt.ParseWithClaims(
    token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
      return []byte(j.ops.SecretKey), nil
    },
  )
  return claims, err
}

func (j *Jwt) ParseWithDst(token string, dst interface{}) (err error) {
  parse, err := j.Parse(token)
  if !parse.Valid {
    err = jwt.ErrInvalidKey
  }
  err = gconv.Scan(parse.Claims, dst)

  if err != nil {
    err = gerror.Wrapf(err, "ParseWithDst")
    fmt.Println(gerror.Stack(err))
  }
  return err
}

func (j *Jwt) Valid(token string) (b bool) {
  parse, _ := j.Parse(token)
  return parse.Valid
}
