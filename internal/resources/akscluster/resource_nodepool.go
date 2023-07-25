/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/go-test/deep"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	aksnodepool "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/akscluster/nodepool"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

var immutableFields = map[string]struct{}{
	"FullName.AksClusterName":    {},
	"FullName.CredentialName":    {},
	"FullName.Name":              {},
	"FullName.ResourceGroupName": {},
	"FullName.SubscriptionID":    {},
	"Spec.Mode":                  {},
	"Spec.VMSize":                {},
	"Spec.OsDiskType":            {},
	"Spec.OsDiskSizeGb":          {},
	"Spec.MaxPods":               {},
	"Spec.VnetSubnetID":          {},
	"Spec.VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig.Enabled": {},
	"Spec.ScaleSetPriority": {},
}

// nodePoolOperations the reconciliation data that will be used to apply nodepool changes.
type nodePoolOperations struct {
	existing []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool
	desired  []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool
}

// createNodepools sends the create request for the given nodepools as part of cluster creation flow.
func createNodepools(ctx context.Context, nodepools []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client aksnodepool.ClientService) error {
	var systemPoolsCreated int

	for _, np := range nodepools {
		if err := createNodepool(ctx, np, client); err == nil && *np.Spec.Mode == aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolModeSYSTEM {
			systemPoolsCreated += 1
		}
	}

	if systemPoolsCreated < 1 {
		return errors.New("no system nodepools were successfully created.")
	}

	return nil
}

// handleNodepoolChanges if nodepool changes are detected delegates to the appropriate node pool operation: `Create`, `Update`, `Delete`.
func handleNodepoolChanges(ctx context.Context, existing []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, data *schema.ResourceData, tc *client.TanzuMissionControl) error {
	npData := nodePoolOperations{
		existing: existing,
		desired:  ConstructNodepools(data),
	}

	return applyUpdates(ctx, npData, tc, getTimeOut(data))
}

func applyUpdates(ctx context.Context, npData nodePoolOperations, tc *client.TanzuMissionControl, timeout time.Duration) error {
	for _, np := range npData.desired {
		// Ignore any nodepools that already exist in the desired state.
		if existingNp := checkIfNodepoolExists(np, npData.existing); existingNp != nil && identical(np, existingNp) {
			// No action required the nodepool exists in the desired config.
			continue
		}

		// Create any nodepools that do not exist.
		if existingNp := checkIfNodepoolExists(np, npData.existing); existingNp == nil && np.FullName.Name != "" {
			if err := addNodepool(ctx, np, tc, timeout); err != nil {
				return err
			}
		}

		// If the node pool already exists but a requested change is immutable delete and recreate the nodepool.
		if existingNp := checkIfNodepoolExists(np, npData.existing); existingNp != nil && hasImmutableChange(np, existingNp) {
			if err := deleteAndRecreateNodepool(ctx, existingNp, tc, timeout, np); err != nil {
				return err
			}
		}

		// Update the nodepool in place if the change is mutable.
		if existingNp := checkIfNodepoolExists(np, npData.existing); existingNp != nil && !hasImmutableChange(np, existingNp) {
			if err := updateNodepool(ctx, np, tc, timeout); err != nil {
				return err
			}
		}

		// delete any existing nodepools that do not appear in the desired state.
		for _, existing := range npData.existing {
			if desired := checkIfNodepoolExists(existing, npData.desired); desired == nil {
				if err := deleteNodepool(ctx, existing, tc, timeout); err != nil {
					return err
				}
			}
		}
	}

	// delete any existing nodepools that do not appear in the desired state.
	for _, existing := range npData.existing {
		if desired := checkIfNodepoolExists(existing, npData.desired); desired == nil {
			if err := deleteNodepool(ctx, existing, tc, timeout); err != nil {
				return err
			}
		}
	}

	return nil
}

// createNodepool creates a nodepool, does not wait for nodepool to become ready. Use for cluster creation.
func createNodepool(_ context.Context, np *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client aksnodepool.ClientService) error {
	req := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest{Nodepool: np}
	if _, err := client.AksNodePoolResourceServiceCreate(req); err != nil {
		return err
	}

	return nil
}

