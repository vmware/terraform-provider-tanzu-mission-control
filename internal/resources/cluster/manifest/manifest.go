/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package manifest

import (
	"fmt"

	"github.com/pkg/errors"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	// importing this to avoid `panic: No Auth Provider found for name "gcp"`
	// ref: https://github.com/kubernetes/client-go/issues/345
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func Create(
	k8sclient *k8sClient.Client,
	k8sManifest string,
	forceClean bool,
) error {
	if k8sclient == nil {
		return errors.New("kubernetes client cannot be empty")
	}

	manifests, err := getManifests(k8sManifest)
	if err != nil {
		return errors.WithMessage(err, "failure to fetch attach manifests")
	}

	toBeCleaned, err := objectsToBeCleaned(k8sclient, manifests, forceClean)
	if err != nil && forceClean {
		return errors.WithMessage(err, "error while cleaning up the resources")
	}

	if len(toBeCleaned) != 0 {
		fmt.Println("Provided kubeconfig cannot be used to attach, reason:")

		for _, cleanup := range toBeCleaned {
			fmt.Println(cleanup)
		}

		return errors.New("please clean up the above mentioned k8s objects or follow cluster detach steps and retry")
	}

	err = createObjects(k8sclient, manifests)
	if err != nil {
		return errors.WithMessage(err, "error while attaching the cluster")
	}

	fmt.Println("TMC resources applied to the cluster successfully")

	return nil
}
