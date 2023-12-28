package controller

import (
	"PM_backend/app/response"
	"PM_backend/model"
	"PM_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func PingHandler(c *gin.Context) {
	response.MyResponse(c, http.StatusOK, "Pong!", nil)
}

func RegisterHandler(c *gin.Context) {
	u := response.User{}
	err := c.BindJSON(&u)
	if err != nil {
		response.MyResponse(c, http.StatusPreconditionFailed, "Bind input fail, "+err.Error(), nil)
		return
	}
	mu := model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: utils.GetHashPwd(u.Password),
	}
	model.DB.Create(&mu)
	response.MyResponse(c, http.StatusOK, "Register success.", nil)
}

func LoginHandler(c *gin.Context) {
	u := response.User{}
	r := u
	err := c.BindJSON(&u)
	if err != nil {
		response.MyResponse(c, http.StatusInternalServerError, "Bind input fail, "+err.Error(), nil)
		logrus.Info("Bind input fail, " + err.Error())
		return
	}
	if u.Name != "" && u.Email == "" {
		model.DB.Table("user").Where("name = ?", u.Name).Find(&r)
		if r.Password == utils.GetHashPwd(u.Password) {
			token, err := utils.GenToken(r.Name, utils.MyKey)
			if err != nil {
				response.MyResponse(c, http.StatusInternalServerError, "Generate token fail, "+err.Error(), nil)
				logrus.Info("Token generation fail, " + err.Error())
				return
			}
			c.SetCookie("PM_backend", token, 36000, "/", "127.0.0.1", false, true)
		} else {
			response.MyResponse(c, http.StatusPreconditionFailed, "User input wrong.", nil)
			return
		}
	} else if u.Name == "" && u.Email != "" {
		model.DB.Table("user").Where("email = ?", u.Email).Find(&r)
		if r.Password == utils.GetHashPwd(u.Password) {
			token, err := utils.GenToken(r.Name, utils.MyKey)
			if err != nil {
				response.MyResponse(c, http.StatusInternalServerError, "Generate token fail, "+err.Error(), nil)
				logrus.Info("Token generation fail, " + err.Error())
				return
			}
			c.SetCookie("PM_backend", token, 36000, "/", "127.0.0.1", false, true)
		} else {
			response.MyResponse(c, http.StatusPreconditionFailed, "User input wrong", nil)
			return
		}
	}
	response.MyResponse(c, http.StatusOK, "Login success.", nil)
}

func LogoutHandler(c *gin.Context) {
	cookie, err := c.Cookie("PM_backend")
	if err != nil {
		response.MyResponse(c, http.StatusPreconditionFailed, "No cookie, no need to logout", nil)
		return
	}
	c.SetCookie("PM_backend", cookie, -1, "/", "127.0.0.1", false, true)
	response.MyResponse(c, http.StatusOK, "Logout success", nil)
}

func SearchSiteHandler(c *gin.Context) {
	s := response.Search{}
	var ss []response.Site
	err := c.BindJSON(&s)
	if err != nil {
		response.MyResponse(c, http.StatusInternalServerError, "Bind input fail, "+err.Error(), nil)
		logrus.Info("Bind input fail, " + err.Error())
		return
	}
	if s.SiteName != "" {
		if s.Continent != "" {
			if s.Country != "" {
				if s.City != "" {
					model.DB.Table("site").Where("continent = ? and country = ? and city = ? and site_name = ?", s.Continent, s.Country, s.City, s.SiteName).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("continent = ? and country = ? and site_name = ?", s.Continent, s.Country, s.SiteName).Find(&ss)
				}

			} else if s.Country == "" {
				if s.City != "" {
					model.DB.Table("site").Where("continent = ? and city = ? and site_name = ?", s.Continent, s.City, s.SiteName).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("continent = ? and site_name = ?", s.Continent, s.SiteName).Find(&ss)
				}
			}

		} else if s.Continent == "" {
			if s.Country != "" {
				if s.City != "" {
					model.DB.Table("site").Where("and country = ? and city = ? and site_name = ?", s.Country, s.City, s.SiteName).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("country = ? and site_name = ?", s.Country, s.SiteName).Find(&ss)
				}

			} else if s.Country == "" {
				if s.City != "" {
					model.DB.Table("site").Where("city = ? and site_name = ?", s.City, s.SiteName).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("site_name = ?", s.SiteName).Find(&ss)
				}
			}
		}

	} else if s.SiteName == "" {
		if s.Continent != "" {
			if s.Country != "" {
				if s.City != "" {
					model.DB.Table("site").Where("continent = ? and country = ? and city = ?", s.Continent, s.Country, s.City).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("continent = ? and country = ?", s.Continent, s.Country).Find(&ss)
				}

			} else if s.Country == "" {
				if s.City != "" {
					model.DB.Table("site").Where("continent = ? and city = ?", s.Continent, s.City).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("continent = ?", s.Continent).Find(&ss)
				}
			}

		} else if s.Continent == "" {
			if s.Country != "" {
				if s.City != "" {
					model.DB.Table("site").Where("and country = ? and city = ?", s.Country, s.City).Find(&ss)
				} else if s.City == "" {
					model.DB.Table("site").Where("country = ?", s.Country).Find(&ss)
				}

			} else if s.Country == "" {
				if s.City != "" {
					model.DB.Table("site").Where("city = ?", s.City).Find(&ss)
				} else if s.City == "" {
					response.MyResponse(c, http.StatusPreconditionFailed, "No enough search queries!", nil)
					return
				}
			}
		}
	}

	response.MyResponse(c, http.StatusOK, "Sites are: ", ss)
}

