package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/registration/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

type getRegisterHandler struct {
	domain.Usecase
}

func NewGetRegisterHandler(usecase domain.Usecase) utils.Handler {
	addRegister := new(getRegisterHandler)
	addRegister.Usecase = usecase
	return addRegister
}

func (Register *getRegisterHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		// check authorization
		role, _ := ctx.GetQuery("role")
		if role == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "role must be filled in query params"})
			return
		}
		if role != utils.Admin && role != utils.HeadMaster {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "role must be headMaster or admin"})
			return
		}

		err, code, result := Register.Usecase.GetRegisters()
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
