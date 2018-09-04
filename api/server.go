package api

import (
	"github.com/gookit/ini"
	"github.com/gookit/redis-viewer/app"
	"github.com/gookit/sux"
)

// BaseAPI controller
type BaseAPI struct {
}

// ServerAPI controller
type ServerAPI struct {
	BaseAPI
}

// AddRoutes for the controller
func (a *ServerAPI) AddRoutes(g *sux.Router) {
	g.GET("", a.Index)
	g.GET("/names", a.Names)
	g.GET(RdsName, a.Get)
	g.POST("", a.Create)
	g.PUT(RdsName, a.Update)
	g.DELETE(RdsName, a.Delete)
}

// Servers get redis server list
func (a *ServerAPI) Index(c *sux.Context) {
	var servers []ini.Section
	for _, name := range app.Names {
		if conf, ok := app.Cfg.StringMap(name); ok {
			conf["name"] = name
			servers = append(servers, conf)
		}
	}

	c.JSONBytes(200, app.JSON(servers))
}

// Names get redis server names from config
func (a *ServerAPI) Names(c *sux.Context) {
	ss, _ := app.Cfg.Strings("servers", ",")

	c.JSONBytes(200, app.JSON(ss))
}

// Get a redis server config by name
func (a *ServerAPI) Get(c *sux.Context) {
	name := c.Param("name")

	if conf, ok := app.Cfg.StringMap(name); ok {
		conf["name"] = name
		c.JSONBytes(200, app.JSON(conf))
	} else {
		c.JSONBytes(404, app.ErrJSON(404, "not found"))
	}
}

// Create new redis server config
func (a *ServerAPI) Create(c *sux.Context) {

}

// Update a redis server config
func (a *ServerAPI) Update(c *sux.Context) {

}

// Delete a redis server config
func (a *ServerAPI) Delete(c *sux.Context) {
	name, ok := c.QueryParam("name")
	if !ok || name == ""{
		bs := app.ErrJSON(2, "invalid request params")
		c.JSONBytes(200, bs)
		return
	}

	app.Cfg.DelSection(name)
	c.JSONBytes(200, app.JSON(nil))
}