func GetSiteAverageMark(c *gin.Context) {

}

func AddRemarkHandler(c *gin.Context) {
	currentUser, ok := c.Get("username")
	if !ok {
		response.MyResponse(c, http.StatusInternalServerError, "Server context error.", nil)
		return
	}
	a := response.AddRemark{}
	err := c.BindJSON(&a)
	if err != nil {
		response.MyResponse(c, http.StatusInternalServerError, "Bind input fail, "+err.Error(), nil)
		logrus.Info("Bind input fail, " + err.Error())
		return
	}
	if a.UserName != currentUser {
		response.MyResponse(c, http.StatusConflict, "You can create others' remark.", nil)
		return
	}
	su := response.SUser{}
	ss := response.SSite{}
	model.DB.Table("user").Where("name = ?", a.UserName).Find(&su)
	model.DB.Table("site").Where("site_name = ?", a.SiteName).Find(&ss)
	if su.Id == 0 || ss.Id == 0 {
		response.MyResponse(c, http.StatusPreconditionFailed, "There is something wrong with userName or siteName. Please check.", nil)
		return
	}
	r := response.Remark{
		Content: a.Content,
		Mark:    a.Mark,
		UserId:  su.Id,
		SiteId:  ss.Id,
	}
	tmp := model.Remark{}
	model.DB.Table("remark").Where("user_id = ? and site_id = ?", su.Id, ss.Id).Find(&tmp)
	if tmp.Id != 0 {
		response.MyResponse(c, http.StatusPreconditionFailed, "You have remarked before.", nil)
		return
	}
	model.DB.Table("remark").Create(&r)
	response.MyResponse(c, http.StatusOK, "Remark success.", nil)
}

func DeleteRemarkHandler(c *gin.Context) {
	currentUser, ok := c.Get("username")
	if !ok {
		response.MyResponse(c, http.StatusInternalServerError, "Server context error.", nil)
		return
	}
	a := response.AddRemark{}
	err := c.BindJSON(&a)
	if err != nil {
		response.MyResponse(c, http.StatusInternalServerError, "Bind input fail, "+err.Error(), nil)
		logrus.Info("Bind input fail, " + err.Error())
		return
	}
	if a.UserName != currentUser {
		response.MyResponse(c, http.StatusConflict, "Not your own remark.", nil)
		return
	}
	su := response.SUser{}
	ss := response.SSite{}
	model.DB.Table("user").Where("name = ?", a.UserName).Find(&su)
	model.DB.Table("site").Where("site_name = ?", a.SiteName).Find(&ss)
	if su.Id == 0 || ss.Id == 0 {
		response.MyResponse(c, http.StatusPreconditionFailed, "Wrong input.", nil)
		return
	}
	r := model.Remark{
		UserId: su.Id,
		SiteId: ss.Id,
	}
	model.DB.Table("remark").Where("user_id = ? and site_id = ?", r.UserId, r.SiteId).Delete(&r)
	response.MyResponse(c, http.StatusOK, "Delete success.", nil)
	return
}

func ModifyRemarkHandler(c *gin.Context) {
	currentUser, ok := c.Get("username")
	if !ok {
		response.MyResponse(c, http.StatusInternalServerError, "Server context error.", nil)
		return
	}
	a := response.AddRemark{}
	err := c.BindJSON(&a)
	if err != nil {
		response.MyResponse(c, http.StatusInternalServerError, "Bind input fail, "+err.Error(), nil)
		logrus.Info("Bind input fail, " + err.Error())
		return
	}
	if a.UserName != currentUser {
		response.MyResponse(c, http.StatusConflict, "Not your own remark.", nil)
		return
	}
	su := response.SUser{}
	ss := response.SSite{}
	model.DB.Table("user").Where("name = ?", a.UserName).Find(&su)
	model.DB.Table("site").Where("site_name = ?", a.SiteName).Find(&ss)
	if su.Id == 0 || ss.Id == 0 {
		response.MyResponse(c, http.StatusPreconditionFailed, "Wrong input.", nil)
		return
	}
	r := model.Remark{
		Content: a.Content,
		Mark:    a.Mark,
		UserId:  su.Id,
		SiteId:  ss.Id,
	}
	model.DB.Table("remark").Where("user_id = ? and site_id = ?", r.UserId, r.SiteId).Updates(&r)
	response.MyResponse(c, http.StatusOK, "Modify success.", nil)
	return
}

func GetUserHistoryRemark(c *gin.Context) {

}
