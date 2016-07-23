/*
 * datumboxmachinelearningapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 07/23/2016
 */

package metrics_pkg

import "datumboxmachinelearningapi_lib/models_pkg"

/*
 * Interface for the METRICS_IMPL
 */
type METRICS interface {
    CreateDocumentSimilarity (string, string, string) ([]*models_pkg.ObjectForDocument, error)
}

/*
 * Factory for the METRICS interaface returning METRICS_IMPL
 */
func NewMETRICS() METRICS {
    return &METRICS_IMPL{}
}