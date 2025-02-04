package scanners

import "github.com/FourCoreLabs/EDRHunt/pkg/resources"

type CrowdstrikeDetection struct{}

func (w *CrowdstrikeDetection) Name() string {
	return "Crowdstrike EDR Solution"
}

func (w *CrowdstrikeDetection) Type() resources.EDRType {
	return resources.SymantecEDR
}

var CrowdstrikeHeuristic = []string{
	"CrowdStrike",
	"%SYSTEMROOT%\\system32\\drivers\\crowdstrike\\CsDeviceControl.inf",
	"%SYSTEMROOT%\\system32\\drivers\\crowdstrike\\CsFirmwareAnalysis.inf",
}

func (w *CrowdstrikeDetection) Detect(data resources.SystemData) (resources.EDRType, bool) {
	_, ok := data.CountMatchesAll(CrowdstrikeHeuristic)
	if !ok {
		return "", false
	}

	return resources.CrowdstrikeEDR, true
}
