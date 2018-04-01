package registry

import (
	"testing"
)

func TestGetDigestDockerHub(t *testing.T) {
	client := New("https://index.docker.io", "", "")

	tags, err := client.Tags("karolisr/keel")
	if err != nil {
		t.Errorf("failed to get tags, error: %s", err)
	}

	if len(tags) == 0 {
		t.Errorf("no tags?")
	}
}

// Supply authentication details to test this
func TestGetDigestArtifactory(t *testing.T) {
	t.Skip()

	client := New("https://keel-docker-local.jfrog.io", "", "")

	tags, err := client.Tags("webhook-demo")
	if err != nil {
		t.Fatalf("failed to get tags, error: %s", err)
	}

	if len(tags) == 0 {
		t.Errorf("no tags?")
	}
}
