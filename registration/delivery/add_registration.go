package delivery

import (
	"github.com/simanasiry/school/registration/domain"
	"github.com/simanasiry/school/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addRegisterHandler struct {
	domain.Usecase
}

func NewRegisterHandler(usecase domain.Usecase) utils.Handler {
	addRegister := new(addRegisterHandler)
	addRegister.Usecase = usecase
	return addRegister
}

func (Register *addRegisterHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		requestBody := domain.AddRegisterRequest{}
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err, code, result := Register.Usecase.AddRegister(&requestBody)
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
