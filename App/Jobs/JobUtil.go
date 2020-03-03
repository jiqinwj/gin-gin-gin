package Jobs

import "github.com/robfig/cron/v3"

var MyCron *cron.Cron
func init()  {
	MyCron=cron.New(cron.WithSeconds())
}