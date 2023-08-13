package environment

import (
	"github.com/hampgoodwin/todo/internal/config"
)

var TestEnvironment = Environment{
	Config: config.Local,
}
