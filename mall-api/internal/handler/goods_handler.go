package handler

import (
	"errors"
	"mall-api/internal/pkg/errcode"
	"mall-api/internal/pkg/response"
	"mall-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsHandler struct {
	svc *service.GoodsService
}

func NewGoodsHandler(svc *service.GoodsService) *GoodsHandler {
	return &GoodsHandler{svc: svc}
}

// 1. 查列表接口
func (h *GoodsHandler) GetGoodsList(c *gin.Context) {
	goodsList, err := h.svc.GetGoods()
	if err != nil {
		response.Fail(c, errcode.CodeInternalServerError, http.StatusInternalServerError, "服务器内部错误")
		return // 【致命细节】必须 return！
	}
	response.Success(c, goodsList)
}

// 2. 查详情接口
func (h *GoodsHandler) GetGoodsByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		response.Fail(c, errcode.CodeInvalidParam, 400, "商品ID格式错误")
		return
	}
	goods, err := h.svc.GetGoodsByID(uint(id))
	if err != nil {
		// 判断是不是业务层面的"商品不存在"
		if errors.Is(err, service.ErrGoodsNotFound) {
			response.Fail(c, errcode.CodeGoodsNotFound, http.StatusNotFound, "商品不存在")
			return // 【致命细节】必须 return！
		}

		// ③其它错误（比如数据库挂了），按服务器内部错误处理
		response.Fail(c, errcode.CodeInternalServerError, http.StatusInternalServerError, "服务器内部错误")
		return // 【致命细节】必须 return！
	}

	//成功，把商品塞进统一外壳返回
	response.Success(c, goods)
}
