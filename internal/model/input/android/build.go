// Package android 定义 Android 构建相关的输入模型
package android

// BuildStartInp 开始构建输入
type BuildStartInp struct {
	BuildName   string            `json:"buildName"   v:"required#请输入构建名称"`
	BuildType   string            `json:"buildType"   v:"required|in:aosp,rom,kernel#请选择构建类型|构建类型错误"`
	Version     string            `json:"version"     v:"required#请输入Android版本"`
	Device      string            `json:"device"      v:"required#请选择目标设备"`
	Variant     string            `json:"variant"     v:"required|in:user,userdebug,eng#请选择构建变体|构建变体错误"`
	CleanBuild  bool              `json:"cleanBuild"`
	SyncSource  bool              `json:"syncSource"`
	CustomConfig map[string]string `json:"customConfig"`
}

// BuildListInp 构建列表查询输入
type BuildListInp struct {
	Page      int    `json:"page"      v:"required|min:1#请输入页码|页码最小为1"`
	PageSize  int    `json:"pageSize"  v:"required|between:1,100#请输入每页数量|每页数量在1-100之间"`
	BuildType string `json:"buildType" v:"in:aosp,rom,kernel#构建类型错误"`
	Status    string `json:"status"    v:"in:pending,building,success,failed#状态错误"`
	Device    string `json:"device"`
}

// BuildCancelInp 取消构建输入
type BuildCancelInp struct {
	BuildId int `json:"buildId" v:"required|min:1#请输入构建ID|构建ID无效"`
}

// ApkAnalyzeInp APK 分析输入
type ApkAnalyzeInp struct {
	ApkPath          string `json:"apkPath"          v:"required#请输入APK路径"`
	DecompileSource  bool   `json:"decompileSource"`
	ExtractResources bool   `json:"extractResources"`
	ExtractAssets    bool   `json:"extractAssets"`
	SaveResult       bool   `json:"saveResult"`
}

// ApkUploadInp APK 上传输入
type ApkUploadInp struct {
	FileName string `json:"fileName" v:"required#请输入文件名"`
	FileSize int64  `json:"fileSize" v:"required|min:1#请输入文件大小|文件大小无效"`
}

// DeviceListInp 设备列表查询输入
type DeviceListInp struct {
	Status       string `json:"status" v:"in:online,offline,unauthorized#状态错误"`
	Manufacturer string `json:"manufacturer"`
}

// DeviceConnectInp 设备连接输入
type DeviceConnectInp struct {
	IpAddress string `json:"ipAddress" v:"required|ip#请输入IP地址|IP地址格式错误"`
	Port      int    `json:"port"      v:"required|between:1,65535#请输入端口|端口范围1-65535"`
}

// DeviceCommandInp 设备命令输入
type DeviceCommandInp struct {
	SerialNumber string   `json:"serialNumber" v:"required#请输入设备序列号"`
	Command      string   `json:"command"      v:"required#请输入命令"`
	Args         []string `json:"args"`
	Timeout      int      `json:"timeout"      v:"between:1,300#超时时间在1-300秒之间"`
}

// DeviceFlashInp 设备刷机输入
type DeviceFlashInp struct {
	SerialNumber string   `json:"serialNumber" v:"required#请输入设备序列号"`
	ImagePath    string   `json:"imagePath"    v:"required#请输入镜像路径"`
	Partitions   []string `json:"partitions"   v:"required#请选择分区"`
	RebootAfter  bool     `json:"rebootAfter"`
}

// TaskCreateInp 创建任务输入
type TaskCreateInp struct {
	TaskName string                 `json:"taskName" v:"required#请输入任务名称"`
	TaskType string                 `json:"taskType" v:"required|in:build,flash,analyze#请选择任务类型|任务类型错误"`
	Priority int                    `json:"priority" v:"between:1,10#优先级在1-10之间"`
	Config   map[string]interface{} `json:"config"   v:"required#请输入任务配置"`
}

// TaskListInp 任务列表查询输入
type TaskListInp struct {
	Page     int    `json:"page"     v:"required|min:1#请输入页码|页码最小为1"`
	PageSize int    `json:"pageSize" v:"required|between:1,100#请输入每页数量|每页数量在1-100之间"`
	TaskType string `json:"taskType" v:"in:build,flash,analyze#任务类型错误"`
	Status   string `json:"status"   v:"in:queued,running,completed,failed,cancelled#状态错误"`
}
