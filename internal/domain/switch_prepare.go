package domain

// SwitchPreparation 切换准备项
type SwitchPreparation struct {
	Version   SwitchVersion
	BizDomain string
	Name      string
	Status    string // 准备完成、未完成
	Sequence  int    // 顺序，int从大到小，按顺序执行
}

// BreakDown cat故障
type BreakDown struct {
	Version SwitchVersion
	Content string
	File    string // cat截图
}
