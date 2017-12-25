package baseController

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"self/commons/g"
)

func (this *Base) MustUsername(c *gin.Context) string {
	ret := sessions.Default(c).Get("username").(string)
	if ret == "" {
		errInfo := fmt.Sprintf("用户名为空")
		panic(g.ParamError(errInfo))
	}
	return ret
}
