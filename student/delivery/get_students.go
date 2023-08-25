package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/student/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

type getStudentHandler struct {
	domain.Usecase
}

func NewGetStudentsHandler(usecase domain.Usecase) utils.Handler {
	addstudent := new(getStudentHandler)
	addstudent.Usecase = usecase
	return addstudent
}

func (student *getStudentHandler) Handle() func(ctx *gin.Context) {
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

		err, code, result := student.Usecase.GetStudents()
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
