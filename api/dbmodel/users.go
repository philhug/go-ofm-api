package dbmodel

import (
	"database/sql"
	"encoding/json"
	"github.com/philhug/go-ofm-api/gen/models"
)

func NewUsersDBModel(db *sql.DB) (*DbModel, error) {
	config := &DbConfig{
		Pk:                   "UserID",
		Table:                "U1T",
		ExtraColumns:         []string{"Username", "Password"},
		TypeTable:            "U2T",
		TypeTableId:          "UserPropertiesTypeID",
		TypeTableDescription: "UserPropertiesTypeDescription",
		TypeTableFormat:      "UserPropertiesTypeFormat",
		TypeTableMultipleUse: "multipleUse",
		TypeTableHidden:      "hidden",
	}

	return NewDBModel(db, config)
}

func (dbm *DbModel) LoadUser(id string) (*models.User, error) {

	user := &models.User{}
	var username, password string
	properties, err := dbm.LoadByPK(nil, id, &user.ID, &username, &password)
	if err != nil {
		return nil, err
	}
	j, err := json.Marshal(properties)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &user)
	return user, err
}
