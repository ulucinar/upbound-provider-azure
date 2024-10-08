// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this HPCCache.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *HPCCache) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetIDRef,
			Selector:     mg.Spec.ForProvider.SubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetIDRef,
			Selector:     mg.Spec.InitProvider.SubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HPCCacheAccessPolicy.
func (mg *HPCCacheAccessPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HPCCacheID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.HPCCacheIDRef,
			Selector:     mg.Spec.ForProvider.HPCCacheIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HPCCacheID")
	}
	mg.Spec.ForProvider.HPCCacheID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HPCCacheIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HPCCacheID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.HPCCacheIDRef,
			Selector:     mg.Spec.InitProvider.HPCCacheIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HPCCacheID")
	}
	mg.Spec.InitProvider.HPCCacheID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HPCCacheIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HPCCacheBlobNFSTarget.
func (mg *HPCCacheBlobNFSTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CacheName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CacheNameRef,
			Selector:     mg.Spec.ForProvider.CacheNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CacheName")
	}
	mg.Spec.ForProvider.CacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CacheNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CacheName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CacheNameRef,
			Selector:     mg.Spec.InitProvider.CacheNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CacheName")
	}
	mg.Spec.InitProvider.CacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CacheNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HPCCacheBlobTarget.
func (mg *HPCCacheBlobTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CacheName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CacheNameRef,
			Selector:     mg.Spec.ForProvider.CacheNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CacheName")
	}
	mg.Spec.ForProvider.CacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CacheNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageContainerID),
			Extract:      resource.ExtractParamPath("resource_manager_id", true),
			Reference:    mg.Spec.ForProvider.StorageContainerIDRef,
			Selector:     mg.Spec.ForProvider.StorageContainerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageContainerID")
	}
	mg.Spec.ForProvider.StorageContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageContainerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CacheName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CacheNameRef,
			Selector:     mg.Spec.InitProvider.CacheNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CacheName")
	}
	mg.Spec.InitProvider.CacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CacheNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageContainerID),
			Extract:      resource.ExtractParamPath("resource_manager_id", true),
			Reference:    mg.Spec.InitProvider.StorageContainerIDRef,
			Selector:     mg.Spec.InitProvider.StorageContainerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageContainerID")
	}
	mg.Spec.InitProvider.StorageContainerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageContainerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HPCCacheNFSTarget.
func (mg *HPCCacheNFSTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CacheName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CacheNameRef,
			Selector:     mg.Spec.ForProvider.CacheNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CacheName")
	}
	mg.Spec.ForProvider.CacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CacheNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storagecache.azure.upbound.io", "v1beta2", "HPCCache", "HPCCacheList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CacheName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CacheNameRef,
			Selector:     mg.Spec.InitProvider.CacheNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CacheName")
	}
	mg.Spec.InitProvider.CacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CacheNameRef = rsp.ResolvedReference

	return nil
}
