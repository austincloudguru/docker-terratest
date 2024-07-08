package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerfile(t *testing.T) {
    tag := "austincloud/terratest"
    // This still needs to be the actual version of terraform
    tfVersion := "1.9.1"
    goVersion := "1.22"

    buildOptions := &docker.BuildOptions{
        Tags: []string{tag},
    }

    docker.Build(t, "../", buildOptions)

    goOpts := &docker.RunOptions{Command: []string{"go", "version"}}
    goOutput := docker.Run(t, tag, goOpts)
    assert.Contains(t, goOutput, goVersion)

    tfOpts := &docker.RunOptions{Command: []string{"terraform", "--version"}}
    tfOutput := docker.Run(t, tag, tfOpts)
    assert.Contains(t, tfOutput, tfVersion)
}
