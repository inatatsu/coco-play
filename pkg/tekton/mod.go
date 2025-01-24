/*
Copyright Confidential Containers Contributors
SPDX-License-Identifier: Apache-2.0
*/
package tekton

import (
	"fmt"

	"github.com/wainersm/coco-play/pkg/cluster"
)

func Install(version string) error {
	namespace := "tekton-pipelines"

	fmt.Println("Installing Tekton Pipelines...")
	out, err := cluster.Kubectl("apply", "-f", "https://storage.googleapis.com/tekton-releases/pipeline/previous/"+version+"/release.yaml")
	if err != nil {
		return fmt.Errorf("Failed to install Tekton Pipelines: %v", err)
	}
	fmt.Println(out)

	out, err = cluster.Kubectl("rollout", "status", "-w", "deployment/tekton-pipelines-controller", "-n", namespace)
	if err != nil {
		return fmt.Errorf("Controller is not ready: %v", err)
	}
	fmt.Println(out)

	return nil
}
