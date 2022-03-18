package v1

import (
	"adcenter/app/models/department"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type DepartmentsController struct {
	BaseAPIController
}

func (ctrl *DepartmentsController) Index(c *gin.Context) {
	departments := department.All()
	response.Data(c, departments)
}

func (ctrl *DepartmentsController) Show(c *gin.Context) {
	departmentModel := department.Get(c.Param("id"))
	if departmentModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, departmentModel)
}

func (ctrl *DepartmentsController) Store(c *gin.Context) {

	request := requests.DepartmentRequest{}
	if ok := requests.Validate(c, &request, requests.DepartmentSave); !ok {
		return
	}

	departmentModel := department.Department{

		DepartmentName: request.DepartmentName,
		ParentId:       request.ParentId,
	}
	departmentModel.Create()
	if departmentModel.ID > 0 {
		response.Created(c, departmentModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *DepartmentsController) Update(c *gin.Context) {

	departmentModel := department.Get(c.Param("id"))
	if departmentModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyDepartment(c, departmentModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.DepartmentRequest{}
	bindOk := requests.Validate(c, &request, requests.DepartmentSave)
	if !bindOk {
		return
	}

	departmentModel.DepartmentName = request.DepartmentName
	departmentModel.ParentId = request.ParentId
	rowsAffected := departmentModel.Save()
	if rowsAffected > 0 {
		response.Data(c, departmentModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *DepartmentsController) Delete(c *gin.Context) {

	departmentModel := department.Get(c.Param("id"))
	if departmentModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyDepartment(c, departmentModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := departmentModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
