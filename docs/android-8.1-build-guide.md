# Android 8.1 AOSP 定制构建完整指南

> **作者**: Android Build Team  
> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **Android 版本**: 8.1.0 Oreo (API Level 27)

---

## 📋 目录

- [1. 概述](#1-概述)
- [2. 系统要求](#2-系统要求)
- [3. 环境准备](#3-环境准备)
- [4. 下载 AOSP 源码](#4-下载-aosp-源码)
- [5. 构建配置](#5-构建配置)
- [6. 编译构建](#6-编译构建)
- [7. 定制修改](#7-定制修改)
- [8. 刷机部署](#8-刷机部署)
- [9. 调试技巧](#9-调试技巧)
- [10. 常见问题](#10-常见问题)

---

## 1. 概述

### 1.1 什么是 Android 8.1 AOSP

AOSP (Android Open Source Project) 是 Google 提供的 Android 开源项目。Android 8.1 (Oreo) 是 2017 年 12 月发布的版本，具有以下特性：

- **API Level**: 27
- **Linux Kernel**: 4.4+
- **构建系统**: Soong + Make
- **主要特性**:
  - Neural Networks API
  - 自动填充框架增强
  - 通知渠道改进
  - 画中画模式优化
  - 后台执行限制

### 1.2 定制构建的用途

- **设备适配**: 为特定硬件设备定制 ROM
- **功能增强**: 添加自定义功能和服务
- **系统优化**: 性能优化、精简系统
- **安全加固**: 添加安全机制和权限控制
- **研究学习**: 深入理解 Android 系统架构

### 1.3 构建流程概览

```
环境准备 → 下载源码 → 配置设备 → 编译构建 → 生成镜像 → 刷入设备
```

---

## 2. 系统要求

### 2.1 硬件要求

| 组件 | 最低要求 | 推荐配置 |
|------|----------|----------|
| **CPU** | 4 核心 | 8 核心以上 |
| **内存** | 16 GB RAM | 32 GB RAM 或更多 |
| **硬盘** | 250 GB 可用空间 | 500 GB SSD |
| **网络** | 稳定的互联网连接 | 高速网络（源码下载约 100+ GB） |

### 2.2 操作系统要求

**推荐系统**:
- Ubuntu 18.04 LTS (64-bit)
- Ubuntu 20.04 LTS (64-bit)
- Debian 10 (64-bit)

**其他支持的系统**:
- macOS 10.13+ (High Sierra 或更高)
- Windows 10/11 + WSL2 (不推荐，构建速度慢)

### 2.3 软件版本要求

```bash
# Java
OpenJDK 8 (必需)

# Python
Python 2.7 (构建系统使用)
Python 3.6+ (部分工具使用)

# Make
GNU Make 3.81-3.82

# Git
Git 2.17+
```

---

## 3. 环境准备

### 3.1 安装系统依赖

#### Ubuntu/Debian 系统

```bash
# 更新软件包列表
sudo apt-get update

# 安装必需的依赖包
sudo apt-get install -y \
    git-core gnupg flex bison build-essential zip curl \
    zlib1g-dev gcc-multilib g++-multilib libc6-dev-i386 \
    libncurses5 lib32ncurses5-dev x11proto-core-dev libx11-dev \
    lib32z1-dev libgl1-mesa-dev libxml2-utils xsltproc unzip \
    fontconfig openjdk-8-jdk python-dev python3-dev \
    rsync schedtool ccache lzop bc libssl-dev imagemagick \
    repo android-tools-adb android-tools-fastboot
```

### 3.2 配置 Java 环境

```bash
# 安装 OpenJDK 8
sudo apt-get install openjdk-8-jdk

# 设置默认 Java 版本
sudo update-alternatives --config java
sudo update-alternatives --config javac

# 验证 Java 版本
java -version
javac -version

# 应该显示: openjdk version "1.8.0_xxx"
```

### 3.3 配置 Git

```bash
# 配置 Git 用户信息
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# 配置 Git 颜色输出
git config --global color.ui true
```

### 3.4 安装 Repo 工具

```bash
# 创建 bin 目录
mkdir -p ~/bin
export PATH=~/bin:$PATH

# 下载 Repo 工具
curl https://storage.googleapis.com/git-repo-downloads/repo > ~/bin/repo
chmod a+x ~/bin/repo

# 将 PATH 添加到 ~/.bashrc
echo 'export PATH=~/bin:$PATH' >> ~/.bashrc
source ~/.bashrc

# 验证 Repo 安装
repo version
```

### 3.5 配置 ccache (加速编译)

```bash
# 安装 ccache
sudo apt-get install ccache

# 设置 ccache 大小 (推荐 50-100GB)
ccache -M 50G

# 配置环境变量
echo 'export USE_CCACHE=1' >> ~/.bashrc
echo 'export CCACHE_DIR=~/.ccache' >> ~/.bashrc
source ~/.bashrc

# 查看 ccache 配置
ccache -s
```

---

## 4. 下载 AOSP 源码

### 4.1 创建工作目录

```bash
# 创建 AOSP 源码目录
mkdir -p ~/aosp/android-8.1
cd ~/aosp/android-8.1
```

### 4.2 初始化 Repo 仓库

```bash
# 初始化 Android 8.1 分支
repo init -u https://android.googlesource.com/platform/manifest -b android-8.1.0_r81

# 使用清华大学镜像源 (中国用户推荐)
repo init -u https://mirrors.tuna.tsinghua.edu.cn/git/AOSP/platform/manifest -b android-8.1.0_r81
```

### 4.3 同步源码

```bash
# 开始同步源码 (使用多线程下载，j4 表示 4 线程)
repo sync -c -j4 --no-tags --no-clone-bundle

# 说明:
# -c: 只下载当前分支
# -j4: 使用 4 个并行任务
# --no-tags: 不下载 tag
# --no-clone-bundle: 不使用 clone bundle (某些镜像源不支持)
```

**注意**: 源码下载非常耗时，完整的 AOSP 源码约 100+ GB，根据网络速度可能需要几小时到几天。

### 4.4 源码目录结构

下载完成后，目录结构如下：

```
aosp/android-8.1/
├── art/                    # Android Runtime (ART)
├── bionic/                 # C 库
├── bootable/               # 启动加载器
├── build/                  # 构建系统
├── cts/                    # 兼容性测试套件
├── dalvik/                 # Dalvik 虚拟机
├── developers/             # 开发者资源
├── development/            # 开发工具
├── device/                 # 设备特定配置
├── external/               # 外部开源项目
├── frameworks/             # Android 框架
├── hardware/               # 硬件抽象层 (HAL)
├── kernel/                 # Linux 内核
├── packages/               # 系统应用和服务
├── platform_testing/       # 平台测试
├── prebuilts/              # 预编译工具
├── sdk/                    # Android SDK
├── system/                 # 系统核心组件
├── toolchain/              # 工具链
├── tools/                  # 构建和开发工具
└── vendor/                 # 厂商特定代码
```

---

## 5. 构建配置

### 5.1 初始化构建环境

```bash
# 进入源码目录
cd ~/aosp/android-8.1

# 初始化环境变量
source build/envsetup.sh

# 显示可用的 lunch 选项
lunch
```

### 5.2 选择目标设备

#### 常用的构建目标

| 目标名称 | 说明 | 用途 |
|---------|------|------|
| `aosp_arm64-eng` | ARM64 工程版本 | 开发和调试 |
| `aosp_arm-eng` | ARM 32位工程版本 | 旧设备开发 |
| `aosp_x86_64-eng` | x86_64 工程版本 | 模拟器使用 |
| `aosp_arm64-userdebug` | ARM64 用户调试版本 | 测试和调试 |
| `aosp_arm64-user` | ARM64 用户版本 | 正式发布 |

#### 构建变体说明

- **eng (Engineering)**: 开发版本，包含完整的调试工具，root 权限
- **userdebug**: 用户调试版本，类似用户版本但可以 root，包含调试工具
- **user**: 正式发布版本，移除调试功能，禁用 root

```bash
# 选择 ARM64 工程版本
lunch aosp_arm64-eng

# 或者直接指定
lunch 23  # 假设 23 是 aosp_arm64-eng 的编号
```

### 5.3 配置设备特定选项

如果要为特定设备构建，需要：

1. **添加设备配置文件** (device.mk)
2. **配置内核** (kernel config)
3. **设置硬件抽象层** (HAL)

示例：为 Nexus 5X 构建

```bash
# 下载设备特定文件
git clone https://android.googlesource.com/device/lge/bullhead device/lge/bullhead
git clone https://android.googlesource.com/device/lge/bullhead-kernel device/lge/bullhead-kernel

# 下载供应商二进制文件
./device/lge/bullhead/download-vendor.sh

# 选择设备
lunch aosp_bullhead-userdebug
```

---

## 6. 编译构建

### 6.1 完整构建

```bash
# 开始编译 (使用所有 CPU 核心)
make -j$(nproc)

# 或指定线程数 (例如 8 线程)
make -j8

# 如果遇到问题，使用单线程构建 (便于查看错误)
make
```

**构建时间**: 根据硬件配置，首次编译可能需要 1-8 小时。

### 6.2 增量构建

```bash
# 清理输出目录
make clean

# 只清理当前目标
make installclean

# 重新编译特定模块
mmm packages/apps/Settings

# 编译并安装到设备
mmm packages/apps/Settings -j8
adb install -r out/target/product/generic_arm64/system/app/Settings/Settings.apk
```

### 6.3 构建特定组件

```bash
# 只编译 framework
make framework

# 只编译 system image
make systemimage

# 只编译 boot image
make bootimage

# 编译所有镜像
make droid
```

### 6.4 查看构建输出

构建完成后，输出文件位于：

```bash
out/target/product/<device_name>/

# 主要输出文件:
├── system.img          # 系统分区镜像
├── boot.img            # 启动镜像
├── userdata.img        # 用户数据镜像
├── ramdisk.img         # RAM 磁盘镜像
├── recovery.img        # 恢复模式镜像
├── vendor.img          # 厂商分区镜像 (如果存在)
└── *.zip              # OTA 更新包
```

---

## 7. 定制修改

### 7.1 修改系统应用

#### 移除系统应用

编辑 `device/<manufacturer>/<device>/device.mk`:

```makefile
# 移除不需要的系统应用
PRODUCT_PACKAGES_REMOVE := \
    Browser \
    Email \
    Gallery2 \
    Music

# 或者在 build/core/product.mk 中修改
```

#### 添加自定义应用

```makefile
# 在 device.mk 中添加
PRODUCT_PACKAGES += \
    MyCustomApp \
    AnotherApp

# 将 APK 放到 vendor/<manufacturer>/<device>/apps/ 目录
```

### 7.2 修改系统属性

编辑 `build.prop` 或 `system.prop`:

```bash
# device/<manufacturer>/<device>/system.prop

# 修改 Build 信息
ro.build.display.id=My Custom ROM
ro.build.version.release=8.1.0
ro.build.version.incremental=MyBuild-001

# 性能优化
dalvik.vm.heapsize=512m
persist.sys.use_dithering=1

# 调试选项
persist.service.adb.enable=1
persist.service.debuggable=1
persist.sys.usb.config=mtp,adb
```

### 7.3 修改启动动画

```bash
# 创建自定义启动动画
cd device/<manufacturer>/<device>

# 启动动画目录结构
bootanimation/
├── desc.txt           # 描述文件
├── part0/             # 第一部分动画
│   ├── 00000.png
│   ├── 00001.png
│   └── ...
└── part1/             # 第二部分动画
    └── ...

# desc.txt 格式:
# 宽度 高度 帧率
# p 循环次数 延迟 目录名
1080 1920 30
p 1 0 part0
p 0 0 part1

# 打包启动动画
cd bootanimation
zip -0qry -i \*.txt \*.png @ ../bootanimation.zip *.txt part*

# 在 device.mk 中指定
PRODUCT_COPY_FILES += \
    device/<manufacturer>/<device>/bootanimation.zip:system/media/bootanimation.zip
```

### 7.4 添加系统补丁

```bash
# 应用补丁文件
cd frameworks/base
patch -p1 < ~/patches/custom-feature.patch

# 查看补丁内容
git diff > ~/my-changes.patch

# 撤销补丁
patch -p1 -R < ~/patches/custom-feature.patch
```

### 7.5 修改系统权限

编辑 `frameworks/base/core/res/AndroidManifest.xml`:

```xml
<!-- 添加自定义权限 -->
<permission android:name="com.custom.permission.ACCESS_CUSTOM_SERVICE"
    android:label="@string/custom_permission_label"
    android:description="@string/custom_permission_desc"
    android:protectionLevel="signature" />
```

### 7.6 内核定制

```bash
# 进入内核目录
cd kernel/<manufacturer>/<device>

# 配置内核
make menuconfig

# 或使用已有配置
make <device>_defconfig

# 编译内核
make -j$(nproc)

# 内核输出
arch/arm64/boot/Image.gz-dtb
```

---

## 8. 刷机部署

### 8.1 准备设备

#### 解锁 Bootloader

**警告**: 解锁 Bootloader 会清除所有数据！

```bash
# 重启到 Bootloader
adb reboot bootloader

# 解锁 Bootloader
fastboot oem unlock

# 或者 (新设备)
fastboot flashing unlock
```

#### 备份数据

```bash
# 备份用户数据
adb backup -all -f backup.ab

# 备份分区 (需要 root)
adb shell su -c "dd if=/dev/block/bootdevice/by-name/boot of=/sdcard/boot.img"
adb pull /sdcard/boot.img
```

### 8.2 刷入镜像

#### 方法 1: 使用 Fastboot

```bash
# 确保设备在 Fastboot 模式
adb reboot bootloader

# 检查设备连接
fastboot devices

# 刷入各个分区
fastboot flash boot out/target/product/<device>/boot.img
fastboot flash system out/target/product/<device>/system.img
fastboot flash userdata out/target/product/<device>/userdata.img
fastboot flash vendor out/target/product/<device>/vendor.img

# 清除缓存
fastboot erase cache

# 重启设备
fastboot reboot
```

#### 方法 2: 刷入完整镜像包

```bash
# 使用 flashall 命令 (需要 Android.mk 中配置)
cd out/target/product/<device>
fastboot flashall

# 或使用脚本
./flash-all.sh
```

#### 方法 3: 使用 Recovery 刷入

```bash
# 构建 OTA 包
make otapackage

# 重启到 Recovery
adb reboot recovery

# 在 Recovery 中选择 "Apply update from ADB"
adb sideload out/target/product/<device>/aosp_<device>-ota-*.zip
```

### 8.3 验证刷机结果

```bash
# 检查设备启动
adb wait-for-device

# 查看系统信息
adb shell getprop ro.build.display.id
adb shell getprop ro.build.version.release
adb shell getprop ro.product.model

# 查看系统日志
adb logcat | grep "System"
```

---

## 9. 调试技巧

### 9.1 ADB 调试

```bash
# 连接设备
adb devices

# 查看日志
adb logcat
adb logcat | grep "ActivityManager"
adb logcat *:E  # 只显示错误

# 进入 Shell
adb shell

# 安装应用
adb install app.apk
adb install -r app.apk  # 覆盖安装

# 文件操作
adb push local_file /sdcard/
adb pull /sdcard/file local_file

# 重启设备
adb reboot
adb reboot bootloader
adb reboot recovery
```

### 9.2 系统日志分析

```bash
# 导出完整日志
adb logcat -d > logcat.txt

# 查看内核日志
adb shell dmesg

# 查看系统事件
adb shell dumpsys

# 查看特定服务
adb shell dumpsys activity
adb shell dumpsys package
adb shell dumpsys window
```

### 9.3 性能分析

```bash
# CPU 使用率
adb shell top

# 内存使用
adb shell dumpsys meminfo

# 进程信息
adb shell ps
adb shell ps | grep <package_name>

# 性能跟踪
adb shell atrace --async_start gfx view wm
# ... 执行操作 ...
adb shell atrace --async_stop -z -c > trace.html
```

### 9.4 网络调试

```bash
# 查看网络接口
adb shell ifconfig

# 查看网络连接
adb shell netstat

# Ping 测试
adb shell ping 8.8.8.8

# 通过 Wi-Fi 连接 ADB
adb tcpip 5555
adb connect <device_ip>:5555
```

---

## 10. 常见问题

### 10.1 编译错误

#### 错误: "Out of memory"

**解决方案**:
```bash
# 减少并行任务数
make -j4

# 或增加交换空间
sudo fallocate -l 8G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

#### 错误: "ninja: build stopped"

**解决方案**:
```bash
# 清理构建输出
make clean

# 重新初始化环境
source build/envsetup.sh
lunch <target>

# 重新编译
make -j8
```

#### 错误: "Java version incorrect"

**解决方案**:
```bash
# 确保使用 Java 8
sudo update-alternatives --config java
java -version  # 应显示 1.8.0

# 设置 JAVA_HOME
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
```

### 10.2 源码同步问题

#### 错误: "repo sync 失败"

**解决方案**:
```bash
# 使用镜像源
repo init -u https://mirrors.tuna.tsinghua.edu.cn/git/AOSP/platform/manifest

# 减少并行数
repo sync -j1

# 强制同步
repo sync --force-sync
```

#### 错误: "空间不足"

**解决方案**:
```bash
# 清理 Git 缓存
repo forall -c 'git gc'

# 删除未使用的分支
repo prune
```

### 10.3 刷机问题

#### 错误: "fastboot: device not found"

**解决方案**:
```bash
# 检查 USB 连接
lsusb

# 添加 udev 规则
sudo nano /etc/udev/rules.d/51-android.rules
# 添加: SUBSYSTEM=="usb", ATTR{idVendor}=="18d1", MODE="0666"

# 重新加载规则
sudo udevadm control --reload-rules
```

#### 错误: "boot loop (循环重启)"

**解决方案**:
```bash
# 清除 userdata 和 cache
fastboot erase userdata
fastboot erase cache

# 重新刷入完整镜像
fastboot flashall -w  # -w 会清除数据
```

### 10.4 性能优化

#### 编译速度慢

**优化方案**:
```bash
# 使用 ccache
export USE_CCACHE=1
ccache -M 100G

# 使用预编译工具链
export USE_PREBUILT_TOOLCHAIN=1

# 使用 SSD 存储源码
# 增加内存和 CPU 核心数
```

#### 生成的镜像太大

**优化方案**:
```bash
# 构建 user 版本 (移除调试工具)
lunch aosp_<device>-user

# 移除不需要的应用
# 编辑 device.mk, 添加 PRODUCT_PACKAGES_REMOVE

# 启用压缩
# 在 BoardConfig.mk 中:
# BOARD_SYSTEMIMAGE_PARTITION_SIZE := 2147483648
# BOARD_SYSTEMIMAGE_FILE_SYSTEM_TYPE := ext4
```

---

## 附录

### A. 有用的命令速查表

```bash
# 环境初始化
source build/envsetup.sh
lunch <target>

# 编译
make -j$(nproc)           # 完整构建
make systemimage          # 只构建系统镜像
mmm <path>               # 编译指定模块

# 刷机
fastboot flashall        # 刷入所有分区
fastboot flash boot boot.img  # 刷入 boot 分区

# 调试
adb logcat               # 查看日志
adb shell                # 进入设备 Shell
adb reboot bootloader    # 重启到 Bootloader
```

### B. 重要文件和目录

| 路径 | 说明 |
|------|------|
| `build/envsetup.sh` | 环境初始化脚本 |
| `build/core/` | 构建系统核心 |
| `device/<manufacturer>/<device>/` | 设备配置 |
| `out/target/product/<device>/` | 构建输出 |
| `frameworks/base/` | Android 框架 |
| `system/core/` | 核心系统组件 |
| `packages/apps/` | 系统应用 |

### C. 参考资源

- **官方文档**: https://source.android.com/setup/build
- **AOSP 源码**: https://android.googlesource.com
- **镜像源 (清华)**: https://mirrors.tuna.tsinghua.edu.cn/help/AOSP/
- **XDA 论坛**: https://forum.xda-developers.com
- **LineageOS 文档**: https://wiki.lineageos.org

---

## 总结

本指南涵盖了 Android 8.1 AOSP 构建的完整流程，从环境准备到最终刷机部署。按照本指南操作，您应该能够：

✅ 配置完整的构建环境  
✅ 下载和编译 AOSP 源码  
✅ 进行系统定制和修改  
✅ 刷入定制系统到设备  
✅ 进行调试和问题排查

**下一步**: 学习[逆向工程工具使用](./reverse-engineering-tools.md)，深入分析 Android 系统和应用。

---

**祝您构建成功！** 🎉
