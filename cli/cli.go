package cli

import (
	"fmt"
	"os"
)

func New(app *App) *App {
	newApp := new(App)
	newApp = app
	return newApp

}

func (app *App) Run(args []string) error {
	if len(args) == 1 {
		app.PrintHelp()
	}
	task := args[1]
	data := make([]string, 0)
	save := false
	for i := 0; i < len(args); i++ {
		if args[i] == task && !save {
			save = true
		} else {
			for _, b := range app.Sub {
				if b.Use == args[i] {
					save = false
					break
				}
			}
			if args[i][0] == '-' {
				break
			}
			if save {
				data = append(data, args[i])
			}
		}

	}

	switch task {
	case "-h", "--h", "help":
		app.PrintHelp()
	default:
		for _, a := range app.Sub {

			if a.Use == task {
				return a.Func(data)
			}
		}
		app.PrintHelp()
	}
	return fmt.Errorf("unknown task %s", task)
}

func (app App) PrintHelp() {
	fmt.Printf("%s %v.%c %v\n", app.Name, app.Version, app.Release[0], app.Authors)
	fmt.Println("\nDescription:", app.Desc)
	fmt.Println("\nUSAGE:", app.Usage)
	fmt.Println("\nSub:")
	for _, s := range app.Sub {
		fmt.Println("  ", s.Use, "\t->\t", s.Usage)
	}
	os.Exit(1)
}

func (app App) getFlag(fl string, args []string) string {
	for i, a := range args {
		if a == fl {
			return args[i+1]
		}
	}
	return ""
}
