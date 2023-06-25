package controllers

import (
	"covid/db"
	"covid/entity"
	"covid/models"
	"covid/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type CovidController struct{}
type groupProvince struct {
	Province  string
	Cprovince int
}

type groupAge struct {
	Age  int
	Cage int
}

type response struct {
	Provice map[string]int
	Age     map[string]int
}

func (c CovidController) CovidSummary(ctx *gin.Context) {
	c.GetCovidCases(ctx)
	var gProvince []groupProvince
	var gAge []groupAge

	g := new(errgroup.Group)

	g.Go(func() error {
		if err := db.Conn.Model(&entity.CovidSummary{}).Select("province, count(province) as cprovince").Group("province").Find(&gProvince).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}
		return nil
	})
	g.Go(func() error {
		if err := db.Conn.Model(&entity.CovidSummary{}).Select("age, count(age) as cage").Group("age").Find(&gAge).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}

	gp := map[string]int{}
	for _, v := range gProvince {
		if v.Province == "" {
			v.Province = "N/A"
		}
		gp[v.Province] = v.Cprovince
	}

	ga := map[string]int{}
	for _, v := range gAge {
		if v.Age <= 30 {
			ga["0-30"] += v.Cage
		} else if v.Age <= 60 {
			ga["31-60"] += v.Cage
		} else {
			ga["60+"] += v.Cage
		}
	}

	response := response{
		Provice: gp,
		Age:     ga,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c CovidController) GetCovidCases(ctx *gin.Context) {
	var data models.DataCovidResponse
	g := new(errgroup.Group)

	g.Go(func() error {
		c.delteData(&entity.CovidSummary{}, ctx)
		return nil
	})

	g.Go(func() error {
		// call API
		url := "https://static.wongnai.com/devinterview/covid-cases.json"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error sending GET request: %v\n", err)
			return nil
		}
		defer resp.Body.Close()

		// Read the response body
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return nil
		}

		if err := json.Unmarshal(respBody, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return nil
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}

	c.goToCreate(data, ctx)

}

func (c CovidController) goToCreate(md models.DataCovidResponse, ctx *gin.Context) {

	var dto []entity.CovidSummary
	if len(md.Data) > 0 {
		for i, d := range md.Data {
			// c.goToCreate(d)

			dto = append(dto, entity.CovidSummary{
				No:             d.No,
				Age:            d.Age,
				Gender:         d.Gender,
				GenderEn:       d.GenderEn,
				Nation:         d.Nation,
				NationEn:       d.NationEn,
				Province:       d.Province,
				ProvinceID:     d.ProvinceID,
				ProvinceEn:     d.ProvinceEn,
				District:       d.District,
				StatQuarantine: d.StatQuarantine,
			})
			if d.ConfirmDate != "" {
				date, err := time.Parse("2006-01-02", d.ConfirmDate)
				if err != nil {
					fmt.Println("date: ", date, err)
					return
				}

				dto[i].ConfirmDate = int64(date.UnixMilli())

			}
		}
		repo := repository.Reposirory{}
		if err := repo.CreateCovid(&dto); err != nil {
			log.Println("Cannot create data covid summary: ", err)
			ctx.JSON(http.StatusBadRequest, err)
		}
	}
}

func (c CovidController) delteData(data interface{}, ctx *gin.Context) {
	repo := repository.Reposirory{}
	if err := repo.DeleteCovid(&data); err != nil {
		log.Println("Cannot delete data covid summary: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	return
}
