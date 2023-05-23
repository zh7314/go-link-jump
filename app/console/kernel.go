package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"
	"goravel/app/console/commands"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		facades.Schedule.Command("scanLink").EveryMinute(),
		//facades.Schedule.Call(func() {
		//	services.NewApiService().ScanLink()
		//}).EveryMinute(),
	}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&commands.ScanLink{},
	}
}
