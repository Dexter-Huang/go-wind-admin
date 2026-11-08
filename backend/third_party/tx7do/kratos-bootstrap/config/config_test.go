package config

import (
	"os"
	"path/filepath"
	"testing"

	kratosconfig "github.com/go-kratos/kratos/v2/config"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

func TestConfigurationSourceIsolation(t *testing.T) {
	dir := t.TempDir()
	write := func(name, content string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	write("server.yaml", "server:\n  rest:\n    addr: ':19990'\nlocal_only: true\n")
	write("config.yaml", "config:\n  type: nacos\n  nacos:\n    address: unreachable.invalid\n")
	calls := 0
	remote := filepath.Join(t.TempDir(), "remote.yaml")
	if err := os.WriteFile(remote, []byte("server:\n  rest:\n    addr: ':19991'\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := RegisterFactory(TypeNacos, func(*conf.RemoteConfig) (kratosconfig.Source, error) {
		calls++
		return NewFileConfigSource(remote), nil
	}); err != nil {
		t.Fatal(err)
	}
	for _, mode := range []string{"local", "nacos"} {
		t.Run(mode, func(t *testing.T) {
			cfg, err := NewConfigProvider(dir, mode)
			if err != nil {
				t.Fatal(err)
			}
			defer cfg.Close()
			if err := cfg.Load(); err != nil {
				t.Fatal(err)
			}
			addr, err := cfg.Value("server.rest.addr").String()
			if err != nil {
				t.Fatal(err)
			}
			if mode == "local" {
				if calls != 0 || addr != ":19990" {
					t.Fatalf("local connected remotely or returned wrong address: calls=%d addr=%s", calls, addr)
				}
			} else {
				if calls != 1 || addr != ":19991" {
					t.Fatalf("unexpected remote result: calls=%d addr=%s", calls, addr)
				}
				if _, err := cfg.Value("local_only").Bool(); err == nil {
					t.Fatal("local business config leaked into nacos mode")
				}
			}
		})
	}
	if _, err := NewConfigProvider(dir, "nacso"); err == nil {
		t.Fatal("invalid source accepted")
	}
	if _, err := NewConfigProvider(t.TempDir(), "nacos"); err == nil {
		t.Fatal("missing bootstrap accepted")
	}
	write("config.yaml", "config:\n  type: consul\n")
	if _, err := NewConfigProvider(dir, "nacos"); err == nil {
		t.Fatal("wrong provider accepted")
	}
}
