package configloader

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hampgoodwin/errors"
	"github.com/hampgoodwin/todo/internal/config"
	"github.com/matryer/is"
)

func TestLoad(t *testing.T) {
	testCases := []struct {
		description string
		cfg         config.Config
		filepath    string
		envVars     map[string]string
		expected    config.Config
		err         error // TODO add specific err case catches
	}{
		{
			description: "json-file-empty-vars-empty-config-error-decoding",
			filepath:    "../../test/data/configloader/json.env.toml",
			err:         errors.New("(1, 1): parsing error: keys cannot contain { character"),
		},
		{
			description: "empty-file-empty-vars-empty-config-err-validation",
			filepath:    "../../test/data/configloader/empty.env.toml",
			err:         validator.ValidationErrors{},
		},
		{
			description: "full-file-empty-vars-empty-config",
			filepath:    "../../.env.toml.example",
			expected: config.Config{
				Environment: config.Environment{Type: "LOCAL"},
				Database: config.Database{
					Host:     "127.0.0.1",
					Port:     "5432",
					User:     "user",
					Pass:     "password",
					Database: "todo",
					SSLMode:  "disable",
				},
				GRPCServer: config.GRPCServer{
					Host: "localhost",
					Port: "5000",
				},
				NATS: config.NATS{
					Host: "localhost",
					Port: "4222",
					Wiretap: config.NATSWiretap{
						Enable: true,
						Host:   "localhost",
						Port:   "4222",
					},
				},
			},
		},
		{
			description: "full-file-full-vars-empty-config-overwrite-file",
			filepath:    "../../.env.toml.example",
			envVars: map[string]string{
				EnvType:        "DEV",
				DBHost:         "TODO_DB_HOST",
				DBPort:         "TODO_DB_PORT",
				DBUser:         "TODO_DB_USER",
				DBPass:         "TODO_DB_PASS",
				DBDatabase:     "TODO_DB_DATABASE",
				DBSSLMode:      "TODO_DB_SSLMODE",
				HTTPServerHost: "TODO_HTTP_SERVER_HOST",
				HTTPServerPort: "TODO_HTTP_SERVER_PORT",
				GRPCServerHost: "TODO_GRPC_SERVER_HOST",
				GRPCServerPort: "TODO_GRPC_SERVER_PORT",
				NATSHost:       "TODO_NATS_HOST",
				NATSPort:       "TODO_NATS_PORT",
				WiretapEnable:  "TODO_WIRETAP_ENABLE",
				WiretapHost:    "TODO_WIRETAP_HOST",
				WiretapPort:    "TODO_WIRETAP_PORT",
			},
			expected: config.Config{
				Environment: config.Environment{Type: "DEV"},
				Database: config.Database{
					Host:     "TODO_DB_HOST",
					Port:     "TODO_DB_PORT",
					User:     "TODO_DB_USER",
					Pass:     "TODO_DB_PASS",
					Database: "TODO_DB_DATABASE",
					SSLMode:  "TODO_DB_SSLMODE",
				},
				GRPCServer: config.GRPCServer{
					Host: "TODO_GRPC_SERVER_HOST",
					Port: "TODO_GRPC_SERVER_PORT",
				},
				NATS: config.NATS{
					Host: "TODO_NATS_HOST",
					Port: "TODO_NATS_PORT",
					Wiretap: config.NATSWiretap{
						Enable: false,
						Host:   "TODO_WIRETAP_HOST",
						Port:   "TODO_WIRETAP_PORT",
					},
				},
			},
		},
		{
			description: "full-file-partial-vars-empty-config-overwrite-file",
			filepath:    "../../.env.toml.example",
			envVars: map[string]string{
				EnvType:        "DEV",
				HTTPServerHost: "TODO_HTTP_SERVER_HOST",
				HTTPServerPort: "TODO_HTTP_SERVER_PORT",
			},
			expected: config.Config{
				Environment: config.Environment{Type: "DEV"},
				Database: config.Database{
					Host:     "127.0.0.1",
					Port:     "5432",
					User:     "user",
					Pass:     "password",
					Database: "todo",
					SSLMode:  "disable",
				},
				GRPCServer: config.GRPCServer{
					Host: "localhost",
					Port: "5000",
				},
				NATS: config.NATS{
					Host: "localhost",
					Port: "4222",
					Wiretap: config.NATSWiretap{
						Enable: true,
						Host:   "localhost",
						Port:   "4222",
					},
				},
			},
		},
		{
			description: "full-file-partial-vars-partial-config-overwrite-file",
			cfg:         config.Config{Environment: config.Environment{Type: "DEV"}},
			filepath:    "../../.env.toml.example",
			envVars: map[string]string{
				EnvType:        "LOCAL",
				HTTPServerHost: "TODO_HTTP_SERVER_HOST",
				HTTPServerPort: "TODO_HTTP_SERVER_PORT",
			},
			expected: config.Config{
				Environment: config.Environment{Type: "LOCAL"},
				Database: config.Database{
					Host:     "127.0.0.1",
					Port:     "5432",
					User:     "user",
					Pass:     "password",
					Database: "todo",
					SSLMode:  "disable",
				},
				GRPCServer: config.GRPCServer{
					Host: "localhost",
					Port: "5000",
				},
				NATS: config.NATS{
					Host: "localhost",
					Port: "4222",
					Wiretap: config.NATSWiretap{
						Enable: true,
						Host:   "localhost",
						Port:   "4222",
					},
				},
			},
		},
		{
			description: "full-file-partial-vars-partial-config-persist-merge",
			cfg:         config.Config{Environment: config.Environment{Type: "DEV"}},
			filepath:    "../../.env.toml.example",
			envVars: map[string]string{
				HTTPServerHost: "TODO_HTTP_SERVER_HOST",
				HTTPServerPort: "TODO_HTTP_SERVER_PORT",
			},
			expected: config.Config{
				Environment: config.Environment{Type: "LOCAL"},
				Database: config.Database{
					Host:     "127.0.0.1",
					Port:     "5432",
					User:     "user",
					Pass:     "password",
					Database: "todo",
					SSLMode:  "disable",
				},
				GRPCServer: config.GRPCServer{
					Host: "localhost",
					Port: "5000",
				},
				NATS: config.NATS{
					Host: "localhost",
					Port: "4222",
					Wiretap: config.NATSWiretap{
						Enable: true,
						Host:   "localhost",
						Port:   "4222",
					},
				},
			},
		},
	}

	a := is.New(t)
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.description), func(t *testing.T) {
			resetApplicationEnvironmentVariables()
			defer resetApplicationEnvironmentVariables()

			for k, v := range tc.envVars {
				os.Setenv(k, v)
			}

			actual, err := Load(tc.cfg, tc.filepath)
			if tc.err != nil {
				a.True(err != nil)
				// Using errors.As because it detects validator.ValidationErrors
				a.True(errors.As(err, &tc.err))
				return
			}
			a.NoErr(err)
			a.Equal(tc.expected, actual)
		})
	}
}

