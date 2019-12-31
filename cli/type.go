package cli

type AppFunc func([]string) error
type App struct {
	Name    string
	Version float32
	Release string
	Desc    string
	Authors []AuthorData
	Usage   string
	Sub     []SubCommand
	Flags   []Flag
	Config  interface{}
}

type AuthorData struct {
	Name  string
	Email string
}

type SubCommand struct {
	Use   string
	Usage string
	Func  AppFunc
}

type Flag struct {
	Alias []string
	Desc  string
}
