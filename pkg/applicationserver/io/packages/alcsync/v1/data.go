package alcsyncv1

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/structpb"
)

var defaultThreshold = time.Duration(4) * time.Second

type packageData struct {
	Threshold time.Duration
}

func (d *packageData) fromStruct(st *structpb.Struct) error {
	fields := st.GetFields()
	value, ok := fields["threshold"]
	if ok {
		numberValue, ok := value.GetKind().(*structpb.Value_NumberValue)
		if !ok {
			return errInvalidFieldType.WithAttributes(
				"field", "threshold",
				"type", "number",
			)
		}
		d.Threshold = time.Duration(numberValue.NumberValue) * time.Second
	}
	return nil
}

func mergePackageData(
	def *ttnpb.ApplicationPackageDefaultAssociation,
	assoc *ttnpb.ApplicationPackageAssociation,
) (*packageData, uint32, error) {
	var defaultData, associationData packageData
	if err := defaultData.fromStruct(def.GetData()); err != nil {
		return nil, 0, errPkgDataMerge.WithCause(err).New()
	}
	if err := associationData.fromStruct(assoc.GetData()); err != nil {
		return nil, 0, errPkgDataMerge.WithCause(err).New()
	}

	merged := &packageData{
		Threshold: defaultThreshold,
	}
	for _, data := range []packageData{defaultData, associationData} {
		if data.Threshold != 0 {
			merged.Threshold = data.Threshold
		}
	}
	fPort := def.GetIds().GetFPort()
	assocFPort := assoc.GetIds().GetFPort()
	if assocFPort != 0 {
		fPort = assocFPort
	}
	return merged, fPort, nil
}
