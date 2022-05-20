/**
* @Author:drl
* @Date: 2022/5/19 13:12
 */
package response

type Response struct {
	StatusCode int32  `json:"status_code"`          //状态码
	StatusMsg  string `json:"status_msg,omitempty"` //状态信息
}
