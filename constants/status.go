package constants

type ProjectStatus int

const (
	ProjectInactive        ProjectStatus = 0
	ProjectCancelled       ProjectStatus = 1
	ProjectMilestoneFailed ProjectStatus = 2
	ProjectEnded           ProjectStatus = 3
	ProjectError           ProjectStatus = 4
	ProjectDeployed        ProjectStatus = 5
	ProjectMilestonePhase  ProjectStatus = 6
	ProjectModerationPhase ProjectStatus = 7
	ProjectReadyToCancel   ProjectStatus = 8
)
