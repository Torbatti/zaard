package zaard

var Version = "(untracked)"

type Zaard struct{}
type Config struct{}

func New() *Zaard           { return &Zaard{} }
func NewWithConfig() *Zaard { return &Zaard{} }

func (zrd *Zaard) Start()   {}
func (zrd *Zaard) Execute() {}
