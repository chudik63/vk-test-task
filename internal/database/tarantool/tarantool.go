package tarantool

import (
	"fmt"
	"vk-test-task/internal/config"

	"github.com/tarantool/go-tarantool"
)

type Tarantool struct {
	conn *tarantool.Connection
}

func New(cfg config.TarantoolConfig) (*Tarantool, error) {
	conn, err := tarantool.Connect(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), tarantool.Opts{
		User: cfg.User,
		Pass: cfg.Password,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Tarantool: %w", err)
	}

	return &Tarantool{conn}, nil
}

func (t *Tarantool) Close() error {
	return t.conn.Close()
}