func TestLoadConfigurationFile(t *testing.T) {
	testCases := []struct {
		description string
		filepath    string
		expected    config.Config
		err         error
	}{
		{
			description: "not-toml-file",
			filepath:    "../../test/data/configloader/json.env.toml",
			err:         errors.New(""),
		},
		{
			description: "empty-file-empty-config",
			filepath:    "../../test/data/configloader/empty.env.toml",
		},
		{
			description: "full-file-full-config",
			filepath:    "../../.env.toml.example",
			expected: config.Config{
				Environment: config.Environment{Type: "LOCAL"},
				Database: config.Database{
					Host:     "127.0.0.1",
					Port:     "5432",
					User:     "user",
					Pass:     "password",
					Database: "todo",
					SSLMode:  "disable",
				},
				GRPCServer: config.GRPCServer{
					Host: "localhost",
					Port: "5000",
				},
				NATS: config.NATS{
					Host: "localhost",
					Port: "4222",
					Wiretap: config.NATSWiretap{
						Enable: true,
						Host:   "localhost",
						Port:   "4222",
					},
				},
			},
		},
		{
			description: "partial-file-partial-config",
			filepath:    "../../test/data/configloader/partial.env.toml",
			expected: config.Config{
				Environment: config.Environment{Type: "DEV"},
				GRPCServer: config.GRPCServer{
					Host: "localhost",
					Port: "5000",
				},
			},
		},
	}

	a := is.New(t)

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.description), func(t *testing.T) {
			actual, err := loadConfigurationFile(tc.filepath)
			if tc.err != nil {
				a.True(err != nil)
				a.True(errors.As(err, &tc.err))
				return
			}
			a.NoErr(err)
			a.Equal(tc.expected, actual)
		})
	}
}

