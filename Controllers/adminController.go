package Controllers

import (
	"net/http"
	"strconv"

	"github.com/bigby/project/Services"
	"github.com/bigby/project/Utils"
	"github.com/gin-gonic/gin"
)

// @Summary 管理员获取学生信息列表
// @Description 管理员获取所有学生信息列表
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Student}
// @Router /app/admin/get_student_list [get]
func StudentListQueryFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	result, err := adminService.QueryStudentList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员获取学生信息列表（排序后）
// @Description 管理员获取所有学生信息列表（排序后）
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Services.Student}
// @Router /app/admin/get_student_sorted_list [get]
func StudentListQuerySortedFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	sortName := c.Query("sortName")
	sortDir := c.Query("sortDir")

	result, err := adminService.QueryStudentSortedList(sortName, sortDir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员获取一个学生信息
// @Description 管理员获取一个学生信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Success 200 {object} Res{data=Services.Student}
// @Router /app/admin/get_student [get]
func StudentQueryFromAdmin(c *gin.Context) {
	var studentService Services.Student

	var err error

	studentService.Sid, err = strconv.ParseInt(c.Query("Sid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := studentService.QueryBySid(studentService.Sid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员增加学生
// @Description 管理员增加一个学生信息，注册学生
// @Tags Admin
// @Accept json
// @Produce  json
// @Param studentData formData Services.Student true "学生数据"
// @Success 200 {object} Res{data=Services.Student}
// @Router /app/admin/insert_student [POST]
func StudentInsertFromAdmin(c *gin.Context) {
	var studentService Services.Student

	err := c.ShouldBindJSON(&studentService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	studentService.Sgrade = Utils.GradeStringToInt(studentService.SgradeName)
	if studentService.Sicon == "" {
		studentService.Sicon = "default_boy.png"
	}

	result, err := studentService.Register()
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 管理员编辑学生
// @Description 管理员编辑一个学生信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param studentData formData Services.Student true "学生数据"
// @Success 200 {object} Res{data=Services.Student}
// @Router /app/admin/update_student [POST]
func StudentUpdateFromAdmin(c *gin.Context) {
	var studentService Services.Student

	err := c.ShouldBindJSON(&studentService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	studentService.Sgrade = Utils.GradeStringToInt(studentService.SgradeName)
	if studentService.Sicon == "" {
		studentService.Sicon = "default_boy.png"
	}

	result, err := studentService.Update(studentService.Sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 删除学生
// @Description 删除学生数据
// @Tags Admin
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Success 200 {object} Res {"code":200,"msg":"Delete() success!","data":500001}
// @Router /app/student/delete_student/:id [DELETE]
func StudentDeleteFromAdmin(c *gin.Context) {
	var studentService Services.Student

	var err error

	studentService.Sid, err = strconv.ParseInt(c.Query("Sid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := studentService.Delete(studentService.Sid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Delete Success!",
		Data: result.Sid,
	})
}

// @Summary 管理员获取教师信息列表
// @Description 管理员获取所有教师信息列表
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Services.Teacher}
// @Router /app/admin/get_teacher_list [get]
func TeacherListQueryFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	result, err := adminService.QueryTeacherList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员获取教师信息列表（排序后）
// @Description 管理员获取所有教师信息列表（排序后）
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Services.Teacher}
// @Router /app/admin/get_teacher_sorted_list [get]
func TeacherListQuerySortedFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	sortName := c.Query("sortName")
	sortDir := c.Query("sortDir")

	result, err := adminService.QueryTeacherSortedList(sortName, sortDir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员获取一个教师信息
// @Description 管理员获取一个教师信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param Tid query string true "教师id"
// @Success 200 {object} Res{data=Services.Teacher}
// @Router /app/admin/get_teacher [get]
func TeacherQueryFromAdmin(c *gin.Context) {
	var teacherService Services.Teacher

	var err error
	teacherService.Tid, err = strconv.ParseInt(c.Query("Tid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := teacherService.QueryByTid(teacherService.Tid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员增加教师
// @Description 管理员增加一个教师信息，注册教师
// @Tags Admin
// @Accept json
// @Produce  json
// @Param teacherData formData Services.Teacher true "教师数据"
// @Success 200 {object} Res{data=Services.Teacher}
// @Router /app/admin/insert_teacher [POST]
func TeacherInsertFromAdmin(c *gin.Context) {
	var teacherService Services.Teacher

	err := c.ShouldBindJSON(&teacherService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	if teacherService.Ticon == "" {
		teacherService.Ticon = "default_teacher.png"
	}

	result, err := teacherService.Register()
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 管理员编辑教师
// @Description 管理员编辑一个教师信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param teacherData formData Services.Teacher true  "教师数据"
// @Success 200 {object} Res{data=Services.Teacher}
// @Router /app/admin/update_teacher [POST]
func TeacherUpdateFromAdmin(c *gin.Context) {
	var teacherService Services.Teacher

	err := c.ShouldBindJSON(&teacherService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	if teacherService.Ticon == "" {
		teacherService.Ticon = "default_teacher.png"
	}

	result, err := teacherService.Update(teacherService.Tid)
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 删除教师
// @Description 删除教师数据
// @Tags Admin
// @Accept json
// @Produce  json
// @Param Tid query string true "教师id"
// @Success 200 {object} Res {"code":200,"msg":"Delete() success!","data":500001}
// @Router /app/teacher/delete_teacher/:id [DELETE]
func TeacherDeleteFromAdmin(c *gin.Context) {
	var teacherService Services.Teacher

	var err error

	teacherService.Tid, err = strconv.ParseInt(c.Query("Tid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := teacherService.Delete(teacherService.Tid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Delete Success!",
		Data: result.Tid,
	})
}

// @Summary 管理员获取校区信息列表
// @Description 管理员获取所有校区信息列表
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Campus}
// @Router /app/admin/get_campus_list [get]
func CampusListQueryFromAdmin(c *gin.Context) {
	var campusService Services.Campus

	result, err := campusService.QueryAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员增加校区信息
// @Description 管理员增加校区信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/admin/insert_campus [POST]
func CampusInsertFromAdmin(c *gin.Context) {
	var campusService Services.Campus

	err := c.ShouldBindJSON(&campusService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := campusService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 管理员编辑校区信息
// @Description 管理员编辑校区信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/admin/update_campus [PUT]
func CampusUpdateFromAdmin(c *gin.Context) {
	var campusService Services.Campus

	var campusArray []Services.Campus

	err := c.ShouldBindJSON(&campusArray)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := campusService.UpdateFromArray(campusArray)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Edit() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Edit Success!",
		Data: result,
	})
}

// @Summary 删除校区
// @Description 删除校区数据
// @Tags Admin
// @Accept json
// @Produce  json
// @Param Cid query string true "校区id"
// @Success 200 {object} Res {"code":200,"msg":"Delete() success!","data":500001}
// @Router /app/campus/delete_campus/:id [DELETE]
func CampusDeleteFromAdmin(c *gin.Context) {
	var campusService Services.Campus

	var err error

	campusService.Cid, err = strconv.ParseInt(c.Query("Cid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := campusService.Delete(campusService.Cid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Delete Success!",
		Data: result.Cid,
	})
}

// @Summary 预约家教列表
// @Description 获取预约家教列表
// @Tags OrderTeacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.OrderTeacher}
// @Router /app/order/get_order_teacher_list [get]
func OrderTeacherListQueryFromAdmin(c *gin.Context) {
	var otService Services.OrderTeacher

	result, err := otService.QueryListByDate()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusOK, Res{
			Code: 2,
			Msg:  "No Order Teacher!",
			Data: result,
		})
	} else {
		c.JSON(http.StatusOK, Res{
			Code: 1,
			Msg:  "Query Success!",
			Data: result,
		})
	}

}

// @Summary 预约家教信息
// @Description 获取某个预约家教信息
// @Tags OrderTeacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=Models.OrderTeacher}
// @Router /app/order/get_order_teacher [get]
func OrderTeacherQueryFromAdmin(c *gin.Context) {
	var otService Services.OrderTeacher
	var err error

	otService.OTid, err = strconv.ParseInt(c.Query("OTid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := otService.QueryByOTid(otService.OTid)

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})

}

// @Summary 管理员获取管理员列表
// @Description 管理员获取管理员列表
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/admin/get_admin_list [GET]
func AdminListQueryFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	result, err := adminService.QueryAdminList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员获取管理员
// @Description 管理员获取一个管理员信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/admin/get_admin [GET]
func AdminQueryFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	var err error
	adminService.Aid, err = strconv.ParseInt(c.Query("Aid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := adminService.QueryAdmin(adminService.Aid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 管理员增加管理员
// @Description 管理员增加管理员
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/admin/insert_admin [POST]
func AdminInsertFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	err := c.ShouldBindJSON(&adminService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := adminService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Insert Failed!: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 管理员编辑管理员
// @Description 管理员编辑管理员信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/admin/update_admin [PUT]
func AdminUpdateFromAdmin(c *gin.Context) {
	var adminService Services.Admin

	err := c.ShouldBindJSON(&adminService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := adminService.Update()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Update Failed!: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
