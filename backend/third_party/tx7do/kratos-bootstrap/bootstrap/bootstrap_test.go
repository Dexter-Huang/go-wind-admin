package bootstrap

import (
	"context"
	"errors"
	"testing"

	"github.com/go-kratos/kratos/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

func initApp(ctx *Context) (*kratos.App, func(), error) {
	app := NewApp(ctx)
	return app, func() {
	}, nil
}

func TestBootstrapWithNameVersion(t *testing.T) {
	serviceName := "test"
	version := "v0.0.1"

	ctx := NewContext(context.Background(), &conf.AppInfo{
		Project: "",
		AppId:   serviceName,
		Version: version,
	})

	// 独立配置目录，验证参数传递后在启动服务器前结束。 / Isolate config and stop before serving.
	stop := errors.New("test bootstrap completed")
	err := RunApp(ctx, func(ctx *Context) (*kratos.App, func(), error) {
		assert.Equal(t, serviceName, ctx.appInfo.AppId)
		return nil, nil, stop
	}, func(cmd *cobra.Command) {
		cmd.SetArgs([]string{"--config-source=local", "--conf", t.TempDir()})
	})
	assert.ErrorIs(t, err, stop)
}

func TestNewInstanceId(t *testing.T) {
	instanceId := NewInstanceId("gowind-test-service", "1.0.0", "127.0.0.1", "8000")
	t.Logf("InstanceId: %s", instanceId)
}
