/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package manifest

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
)

type manifest struct {
	namespacedName types.NamespacedName
	gvk            *runtimeSchema.GroupVersionKind
	usObj          map[string]interface{}
}

const ( // yamlSeparator is separator for multi-YAML resource files
	yamlSeparator = "\n---\n"
	interval      = 5 * time.Second
	retries       = 3
)

func getManifests(manifestsBlob string) (manifests []manifest, err error) {
	// Add v1 api-extensions to the set of default schemes to support CRDs.
	err = apiextensionsv1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	err = apiextensionsv1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	deserializer := serializer.NewCodecFactory(scheme.Scheme).UniversalDeserializer()

	manifestsData := strings.Split(strings.TrimSpace(manifestsBlob), yamlSeparator)
	for _, manifestData := range manifestsData {
		if invalidManifestData := strings.TrimSpace(manifestData); invalidManifestData == "---" || invalidManifestData == "" {
			continue
		}

		manifest := manifest{}

		obj, gvk, err := deserializer.Decode([]byte(manifestData), nil, nil)
		if err != nil {
			return nil, err
		}

		usObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj.(metav1.Object))
		if err != nil {
			return nil, err
		}

		manifest.gvk = gvk
		manifest.usObj = usObj

		// NestedString will always return the value if present or a empty string
		name, ok, err := unstructured.NestedString(usObj, "metadata", "name")
		if err != nil {
			return nil, fmt.Errorf("provided value for name in the metadata of kind %v is not of type string, error :%v", gvk, err)
		}

		if !ok {
			return nil, fmt.Errorf("provided value for name in the metadata of kind %v is empty", gvk)
		}

		namespace, ok, err := unstructured.NestedString(usObj, "metadata", "namespace")

		if err != nil {
			return nil, fmt.Errorf("provided value for namespace in the metadata of kind %v is not of type string, error :%v", gvk, err)
		}

		if !ok {
			// for cluster scoped k8s objects namespace will be empty
			namespace = ""
		}

		manifest.namespacedName = types.NamespacedName{
			Namespace: namespace,
			Name:      name,
		}

		manifests = append(manifests, manifest)
	}

	return manifests, nil
}

func objectsToBeCleaned(k8sclient k8sClient.Client, manifests []manifest, clean bool) (tobeCleaned []string, err error) {
	for _, manifest := range manifests {
		unstruct := &unstructured.Unstructured{}
		unstruct.SetGroupVersionKind(*manifest.gvk)

		err := k8sclient.Get(context.Background(), manifest.namespacedName, unstruct)
		if err == nil {
			if clean {
				err := ensureObjectDeleted(k8sclient, unstruct)
				if err != nil {
					return tobeCleaned, fmt.Errorf("failed to delete object %v of type %v, error:%v", manifest.namespacedName, manifest.gvk, err)
				}
			} else {
				tobeCleaned = append(tobeCleaned, fmt.Sprintf("object %v of type %v is already present", manifest.namespacedName, manifest.gvk))
			}
		}
	}

	return
}

func createObjects(k8sclient k8sClient.Client, manifests []manifest) error {
	for _, manifest := range manifests {
		err := k8sclient.Create(context.Background(), &unstructured.Unstructured{Object: manifest.usObj})
		if err != nil {
			return fmt.Errorf("error creating object with namespaced:%+v and gvk:%+v, error :%v", manifest.namespacedName, manifest.gvk, err)
		}
	}

	return nil
}

func ensureObjectDeleted(k8sclient k8sClient.Client, object *unstructured.Unstructured) (err error) {
	deleteFn := func() (bool, error) {
		err = k8sclient.Delete(context.Background(), object)
		if k8serrors.IsNotFound(err) || err == nil {
			return false, nil
		}

		if k8serrors.IsServerTimeout(err) {
			return false, err
		}

		return true, err
	}

	if _, err := helper.Retry(deleteFn, interval, retries); err != nil {
		return err
	}

	return nil
}

func GetK8sManifest(depLink string) ([]byte, error) {
	client := http.Client{Timeout: 60 * time.Second}

	resp, err := client.Get(depLink)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.TrimLeft(respData, "\n\t "), nil
}
