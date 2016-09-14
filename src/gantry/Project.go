package gantry

import (
	"os"
	"regexp"
	"strings"
	"path"
	"strconv"
	"fmt"
)

func getProjectName() string {
	project := os.Getenv("GANTRY_PROJECT");
	if len(project) > 0 {
		return project
	}
	project = os.Getenv("COMPOSE_PROJECT_NAME");
	if len(project) > 0 {
		return project
	}
	project, err := os.Getwd();
	if err == nil {
		re := regexp.MustCompile("[^a-z]")
		return re.ReplaceAllString(strings.ToLower(path.Base(project)), "");
	}
	return ""
}
func getApplicationEnv() string {
	return os.Getenv("APP_ENV");
}
func getHTTPPort() int {
	port, _ := strconv.ParseInt(os.Getenv("DOCKER_HTTP_PORT"), 10, 32);
	return (int)(port)
}
func getProjectBasePath() string {
	return os.Getenv("PWD");
}

func getComposeFilesArgs() []string {
	var composeArgs []string
	base := getProjectBasePath()
	compose := path.Join(base, "docker-compose.yml")
	if _, err := os.Stat(compose); err == nil {
		composeArgs = append(composeArgs, "-f", compose)
	}
	env := getApplicationEnv()
	compose = path.Join(base, fmt.Sprintf("docker-compose-%s.yml", env))
	if _, err := os.Stat(compose); err == nil {
		composeArgs = append(composeArgs, "-f", compose)

	}
	return composeArgs
}