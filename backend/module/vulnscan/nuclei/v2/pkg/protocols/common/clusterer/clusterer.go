package clusterer

import (
	templates2 "backend/module/vulnscan/nuclei/v2/pkg/templates"
)

// Cluster clusters a list of templates into a lesser number if possible based
// on the similarity between the sent requests.
//
// If the attributes match, multiple requests can be clustered into a single
// request which saves time and network resources during execution.
func Cluster(list map[string]*templates2.Template) [][]*templates2.Template {
	final := [][]*templates2.Template{}

	// Each protocol that can be clustered should be handled here.
	for key, template := range list {
		// We only cluster http requests as of now.
		// Take care of requests that can't be clustered first.
		if len(template.RequestsHTTP) == 0 {
			delete(list, key)
			final = append(final, []*templates2.Template{template})
			continue
		}

		delete(list, key) // delete element first so it's not found later.
		// Find any/all similar matching request that is identical to
		// this one and cluster them together for http protocol only.
		if len(template.RequestsHTTP) == 1 {
			cluster := []*templates2.Template{}

			for otherKey, other := range list {
				if len(other.RequestsHTTP) == 0 {
					continue
				}
				if template.RequestsHTTP[0].CanCluster(other.RequestsHTTP[0]) {
					delete(list, otherKey)
					cluster = append(cluster, other)
				}
			}
			if len(cluster) > 0 {
				cluster = append(cluster, template)
				final = append(final, cluster)
				continue
			}
		}
		final = append(final, []*templates2.Template{template})
	}
	return final
}
