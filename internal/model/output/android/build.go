// Package android 定义 Android 构建相关的输出模型
package android

import "time"

// BuildStartOut 开始构建输出
type BuildStartOut struct {
	BuildId   int    `json:"buildId"`
	BuildName string `json:"buildName"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

// BuildDetailOut 构建详情输出
type BuildDetailOut struct {
	Id          int       `json:"id"`
	BuildName   string    `json:"buildName"`
	BuildType   string    `json:"buildType"`
	Version     string    `json:"version"`
	Device      string    `json:"device"`
	Variant     string    `json:"variant"`
	Status      string    `json:"status"`
	Progress    int       `json:"progress"`
	LogPath     string    `json:"logPath"`
	OutputPath  string    `json:"outputPath"`
	ErrorMsg    string    `json:"errorMsg"`
	StartedAt   time.Time `json:"startedAt"`
	FinishedAt  time.Time `json:"finishedAt"`
	Duration    int64     `json:"duration"` // 秒
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// BuildListOut 构建列表输出
type BuildListOut struct {
	List  []BuildDetailOut `json:"list"`
	Total int              `json:"total"`
	Page  int              `json:"page"`
	PageSize int           `json:"pageSize"`
}

// ApkAnalysisOut APK 分析输出
type ApkAnalysisOut struct {
	Id              int                    `json:"id"`
	ApkName         string                 `json:"apkName"`
	PackageName     string                 `json:"packageName"`
	VersionName     string                 `json:"versionName"`
	VersionCode     int                    `json:"versionCode"`
	MinSdkVersion   int                    `json:"minSdkVersion"`
	TargetSdkVersion int                   `json:"targetSdkVersion"`
	Permissions     []string               `json:"permissions"`
	Activities      []string               `json:"activities"`
	Services        []string               `json:"services"`
	Receivers       []string               `json:"receivers"`
	FileSize        int64                  `json:"fileSize"`
	FileMd5         string                 `json:"fileMd5"`
	DecompiledPath  string                 `json:"decompiledPath"`
	ManifestData    map[string]interface{} `json:"manifestData"`
	CreatedAt       time.Time              `json:"createdAt"`
}

// DeviceInfoOut 设备信息输出
type DeviceInfoOut struct {
	Id              int       `json:"id"`
	SerialNumber    string    `json:"serialNumber"`
	Model           string    `json:"model"`
	Manufacturer    string    `json:"manufacturer"`
	AndroidVersion  string    `json:"androidVersion"`
	ApiLevel        int       `json:"apiLevel"`
	BuildId         string    `json:"buildId"`
	Status          string    `json:"status"`
	AdbMode         string    `json:"adbMode"`
	IpAddress       string    `json:"ipAddress"`
	Port            int       `json:"port"`
	LastSeen        time.Time `json:"lastSeen"`
	BatteryLevel    int       `json:"batteryLevel,omitempty"`
	ScreenSize      string    `json:"screenSize,omitempty"`
	RootStatus      bool      `json:"rootStatus,omitempty"`
}

// DeviceListOut 设备列表输出
type DeviceListOut struct {
	List  []DeviceInfoOut `json:"list"`
	Total int             `json:"total"`
}

// DeviceCommandOut 设备命令输出
type DeviceCommandOut struct {
	Success    bool   `json:"success"`
	Output     string `json:"output"`
	ErrorMsg   string `json:"errorMsg"`
	ExitCode   int    `json:"exitCode"`
	Duration   int64  `json:"duration"` // 毫秒
}

// TaskDetailOut 任务详情输出
type TaskDetailOut struct {
	Id          int                    `json:"id"`
	TaskName    string                 `json:"taskName"`
	TaskType    string                 `json:"taskType"`
	Priority    int                    `json:"priority"`
	Status      string                 `json:"status"`
	Config      map[string]interface{} `json:"config"`
	Result      map[string]interface{} `json:"result"`
	ErrorMsg    string                 `json:"errorMsg"`
	StartedAt   time.Time              `json:"startedAt"`
	FinishedAt  time.Time              `json:"finishedAt"`
	Duration    int64                  `json:"duration"` // 秒
	CreatedAt   time.Time              `json:"createdAt"`
}

// TaskListOut 任务列表输出
type TaskListOut struct {
	List     []TaskDetailOut `json:"list"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

// BuildLogOut 构建日志输出
type BuildLogOut struct {
	BuildId   int      `json:"buildId"`
	LogLines  []string `json:"logLines"`
	TotalLines int     `json:"totalLines"`
	HasMore   bool     `json:"hasMore"`
}

// SystemInfoOut 系统信息输出
type SystemInfoOut struct {
	JavaVersion    string            `json:"javaVersion"`
	PythonVersion  string            `json:"pythonVersion"`
	AdbVersion     string            `json:"adbVersion"`
	ToolsInstalled map[string]bool   `json:"toolsInstalled"`
	DiskSpace      map[string]int64  `json:"diskSpace"` // 字节
	SystemLoad     map[string]float64 `json:"systemLoad"`
}
