package delivery

import (
	"github.com/simanasiry/school/teacher/domain"
	"github.com/simanasiry/school/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getStudentsHandler struct {
	domain.Usecase
}

func NewStudentHandler(usecase domain.Usecase) utils.Handler {
	addTeacher := new(getStudentsHandler)
	addTeacher.Usecase = usecase
	return addTeacher
}

func (teacher *getStudentsHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		teacherIdStr := ctx.Param("id")
		teacherId, _ := strconv.ParseUint(teacherIdStr, 10, 64)

		err, code, result := teacher.Usecase.GetStudents(teacherId)
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
