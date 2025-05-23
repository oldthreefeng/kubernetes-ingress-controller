<!-- This document is generated by KIC's 'generate.docs' make target, DO NOT EDIT -->

## Packages
- [incubator.ingress-controller.konghq.com/v1alpha1](#incubatoringress-controllerkonghqcomv1alpha1)


## incubator.ingress-controller.konghq.com/v1alpha1

Package v1alpha1 contains API Schema definitions for the incubator.ingress-controller.konghq.com v1alpha1 API group.

- [KongServiceFacade](#kongservicefacade)
### KongServiceFacade


KongServiceFacade allows creating separate Kong Services for a single Kubernetes
Service. It can be used as Kubernetes Ingress' backend (via its path's `backend.resource`
field). It's designed to enable creating two "virtual" Services in Kong that will point
to the same Kubernetes Service, but will have different configuration (e.g. different
set of plugins, different load balancing algorithm, etc.).<br /><br />
KongServiceFacade requires `kubernetes.io/ingress.class` annotation with a value
matching the ingressClass of the Kong Ingress Controller (`kong` by default) to be reconciled.

<!-- kong_service_facade description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `incubator.ingress-controller.konghq.com/v1alpha1`
| `kind` _string_ | `KongServiceFacade`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KongServiceFacadeSpec](#kongservicefacadespec)_ |  |



### Types

In this section you will find types that the CRDs rely on.
#### KongServiceFacadeBackend


KongServiceFacadeBackend is a reference to a Kubernetes Service
that is used as a backend for a Kong Service Facade.



| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the referenced Kubernetes Service. |
| `port` _integer_ | Port is the port of the referenced Kubernetes Service. |


_Appears in:_
- [KongServiceFacadeSpec](#kongservicefacadespec)

#### KongServiceFacadeSpec


KongServiceFacadeSpec defines the desired state of KongServiceFacade.



| Field | Description |
| --- | --- |
| `backendRef` _[KongServiceFacadeBackend](#kongservicefacadebackend)_ | Backend is a reference to a Kubernetes Service that is used as a backend for this Kong Service Facade. |


_Appears in:_
- [KongServiceFacade](#kongservicefacade)



