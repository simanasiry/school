package app

import (
	"github.com/simanasiry/school/registration"
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/student"
	"github.com/simanasiry/school/teacher"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	config config

	registry   *viper.Viper
	gin        *gin.Engine
	httpServer *http.Server
	db         *gorm.DB

	schoolModule       *school.Module
	teacherModule      *teacher.Module
	studentModule      *student.Module
	registrationModule *registration.Module
}

func New(registry *viper.Viper) *App {
	return &App{registry: registry}
}

func (a *App) InitCore() {
	a.initConfig()
	a.initEngine()
	a.initDB()
}

func (a *App) InitModules() {
	a.initSchools()
	a.initTeachers()
	a.initStudents()
	a.initRegistrations()

}

func (a *App) panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
