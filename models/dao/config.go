package dao

type Configuration struct {
	Cuid string
	Items []ConfigItem
}

type ConfigItem struct {
	Type string
	Trigger string
	TargetProtocol string
	TargetHost string
	TargetPort int
	TargetPath string
}