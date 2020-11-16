package jwt

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

//SetToRedis add token to redis
func SetToRedis(username, token string) error {
	redis, err := RedisNewClient()
	if err != nil {
		return err
	}
	err = redis.Set(username, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetFromRedis retun key val
func GetFromRedis(username string) (string, error) {
	redis, _ := RedisNewClient()
	token, err := redis.Get(username).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

// GenerateJWTSigned return Decode payload
func GenerateJWTSigned(privateClim interface{}) (string, error) {
	key := []byte("secret")
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "issuer",
		NotBefore: jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	}

	raw, err := jwt.Signed(sig).Claims(cl).Claims(privateClim).CompactSerialize()
	if err != nil {
		return "", err
	}

	return raw, nil
}

// ParseJSONWebTokenClaims test
func ParseJSONWebTokenClaims(tokP string, out2 interface{}) (jwt.Claims, error) {
	var sharedKey = []byte("secret")
	raw := tokP
	out1 := jwt.Claims{}
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return jwt.Claims{}, err
	}

	if err := tok.Claims(sharedKey, &out1, &out2); err != nil {
		return jwt.Claims{}, err
	}
	//fmt.Printf("iss: %s, sub: %s, scopes: %s\n", out.Issuer, out.Subject, strings.Join(out2.Scopes, ","))
	// Output: iss: issuer, sub: subject, scopes: foo,bar
	return out1, nil
}

//RedisNewClient return cl to con
func RedisNewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	//fmt.Println(pong, err)
	// Output: PONG <nil>
	return client, nil
}

//RedidTestClient simple test to check Connection Work or not
func RedidTestClient() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get("key").Result()
	if err != nil {
		return err
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		return err
	} else {
		fmt.Println("key2", val2)
	}
	res := client.HSet("Mylist", "f1", "1")
	if res.Err() != nil {
		return res.Err()
	}
	valHget := client.HGet("Mylist", "f1")

	if valHget.Err() != nil {
		return valHget.Err()
	}
	fmt.Println("my list f1:", valHget.Val())
	return nil
}
