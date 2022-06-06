package main

import (
	"context"

	myCLI "github.com/alewkinr/pingo/cmd"
	"github.com/alewkinr/pingo/internal/config"
	"github.com/alewkinr/pingo/pkg/log"
	"github.com/alewkinr/pingo/pkg/validation"
	"github.com/segmentio/cli"
)

// CmdSendMessage — кодманда отправки сообщения без шаблона
const CmdSendMessage = "send"

// CmdSendMessageArgs — аргументы команды CmdSendMessage
// TODO: не придумал, как это унести в cmd.CLI
type CmdSendMessageArgs struct {
	_ struct{} `help:"Send custom message via command line"`
	// Destination — пункт назначения для отправленного сообщения
	Destination string `flag:"-d,--destination" help:"Destination for your message" validate:"required"` // TODO: добавить кастомный тег валидации, проверяющий формат. Его также можно будет юзать в config.Config
	// Message — текст сообщения для отправки
	Message string `flag:"-m,--message" help:"Message text" validate:"required_with=Destination,min=1"`
}

// CmdSendTemplateMessage — команда для отправки сообщения по шаблону
const CmdSendTemplateMessage = "send-template"

// CmdSendTemplateMessageArgs — аргументы команды CmdSendTemplateMessage
type CmdSendTemplateMessageArgs struct {
	_ struct{} `help:"Reads templates config file and send template by its name"`
	// TemplatesConfigFile — путь до файла конфигурации шаблонов
	TemplatesConfigFile string `flag:"-c,--config" help:"Templates config file" default:"templates.yaml" validate:"required,file"`
	// TemplateName — название шаблона для отправки
	TemplateName string `flag:"-n,--name" help:"Template name" validate:"required_with=TemplatesConfigFile,min=1"`
}

// runCLI — запускаем CLI
func runCLI(ctx context.Context) {
	lgr := log.SetUpLogging()
	v := validation.NewPlayground()
	settings := config.MustInitConfig()

	app := myCLI.NewCLI(lgr, v, settings)
	app.L().Debugln("🎉 New CLI app initialized!")

	// ref: https://pkg.go.dev/github.com/segmentio/cli#example-CommandSet-Help2
	cli.ExecContext(ctx, cli.CommandSet{
		CmdSendMessage: cli.Command(func(args CmdSendMessageArgs) (int, error) {
			if validateArgsErr := v.Validate(args); validateArgsErr != nil {
				return myCLI.ExitWithErr, validateArgsErr
			}
			return app.Send(args.Destination, args.Message)
		}),
		CmdSendTemplateMessage: cli.Command(
			func(args CmdSendTemplateMessageArgs) (int, error) {
				if validateArgsErr := v.Validate(args); validateArgsErr != nil {
					return myCLI.ExitWithErr, validateArgsErr
				}
				return app.SendTemplate(args.TemplatesConfigFile, args.TemplateName)
			}),
	})
}

func main() {
	ctx := context.TODO()
	runCLI(ctx)
}
