package commands

import (
	"fmt"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"goravel/app/services"
)

type ScanLink struct {
}

// 命令名称
func (receiver *ScanLink) Signature() string {
	return "scanLink"
}

// 命令描述
func (receiver *ScanLink) Description() string {
	return "scanLink"
}

// Extend The console command extend.
func (receiver *ScanLink) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *ScanLink) Handle(ctx console.Context) error {

	err := services.NewJumpService().ScanLink()
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
