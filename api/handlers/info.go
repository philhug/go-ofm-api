package handlers

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/go-openapi/runtime/middleware"
	ofmdb "github.com/philhug/go-ofm-api"
	"github.com/philhug/go-ofm-api/gen/models"
	"github.com/philhug/go-ofm-api/gen/restapi/operations/information"
	"log"
)

const (
	AMMNT_FIR = "AMMNT_FIR"
)

func NewInformationGetRegionsHandler(rt *ofmdb.Runtime) information.GetRegionsHandler {
	return &infoGetRegionHandler{rt: rt}
}

type infoGetRegionHandler struct {
	rt *ofmdb.Runtime
}

func (d *infoGetRegionHandler) Handle(params information.GetRegionsParams, more interface{}) middleware.Responder {
	log := d.rt.Logger()
	db := d.rt.DB()

	rows, err := sq.Select("PK", "Name", "IcaoCode").From(AMMNT_FIR).RunWith(db).Query()
	if err != nil {
		// TODO
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	rl := &models.RegionList{}
	for rows.Next() {
		r := &models.Region{}
		if err := rows.Scan(&r.IDInternal, &r.Name, &r.ID); err != nil {
			log.Fatal(err)
		}
		rl.Items = append(rl.Items, r)
	}

	return information.NewGetRegionsOK().WithPayload(rl)
}

type infoGetUserInfoHandler struct {
	rt *ofmdb.Runtime
}

func NewInformationGetUserInfoHandler(rt *ofmdb.Runtime) information.GetUserInfoHandler {
	return &infoGetUserInfoHandler{rt: rt}
}

func (d *infoGetUserInfoHandler) Handle(params information.GetUserInfoParams, more interface{}) middleware.Responder {
	//log := d.rt.Logger()
	//db := d.rt.DB()
	usermodel := d.rt.Users()

	user, err := usermodel.LoadUser("0")
	if err != nil {
		// TODO
		log.Fatal(err)
	}
	return information.NewGetUserInfoOK().WithPayload(user)
}
