// Package apitest contains common API definition test utilities.
package apitest

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// RunTestFieldMaskSpecified runs a test that checks whether all RPCs have their allowed field mask paths set.
func RunTestFieldMaskSpecified(t *testing.T, pkg protoreflect.FullName, paths map[string]ttnpb.RPCFieldMaskPathValue) {
	t.Helper()
	a := assertions.New(t)
	protoregistry.GlobalFiles.RangeFilesByPackage(pkg, func(fd protoreflect.FileDescriptor) bool {
		t.Helper()
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			methods := services.Get(i).Methods()
			for j := 0; j < methods.Len(); j++ {
				method := methods.Get(j)
				fields := method.Input().Fields()
				for k := 0; k < fields.Len(); k++ {
					field := fields.Get(k)
					if field.Name() != "field_mask" {
						continue
					}
					message := field.Message()
					if message == nil || message.FullName() != "google.protobuf.FieldMask" {
						continue
					}
					a.So(
						paths,
						should.ContainKey,
						fmt.Sprintf("/%s/%s", method.FullName().Parent(), method.Name()),
					)
				}
			}
		}
		return true
	})
}
