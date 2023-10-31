package alcsyncv1

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// PackageName is the name of the package.
const PackageName = "alcsync-v1"

type alcsyncpkg struct {
	server   io.Server
	registry packages.Registry
}

// HandleUp implements packages.ApplicationPackageHandler.
func (a *alcsyncpkg) HandleUp(
	ctx context.Context,
	def *ttnpb.ApplicationPackageDefaultAssociation,
	assoc *ttnpb.ApplicationPackageAssociation,
	up *ttnpb.ApplicationUp,
) (err error) {
	ctx = log.NewContextWithField(ctx, "namespace", "applicationserver/io/packages/alcsync/v1")
	logger := log.FromContext(ctx)

	if def == nil && assoc == nil {
		logger.Error("No association available")
		return errNoAssociation.New()
	}

	logger.Debug("Handle uplink")

	msg := up.GetUplinkMessage()
	if msg == nil {
		logger.Debug("Uplink is not an uplink message")
		return nil
	}

	eventBuilders := make(events.Builders, 0)

	defer func(ids *ttnpb.EndDeviceIdentifiers) {
		if err != nil {
			eventBuilders = append(eventBuilders, EvtPkgFail.With(
				events.WithIdentifiers(ids),
				events.WithData(err),
			))
		}
		publishEvents(ctx, eventBuilders...)
	}(up.GetEndDeviceIds())

	data, fPort, err := mergePackageData(def, assoc)
	if err != nil {
		logger.WithError(err).Debug("Failed to merge package data")
		return err
	}

	if msg.GetFPort() != fPort {
		logger.WithFields(log.Fields(
			"expected_fport", fPort,
			"received_fport", msg.GetFPort(),
		)).Debug("Uplink received on unhadled FPort")
		return nil
	}

	if len(msg.GetFrmPayload()) == 0 {
		logger.Debug("Uplink message has no payload")
		return nil
	}

	commands, evts, err := MakeCommands(msg, fPort, data)
	eventBuilders = append(eventBuilders, evts...)
	if err != nil {
		logger.WithError(err).Debug("Failed to parse frame payload into commands")
		return err
	}

	results := make([]Result, 0, len(commands))
	for _, cmd := range commands {
		result, err := cmd.Execute()
		if errors.IsUnavailable(err) {
			continue
		}
		if err != nil {
			logger.WithError(err).WithField("command_id", cmd.Code()).Debug("Failed to execute command")
			eventBuilders = append(eventBuilders, EvtPkgFail.With(
				events.WithIdentifiers(up.GetEndDeviceIds()),
				events.WithData(err)),
			)
			continue
		}
		if result != nil {
			results = append(results, result)
			eventBuilders = append(eventBuilders, result.AnswerEnqueuedEventBuilder())
		}
	}
	downlink, err := MakeDownlink(results, fPort)
	if err != nil {
		logger.WithError(err).Debug("Failed to create downlink from results")
		return err
	}
	if err := a.server.DownlinkQueuePush(ctx, up.EndDeviceIds, []*ttnpb.ApplicationDownlink{downlink}); err != nil {
		logger.WithError(err).Debug("Failed to push downlinks to queue")
		return err
	}
	return nil
}

// Package implements packages.ApplicationPackageHandler.
func (*alcsyncpkg) Package() *ttnpb.ApplicationPackage {
	return &ttnpb.ApplicationPackage{
		Name:         PackageName,
		DefaultFPort: 202,
	}
}

// New returns a new ALCSync package.
func New(server io.Server, registry packages.Registry) packages.ApplicationPackageHandler {
	return &alcsyncpkg{
		server:   server,
		registry: registry,
	}
}
