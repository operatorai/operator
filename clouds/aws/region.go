package aws

import (
	"encoding/json"

	"github.com/operatorai/kettle-cli/command"
	"github.com/operatorai/kettle-cli/config"
)

func SetDeploymentRegion(settings *config.Settings) error {
	if settings.DeploymentRegion != "" {
		return nil
	}

	regions, err := getAWSRegions()
	if err != nil {
		return err
	}

	region, err := command.PromptForValue("Deployment region", regions, false)
	if err != nil {
		return err
	}

	settings.DeploymentRegion = region
	return nil
}

// aws ec2 describe-regions --output json
func getAWSRegions() (map[string]string, error) {
	output, err := command.ExecuteWithResult("aws", []string{
		"ec2",
		"describe-regions",
		"--output", "json",
	}, "Collecting ec2 regions")
	if err != nil {
		return nil, err
	}

	var result struct {
		Regions []struct {
			RegionName string `json:"RegionName"`
		} `json:"Regions"`
	}
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	regions := map[string]string{}
	for _, region := range result.Regions {
		regions[region.RegionName] = region.RegionName
	}
	return regions, nil
}
