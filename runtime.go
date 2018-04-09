package go_ofmdbapi

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/casualjim/go-app"
	"github.com/casualjim/go-app/tracing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/philhug/go-ofm-api/api/dbmodel"
	"github.com/spf13/viper"
)

// NewRuntime creates a new application level runtime that encapsulates the shared services for this application
func NewRuntime(app app.Application) (*Runtime, error) {
	dsn := app.Config().GetString("dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	services, err := dbmodel.NewServicesDBModel(db)
	if err != nil {
		return nil, err
	}
	serviceInstances, err := dbmodel.NewServiceInstancesDBModel(db)
	if err != nil {
		return nil, err
	}
	users, err := dbmodel.NewUsersDBModel(db)
	if err != nil {
		return nil, err
	}

	return &Runtime{
		db:    db,
		app:   app,
		services: services,
		serviceInstances: serviceInstances,
		users: users,
	}, nil
}

// Runtime encapsulates the shared services for this application
type Runtime struct {
	db    *sql.DB
	app   app.Application
	users *dbmodel.DbModel
	services *dbmodel.DbModel
	serviceInstances *dbmodel.DbModel
}

// DB returns the users model
func (r *Runtime) Users() *dbmodel.DbModel {
	return r.users
}
func (r *Runtime) Services() *dbmodel.DbModel {
	return r.services
}
func (r *Runtime) ServiceInstances() *dbmodel.DbModel {
	return r.serviceInstances
}

// DB returns the persistent store
func (r *Runtime) DB() *sql.DB {
	return r.db
}

// Tracer returns the root tracer, this is typically the only one you need
func (r *Runtime) Tracer() tracing.Tracer {
	return r.app.Tracer()
}

// Logger gets the root logger for this application
func (r *Runtime) Logger() logrus.FieldLogger {
	return r.app.Logger()
}

// NewLogger creates a new named logger for this application
func (r *Runtime) NewLogger(name string, fields logrus.Fields) logrus.FieldLogger {
	return r.app.NewLogger(name, fields)
}

// Config returns the viper config for this application
func (r *Runtime) Config() *viper.Viper {
	return r.app.Config()
}
