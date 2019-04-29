/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lbcfcontroller

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	"git.tencent.com/tke/lb-controlling-framework/pkg/apis/lbcf.tke.cloud.tencent.com/v1beta1"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type LBDriverType string

const (
	LBDriverWebhook LBDriverType = "Webhook"
)

const (
	DefaultWebhookTimeout = 10 * time.Second
	SystemDriverPrefix    = "lbcf-"
)

func ValidateLoadBalancerDriver(raw *v1beta1.LoadBalancerDriver) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, validateDriverName(raw.Name, raw.Namespace, field.NewPath("metadata").Child("name"))...)
	allErrs = append(allErrs, validateDriverType(raw.Spec.DriverType, field.NewPath("spec").Child("driverType"))...)
	//allErrs = append(allErrs, validateDriverUrl(raw.Spec.Url, field.NewPath("spec").Child("url"))...)
	if raw.Spec.Webhooks != nil {
		allErrs = append(allErrs, validateDriverWebhooks(raw.Spec.Webhooks, field.NewPath("spec"))...)
	}
	return allErrs
}

func validateDriverName(name string, namespace string, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if namespace == "kube-system" {
		if !strings.HasPrefix(name, SystemDriverPrefix) {
			allErrs = append(allErrs, field.Invalid(path, name, fmt.Sprintf("metadata.name must start with %q for drivers in namespace %q", SystemDriverPrefix, "kube-system")))
		}
		return allErrs
	}
	if strings.HasPrefix(name, SystemDriverPrefix) {
		allErrs = append(allErrs, field.Invalid(path, name, fmt.Sprintf("metaname.name must not start with %q for drivers not in namespace %q", SystemDriverPrefix, "kube-system")))
	}
	return allErrs
}

func validateDriverType(raw string, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if raw != string(LBDriverWebhook) {
		allErrs = append(allErrs, field.Invalid(path, raw, fmt.Sprintf("driverType must be %v", LBDriverWebhook)))

	}
	return allErrs
}

func validateDriverWebhooks(raw []v1beta1.WebhookConfig, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	webhookSet := sets.NewString()
	for i, wh := range raw {
		curPath := path.Child(fmt.Sprintf("webhooks[%d]", i))
		if !knownWebhooks.Has(wh.Name) {
			allErrs = append(allErrs, field.NotSupported(curPath.Child("name"), wh, knownWebhooks.List()))
			continue
		}
		if webhookSet.Has(wh.Name) {
			allErrs = append(allErrs, field.Duplicate(curPath.Child("name"), wh.Name))
			continue
		}
		webhookSet.Insert(wh.Name)
		if wh.Timeout != nil {
			if wh.Timeout.Nanoseconds() > (1 * time.Minute).Nanoseconds() {
				allErrs = append(allErrs, field.Invalid(curPath.Child("timeout"), *wh.Timeout, fmt.Sprintf("webhook %s invalid, timeout of must be less than or equal to 1m", wh.Name)))
				continue
			}
		}
	}
	return allErrs
}

func DriverUpdatedFieldsAllowed(cur *v1beta1.LoadBalancerDriver, old *v1beta1.LoadBalancerDriver) bool {
	if cur.Name != old.Name {
		return false
	}
	if old.Spec.Url != cur.Spec.Url {
		return false
	}
	if old.Spec.DriverType != cur.Spec.DriverType {
		return false
	}
	return true
}

func LBUpdatedFieldsAllowed(cur *v1beta1.LoadBalancer, old *v1beta1.LoadBalancer) bool {
	if cur.Name != old.Name {
		return false
	}
	if cur.Spec.LBDriver != old.Spec.LBDriver {
		return false
	}
	if !reflect.DeepEqual(cur.Spec.LBSpec, old.Spec.LBSpec) {
		return false
	}
	return true
}

func ValidateBackendGroup(raw *v1beta1.BackendGroup) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validateBackends(&raw.Spec, field.NewPath("spec"))...)
	return allErrs
}

func validateBackends(raw *v1beta1.BackendGroupSpec, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if raw.Service != nil {
		if raw.Pods != nil {
			allErrs = append(allErrs, field.Invalid(path.Child("pods"), raw.Pods, "only one of \"service, pods, static\" is allowed"))
		} else if raw.Static != nil {
			allErrs = append(allErrs, field.Invalid(path.Child("static"), raw.Pods, "only one of \"service, pods, static\" is allowed"))
		} else {
			allErrs = append(allErrs, validateServiceBackend(raw.Service, path.Child("service"))...)
		}
		return allErrs
	}

	if raw.Pods != nil {
		if raw.Static != nil {
			allErrs = append(allErrs, field.Invalid(path.Child("static"), raw.Pods, "only one of \"service, pods, static\" is allowed"))
		} else {
			allErrs = append(allErrs, validatePodBackend(raw.Pods, path.Child("pods"))...)
		}
		return allErrs
	}

	if raw.Static == nil {
		allErrs = append(allErrs, field.Required(path.Child("service/pods/static"), "one of \"service, pods, static\" must be specified"))
	}
	return allErrs
}

func validateServiceBackend(raw *v1beta1.ServiceBackend, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validatePortSelector(raw.Port, path.Child("port"))...)
	return allErrs
}

func validatePodBackend(raw *v1beta1.PodBackend, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validatePortSelector(raw.Port, path.Child("port"))...)
	if raw.ByLabel != nil {
		if raw.ByName != nil {
			allErrs = append(allErrs, field.Invalid(path.Child("byName"), raw.ByName, "only one of \"byLabel, byName\" is allowed"))
		}
	}

	if raw.ByName == nil {
		allErrs = append(allErrs, field.Required(path.Child("byLabel/byName"), "one of \"byLabel, byName\" must be specified"))
	}
	return allErrs
}

func validatePortSelector(raw v1beta1.PortSelector, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if raw.PortNumber <= 0 || raw.PortNumber > 65535 {
		allErrs = append(allErrs, field.Invalid(path.Child("portNumber"), raw.PortNumber, "portNumber must be greater than 0 and less than 65536"))
	}

	if raw.Protocol != nil {
		p := strings.ToUpper(*raw.Protocol)
		if p != string(v1.ProtocolTCP) && p != string(v1.ProtocolUDP) {
			allErrs = append(allErrs, field.Invalid(path.Child("protocol"), raw.Protocol, "portNumber must be \"TCP\" or \"UDP\""))
		}
	}
	return allErrs
}

func BackendGroupUpdateFieldsAllowed(cur *v1beta1.BackendGroup, old *v1beta1.BackendGroup) bool {
	if getBackendType(cur) != getBackendType(old) {
		return false
	}
	return true
}