func TestLoadEnvironmentVariables(t *testing.T) {
	testCases := []struct {
		description string
		envVars     map[string]string
		expected    config.Config
	}{
		{
			description: "empty-vars",
			envVars: map[string]string{
				EnvType:        "",
				DBHost:         "",
				DBUser:         "",
				DBPass:         "",
				DBDatabase:     "",
				DBPort:         "",
				HTTPServerHost: "",
				HTTPServerPort: "",
				GRPCServerHost: "",
				GRPCServerPort: "",
				NATSHost:       "",
				NATSPort:       "",
				WiretapEnable:  "",
				WiretapHost:    "",
				WiretapPort:    "",
			},
		},
		{
			description: "filled-vars",
			envVars: map[string]string{
				EnvType:        "TODO_ENVTYPE",
				DBHost:         "TODO_DBHOST",
				DBPort:         "TODO_DBPORT",
				DBUser:         "TODO_DBUSER",
				DBPass:         "TODO_DBPASS",
				DBDatabase:     "TODO_DBDATABASE",
				HTTPServerHost: "TODO_HTTP_SERVER_HOST",
				HTTPServerPort: "TODO_HTTP_SERVER_PORT",
				GRPCServerHost: "TODO_GRPC_SERVER_HOST",
				GRPCServerPort: "TODO_GRPC_SERVER_PORT",
				NATSHost:       "TODO_NATS_HOST",
				NATSPort:       "TODO_NATS_PORT",
				WiretapEnable:  "TODO_WIRETAP_ENABLE",
				WiretapHost:    "TODO_WIRETAP_HOST",
				WiretapPort:    "TODO_WIRETAP_PORT",
			},
			expected: config.Config{
				Environment: config.Environment{Type: "TODO_ENVTYPE"},
				Database: config.Database{
					Host:     "TODO_DBHOST",
					Port:     "TODO_DBPORT",
					User:     "TODO_DBUSER",
					Pass:     "TODO_DBPASS",
					Database: "TODO_DBDATABASE",
				},
				GRPCServer: config.GRPCServer{
					Host: "TODO_GRPC_SERVER_HOST",
					Port: "TODO_GRPC_SERVER_PORT",
				},
				NATS: config.NATS{
					Host: "TODO_NATS_HOST",
					Port: "TODO_NATS_PORT",
					Wiretap: config.NATSWiretap{
						Enable: false,
						Host:   "TODO_WIRETAP_HOST",
						Port:   "TODO_WIRETAP_PORT",
					},
				},
			},
		},
		{
			description: "partial-vars",
			envVars: map[string]string{
				EnvType:        "TODO_ENVTYPE",
				HTTPServerHost: "TODO_HTTP_SERVER_HOST",
				HTTPServerPort: "TODO_HTTP_SERVER_PORT",
			},
			expected: config.Config{
				Environment: config.Environment{Type: "TODO_ENVTYPE"},
				GRPCServer: config.GRPCServer{
					Host: "TODO_HTTP_SERVER_HOST",
					Port: "TODO_HTTP_SERVER_PORT",
				},
			},
		},
	}

	a := is.New(t)
	for i, tc := range testCases {
		// clean the
		t.Run(fmt.Sprintf("%d:%s", i, tc.description), func(t *testing.T) {
			resetApplicationEnvironmentVariables()
			defer resetApplicationEnvironmentVariables()

			for k, v := range tc.envVars {
				os.Setenv(k, v)
			}

			actual := loadAndMergeEnvironmentVariables(config.Config{})
			a.Equal(tc.expected, actual)
		})

	}
}

func resetApplicationEnvironmentVariables() {
	for _, k := range EnvironmentVariableKeys {
		os.Unsetenv(k)
	}
}
