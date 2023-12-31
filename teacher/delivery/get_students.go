package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/teacher/domain"
	"github.com/simanasiry/school/utils"
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

		teacherIdStr := ctx.Param("teacherId")
		teacherId, _ := strconv.ParseUint(teacherIdStr, 10, 64)

		err, code, result := teacher.Usecase.GetStudents(teacherId)
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
