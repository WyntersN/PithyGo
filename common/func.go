/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-05-24 18:47:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-08-08 18:16:11
 */
package common

import (
	"PithyGo/addons/tencentyun/cos"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ebar-go/curl"
	"github.com/kataras/iris/v12/sessions"
	"github.com/r3labs/sse/v2"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var (
	DataPrefix  string
	SessManager *sessions.Sessions
	Sse         = sse.New()
	SessStart   *sessions.Session
)

func init() {
	Sse.CreateStream("messages")
}

func SsePublish(id string, str string) {
	Sse.Publish(id, &sse.Event{
		Data: []byte(str),
	})
}

func ArrayUintToDB(array []int) string {
	str := "{"
	for _, v := range array {
		str += strconv.Itoa(v) + ","
		// if i+1 != len(array) {
		// 	str += ","
		// }
	}
	return str[:len(str)-1] + "}"
}

func GetNowRandOrder(n int) string {
	timeStr := time.Now()
	return timeStr.Format("060102150405") + GetRandomStringByNum(n)
}

func StrFilterNonChinese(src string) string {
	var (
		hzRegexp = regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5()（）【】\\[\\],，。\\.;；'‘’<>《》]+$")
		strn     = ""
	)
	for _, c := range src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}

	if strn == "" {
		return src
	}
	return strn
}

func StrFilterNonNumber(src string) string {
	var (
		hzRegexp = regexp.MustCompile("^[0-9]+$")
		strn     = ""
	)
	for _, c := range src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}

	if strn == "" {
		return src
	}
	return strn
}

func ReturnErrorToDB(tips string, err error) error {
	errStr := err.Error()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errStr = "未检测到任何数据"
	}
	return errors.New(tips + "：" + errStr)
}

// 利用正则表达式压缩字符串，去除空格或制表符
func CompresStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+|\\n+")
	return strings.Trim(reg.ReplaceAllString(str, ""), "")
}

//@author: [piexlmax]
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 计算SHA1值
func Sha(mode, data string) string {
	var s = md5.New()
	switch mode {
	case "sha1":
		s = sha1.New()
	case "sha256":
		s = sha256.New()
	}
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum(nil))
}

// 计算HmacSha256值
func HmacSha256(message, secret string) string {
	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte(message))
	return hex.EncodeToString(m.Sum(nil))
}

func FileNameFind(fileName string) (string, string) {
	filenameall := path.Base(fileName)
	filesuffix := path.Ext(fileName)
	return filenameall, filesuffix[1:]
}

// CheckMobile 检验手机号
func CheckMobile(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	// 返回 MatchString 是否匹配
	return regexp.MustCompile("^1[345789]{1}\\d{9}$").MatchString(phone)
}

// 保留2位小数 不四舍五入
func NoRoundNum(f float64, n int32) float64 {
	a, _ := decimal.NewFromFloat(f).RoundFloor(n).Float64()
	return a
}

func GetIntervalTopDate(BeforeDate string) int {
	current := time.Now().Unix()
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02", BeforeDate, loc)
	timestamp := tmp.Unix()              //转化为时间戳 类型是int64
	res := (current - timestamp) / 86400 //相差值
	return int(res)
}

// 下载网络内容并上传
func DownloadNetTenFile(urlFile string, finename string) error {

	respNet, err := curl.Get("https://cos-ows-1302225462.cos.ap-shanghai.myqcloud.com/" + finename)
	if err != nil {
		return err
	}
	if !strings.Contains(respNet.String(), "NoSuchKey") {
		return nil
	}
	// Get the data
	resp, err := curl.Get(urlFile)
	if err != nil {
		return err
	}
	return cos.UploadNetPut(finename, resp.Reader())
}

func ExpressNameToCode(name string) string {

	switch name {
	case "百世快递":
		return "huitongkuaidi"
	case "申通快递":
		return "shentong"
	case "中通快递":
		return "zhongtong"
	case "邮政快递包裹":
		return "youzhengguonei"
	case "圆通速递":
		return "yuantong"
	case "韵达速递":
		return "yunda"
	case "顺丰":
		return "sf"
	}
	return ""
}

// 计算SHA1值
func EnSha1(data string) string {
	s := sha1.New()
	s.Write([]byte(data + "pjhXe^P#zgkCSKNn0ZZdAwyNZSJD@W@s"))
	return hex.EncodeToString(s.Sum(nil))
}

// 计算md5值
func Md5(data string) string {
	s := md5.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum(nil))
}

// 计算md5值
func Md5Byte(data string) []byte {
	s := md5.New()
	s.Write([]byte(data))
	return s.Sum(nil)
}

/**
 * @summary: 获取当前项目目录
 * @description:
 * @param {*}
 * @Author: Wynters
 */
func GetCurrentPath() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

/**
 * @summary: 获取USERID 已废除
 * @description:
 * @param {*}
 * @Author: Wynters
 */
// func GetUserID() uint {
// 	return SessStart.Get("USER_ID").(uint)
// }

// func GetUserIDToString() string {
// 	return strconv.Itoa(int(GetUserID()))
// }
// func GetUserRealName() string {
// 	return SessStart.GetString("REALNAME")
// }

/**
 * @summary:
 * @description:
 * @param {*}
 * @Author: Wynters
 */

func AutoCreatePath(path string) error {
	if _, err := os.Stat(path); err == nil {
		//	fmt.Println("path exists 1", path)
	} else {
		//fmt.Println("path not exists ", path)
		err := os.MkdirAll(path, 0775)

		if err != nil {
			return err
		}
	}

	// check again
	_, err := os.Stat(path)
	return err
}
