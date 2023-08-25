package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

type addSchoolHandler struct {
	domain.Usecase
}

func NewSchoolHandler(usecase domain.Usecase) utils.Handler {
	addSchool := new(addSchoolHandler)
	addSchool.Usecase = usecase
	return addSchool
}

func (school *addSchoolHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		requestBody := domain.AddSchoolRequest{}
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err, code, result := school.Usecase.AddSchool(&requestBody)
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