// addNodepool adds a nodepool to an existing cluster and waits for the nodepool to be ready as part of an inplace
// update flow.
func addNodepool(ctx context.Context, np *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client *client.TanzuMissionControl, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest{Nodepool: np}
	if _, err := client.AKSNodePoolResourceService.AksNodePoolResourceServiceCreate(req); err != nil {
		return err
	}

	return pollUntilNodepoolReady(ctx, np.FullName, client.AKSNodePoolResourceService, getPollInterval(ctx))
}

// updateNodepool updates the configuration of an existing nodepool and waits for the nodepool to be ready as part of an
// inplace update flow.
func updateNodepool(ctx context.Context, np *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client *client.TanzuMissionControl, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest{Nodepool: np}
	if _, err := client.AKSNodePoolResourceService.AksNodePoolResourceServiceUpdate(req); err != nil {
		return err
	}

	return pollUntilNodepoolReady(ctx, np.FullName, client.AKSNodePoolResourceService, getPollInterval(ctx))
}

// deleteNodepool deletes an existing nodepool and waits until the nodepool has been successfully removed as part of an
// inplace update flow.
func deleteNodepool(ctx context.Context, nodepool *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client *client.TanzuMissionControl, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := client.AKSNodePoolResourceService.AksNodePoolResourceServiceDelete(nodepool.FullName); err != nil && !clienterrors.IsNotFoundError(err) {
		return err
	}

	return pollUntilNodepoolDeleted(ctx, nodepool.FullName, client.AKSNodePoolResourceService, getPollInterval(ctx))
}

// deleteAndRecreateNodepool delete and recreate a nodepool in the event of an immutable change.
func deleteAndRecreateNodepool(ctx context.Context, existingNp *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, tc *client.TanzuMissionControl, timeout time.Duration, np *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) error {
	if delErr := deleteNodepool(ctx, existingNp, tc, timeout); delErr != nil {
		return delErr
	}

	if createErr := addNodepool(ctx, np, tc, timeout); createErr != nil {
		return createErr
	}

	return nil
}

// pollUntilNodepoolReady calls get nodepool endpoint based on the provided interval and returns an error the resources
// is not found, has error conditions, or did not become ready before the expected timeout.
func pollUntilNodepoolReady(ctx context.Context, npFn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName, client aksnodepool.ClientService, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.New("timed out waiting for delete")
		case <-ticker.C:
			npResp, err := client.AksNodePoolResourceServiceGet(npFn)
			if clienterrors.IsNotFoundError(err) {
				return errors.Errorf("Unable to get Tanzu Mission Control AKS cluster entry, name : %s", npResp.Nodepool.FullName.Name)
			}

			if nodepoolHasFatalError(npResp) {
				return errors.Errorf("Nodepool creation failed: %s", getErrorReason(npResp.Nodepool.Status.Conditions))
			}

			if nodepoolIsReady(npResp) {
				return nil
			}
		}
	}
}

// pollUntilNodepoolDeleted calls get nodepool endpoint based on the provided interval until a `404` is received.
// A `404` signals the nodepool has been successfully deleted returns an error if operation timesout.
func pollUntilNodepoolDeleted(ctx context.Context, npFn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName, client aksnodepool.ClientService, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.New("timed out waiting for delete")
		case <-ticker.C:
			_, err := client.AksNodePoolResourceServiceGet(npFn)
			if clienterrors.IsNotFoundError(err) {
				return nil
			}
		}
	}
}

func checkIfNodepoolExists(new *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, existing []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool {
	for _, old := range existing {
		if new.FullName.Name == old.FullName.Name {
			return old
		}
	}

	return nil
}

func identical(new *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, old *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) bool {
	return reflect.DeepEqual(new.Spec, old.Spec)
}

func hasImmutableChange(new *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, old *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) bool {
	changes := deep.Equal(new, old)
	for _, c := range changes {
		key := strings.Split(c, ":")[0]

		if _, found := immutableFields[key]; found {
			return true
		}
	}

	return false
}
