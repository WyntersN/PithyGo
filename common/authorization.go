/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-30 22:07:02
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-08-05 17:54:13
 */
package common

import (
	"errors"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 解析token
/*func ParseToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte("YCOAOS-p#M*JW1bl2!ON#IHunorILFjVX*u8CZa$IPAnvw6@&unHc&nkMs8BYM9yXuslC5Vlxv6rbA74Cb3AaG7k6pFjyM0*7GmFmjSY1Z")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}*/

func ParseToken(secretKey, authString string) (jwt.MapClaims, error) {

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return nil, errors.New("这是个无效的令牌")
	}
	tokenString := kv[1]
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(secretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errors.New("令牌已过期或尚未激活")
			} else {
				// Couldn't handle this token
				return nil, errors.New("无法处理此令牌")
			}
		} else {
			// Couldn't handle this token
			return nil, errors.New("无法处理此令牌")
		}
	}
	if !token.Valid {
		return nil, errors.New("令牌无效")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
func GenerateAppPassWord(password string, salt string) string {
	return Sha("sha1", "YCOAoS-"+password+"-"+salt)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz@#$%^&*()<>.?ABCDEFGHIGKLMNOPQRSDUVWXYZ/*~`"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func GetRandomStringByNum(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return "127.0.0.1"
}
