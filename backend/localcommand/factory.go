package localcommand

import (
	"fmt"
	"strings"
	"syscall"
	"time"

	"github.com/yudai/gotty/server"
)

type Options struct {
	CloseSignal  int `hcl:"close_signal" flagName:"close-signal" flagSName:"" flagDescribe:"Signal sent to the command process when gotty close it (default: SIGHUP)" default:"1"`
	CloseTimeout int `hcl:"close_timeout" flagName:"close-timeout" flagSName:"" flagDescribe:"Time in seconds to force kill process after client is disconnected (default: -1)" default:"-1"`
}

type Factory struct {
	command string
	argv    []string
	options *Options
	opts    []Option
}

func NewFactory(command string, argv []string, options *Options) (*Factory, error) {
	opts := []Option{WithCloseSignal(syscall.Signal(options.CloseSignal))}
	if options.CloseTimeout >= 0 {
		opts = append(opts, WithCloseTimeout(time.Duration(options.CloseTimeout)*time.Second))
	}

	return &Factory{
		command: command,
		argv:    argv,
		options: options,
		opts:    opts,
	}, nil
}

func (factory *Factory) Name() string {
	return "local command"
}

func (factory *Factory) New(params map[string][]string) (server.Slave, error) {
	argv := make([]string, len(factory.argv))
	copy(argv, factory.argv)

	if params["arg"] != nil && len(params["arg"]) > 0 {
		for _, v := range params["arg"] {
			key := strings.Split(v, "=")[0]
			if key == "user" || key == "git" {
				continue // prevent manual set in url to overwrite these value
			}
			argv = append(argv, v)
		}
	}
	// fmt.Printf("arg after argv: %#v, params: %#v\n", argv, params)

	if params["user"] != nil && len(params["user"]) > 0 {
		argv = append(argv, params["user"][0])
	}
	if params["git"] != nil && len(params["git"]) > 0 {
		argv = append(argv, params["git"][0])
	}
	if params["env"] != nil && len(params["env"]) > 0 {
		env := params["env"][0]
		argv = append(argv, "env="+env)
	}

	fmt.Printf("setting args: %v\n", argv)
	return New(factory.command, argv, factory.opts...)
}
