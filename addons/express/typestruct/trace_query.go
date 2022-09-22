/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-06-13 09:35:35
 * @LastEditTime: 2022-06-13 09:39:30
 * @FilePath: \PithyGo\addons\express\typestruct\trace_query.go
 */
package typestruct

type TraceQueryStruct struct {
	Brand string                 `json:"brand"`
	Data  []TraceQueryStructData `json:"data"`
	No    string                 `json:"no"`
	Order string                 `json:"order"`
}
type TraceQueryStructData struct {
	Context string `json:"context"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}
