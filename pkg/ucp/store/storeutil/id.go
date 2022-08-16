// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package storeutil

import (
	"strings"

	"github.com/project-radius/radius/pkg/ucp/resources"
	"github.com/project-radius/radius/pkg/ucp/store"
)

const (
	ScopePrefix    = "scope"
	ResourcePrefix = "resource"
)

// ExtractStorageParts extracts the main components of the resource id in a way that easily
// supports our storage abstraction. Returns a tuple of the (prefix, rootScope, routingScope, resourceType)
func ExtractStorageParts(id resources.ID) (string, string, string, string) {
	if id.IsScope() {
		// For a scope we encode the last scope segment as the routing scope, and the previous
		// scope segments as the root scope. This gives us the most desirable behavior for
		// queries and recursion.

		prefix := ScopePrefix
		rootScope := NormalizePart(id.Truncate().RootScope())

		last := resources.ScopeSegment{}
		if len(id.ScopeSegments()) > 0 {
			last = id.ScopeSegments()[len(id.ScopeSegments())-1]
		}
		routingScope := NormalizePart(last.Type + resources.SegmentSeparator + last.Name)
		resourceType := strings.ToLower(last.Type)

		return prefix, rootScope, routingScope, resourceType
	} else {
		prefix := ResourcePrefix
		rootScope := NormalizePart(id.RootScope())
		routingScope := NormalizePart(id.RoutingScope())
		resourceType := strings.ToLower(id.Type())

		return prefix, rootScope, routingScope, resourceType
	}
}

func IDMatchesQuery(id resources.ID, query store.Query) bool {
	prefix, rootScope, routingScope, resourceType := ExtractStorageParts(id)
	if query.IsScopeQuery && !strings.EqualFold(prefix, ScopePrefix) {
		return false
	} else if !query.IsScopeQuery && !strings.EqualFold(prefix, ResourcePrefix) {
		return false
	}

	if query.ScopeRecursive && !strings.HasPrefix(rootScope, NormalizePart(query.RootScope)) {
		return false
	} else if !query.ScopeRecursive && !strings.EqualFold(rootScope, NormalizePart(query.RootScope)) {
		return false
	}

	if query.RoutingScopePrefix != "" && !strings.HasPrefix(routingScope, NormalizePart(query.RoutingScopePrefix)) {
		return false
	}

	if query.ResourceType != "" && !strings.EqualFold(resourceType, query.ResourceType) {
		return false
	}

	return true
}

func NormalizePart(part string) string {
	if len(part) == 0 {
		return ""
	}
	if !strings.HasPrefix(part, resources.SegmentSeparator) {
		part = resources.SegmentSeparator + part
	}
	if !strings.HasSuffix(part, resources.SegmentSeparator) {
		part = part + resources.SegmentSeparator
	}

	return strings.ToLower(part)
}