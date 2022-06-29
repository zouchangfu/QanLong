package request

type Task struct {
	Id        string `json:"id" `
	Name      string `json:"name"`
	Target    string `json:"target"`
	PortRange string `json:"portRange"`
	Ports     string `json:"ports"`
	Status    string `json:"status" `
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
