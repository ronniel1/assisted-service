// Code generated by from events definition; DO NOT EDIT.
package events

import (
    "fmt"
    "strings"

    "github.com/go-openapi/strfmt"
)

//
// Event cancel_install_failed_start
//
type CancelInstallFailedStartEvent struct {
    ClusterId strfmt.UUID
}

var CancelInstallFailedStartEvent_Id string = "CIE00001"

func NewCancelInstallFailedStartEvent(
    clusterId strfmt.UUID,
) *CancelInstallFailedStartEvent {
    return &CancelInstallFailedStartEvent {
        ClusterId: clusterId,
    }
}
    
func (e *CancelInstallFailedStartEvent) GetId() string {
    return CancelInstallFailedStartEvent_Id
}

func (e *CancelInstallFailedStartEvent) GetSeverity() string {
    return "error"
}

func (e *CancelInstallFailedStartEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

func (e *CancelInstallFailedStartEvent) format(message *string) string {
    r := strings.NewReplacer(
    "{cluster_id}", fmt.Sprint(e.ClusterId),
                  )
          return r.Replace(*message)
}

func (e *CancelInstallFailedStartEvent) FormatMessage() string {
    s := "Failed to cancel installation: error starting DB transaction"
    return e.format(&s)
}

func (e *CancelInstallFailedStartEvent) FormatLogMessage() string {

    return e.FormatMessage()

}

//
// Event cancel_install_failed_commit
//
type CancelInstallFailedCommitEvent struct {
    ClusterId strfmt.UUID
}

var CancelInstallFailedCommitEvent_Id string = "CIE00002"

func NewCancelInstallFailedCommitEvent(
    clusterId strfmt.UUID,
) *CancelInstallFailedCommitEvent {
    return &CancelInstallFailedCommitEvent {
        ClusterId: clusterId,
    }
}
    
func (e *CancelInstallFailedCommitEvent) GetId() string {
    return CancelInstallFailedCommitEvent_Id
}

func (e *CancelInstallFailedCommitEvent) GetSeverity() string {
    return "error"
}

func (e *CancelInstallFailedCommitEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

func (e *CancelInstallFailedCommitEvent) format(message *string) string {
    r := strings.NewReplacer(
    "{cluster_id}", fmt.Sprint(e.ClusterId),
                  )
          return r.Replace(*message)
}

func (e *CancelInstallFailedCommitEvent) FormatMessage() string {
    s := "Failed to cancel installation: error committing DB transaction"
    return e.format(&s)
}

func (e *CancelInstallFailedCommitEvent) FormatLogMessage() string {

    return e.FormatMessage()

}

//
// Event iso_gen_failed_format_ign
//
type IsoGenFailedFormatIgnEvent struct {
    ClusterId strfmt.UUID
}

var IsoGenFailedFormatIgnEvent_Id string = "GIF00003"

func NewIsoGenFailedFormatIgnEvent(
    clusterId strfmt.UUID,
) *IsoGenFailedFormatIgnEvent {
    return &IsoGenFailedFormatIgnEvent {
        ClusterId: clusterId,
    }
}
    
func (e *IsoGenFailedFormatIgnEvent) GetId() string {
    return IsoGenFailedFormatIgnEvent_Id
}

func (e *IsoGenFailedFormatIgnEvent) GetSeverity() string {
    return "error"
}

func (e *IsoGenFailedFormatIgnEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

func (e *IsoGenFailedFormatIgnEvent) format(message *string) string {
    r := strings.NewReplacer(
    "{cluster_id}", fmt.Sprint(e.ClusterId),
                  )
          return r.Replace(*message)
}

func (e *IsoGenFailedFormatIgnEvent) FormatMessage() string {
    s := "Failed to generate image: error formatting ignition file"
    return e.format(&s)
}

func (e *IsoGenFailedFormatIgnEvent) FormatLogMessage() string {

    return e.FormatMessage()

}

//
// Event iso_gen_failed
//
type IsoGenFailedEvent struct {
    ClusterId strfmt.UUID
}

var IsoGenFailedEvent_Id string = "GIF00004"

func NewIsoGenFailedEvent(
    clusterId strfmt.UUID,
) *IsoGenFailedEvent {
    return &IsoGenFailedEvent {
        ClusterId: clusterId,
    }
}
    
func (e *IsoGenFailedEvent) GetId() string {
    return IsoGenFailedEvent_Id
}

func (e *IsoGenFailedEvent) GetSeverity() string {
    return "error"
}

func (e *IsoGenFailedEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

func (e *IsoGenFailedEvent) format(message *string) string {
    r := strings.NewReplacer(
    "{cluster_id}", fmt.Sprint(e.ClusterId),
                  )
          return r.Replace(*message)
}

func (e *IsoGenFailedEvent) FormatMessage() string {
    s := "Failed to generate minimal ISO"
    return e.format(&s)
}

func (e *IsoGenFailedEvent) FormatLogMessage() string {

    s := "Failed to generate minimal ISO for cluster {cluster_id}"
    return e.format(&s)

}

