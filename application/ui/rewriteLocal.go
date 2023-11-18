package ui

import (
	"errors"
	"log/slog"
	"os"
	"runtime"

	cp "github.com/otiai10/copy"
)

type RewriteLocalCommand struct {
	FrontendProject string
	basePath        string
}

func NewRewriteLocalCommand(frontendProject string) (*RewriteLocalCommand, error) {
	preSetup()

	command := &RewriteLocalCommand{
		FrontendProject: frontendProject,
		basePath:        os.Getenv("PAYTRACK_CLI_SOURCE_PATH"),
	}

	err := command.validateRequirements()
	if err != nil {
		return nil, err
	}
	return command, nil
}

func preSetup() {
	basePath := os.Getenv("PAYTRACK_CLI_SOURCE_PATH")
	if basePath != "" {
		return
	}

	slog.Info("variável de ambiente PAYTRACK_CLI_SOURCE_PATH não definida")
	slog.Info("definindo variável de ambiente PAYTRACK_CLI_SOURCE_PATH")
	system := runtime.GOOS
	switch system {
	case "windows":
		os.Setenv("PAYTRACK_CLI_SOURCE_PATH", "C:/git/paytrack/fontes/")
	case "darwin":
		homeDir, _ := os.UserHomeDir()
		os.Setenv("PAYTRACK_CLI_SOURCE_PATH", homeDir+"/git/paytrack/fontes/")
	}
}

func (r RewriteLocalCommand) validateRequirements() error {
	slog.Info("validando projeto frontend informado")
	err := r.validateFrontendProject()
	if err != nil {
		return err
	}

	slog.Info("validando projeto paytrack-ui-library")
	err = r.validateUILibraryProject()
	if err != nil {
		return err
	}

	return nil
}

func (r RewriteLocalCommand) validateFrontendProject() error {
	if r.FrontendProject == "" {
		err := errors.New("o parâmetro --frontend-project (-f) é obrigatório")
		slog.Error(err.Error())
		return err
	}

	projectFolder, err := os.Stat(r.basePath + r.FrontendProject)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	if !projectFolder.IsDir() {
		err := errors.New("o projeto frontend informado não existe")
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (r RewriteLocalCommand) validateUILibraryProject() error {
	projectFolder, err := os.Stat(r.basePath + "paytrack-ui-library")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	if !projectFolder.IsDir() {
		err = errors.New("o projeto paytrack-ui-library não existe")
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (r RewriteLocalCommand) Execute() error {
	slog.Info("removendo pasta dist da dependência da UI Library no projeto frontend local")
	err := os.RemoveAll(r.basePath + r.FrontendProject + "/node_modules/paytrack-ui-library/dist")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	srcPath := r.basePath + "paytrack-ui-library/dist"
	dstPath := r.basePath + r.FrontendProject + "/node_modules/paytrack-ui-library/dist"
	slog.Info("copiando pasta dist da UI Library para a dependência no projeto frontend local")
	err = cp.Copy(srcPath, dstPath)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
