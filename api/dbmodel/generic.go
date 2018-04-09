package dbmodel

import (
	sq "github.com/Masterminds/squirrel"

	"database/sql"
	"errors"
	"fmt"
	"github.com/philhug/go-ofm-api/gen/models"
	"log"
)

const ()

func NewDBModel(db *sql.DB, config *DbConfig) (*DbModel, error) {
	dbmodel := &DbModel{db: sq.NewStmtCacher(db), config: config}
	return dbmodel, dbmodel.LoadTypes()
}

type DbTypes map[string]DbFieldType

type DbModel struct {
	//db *sql.DB
	db       sq.DBProxyContext
	config   *DbConfig
	idToType DbTypes
	typeToId DbTypes
}

type DbConfig struct {
	// Service Instancce
	Table                string
	Pk                   string
	ExtraColumns         []string
	ParentColumn         string
	PropertyPk           string
	PropertyIdColumn     string
	PropertyValueColumn  string
	DataTables           map[int]string
	TypeTable            string
	TypeTableId          string
	TypeTableFormat      string
	TypeTableDescription string
	TypeTableMultipleUse string
	TypeTableHidden      string

	// General
	BlobEmbed bool
}

type DbFieldType struct {
	Id              int
	Description     string
	Format          int
	MultipleUse     bool
	ServiceCategory int
	Hidden          bool
}

func (dbm *DbModel) LoadByPK(parentId *string, id string, extraFields ...interface{}) (map[string]interface{}, error) {
	db := dbm.db
	c := dbm.config

	// fetch main record
	// TODO: revision + deleted flag
	pred := sq.Eq{c.ParentColumn: parentId, c.Pk: id}
	columns := append([]string{c.Pk}, c.ExtraColumns...)
	rows, err := sq.Select(columns...).From(c.Table).Where(pred).RunWith(db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Not found")
	}
	if err := rows.Scan(extraFields...); err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	// fetch properties
	for i, tableName := range c.DataTables {
		if tableName == "" {
			continue
		}

		var colums = []string{c.PropertyIdColumn, c.PropertyValueColumn}
		if i == 6 {
			// skip data for blob
			colums[1] = c.PropertyPk
		}
		pred2 := sq.And{pred, sq.NotEq{c.PropertyIdColumn: 19}} // ignore Data Log Entry
		rows, err = sq.Select(colums...).From(tableName).Where(pred2).RunWith(db).Query()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var key int
			var value string
			if err := rows.Scan(&key, &value); err != nil {
				return nil, err
			}
			if i == 6 {
				value = fmt.Sprintf("http://127.0.0.1:8081/0.1/blobstore/%s", value)
			}
			stype := dbm.idToType[string(key)]
			skey := stype.Description
			if stype.MultipleUse {
				multi := []string{}
				if result[skey] != nil {
					multi = result[skey].([]string)
				}
				result[skey] = append(multi, value)
			} else {
				result[skey] = value
			}
		}
	}

	return result, nil
}

func (dbm *DbModel) Query(parentId string, query map[string]interface{}, extraProperty *string) (*models.NodeList, error) {
	db := dbm.db
	c := dbm.config
	pred := sq.And{sq.Eq{c.Table + ".deleted": false}, sq.Eq{c.Table + "." + c.ParentColumn: parentId}}

	var extraField DbFieldType
	qs := []string{c.Table + "." + c.Pk, c.Table + ".deleted", c.Table + ".Revision", c.Table + ".dateOfCreation"}
	if extraProperty != nil {
		if *extraProperty != "xml" {
			return nil, errors.New("invalid extra property")
		}
		extraField = dbm.typeToId[*extraProperty]
		qs = append(qs, "d."+c.PropertyValueColumn)
	}
	q := sq.Select(qs...).From(c.Table).Where(pred)

	if extraProperty != nil {
		dformat := extraField.Format
		jointable := c.DataTables[dformat]
		join := fmt.Sprintf("%s AS %s USING (%s)", jointable, "d", c.Pk)
		q = q.LeftJoin(join).Where(sq.Eq{"d." + c.PropertyIdColumn: extraField.Id})
	}

	var ctr int
	for k, value := range query {
		if value == nil {
			continue
		}
		dformat := dbm.typeToId[k].Format
		dId := dbm.typeToId[k].Id
		jointable := c.DataTables[dformat]
		joinas := fmt.Sprintf("j%d", ctr)
		join := fmt.Sprintf("%s AS %s USING (%s)", jointable, joinas, c.Pk)
		q = q.Join(join).Where(sq.And{sq.Eq{joinas + "." + c.PropertyIdColumn: dId}, sq.Eq{joinas + "." + c.PropertyValueColumn: value}})
		ctr++
	}

	log.Println(q.ToSql())
	rows, err := q.RunWith(db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result models.NodeList
	for rows.Next() {
		var node models.Node
		scanFields := []interface{}{&node.ID, &node.Deleted, &node.Rev, &node.Updated}
		if extraProperty != nil {
			scanFields = append(scanFields, &node.XML)
		}
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}
		result.Items = append(result.Items, &node)
	}
	return &result, nil
}

func (dbm *DbModel) LoadTypes() error {
	db := dbm.db
	c := dbm.config

	if c.TypeTableId == "" {
		return nil
	}
	rows, err := sq.Select(c.TypeTableId, c.TypeTableFormat, c.TypeTableDescription, c.TypeTableMultipleUse).From(c.TypeTable).RunWith(db).Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	idToType := DbTypes{}
	typeToId := DbTypes{}

	for rows.Next() {
		var fieldType DbFieldType
		if err := rows.Scan(&fieldType.Id, &fieldType.Format, &fieldType.Description, &fieldType.MultipleUse); err != nil {
			return err
		}
		idToType[string(fieldType.Id)] = fieldType
		typeToId[fieldType.Description] = fieldType
	}
	log.Println(idToType)
	dbm.idToType = idToType
	dbm.typeToId = typeToId

	return nil
}
