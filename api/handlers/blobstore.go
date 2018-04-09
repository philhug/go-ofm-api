package handlers

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/go-openapi/runtime/middleware"
	ofmdb "github.com/philhug/go-ofm-api"
	"github.com/philhug/go-ofm-api/gen/restapi/operations/native_client"
	"bytes"
	"io/ioutil"
)


func NewNativeClientGetBlobHandler(rt *ofmdb.Runtime) native_client.GetBlobHandler {
	return &nativeClientGetBlobHandler{rt: rt}
}

type nativeClientGetBlobHandler struct {
	rt *ofmdb.Runtime
}

func (d *nativeClientGetBlobHandler) Handle(params native_client.GetBlobParams, more interface{}) middleware.Responder {
	log := d.rt.Logger()
	db := d.rt.DB()

	pred := sq.Eq{"PK": params.ID}
	rows, err := sq.Select("ServiceEntityPropertiesTypeValue").From("S3A6").Where(pred).RunWith(db).Query()
	if err != nil {
		// TODO
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	var blob []byte
	if rows.Next() {
		if err := rows.Scan(&blob); err != nil {
			log.Fatal(err)
		}
	} else {
		// TODO
		log.Fatal("Record not found")
	}

	return native_client.NewGetBlobOK().WithPayload(ioutil.NopCloser(bytes.NewReader(blob)))
}

