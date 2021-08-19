package test

import (
        "testing"
        "github.com/stretchr/testify/assert"
        "github.com/gruntwork-io/terratest/modules/docker"
)

func TestDockerfile(t *testing.T) {
    tag := "austincloud/terratest"
    tfVersion := "1.0.4"
    goVersion := "1.16"

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
