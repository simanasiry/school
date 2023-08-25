package delivery

import (
	"github.com/simanasiry/school/student/domain"
	"github.com/simanasiry/school/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addStudentHandler struct {
	domain.Usecase
}

func NewStudentHandler(usecase domain.Usecase) utils.Handler {
	addstudent := new(addStudentHandler)
	addstudent.Usecase = usecase
	return addstudent
}

func (student *addStudentHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		requestBody := domain.AddStudentRequest{}
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err, code, result := student.Usecase.AddStudent(&requestBody)
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
