package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerfile(t *testing.T) {
    tag := "austincloud/terratest"
    tfVersion := "1.3.4"
    goVersion := "1.18"

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
