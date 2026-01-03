// Package entity 定义 Android 构建相关的数据实体
package entity

import "time"

// AndroidBuild Android 构建记录
type AndroidBuild struct {
	Id          int       `json:"id"          orm:"id,primary"    description:"构建ID"`
	BuildName   string    `json:"buildName"   orm:"build_name"    description:"构建名称"`
	BuildType   string    `json:"buildType"   orm:"build_type"    description:"构建类型(aosp/rom/kernel)"`
	Version     string    `json:"version"     orm:"version"       description:"Android版本"`
	Device      string    `json:"device"      orm:"device"        description:"目标设备"`
	Variant     string    `json:"variant"     orm:"variant"       description:"构建变体(user/userdebug/eng)"`
	Status      string    `json:"status"      orm:"status"        description:"构建状态(pending/building/success/failed)"`
	Progress    int       `json:"progress"    orm:"progress"      description:"构建进度(0-100)"`
	LogPath     string    `json:"logPath"     orm:"log_path"      description:"构建日志路径"`
	OutputPath  string    `json:"outputPath"  orm:"output_path"   description:"输出文件路径"`
	ErrorMsg    string    `json:"errorMsg"    orm:"error_msg"     description:"错误信息"`
	StartedAt   time.Time `json:"startedAt"   orm:"started_at"    description:"开始时间"`
	FinishedAt  time.Time `json:"finishedAt"  orm:"finished_at"   description:"完成时间"`
	CreatedAt   time.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`
	UpdatedAt   time.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`
	DeletedAt   *time.Time `json:"deletedAt"   orm:"deleted_at"    description:"删除时间"`
}

// ApkAnalysis APK 分析记录
type ApkAnalysis struct {
	Id              int       `json:"id"              orm:"id,primary"        description:"分析ID"`
	ApkName         string    `json:"apkName"         orm:"apk_name"          description:"APK名称"`
	PackageName     string    `json:"packageName"     orm:"package_name"      description:"包名"`
	VersionName     string    `json:"versionName"     orm:"version_name"      description:"版本名称"`
	VersionCode     int       `json:"versionCode"     orm:"version_code"      description:"版本号"`
	MinSdkVersion   int       `json:"minSdkVersion"   orm:"min_sdk_version"   description:"最小SDK版本"`
	TargetSdkVersion int      `json:"targetSdkVersion" orm:"target_sdk_version" description:"目标SDK版本"`
	Permissions     string    `json:"permissions"     orm:"permissions"       description:"权限列表(JSON)"`
	Activities      string    `json:"activities"      orm:"activities"        description:"Activity列表(JSON)"`
	Services        string    `json:"services"        orm:"services"          description:"Service列表(JSON)"`
	Receivers       string    `json:"receivers"       orm:"receivers"         description:"Receiver列表(JSON)"`
	FilePath        string    `json:"filePath"        orm:"file_path"         description:"APK文件路径"`
	FileSize        int64     `json:"fileSize"        orm:"file_size"         description:"文件大小(字节)"`
	FileMd5         string    `json:"fileMd5"         orm:"file_md5"          description:"文件MD5"`
	DecompiledPath  string    `json:"decompiledPath"  orm:"decompiled_path"   description:"反编译输出路径"`
	Status          string    `json:"status"          orm:"status"            description:"分析状态(pending/analyzing/success/failed)"`
	ErrorMsg        string    `json:"errorMsg"        orm:"error_msg"         description:"错误信息"`
	CreatedAt       time.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`
	UpdatedAt       time.Time `json:"updatedAt"       orm:"updated_at"        description:"更新时间"`
	DeletedAt       *time.Time `json:"deletedAt"       orm:"deleted_at"        description:"删除时间"`
}

// DeviceInfo 设备信息
type DeviceInfo struct {
	Id           int       `json:"id"           orm:"id,primary"     description:"设备ID"`
	SerialNumber string    `json:"serialNumber" orm:"serial_number"  description:"设备序列号"`
	Model        string    `json:"model"        orm:"model"          description:"设备型号"`
	Manufacturer string    `json:"manufacturer" orm:"manufacturer"   description:"制造商"`
	AndroidVersion string  `json:"androidVersion" orm:"android_version" description:"Android版本"`
	ApiLevel     int       `json:"apiLevel"     orm:"api_level"      description:"API级别"`
	BuildId      string    `json:"buildId"      orm:"build_id"       description:"构建ID"`
	Status       string    `json:"status"       orm:"status"         description:"连接状态(online/offline/unauthorized)"`
	AdbMode      string    `json:"adbMode"      orm:"adb_mode"       description:"ADB模式(usb/tcpip)"`
	IpAddress    string    `json:"ipAddress"    orm:"ip_address"     description:"IP地址"`
	Port         int       `json:"port"         orm:"port"           description:"ADB端口"`
	LastSeen     time.Time `json:"lastSeen"     orm:"last_seen"      description:"最后在线时间"`
	CreatedAt    time.Time `json:"createdAt"    orm:"created_at"     description:"创建时间"`
	UpdatedAt    time.Time `json:"updatedAt"    orm:"updated_at"     description:"更新时间"`
	DeletedAt    *time.Time `json:"deletedAt"    orm:"deleted_at"     description:"删除时间"`
}

// BuildTask 构建任务
type BuildTask struct {
	Id          int       `json:"id"          orm:"id,primary"    description:"任务ID"`
	TaskName    string    `json:"taskName"    orm:"task_name"     description:"任务名称"`
	TaskType    string    `json:"taskType"    orm:"task_type"     description:"任务类型(build/flash/analyze)"`
	Priority    int       `json:"priority"    orm:"priority"      description:"优先级(1-10)"`
	Status      string    `json:"status"      orm:"status"        description:"任务状态(queued/running/completed/failed/cancelled)"`
	Config      string    `json:"config"      orm:"config"        description:"任务配置(JSON)"`
	Result      string    `json:"result"      orm:"result"        description:"任务结果(JSON)"`
	ErrorMsg    string    `json:"errorMsg"    orm:"error_msg"     description:"错误信息"`
	StartedAt   time.Time `json:"startedAt"   orm:"started_at"    description:"开始时间"`
	FinishedAt  time.Time `json:"finishedAt"  orm:"finished_at"   description:"完成时间"`
	CreatedAt   time.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`
	UpdatedAt   time.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`
	DeletedAt   *time.Time `json:"deletedAt"   orm:"deleted_at"    description:"删除时间"`
}
