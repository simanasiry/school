package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/simanasiry/school/teacher/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

type getTeachersHandler struct {
	domain.Usecase
}

func NewGetTeachersHandler(usecase domain.Usecase) utils.Handler {
	addTeacher := new(getTeachersHandler)
	addTeacher.Usecase = usecase
	return addTeacher
}

func (teacher *getTeachersHandler) Handle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		role, _ := ctx.GetQuery("role")
		if role == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "role must be filled in query params"})
			return
		}
		if role != utils.Admin && role != utils.HeadMaster {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "role must be headMaster or admin"})
			return
		}
		err, code, result := teacher.Usecase.GetTeachers()
		if err != nil {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
