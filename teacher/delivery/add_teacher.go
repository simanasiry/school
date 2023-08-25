package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/teacher/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

type addTeacherHandler struct {
	domain.Usecase
}

func NewTeacherHandler(usecase domain.Usecase) utils.Handler {
	addTeacher := new(addTeacherHandler)
	addTeacher.Usecase = usecase
	return addTeacher
}

func (teacher *addTeacherHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		requestBody := domain.AddTeacherRequest{}
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err, code, result := teacher.Usecase.AddTeacher(&requestBody)
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
