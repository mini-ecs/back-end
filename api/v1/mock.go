package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"net/http"
)

//| 注释                 | 描述                                                                                                    |
//| -------------------- | ------------------------------------------------------------------------------------------------------ |
//| description          | 操作行为的详细说明。                                                                                      |
//| description.markdown | 应用程序的简短描述。该描述将从名为`endpointname.md`的文件中读取。                                              |
//| id                   | 用于标识操作的唯一字符串。在所有API操作中必须唯一。                                                            |
//| tags                 | 每个API操作的标签列表，以逗号分隔。                                                                         |
//| summary              | 该操作的简短摘要。                                                                                       |
//| accept               | API 可以使用的 MIME 类型列表。 请注意，Accept 仅影响具有请求正文的操作，例如 POST、PUT 和 PATCH。 值必须如“[Mime类型](#mime-types)”中所述。
//   e.g. json, xml, plain...  https://github.com/swaggo/swag/blob/master/README_zh-CN.md#mime%E7%B1%BB%E5%9E%8B
//| produce              | API可以生成的MIME类型的列表。值必须如“[Mime类型](#mime-types)”中所述。
//   e.g. 同 accept
//| param                | 用空格分隔的参数。`param name`,`param type`,`data type`,`is mandatory?`,`comment` `attribute(optional)`
//   e.g. param type: 	query, path, header, body, formData
//        data type: 	string (string)
//						integer (int, uint, uint32, uint64)
//						number (float32)
//						boolean (bool)
// 						object
//        is mandatory: true / false
//        attribute: https://github.com/swaggo/swag/blob/master/README_zh-CN.md#%E5%B1%9E%E6%80%A7
//user defined struct
//| security             | 每个API操作的[安全性](#安全性)。                                                                      |
//| success              | 以空格分隔的成功响应。`return code`,`{param type}`,`data type`,`comment`                                |
//| failure              | 以空格分隔的故障响应。`return code`,`{param type}`,`data type`,`comment`                                |
//| response             | 与success、failure作用相同                                                                               |
//| header               | 以空格分隔的头字段。 `return code`,`{param type}`,`data type`,`comment`                                 |
//| router               | 以空格分隔的路径定义。 `path`,`[httpMethod]`                                                            |
//| x-name               | 扩展字段必须以`x-`开头，并且只能使用json值。                                                            |

// Welcome godoc
// @Summary      该操作的简短摘要
// @Description  操作行为的详细说明。
// @Tags         example-tag
// @Accept       json
// @Produce      json
// @Param        val1  query      int           true  "这是评论"
// @Success      200   {integer}  string        "这是评论1"
// @Failure      400   {object}   response.Msg  "这是评论2"
// @Router       /welcome [get]
func Welcome(c *gin.Context) {
	logger.Infof("Welcome page")
	c.JSON(http.StatusOK, response.SuccessMsg("Welcome to mini-ecs, the service is ok"))
}
