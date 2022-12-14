package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func GetSwiperListService() dto.RetStruct {
	db := common.InitDB()
	defer db.Close()

	var swiperListVos []vo.SwiperListVo
	db.Model(&model.SwiperList{}).Select("id, uid, img_url, video_href").Scan(&swiperListVos)
	return dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"popRecommend":  "首页推荐",
			"classicReview": "大家常看",
			"swiperList":    swiperListVos,
		},
	}
}
