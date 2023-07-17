/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	aksnodepool "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/akscluster/nodepool"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

// createNodepools sends the create request for the given nodepools as part of cluster creation flow.
func createNodepools(ctx context.Context, nodepools []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client aksnodepool.ClientService) error {
	for _, np := range nodepools {
		if err := createNodepool(ctx, np, client); err != nil {
			return err
		}
	}

	return nil
}

// handleNodepoolChanges if nodepool changes are detected delegates to the appropriate node pool operation: `Create`, `Update`, `Delete`.
func handleNodepoolChanges(ctx context.Context, existing []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, data *schema.ResourceData, tc *client.TanzuMissionControl) error {
	var created, updated, deleted []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool

	cfn := extractClusterFullName(data)
	npCount := changedNodeCount(data)

	for i := 0; i <= npCount; i++ {
		key := fmt.Sprintf("spec.0.nodepool.%v", i)
		if data.HasChange(key) {
			oldResource, newResource := data.GetChange(key)

			// The existence of an old and new state signals and update operation.
			if !isEmpty(oldResource) && !isEmpty(newResource) {
				np := constructNodepool(cfn, newResource.(map[string]any))
				updated = append(updated, np)

				continue
			}

			// If the old resource is empty this signals a new nodepool has been created.
			if isEmpty(oldResource) {
				np := constructNodepool(cfn, newResource.(map[string]any))
				created = append(created, np)

				continue
			}

			// If the new version is empty this signals a node pool has been deleted.
			if isEmpty(newResource) {
				np := constructNodepool(cfn, oldResource.(map[string]any))
				deleted = append(deleted, np)

				continue
			}
		}
	}

	npData := map[string][]*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		"created":  created,
		"updated":  updated,
		"deleted":  deleted,
		"existing": existing,
		"desired":  ConstructNodepools(data),
	}

	return applyUpdates(ctx, npData, tc, getTimeOut(data))
}

func applyUpdates(ctx context.Context, npData map[string][]*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, tc *client.TanzuMissionControl, timeout time.Duration) error {
	for _, np := range npData["created"] {
		// If the node pool already exists it may have been reordered.
		if old := checkIfNodepoolExists(np, npData["existing"]); old != nil {
			if !identical(np, old) {
				// If the reordered nodepool doesn't match the existing one it is actually an update.
				npData["updated"] = append(npData["updated"], np)
			} else {
				// No action required nodepool already exists in desired state.
				continue
			}
		}

		if err := addNodepool(ctx, np, tc, timeout); err != nil {
			return err
		}
	}

	for _, np := range npData["updated"] {
		// If the node pool already exists it may have been reordered.
		if existingNp := checkIfNodepoolExists(np, npData["existing"]); existingNp != nil && identical(np, existingNp) {
			// No action required the nodepool exists in the request config.
			continue
		}

		if err := updateNodepool(ctx, np, tc, timeout); err != nil {
			return err
		}
	}

	for _, np := range npData["deleted"] {
		if desired := checkIfNodepoolExists(np, npData["desired"]); desired != nil {
			// The nodepool exists in the desired state we should not delete it.
			continue
		}

		if err := deleteNodepool(ctx, np, tc, timeout); err != nil {
			return err
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
// inplace update flow
func deleteNodepool(ctx context.Context, nodepool *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, client *client.TanzuMissionControl, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := client.AKSNodePoolResourceService.AksNodePoolResourceServiceDelete(nodepool.FullName); err != nil && !clienterrors.IsNotFoundError(err) {
		return err
	}

	return pollUntilNodepoolDeleted(ctx, nodepool.FullName, client.AKSNodePoolResourceService, getPollInterval(ctx))
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

func changedNodeCount(data *schema.ResourceData) int {
	o, n := data.GetChange("spec.0.nodepool.#")
	npCount := max(o.(int), n.(int))

	return npCount
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
