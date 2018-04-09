package dbmodel

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/philhug/go-ofm-api/gen/models"
	"log"
)

func NewServicesDBModel(db *sql.DB) (*DbModel, error) {
	config := &DbConfig{
		Table:        "S1T",
		Pk:           "ServiceID",
		ExtraColumns: []string{"ServiceName", "deveoplerDescription"},
		DataTables: map[int]string{
			1: "S2A1",
			2: "S2A2",
			3: "S2A3",
			4: "S2A4",
			5: "S2A5",
			6: "S2A6",
			7: "S2A7",
		},
		TypeTable:            "S2T",
		TypeTableId:          "ServicePropertiesTypeID",
		TypeTableDescription: "ServicePropertiestypeDescription",
		TypeTableFormat:      "ServicePropertiesTypeFormat",
		TypeTableMultipleUse: "multipleUse",
		TypeTableHidden:      "",
	}

	return NewDBModel(db, config)
}

func NewServiceInstancesDBModel(db *sql.DB) (*DbModel, error) {
	config := &DbConfig{
		Table:        "S4",
		Pk:           "ServiceEntityID",
		ExtraColumns: []string{"searchTags"},
		DataTables: map[int]string{
			1: "S3A1",
			2: "S3A2",
			3: "S3A3",
			4: "S3A4",
			5: "S3A5",
			6: "S3A6",
			7: "S3A7",
		},
		ParentColumn:         "ParentServiceID",
		PropertyPk:           "PK",
		PropertyIdColumn:     "ServiceEntityPropertiesTypeId",
		PropertyValueColumn:  "ServiceEntityPropertiesTypeValue",
		TypeTable:            "S3T",
		TypeTableId:          "ServiceEntityTypeID",
		TypeTableDescription: "ServiceEntityTypeDescription",
		TypeTableFormat:      "ServiceEntityTypeFormat",
		TypeTableMultipleUse: "multipleUse",
		TypeTableHidden:      "hidden",
	}

	return NewDBModel(db, config)
}

func (dbm *DbModel) LoadServiceByName(id string) (string, error) {
	db := dbm.db
	c := dbm.config

	pred := sq.Eq{"ServiceName": id}
	rows, err := sq.Select(c.Pk).From(c.Table).Where(pred).RunWith(db).Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", errors.New("Not found")
	}
	var pk string
	if err := rows.Scan(&pk); err != nil {
		return "", err
	}

	return pk, nil
}

func (dbm *DbModel) LoadNode(parentId, id string) (*models.Node, error) {

	node := &models.Node{}
	var searchTags sql.NullString
	properties, err := dbm.LoadByPK(&parentId, id, &node.ID, &searchTags)
	if err != nil {
		return nil, err
	}
	j, err := json.Marshal(properties)
	if err != nil {
		return nil, err
	}
	log.Println(string(j))
	decoder := json.NewDecoder(bytes.NewReader(j))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&node.NodeProperties)
	return node, err
}

func (dbm *DbModel) QueryNode(dbId string, query *models.NodeQuery, delete *bool, bbox *string, fulltext *string) (*models.NodeList, error) {
	j, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	log.Println(string(j))

	var q map[string]interface{}
	err = json.Unmarshal(j, &q)

	extraProperty := "xml"
	items, err := dbm.Query(dbId, q, &extraProperty)
	return items, err
}
