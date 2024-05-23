package dto

// 参数验证方式和约束条件可参考：https://segmentfault.com/a/1190000023725115
import (
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
)

type ProductRecRequest struct {
	MerchantId string                 `json:"merchantId"`
	Option     string                 `json:"option" binding:"required,oneof=all rec hot"`
	Limit      int                    `json:"limit" binding:"min=1,max=50"`
	Uids       []string               `json:"uids" binding:"validate_uids"`
	MoreInfo   map[string]interface{} `json:"moreInfo"`
}

// 自定义验证规则断言
func validate_uids(fl validator.FieldLevel) bool {
	if uids, ok := fl.Field().Interface().([]string); ok {
		if len(uids) == 0 {
			return false
		}
	}
	return true
}


// 注册验证器
func RegisterValidator() error{
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        err := v.RegisterValidation("validate_uids", validate_uids)
        if err != nil {
            return err
        }
	}
	return nil
}