/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-21 21:37:53
 * @LastEditTime: 2022-07-21 14:24:37
 * @FilePath: \PithyGo\service\websocket\cache\submit.go
 */
package cache

import (
	"PithyGo/service"
	"fmt"
)

const (
	submitAgainPrefix = "acc:submit:again:" // 数据不重复提交
)

/*********************  查询数据是否处理过  ************************/

// 获取数据提交去除key
func getSubmitAgainKey(from string, value string) (key string) {
	key = fmt.Sprintf("%s%s:%s", submitAgainPrefix, from, value)

	return
}

// 重复提交
// return true:重复提交 false:第一次提交
func submitAgain(from string, second int, value string) (isSubmitAgain bool) {

	// 默认重复提交
	isSubmitAgain = true
	key := getSubmitAgainKey(from, value)

	number, err := service.RedisClient.Do("setNx", key, "1").Int()
	if err != nil {
		service.LOG.Sugar().Info("submitAgain", key, number, err)

		return
	}

	if number != 1 {

		return
	}
	// 第一次提交
	isSubmitAgain = false

	service.RedisClient.Do("Expire", key, second)

	return

}

// Seq 重复提交
func SeqDuplicates(seq string) (result bool) {
	result = submitAgain("seq", 12*60*60, seq)

	return
}
