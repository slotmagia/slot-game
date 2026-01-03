# Android 设备刷机与调试完全指南

> **作者**: Android Development Team  
> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **适用版本**: Android 8.1 及其他版本

---

## 📋 目录

- [1. 刷机前准备](#1-刷机前准备)
- [2. Bootloader 解锁](#2-bootloader-解锁)
- [3. 刷入 Recovery](#3-刷入-recovery)
- [4. 刷机方法](#4-刷机方法)
- [5. ADB 调试](#5-adb-调试)
- [6. Logcat 日志分析](#6-logcat-日志分析)
- [7. 系统调试](#7-系统调试)
- [8. 问题排查](#8-问题排查)

---

## 1. 刷机前准备

### 1.1 数据备份

⚠️ **重要警告**: 刷机和解锁 Bootloader 会清除所有数据！

#### 备份用户数据

```bash
# 方法 1: 使用 ADB 备份
adb backup -apk -shared -all -f backup.ab

# 恢复备份
adb restore backup.ab

# 方法 2: 备份特定应用
adb backup -f app_backup.ab -apk com.example.app

# 方法 3: 备份照片和文件
adb pull /sdcard/ ~/phone_backup/
```

#### 备份系统分区

```bash
# 需要 root 权限
# 列出所有分区
adb shell su -c "ls -l /dev/block/platform/*/by-name"

# 备份 boot 分区
adb shell su -c "dd if=/dev/block/bootdevice/by-name/boot of=/sdcard/boot.img"
adb pull /sdcard/boot.img

# 备份 recovery 分区
adb shell su -c "dd if=/dev/block/bootdevice/by-name/recovery of=/sdcard/recovery.img"
adb pull /sdcard/recovery.img

# 备份 system 分区 (较大，需要时间)
adb shell su -c "dd if=/dev/block/bootdevice/by-name/system of=/sdcard/system.img"
adb pull /sdcard/system.img
```

### 1.2 检查设备信息

```bash
# 查看设备型号
adb shell getprop ro.product.model

# 查看 Android 版本
adb shell getprop ro.build.version.release

# 查看 Bootloader 版本
adb shell getprop ro.bootloader

# 查看设备代号
adb shell getprop ro.product.device

# 查看分区布局
adb shell cat /proc/partitions

# 查看存储空间
adb shell df -h
```

### 1.3 准备刷机文件

根据刷机方式准备相应文件：

| 刷机方式 | 需要的文件 |
|---------|-----------|
| **Fastboot 刷机** | boot.img, system.img, vendor.img, userdata.img 等 |
| **Recovery 刷机** | ROM.zip (OTA 包) |
| **线刷包** | flash-all.sh/bat + 镜像文件 |

### 1.4 驱动安装 (Windows)

```powershell
# 下载 Google USB Driver
# https://developer.android.com/studio/run/win-usb

# 或使用第三方工具
# 15 seconds ADB Installer
# Universal ADB Drivers
```

---

## 2. Bootloader 解锁

### 2.1 为什么要解锁 Bootloader

- 刷入第三方 ROM
- 刷入第三方 Recovery (TWRP)
- 获取 Root 权限
- 刷入自定义内核

⚠️ **警告**:
- 解锁会清除所有数据
- 可能失去保修
- 某些设备不支持解锁 (运营商定制机)

### 2.2 解锁步骤

#### 通用方法 (大部分设备)

```bash
# 1. 启用 OEM 解锁
# 设置 -> 关于手机 -> 连续点击版本号 7 次 (启用开发者选项)
# 设置 -> 开发者选项 -> 启用 OEM 解锁

# 2. 启用 USB 调试
# 设置 -> 开发者选项 -> 启用 USB 调试

# 3. 重启到 Bootloader
adb reboot bootloader

# 或手动进入:
# 关机 -> 同时按住 音量下键 + 电源键

# 4. 检查 Fastboot 连接
fastboot devices

# 5. 解锁 Bootloader
fastboot oem unlock

# 或 (新设备):
fastboot flashing unlock

# 6. 在设备上确认解锁 (使用音量键选择, 电源键确认)

# 7. 重启设备
fastboot reboot
```

#### 品牌特定方法

**小米设备**:
```bash
# 1. 绑定账号 (设置 -> 开发者选项 -> 设备解锁状态)
# 2. 等待 168 小时 (7天)
# 3. 使用官方解锁工具
#    下载: http://www.miui.com/unlock/index.html

# 运行解锁工具
# 登录小米账号 -> 连接设备 -> 点击解锁
```

**华为设备**:
```bash
# 华为已停止提供解锁码服务
# 可能需要第三方服务 (不推荐)
```

**一加设备**:
```bash
# 无需等待, 直接解锁
fastboot oem unlock
```

**三星设备**:
```bash
# 无需解锁 Bootloader
# 但需要刷入 TWRP 并禁用 Knox
```

### 2.3 验证解锁状态

```bash
# 进入 Bootloader
adb reboot bootloader

# 查看锁定状态
fastboot getvar unlocked

# 或查看 Bootloader 界面
# 通常会显示 "UNLOCKED" 或红色警告
```

### 2.4 重新锁定 Bootloader (如果需要)

```bash
# ⚠️ 警告: 锁定前确保系统完整, 否则会变砖!

# 进入 Bootloader
adb reboot bootloader

# 锁定
fastboot oem lock

# 或:
fastboot flashing lock
```

---

## 3. 刷入 Recovery

### 3.1 TWRP Recovery

TWRP (Team Win Recovery Project) 是最流行的第三方 Recovery。

#### 下载 TWRP

```bash
# 访问 TWRP 官网
# https://twrp.me/Devices/

# 根据设备型号下载对应的 img 文件
# 例如: twrp-3.7.0-device.img
```

#### 临时启动 TWRP (不刷入)

```bash
# 重启到 Bootloader
adb reboot bootloader

# 临时启动 TWRP
fastboot boot twrp-3.7.0-device.img

# 设备会启动到 TWRP, 但重启后恢复原 Recovery
```

#### 永久刷入 TWRP

```bash
# 方法 1: 使用 Fastboot 刷入
adb reboot bootloader
fastboot flash recovery twrp-3.7.0-device.img
fastboot reboot

# 方法 2: 从 TWRP 内部刷入 (防止被覆盖)
# 1. 临时启动 TWRP: fastboot boot twrp.img
# 2. 将 TWRP img 推送到设备
adb push twrp-3.7.0-device.img /sdcard/
# 3. 在 TWRP 中: Install -> Install Image -> 选择 twrp.img
# 4. 选择 Recovery 分区 -> 滑动确认

# 方法 3: 使用 TWRP APP (需要 root)
# 安装 TWRP Manager 应用
# 在应用中选择设备型号并刷入
```

#### 进入 TWRP

```bash
# 使用 ADB 重启
adb reboot recovery

# 或手动进入:
# 关机 -> 同时按住 音量上键 + 电源键
```

### 3.2 其他 Recovery

**OrangeFox Recovery**:
```bash
# 下载: https://orangefox.download/
# 刷入方法同 TWRP
fastboot flash recovery orangefox.img
```

**PBRP (PitchBlack Recovery)**:
```bash
# 下载: https://pitchblackrecovery.com/
fastboot flash recovery pbrp.img
```

### 3.3 Recovery 基本操作

**TWRP 菜单**:
- **Install**: 刷入 ZIP 包
- **Wipe**: 清除数据
  - Dalvik/ART Cache: 清除缓存
  - Cache: 清除应用缓存
  - Data: 清除用户数据 (恢复出厂设置)
  - System: 清除系统分区
- **Backup**: 备份系统
- **Restore**: 恢复备份
- **Mount**: 挂载分区
- **Advanced**: 高级功能
  - File Manager: 文件管理器
  - Terminal: 终端
  - ADB Sideload: 侧载刷机

---

## 4. 刷机方法

### 4.1 Fastboot 刷机

适用于官方镜像或提取的 img 文件。

#### 刷入完整镜像

```bash
# 1. 重启到 Bootloader
adb reboot bootloader

# 2. 检查连接
fastboot devices

# 3. 刷入各个分区
fastboot flash boot boot.img
fastboot flash system system.img
fastboot flash vendor vendor.img
fastboot flash recovery recovery.img
fastboot flash cache cache.img

# 4. 格式化 userdata (可选, 首次刷机推荐)
fastboot erase userdata
fastboot erase cache

# 5. 重启
fastboot reboot

# 一键刷入 (如果有 flash-all 脚本)
# Linux/macOS:
./flash-all.sh

# Windows:
flash-all.bat
```

#### 刷入单个分区

```bash
# 只刷入 boot 分区 (内核)
fastboot flash boot boot.img

# 只刷入 system 分区
fastboot flash system system.img

# 刷入后重启
fastboot reboot
```

#### 使用 fastboot update

```bash
# 刷入 ZIP 格式的镜像包
fastboot update image.zip

# 带清除数据
fastboot update -w image.zip
```

### 4.2 Recovery 刷机

适用于 OTA 包或第三方 ROM。

#### TWRP 刷机步骤

```bash
# 1. 将 ROM 包推送到设备
adb push rom.zip /sdcard/

# 2. 重启到 Recovery
adb reboot recovery

# 3. 在 TWRP 中操作:
# - 选择 "Wipe" -> "Advanced Wipe"
# - 勾选: Dalvik, Cache, System, Data (首次刷机)
# - 滑动确认清除

# 4. 返回主菜单, 选择 "Install"
# - 浏览到 /sdcard/rom.zip
# - 滑动确认刷入

# 5. 等待刷入完成
# - 选择 "Reboot System"

# 或使用 ADB 命令:
adb shell twrp install /sdcard/rom.zip
```

#### ADB Sideload 刷机

```bash
# 1. 重启到 Recovery
adb reboot recovery

# 2. 在 TWRP 中选择 "Advanced" -> "ADB Sideload"
# 滑动开始

# 3. 在电脑上执行:
adb sideload rom.zip

# 4. 等待完成并重启
```

### 4.3 刷入 GApps (Google 服务)

```bash
# 1. 下载 GApps
# Open GApps: https://opengapps.org/
# 选择: 平台(ARM/ARM64), Android版本(8.1), 变体(nano/micro/mini)

# 2. 在刷入 ROM 后立即刷入 GApps
# TWRP -> Install -> 选择 gapps.zip
# 滑动确认

# 3. 清除 Dalvik/Cache
# Wipe -> Advanced Wipe -> 勾选 Dalvik/ART Cache, Cache

# 4. 重启
```

### 4.4 刷入 Magisk (Root)

```bash
# 1. 下载 Magisk
# GitHub: https://github.com/topjohnwu/Magisk/releases
# 下载: Magisk-v26.1.apk

# 方法 1: 修补 boot.img
# 1. 安装 Magisk APP
adb install Magisk-v26.1.apk

# 2. 推送 boot.img 到设备
adb push boot.img /sdcard/

# 3. 打开 Magisk APP -> 安装 -> 选择并修补文件
# 选择 boot.img, 等待修补完成

# 4. 提取修补后的 boot.img
adb pull /sdcard/Download/magisk_patched_*.img

# 5. 刷入修补后的 boot
adb reboot bootloader
fastboot flash boot magisk_patched_*.img
fastboot reboot

# 方法 2: 在 TWRP 中刷入
adb push Magisk-v26.1.apk /sdcard/Magisk.zip
adb reboot recovery
# TWRP -> Install -> 选择 Magisk.zip
```

---

## 5. ADB 调试

### 5.1 ADB 基础命令

#### 设备管理

```bash
# 列出已连接的设备
adb devices

# 查看设备详细信息
adb devices -l

# 等待设备连接
adb wait-for-device

# 指定设备执行命令 (多设备时)
adb -s DEVICE_SERIAL shell

# 重启设备
adb reboot

# 重启到 Bootloader
adb reboot bootloader

# 重启到 Recovery
adb reboot recovery

# 启动/停止 ADB 服务
adb kill-server
adb start-server
```

#### 文件传输

```bash
# 推送文件到设备
adb push local_file /sdcard/
adb push local_dir/ /sdcard/remote_dir/

# 从设备拉取文件
adb pull /sdcard/file.txt
adb pull /sdcard/dir/ local_dir/

# 批量传输
adb push *.jpg /sdcard/Pictures/
```

#### 应用管理

```bash
# 安装 APK
adb install app.apk

# 覆盖安装 (保留数据)
adb install -r app.apk

# 安装到 SD 卡
adb install -s app.apk

# 允许降级安装
adb install -d app.apk

# 卸载应用
adb uninstall com.example.app

# 卸载但保留数据
adb uninstall -k com.example.app

# 列出已安装的包
adb shell pm list packages

# 列出第三方应用
adb shell pm list packages -3

# 列出系统应用
adb shell pm list packages -s

# 查找特定应用
adb shell pm list packages | grep keyword

# 清除应用数据
adb shell pm clear com.example.app
```

### 5.2 Shell 命令

```bash
# 进入设备 Shell
adb shell

# 执行单个命令
adb shell ls /sdcard/

# 以 root 权限执行 (需要 root)
adb shell su -c "command"

# 常用 Shell 命令:

# 查看系统信息
adb shell getprop
adb shell getprop ro.build.version.release  # Android 版本
adb shell getprop ro.product.model          # 设备型号

# 查看进程
adb shell ps
adb shell ps | grep com.example.app

# 查看内存
adb shell cat /proc/meminfo
adb shell dumpsys meminfo

# 查看CPU
adb shell cat /proc/cpuinfo
adb shell top

# 查看存储
adb shell df -h
adb shell du -h /sdcard/ | sort -hr | head -20

# 查看电池
adb shell dumpsys battery

# 查看屏幕
adb shell wm size    # 分辨率
adb shell wm density # DPI

# 截图
adb shell screencap /sdcard/screen.png
adb pull /sdcard/screen.png

# 录屏 (Android 4.4+)
adb shell screenrecord /sdcard/video.mp4
# Ctrl+C 停止录制
adb pull /sdcard/video.mp4
```

### 5.3 网络 ADB (Wi-Fi 调试)

```bash
# 方法 1: 通过 USB 启用
# 1. USB 连接设备
adb devices

# 2. 查看设备 IP
adb shell ip addr show wlan0 | grep inet

# 3. 启用 TCP/IP 模式
adb tcpip 5555

# 4. 断开 USB, 通过 Wi-Fi 连接
adb connect 192.168.1.100:5555

# 5. 验证连接
adb devices

# 6. 恢复 USB 模式
adb usb

# 方法 2: 直接启用 (需要 root)
adb shell su -c "setprop service.adb.tcp.port 5555"
adb shell su -c "stop adbd"
adb shell su -c "start adbd"
```

---

## 6. Logcat 日志分析

### 6.1 Logcat 基础

```bash
# 查看所有日志
adb logcat

# 清空日志
adb logcat -c

# 查看日志并保存
adb logcat > logcat.txt

# 只显示新日志
adb logcat -c && adb logcat

# 显示日志时间
adb logcat -v time

# 显示详细格式
adb logcat -v long
```

### 6.2 日志过滤

#### 按级别过滤

```bash
# 日志级别: V(Verbose) < D(Debug) < I(Info) < W(Warning) < E(Error) < F(Fatal) < S(Silent)

# 只显示 Error 和 Fatal
adb logcat *:E

# 显示 Warning 及以上
adb logcat *:W

# 显示 Info 及以上
adb logcat *:I
```

#### 按标签过滤

```bash
# 只显示特定标签
adb logcat -s TAG_NAME

# 多个标签
adb logcat -s TAG1 TAG2 TAG3

# 标签和级别组合
adb logcat TAG1:D TAG2:W *:S
```

#### 使用 grep 过滤

```bash
# 包含特定关键词
adb logcat | grep "keyword"

# 排除特定关键词
adb logcat | grep -v "unwanted"

# 多个关键词
adb logcat | grep -E "keyword1|keyword2"

# 忽略大小写
adb logcat | grep -i "keyword"

# 显示上下文
adb logcat | grep -A 5 -B 5 "keyword"  # 前后5行
```

### 6.3 常用日志分析

#### 崩溃日志

```bash
# 查看应用崩溃
adb logcat | grep -E "FATAL|AndroidRuntime"

# 查看 Native 崩溃
adb logcat | grep "DEBUG"

# 查看 ANR (Application Not Responding)
adb logcat | grep "ANR"

# 导出 ANR traces
adb pull /data/anr/traces.txt
```

#### Activity 生命周期

```bash
# 监控 Activity 启动
adb logcat | grep "ActivityManager: START"

# 监控 Activity 生命周期
adb logcat ActivityManager:I *:S
```

#### 系统日志

```bash
# 查看系统服务
adb logcat | grep "SystemServer"

# 查看包管理器
adb logcat | grep "PackageManager"

# 查看窗口管理器
adb logcat | grep "WindowManager"
```

### 6.4 日志格式化

```bash
# 使用 pidcat (彩色日志工具)
# 安装: pip install pidcat
pidcat com.example.app

# 只显示特定进程
adb logcat --pid=$(adb shell pidof -s com.example.app)

# 使用 logcat-color
# 安装: pip install logcat-color
adb logcat -v time | logcat-color
```

---

## 7. 系统调试

### 7.1 dumpsys 系统诊断

```bash
# 查看所有服务
adb shell dumpsys -l

# Activity 信息
adb shell dumpsys activity

# 当前 Activity
adb shell dumpsys activity top

# 包管理器信息
adb shell dumpsys package com.example.app

# 窗口信息
adb shell dumpsys window

# 当前显示的窗口
adb shell dumpsys window | grep "mCurrentFocus"

# 内存信息
adb shell dumpsys meminfo
adb shell dumpsys meminfo com.example.app

# 电池信息
adb shell dumpsys battery

# 网络信息
adb shell dumpsys connectivity

# 通知信息
adb shell dumpsys notification
```

### 7.2 性能分析

```bash
# CPU 使用率
adb shell top -m 10  # 前10个进程

# 实时内存监控
adb shell watch -n 1 "dumpsys meminfo com.example.app | grep TOTAL"

# GPU 渲染
adb shell dumpsys gfxinfo com.example.app

# 帧率统计
adb shell dumpsys SurfaceFlinger --latency

# 网络统计
adb shell dumpsys netstats
```

### 7.3 开发者选项调试

```bash
# 显示 CPU 使用率
adb shell setprop debug.performance.showCpu true

# 显示 FPS
adb shell setprop debug.sf.showfps 1

# 启用严格模式
adb shell setprop persist.sys.strictmode.visual 1

# 显示布局边界
adb shell setprop debug.layout true

# GPU 过度绘制
adb shell setprop debug.hwui.overdraw show

# 显示触摸位置
adb shell setprop debug.pointer_location 1
```

---

## 8. 问题排查

### 8.1 设备无法启动 (Bootloop)

**症状**: 设备卡在启动画面或不断重启

**解决方案**:

```bash
# 方法 1: 清除 Cache
adb reboot recovery
# TWRP -> Wipe -> Advanced Wipe -> Cache, Dalvik

# 方法 2: 恢复出厂设置
# TWRP -> Wipe -> Format Data

# 方法 3: 重新刷入 ROM
adb reboot bootloader
fastboot flash system system.img
fastboot flash boot boot.img
fastboot erase userdata
fastboot erase cache
fastboot reboot

# 方法 4: 刷入官方固件
# 下载官方 ROM 并使用线刷工具刷入
```

### 8.2 设备变砖 (Bricked)

**硬砖 (Hard Brick)**: 完全无法开机, 需要专业维修

**软砖 (Soft Brick)**: 可以进入 Bootloader/Recovery

```bash
# 恢复步骤:

# 1. 尝试进入 Bootloader
# 关机 -> 音量下 + 电源键

# 2. 检查 Fastboot 连接
fastboot devices

# 3. 刷入官方固件
fastboot flash boot boot.img
fastboot flash system system.img
fastboot flash recovery recovery.img
fastboot erase userdata
fastboot reboot

# 4. 如果无法进入 Bootloader
# 尝试 EDL 模式 (Emergency Download Mode)
# 需要专用工具: QFIL, SP Flash Tool 等
```

### 8.3 ADB/Fastboot 无法识别

**解决方案**:

```bash
# Linux:
# 1. 检查 udev 规则
lsusb  # 查找设备 Vendor ID
sudo nano /etc/udev/rules.d/51-android.rules
# 添加规则并重新加载
sudo udevadm control --reload-rules

# 2. 检查 USB 连接
# 更换 USB 线
# 更换 USB 接口 (使用 USB 2.0 接口)

# 3. 重启 ADB 服务
adb kill-server
sudo adb start-server

# Windows:
# 安装正确的驱动
# 使用设备管理器更新驱动
# 或使用 15 seconds ADB Installer
```

### 8.4 刷机后功能异常

**问题**: 指纹、Wi-Fi、相机等功能不正常

**解决方案**:

```bash
# 1. 检查是否刷入了完整的固件
# 确保刷入: boot, system, vendor (重要!)

# 2. 重新刷入 vendor 分区
fastboot flash vendor vendor.img

# 3. 刷入对应的固件版本
# 确保 ROM 版本与设备型号匹配

# 4. 清除所有数据重新刷入
fastboot erase userdata
fastboot erase cache
```

### 8.5 TWRP 无法挂载分区

**问题**: TWRP 提示 "Unable to mount /data"

**解决方案**:

```bash
# 原因: 分区加密或文件系统损坏

# 方法 1: 格式化 Data 分区
# TWRP -> Wipe -> Format Data
# 输入 "yes" 确认

# 方法 2: 使用 Fastboot 格式化
adb reboot bootloader
fastboot erase userdata
fastboot format userdata

# 方法 3: 修复文件系统
adb shell
e2fsck -f /dev/block/bootdevice/by-name/userdata
```

---

## 附录

### A. 常用快捷键组合

| 设备品牌 | Bootloader | Recovery | 强制重启 |
|---------|-----------|----------|---------|
| **Google** | 音量下 + 电源 | 音量上 + 电源 | 长按电源 10s |
| **小米** | 音量下 + 电源 | 音量上 + 电源 | 长按电源 10s |
| **一加** | 音量下 + 电源 | 音量下 + 电源 (选择Recovery) | 长按电源 10s |
| **三星** | 音量下 + Bixby + 电源 | 音量上 + Bixby + 电源 | 长按电源 10s |
| **华为** | 音量下 + 电源 | 音量上 + 电源 | 长按电源 10s |

### B. 分区说明

| 分区名 | 说明 |
|-------|------|
| **boot** | 内核和 ramdisk |
| **system** | Android 系统分区 |
| **vendor** | 厂商特定文件和驱动 |
| **recovery** | Recovery 分区 |
| **userdata/data** | 用户数据和应用 |
| **cache** | 系统缓存 |
| **persist** | 持久化数据 (如 IMEI, MAC) |
| **modem** | 基带固件 |
| **boot loader** | 引导加载程序 |

### C. 调试工具清单

```bash
# 必备工具
- ADB (Android Debug Bridge)
- Fastboot
- TWRP Recovery
- Magisk (Root)

# 日志工具
- Logcat
- pidcat
- logcat-color

# 性能工具
- Systrace
- Perfetto
- Android Profiler

# 网络工具
- Charles Proxy
- Wireshark
- tcpdump

# 逆向工具
- JADX
- APKTool
- Frida
```

---

## 总结

本指南涵盖了 Android 设备刷机和调试的完整流程，从 Bootloader 解锁到系统调试。掌握这些技能后，您可以：

✅ 安全地解锁和刷机  
✅ 安装第三方 ROM 和 Recovery  
✅ 使用 ADB 进行高级调试  
✅ 分析和解决系统问题  
✅ 优化设备性能  

**重要提醒**: 刷机有风险，操作需谨慎！务必备份重要数据！

**祝您刷机顺利！** 🚀
