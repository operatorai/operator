package gcloud

import (
	"fmt"

	"github.com/operatorai/operator/command"
	"github.com/operatorai/operator/config"
)

type GoogleCloudRun struct{}

func (GoogleCloudRun) Deploy(directory string, cfg *config.TemplateConfig) error {
	fmt.Println("🏭  Building: ", cfg.Name, "as a Cloud Run container")
	if err := SetProjectID(cfg.Settings); err != nil {
		return err
	}
	if err := SetDeploymentRegion(cfg.Settings); err != nil {
		return err
	}

	containerTag := fmt.Sprintf("gcr.io/%s/%s", cfg.Settings.ProjectName, cfg.Name)
	// Build the docker container
	// gcloud builds submit --tag gcr.io/PROJECT-ID/helloworld
	err := command.Execute("gcloud", []string{
		"builds",
		"submit",
		"--tag", containerTag,
	}, "Building docker container")
	if err != nil {
		return err
	}

	// Deploy the docker container
	// gcloud run deploy --image gcr.io/PROJECT-ID/helloworld
	fmt.Println("🚢  Deploying ", cfg.Name, fmt.Sprintf("as a %s function", cfg.Settings.DeploymentType))
	return command.Execute("gcloud", []string{
		"run",
		"deploy",
		cfg.Name,
		"--image", containerTag,
		"--platform", "managed",
		"--allow-unauthenticated",
		fmt.Sprintf("--region=%s", cfg.Settings.DeploymentRegion),
	}, "Deploying Cloud Run container")
}
