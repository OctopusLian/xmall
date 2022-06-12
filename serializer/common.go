/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:24:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 19:14:13
 */
package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"` //错误码
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"` //错误信息
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// BuildListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Msg:    "ok",
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
