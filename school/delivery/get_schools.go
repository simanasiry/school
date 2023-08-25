package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

type getSchoolHandler struct {
	domain.Usecase
}

func NewGetSchoolHandler(usecase domain.Usecase) utils.Handler {
	addSchool := new(getSchoolHandler)
	addSchool.Usecase = usecase
	return addSchool
}

func (school *getSchoolHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		role, _ := ctx.GetQuery("role")
		if role == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "role must be set in query params"})
			return
		}
		if role != utils.Admin {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "only admin is allowed to see the schools"})
			return
		}

		err, code, result := school.Usecase.GetSchools()
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
