package controller

import (
	"awesomeProject0511/common"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SwiperList struct {
	Id        string `gorm:"primary_key" json:"id"`
	Uid       string `gorm:"int(100);unique;not null" json:"uid"`
	ImgUrl    string `gorm:"varchar(255);not null" json:"imgUrl"`
	VideoHref string `gorm:"varchar(255);not null" json:"videoHref"`
}

type Data struct {
	SwiperList    []SwiperList `json:"swiperList"`
	PopRecommend  string       `json:"popRecommend"`
	ClassicReview string       `json:"classicReview"`
}

func GetHomeList(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()

	var data SwiperList
	db.AutoMigrate(data)

	var users []string
	res, _ := db.Raw("select id,uid,img_url,video_href from swiper_lists").Rows()
	for res.Next() {
		res.Scan(&data.Id, &data.VideoHref, &data.Uid, &data.ImgUrl)
		// fmt.Printf("data: %v\n", data)
		users = append(users, data.Id, data.VideoHref, data.Uid, "channel/"+data.ImgUrl)
	}
	y := Data{[]SwiperList{
		{users[0], users[1], users[2], users[3]},
		{users[4], users[5], users[6], users[7]},
		{users[8], users[9], users[10], users[11]},
		{users[12], users[13], users[14], users[15]}},
		"热门推荐", "大家常看"}

	//f.SwiperList = append(f.SwiperList, y)
	c.JSON(200, gin.H{
		"ret":  true,
		"data": y,
	})

}

/*
CREATE TABLE `aikan`.`home_tops` ( `id` INT(20) NOT NULL AUTO_INCREMENT , `img_href` VARCHAR(50) NOT NULL , `img_title` VARCHAR(100) NOT NULL , `img_url` VARCHAR(50) NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;
INSERT INTO `home_tops` (`id`, `img_href`, `img_title`, `img_url`) VALUES (NULL, 'cont-1736469-15729657_adpkg-ad_hd', 'test4', 'https://www.awas.ink/image/5601950.jpg');
*/
