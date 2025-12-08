package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"korrectkm/app"
	"korrectkm/checkdbg"
	"korrectkm/config"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/guiconnect"
	"korrectkm/reductor"
	"korrectkm/repo"
	"korrectkm/spaserver"
	"korrectkm/trueclient"
	"korrectkm/zaplog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"korrectkm/dbscan"

	"github.com/labstack/echo/v4"
	"github.com/mechiko/utility"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"go.uber.org/zap"

	"golang.org/x/sync/errgroup"
)

const modError = "main"

// var version = "0.0.0"
var fileExe string
var dir string

// если local true то папка создается локально
var local = flag.Bool("local", false, "")
var dontconnect = flag.Bool("dontconnect", false, "")

func errMessageExit(loger *zap.SugaredLogger, title string, err error) {
	if loger != nil {
		loger.Errorf("%s %v", title, err)
	}
	utility.MessageBox(title, err.Error())
	os.Exit(1)
}

func main() {
	flag.Parse()
	fileExe = os.Args[0]
	var err error
	dir, err = filepath.Abs(filepath.Dir(fileExe))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, groupCtx := errgroup.WithContext(ctx)

	cfg, err := config.New("", !*local)
	if err != nil {
		errMessageExit(nil, "ошибка конфигурации", err)
	}

	debug := strings.ToLower(config.Mode) == "development"
	var logsOutConfig = map[string]zaplog.LogConfig{
		"logger": {
			ErrorOutputPaths: []string{"stdout", filepath.Join(cfg.LogPath(), config.Name)},
			Debug:            debug,
			Console:          true,
			Name:             filepath.Join(cfg.LogPath(), config.Name),
		},
		"echo": {
			ErrorOutputPaths: []string{filepath.Join(cfg.LogPath(), "echo")},
			Debug:            debug,
			Console:          false,
			Name:             filepath.Join(cfg.LogPath(), "echo"),
		},
		"reductor": {
			ErrorOutputPaths: []string{filepath.Join(cfg.LogPath(), "reductor")},
			Debug:            debug,
			Console:          true,
			Name:             filepath.Join(cfg.LogPath(), "reductor"),
		},
	}

	zl, err := zaplog.New(logsOutConfig)
	if err != nil {
		errMessageExit(nil, "ошибка создания логера", err)
	}

	lg, err := zl.GetLogger("logger")
	if err != nil {
		errMessageExit(nil, "ошибка получения логера", err)
	}
	loger := lg.Sugar()
	loger.Debug("zaplog started")
	loger.Infof("mode = %s", config.Mode)
	if cfg.Warning() != "" {
		loger.Infof("pkg:config warning %s", cfg.Warning())
	}

	errProcessExit := func(title string, err error) {
		utility.MessageBox(title, err.Error())
		loger.Errorf("errProcessExit: %s %v", title, err)
		cancel()
		err = group.Wait()
		fmt.Printf("game over! error %v\n", err)
		zl.Shutdown()
		os.Exit(1)
	}

	reductorLogger, err := zl.GetLogger("reductor")
	if err != nil {
		errProcessExit("Ошибка получения логера для редуктора", err)
	}

	if err := reductor.New(reductorLogger.Sugar()); err != nil {
		errProcessExit("Ошибка создания редуктора", err)
	}

	loger.Info("new webapp")
	// создаем приложение с опциями из конфига и логером основным
	app := app.New(cfg, loger, dir)
	app.SetDbSelfPath(cfg.ConfigPath())
	// бд основные находятся в текущем каталоге если не переопределено в настройках
	app.SetDefaultDbPath("")

	// инициализируем пути необходимые приложению
	app.CreatePath()

	loger.Info("start repo")
	// инициализируем REPO

	listDbs := make(dbscan.ListDbInfoForScan)
	listDbs[dbscan.Config] = &dbscan.DbInfo{}
	listDbs[dbscan.Other] = &dbscan.DbInfo{
		File:   "korrectkm.db",
		Name:   "korrectkm",
		Driver: "sqlite",
		Path:   app.DbSelfPath(),
	}
	listDbs[dbscan.TrueZnak] = &dbscan.DbInfo{}

	err = repo.New(listDbs, ".")
	if err != nil {
		loger.Warnf("Ошибка запуска репозитория %v", err)
		// errProcessExit("Ошибка запуска репозитория", err)
	}
	repoStart, err := repo.GetRepository()
	if err != nil {
		errProcessExit("Ошибка получения репозитория", err)
	}

	// создаем редуктор с новой моделью
	modelTcl, err := modeltrueclient.New(app)
	if err != nil {
		errProcessExit("Ошибка создания модели TrueClientModel", err)
	}
	// загружаем сертификаты пользователя
	err = modelTcl.LoadStore(app)
	if err != nil {
		errProcessExit("Ошибка загрузки сертификатов пользователя", err)
	}

	err = reductor.SetModel(modelTcl, false)
	if err != nil {
		errProcessExit("Ошибка записи модели в редуктор", err)
	}

	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			repoStart.Shutdown()
		}()
		return repoStart.Run(groupCtx)
	})
	// тесты
	if err := checkdbg.NewChecks(app).Run(); err != nil {
		errProcessExit("Ошибка checkdbg.NewChecks(app).Run()", err)
	}

	loger.Info("start up webapp")

	port := cfg.Configuration().HostPort
	if port == "" || port == "auto" {
		if portFree, err := utility.GetFreePort(); err == nil {
			port = fmt.Sprintf("%d", portFree)
			// порт не записываем в файл конфигурации остается только в модели приложения
			app.SetOptions("hostport", port)
		}
	}
	loger.Infof("http port %s", port)

	echoLogger, err := zl.GetLogger("echo")
	if err != nil {
		errProcessExit("Ошибка получения логера для http", err)
	}

	// вызываем окно подключения к ЧЗ
	if !*dontconnect {
		_, err := trueclient.NewFromModelSingle(app, modelTcl)
		if err != nil {
			// если ошибка подключения вызываем редактор
			err = guiconnect.StartDialog(app, modelTcl)
			if err != nil {
				errProcessExit("Ошибка подключения к ЧЗ", err)
			}
			err = reductor.SetModel(modelTcl, false)
			if err != nil {
				errProcessExit("Ошибка сохранения модели в приложение", err)
			}
		}
	}

	// тут инициализируются так же модели для всех видов
	httpServer := spaserver.New(app, echoLogger, port, true)
	httpServer.SetActivePage(domain.KMState)
	// запускаем сервер эхо через него SSE работает для флэш сообщений
	httpServer.Start()

	appw := application.New(application.Options{
		Name: "KorrectKM",
		Assets: application.AssetOptions{
			Handler:    httpServer.Echo(),
			Middleware: EchoMiddleware(httpServer.Echo()),
		},
		Logger:  nil,
		Windows: application.WindowsOptions{},
	})
	app.SetWails(appw)
	wv := appw.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Инструменты",
		Width:  1024,
		Height: 768,
	})
	if debug {
		appw.Event.OnApplicationEvent(events.Common.ApplicationStarted, func(event *application.ApplicationEvent) {
			wv.OpenDevTools()
		})
	}
	err = appw.Run()
	if err != nil {
		panic(err)
	}
	cancel()
	// ожидание завершения всех в группе
	if err := group.Wait(); err != nil {
		fmt.Printf("game over! error %s\n", err.Error())
	} else {
		fmt.Println("game over!")
	}
	zl.Shutdown()
}

// creates a middleware that passes requests to Echo if they're not handled by Wails
func EchoMiddleware(echoEngine *echo.Echo) application.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Let Wails handle the `/wails` route
			if strings.HasPrefix(r.URL.Path, "/wails") {
				next.ServeHTTP(w, r)
				return
			}
			// Let Echo handle everything else
			echoEngine.ServeHTTP(w, r)
		})
	}
}
