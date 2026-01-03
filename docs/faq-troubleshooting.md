# Android 8.1 构建与逆向 FAQ 和故障排除

> **作者**: Android Development Team  
> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **适用场景**: AOSP 构建、ROM 定制、逆向工程

---

## 📋 目录

- [1. 编译构建问题](#1-编译构建问题)
- [2. 源码同步问题](#2-源码同步问题)
- [3. 刷机问题](#3-刷机问题)
- [4. 调试问题](#4-调试问题)
- [5. 逆向工程问题](#5-逆向工程问题)
- [6. 工具使用问题](#6-工具使用问题)
- [7. 性能优化](#7-性能优化)
- [8. 进阶问题](#8-进阶问题)

---

## 1. 编译构建问题

### Q1.1: 编译时报错 "You are attempting to build with the incorrect version of java"

**错误信息**:
```
You are attempting to build with the incorrect version of java.
Your version is: java version "11.0.x".
The required version is: "1.8.0"
```

**原因**: Android 8.1 需要 Java 8，而系统使用了其他版本。

**解决方案**:
```bash
# 安装 Java 8
sudo apt-get install openjdk-8-jdk

# 切换到 Java 8
sudo update-alternatives --config java
sudo update-alternatives --config javac

# 设置环境变量
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
export PATH=$JAVA_HOME/bin:$PATH

# 验证
java -version
# 应显示: openjdk version "1.8.0_xxx"
```

---

### Q1.2: 编译时出现 "ninja: build stopped: subcommand failed"

**错误信息**:
```
ninja: build stopped: subcommand failed
build/core/ninja.mk:148: recipe for target 'ninja_wrapper' failed
```

**原因**: 编译过程中某个模块构建失败。

**解决方案**:
```bash
# 1. 查看详细错误信息
# 向上滚动查看具体失败的模块

# 2. 单线程编译以查看详细错误
make -j1

# 3. 清理后重新编译
make clean
make -j$(nproc)

# 4. 如果是特定模块问题
mmm <模块路径>

# 5. 更新构建工具
repo sync build/soong
```

---

### Q1.3: 内存不足导致编译失败 "Out of memory"

**错误信息**:
```
FAILED: out/target/...
java.lang.OutOfMemoryError: Java heap space
```

**解决方案**:
```bash
# 方法 1: 减少并行编译任务
make -j4  # 而不是 -j$(nproc)

# 方法 2: 增加 Java 堆内存
export JACK_SERVER_VM_ARGUMENTS="-Dfile.encoding=UTF-8 -XX:+TieredCompilation -Xmx4g"
./prebuilts/sdk/tools/jack-admin kill-server
./prebuilts/sdk/tools/jack-admin start-server

# 方法 3: 增加系统交换空间
sudo fallocate -l 16G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
# 永久生效
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# 方法 4: 使用 ccache 减少内存占用
export USE_CCACHE=1
ccache -M 50G

# 方法 5: 关闭其他占用内存的程序
```

---

### Q1.4: 磁盘空间不足

**错误信息**:
```
No space left on device
```

**解决方案**:
```bash
# 1. 检查磁盘使用
df -h

# 2. 清理 ccache
ccache -C
ccache -M 30G  # 减小 ccache 大小

# 3. 清理构建输出
cd ~/aosp/android-8.1
make clean

# 4. 清理 repo 缓存
rm -rf .repo/repo
rm -rf .repo/project-objects

# 5. 查找大文件
du -sh * | sort -hr | head -20

# 6. 清理 Docker 或其他不需要的文件
docker system prune -a  # 如果使用 Docker
sudo apt-get autoremove
sudo apt-get clean
```

---

### Q1.5: "Permission denied" 权限错误

**错误信息**:
```
Permission denied: '/path/to/file'
```

**解决方案**:
```bash
# 1. 检查文件权限
ls -la /path/to/file

# 2. 修改权限
chmod +x /path/to/file

# 3. 修改所有者
sudo chown $USER:$USER -R /path/to/dir

# 4. 不要使用 sudo 编译
# AOSP 不应该用 root 权限编译

# 5. 检查 SELinux (如果启用)
getenforce
# 如果是 Enforcing, 临时关闭:
sudo setenforce 0
```

---

### Q1.6: "You are building on a case-insensitive filesystem"

**错误信息**: (macOS 常见)
```
You are building on a case-insensitive filesystem.
```

**解决方案**:
```bash
# macOS 需要创建区分大小写的磁盘映像

# 1. 创建磁盘映像
hdiutil create -type SPARSE -fs 'Case-sensitive Journaled HFS+' \
    -size 250g ~/android.dmg

# 2. 挂载
hdiutil attach ~/android.dmg.sparseimage -mountpoint /Volumes/android

# 3. 在挂载点下工作
cd /Volumes/android
mkdir aosp
cd aosp

# 4. 下载源码
repo init ...
repo sync

# 5. 自动挂载 (可选)
# 添加到 ~/.bash_profile
echo 'hdiutil attach ~/android.dmg.sparseimage -mountpoint /Volumes/android' >> ~/.bash_profile
```

---

## 2. 源码同步问题

### Q2.1: repo sync 失败 "Cannot fetch ..."

**错误信息**:
```
fatal: Cannot fetch AOSP/platform/xxx
error: Cannot fetch platform/xxx
```

**解决方案**:
```bash
# 方法 1: 使用镜像源
export REPO_URL='https://mirrors.tuna.tsinghua.edu.cn/git/git-repo'
repo init -u https://mirrors.tuna.tsinghua.edu.cn/git/AOSP/platform/manifest \
    -b android-8.1.0_r81

# 方法 2: 减少并行数
repo sync -j1 --current-branch --no-tags

# 方法 3: 强制同步
repo sync --force-sync --force-broken

# 方法 4: 同步特定项目
repo sync <项目名>

# 方法 5: 删除问题项目后重新同步
rm -rf <项目路径>
repo sync
```

---

### Q2.2: 网络中断导致同步失败

**解决方案**:
```bash
# 创建自动重试脚本
#!/bin/bash
# auto_sync.sh

#!/bin/bash
MAX_RETRY=5
RETRY_COUNT=0

while [ $RETRY_COUNT -lt $MAX_RETRY ]; do
    repo sync -c -j4 --no-tags --no-clone-bundle
    
    if [ $? -eq 0 ]; then
        echo "Sync completed successfully!"
        exit 0
    else
        RETRY_COUNT=$((RETRY_COUNT + 1))
        echo "Sync failed, retry $RETRY_COUNT/$MAX_RETRY..."
        sleep 10
    fi
done

echo "Sync failed after $MAX_RETRY attempts"
exit 1

# 使用
chmod +x auto_sync.sh
./auto_sync.sh
```

---

### Q2.3: "curl 443 error" 或代理问题

**解决方案**:
```bash
# 方法 1: 配置 Git 代理
git config --global http.proxy http://proxy.example.com:8080
git config --global https.proxy http://proxy.example.com:8080

# 方法 2: 使用 socks5 代理
git config --global http.proxy socks5://127.0.0.1:1080
git config --global https.proxy socks5://127.0.0.1:1080

# 方法 3: 配置 repo 使用代理
export http_proxy=http://proxy.example.com:8080
export https_proxy=http://proxy.example.com:8080

# 方法 4: 取消代理
git config --global --unset http.proxy
git config --global --unset https.proxy
unset http_proxy
unset https_proxy

# 方法 5: 使用 SSH 而不是 HTTPS
# 配置 SSH 密钥后:
repo init -u git@github.com:aosp-mirror/platform_manifest.git
```

---

## 3. 刷机问题

### Q3.1: Bootloader 无法解锁

**问题**: 执行 `fastboot oem unlock` 后设备无反应或报错。

**解决方案**:
```bash
# 1. 确认已启用 OEM 解锁
# 设置 -> 开发者选项 -> OEM 解锁 (开启)

# 2. 尝试不同的解锁命令
fastboot oem unlock
fastboot flashing unlock
fastboot oem unlock-go  # 某些设备

# 3. 检查设备是否支持解锁
fastboot oem device-info

# 4. 品牌特定方法:

# 小米: 使用官方解锁工具
# http://www.miui.com/unlock/

# 华为: 已停止解锁服务 (需第三方)

# 三星: 不需要解锁 Bootloader

# 一加: 直接解锁，无需等待
fastboot oem unlock

# 5. 检查 Fastboot 驱动 (Windows)
# 使用设备管理器安装正确的驱动
```

---

### Q3.2: 刷机后无法开机 (Bootloop)

**症状**: 设备卡在启动动画或不断重启。

**解决方案**:
```bash
# Step 1: 清除缓存
adb reboot recovery
# TWRP -> Wipe -> Advanced Wipe
# 勾选: Dalvik/ART Cache, Cache
# 滑动确认

# Step 2: 格式化 Data
# TWRP -> Wipe -> Format Data
# 输入 "yes" 确认

# Step 3: 重新刷入完整 ROM
adb reboot bootloader
fastboot flash boot boot.img
fastboot flash system system.img
fastboot flash vendor vendor.img  # 重要！
fastboot erase userdata
fastboot erase cache
fastboot reboot

# Step 4: 检查版本兼容性
# 确保 ROM 版本与设备型号匹配

# Step 5: 查看日志
adb logcat > bootloop.log
# 分析日志查找错误原因
```

---

### Q3.3: 刷机后 Wi-Fi/蓝牙/相机不工作

**原因**: 未刷入 vendor 分区或版本不匹配。

**解决方案**:
```bash
# 1. 重新刷入 vendor.img
adb reboot bootloader
fastboot flash vendor vendor.img
fastboot reboot

# 2. 检查是否需要刷入 modem
fastboot flash modem modem.img

# 3. 确保 ROM 版本匹配
# 例如: Android 8.1.0 ROM 需要对应的 vendor 版本

# 4. 从官方 ROM 提取 vendor.img
# 下载官方 ROM
unzip official_rom.zip
# 提取 vendor.img
fastboot flash vendor vendor.img

# 5. 检查权限
adb shell su -c "ls -l /dev/video*"
adb shell su -c "ls -l /dev/rfkill"
```

---

### Q3.4: Fastboot 无法识别设备

**症状**: `fastboot devices` 无输出。

**解决方案**:
```bash
# Linux:
# 1. 检查 USB 连接
lsusb | grep -i google  # 或其他厂商

# 2. 添加 udev 规则
sudo nano /etc/udev/rules.d/51-android.rules
# 添加:
SUBSYSTEM=="usb", ATTR{idVendor}=="18d1", MODE="0666", GROUP="plugdev"

# 3. 重新加载 udev
sudo udevadm control --reload-rules
sudo udevadm trigger

# 4. 重新插拔 USB

# 5. 尝试不同的 USB 接口 (USB 2.0)

# Windows:
# 1. 安装 Fastboot 驱动
# 设备管理器 -> 更新驱动 -> 浏览计算机查找驱动

# 2. 使用 15 seconds ADB Installer

# macOS:
# 1. 安装 Android Platform Tools
brew install android-platform-tools

# 2. 重启设备到 Bootloader
```

---

### Q3.5: 刷入 TWRP 后自动恢复为原厂 Recovery

**原因**: 某些设备在首次启动时会覆盖自定义 Recovery。

**解决方案**:
```bash
# 方法 1: 刷入后直接进入 Recovery
fastboot flash recovery twrp.img
fastboot boot twrp.img  # 直接启动，不要 reboot

# 方法 2: 在 TWRP 内再次刷入
fastboot boot twrp.img  # 临时启动 TWRP
adb push twrp.img /sdcard/
# TWRP -> Install -> Install Image
# 选择 twrp.img -> Recovery 分区

# 方法 3: 禁用 dm-verity (高级)
fastboot flash vbmeta --disable-verity --disable-verification vbmeta.img

# 方法 4: 刷入 Magisk 防止被覆盖
# 在 TWRP 中刷入 Magisk.zip
```

---

## 4. 调试问题

### Q4.1: ADB 显示 "device unauthorized"

**症状**: 
```
List of devices attached
ABC123456789    unauthorized
```

**解决方案**:
```bash
# 1. 在设备上允许 USB 调试
# 设备会弹出授权对话框，勾选"始终允许"并确认

# 2. 撤销之前的授权
adb kill-server
adb start-server

# 3. 删除旧的授权密钥
# Linux/macOS:
rm ~/.android/adbkey*
# Windows:
del %USERPROFILE%\.android\adbkey*

# 4. 重启 ADB 并重新授权
adb kill-server
adb start-server
adb devices
# 在设备上重新授权

# 5. 检查 USB 调试设置
# 设置 -> 开发者选项 -> 撤销 USB 调试授权
# 重新连接设备
```

---

### Q4.2: Logcat 输出乱码或无内容

**解决方案**:
```bash
# 1. 清空日志缓冲区
adb logcat -c

# 2. 检查日志缓冲区大小
adb logcat -g

# 3. 增加缓冲区大小
adb logcat -G 16M

# 4. 指定缓冲区
adb logcat -b all

# 5. 使用不同的格式
adb logcat -v time
adb logcat -v threadtime
adb logcat -v long

# 6. 检查过滤器
adb logcat *:V  # 显示所有级别

# 7. 编码问题 (Windows)
# 修改控制台编码
chcp 65001  # UTF-8
```

---

### Q4.3: 无法通过 Wi-Fi 连接 ADB

**解决方案**:
```bash
# 1. 确保设备和电脑在同一网络

# 2. 启用 TCP/IP 模式 (需要先 USB 连接)
adb tcpip 5555

# 3. 查看设备 IP
adb shell ip addr show wlan0 | grep inet

# 4. 连接
adb connect 192.168.1.100:5555

# 5. 验证连接
adb devices

# 6. 如果连接失败:
# - 检查防火墙设置
# - 确认端口 5555 未被占用
netstat -an | grep 5555

# - 重启 ADB 服务
adb kill-server
adb start-server

# 7. 断开连接
adb disconnect 192.168.1.100:5555

# 8. 恢复 USB 模式
adb usb
```

---

### Q4.4: dumpsys 命令输出过多

**解决方案**:
```bash
# 1. 只输出特定服务
adb shell dumpsys activity

# 2. 使用 grep 过滤
adb shell dumpsys | grep "keyword"

# 3. 输出到文件
adb shell dumpsys > dumpsys_output.txt

# 4. 查看可用服务列表
adb shell dumpsys -l

# 5. 只查看摘要信息
adb shell dumpsys activity top

# 6. 分页查看
adb shell dumpsys | more
adb shell dumpsys | less
```

---

## 5. 逆向工程问题

### Q5.1: APKTool 反编译失败

**错误信息**:
```
Exception in thread "main" brut.androlib.AndrolibException: Could not decode arsc file
```

**解决方案**:
```bash
# 1. 更新 APKTool 到最新版本
wget https://bitbucket.org/iBotPeaches/apktool/downloads/apktool_2.9.3.jar

# 2. 跳过资源反编译
apktool d app.apk --no-res

# 3. 跳过源码反编译
apktool d app.apk --no-src

# 4. 强制覆盖
apktool d app.apk -f

# 5. 使用不同的框架
apktool empty-framework-dir
apktool d app.apk

# 6. 如果是加固应用，先脱壳
# 使用 FART、BlackDex 等工具
```

---

### Q5.2: 重打包后的 APK 无法安装

**错误信息**:
```
Failure [INSTALL_PARSE_FAILED_NO_CERTIFICATES]
```

**解决方案**:
```bash
# 1. 正确的重打包流程:

# 反编译
apktool d original.apk -o app_decompiled

# 修改代码...

# 重新打包
apktool b app_decompiled -o app_modified.apk

# 对齐
zipalign -v 4 app_modified.apk app_aligned.apk

# 签名
# 生成密钥 (首次)
keytool -genkey -v -keystore my-release-key.jks \
    -keyalg RSA -keysize 2048 -validity 10000 \
    -alias my-key-alias

# 签名
apksigner sign --ks my-release-key.jks \
    --out app_signed.apk app_aligned.apk

# 或使用 jarsigner
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 \
    -keystore my-release-key.jks app_aligned.apk my-key-alias

# 验证签名
apksigner verify app_signed.apk

# 2. 卸载旧版本
adb uninstall com.example.app

# 3. 安装新版本
adb install -r app_signed.apk
```

---

### Q5.3: JADX 反编译后代码无法理解

**问题**: 代码中有大量混淆的类名和变量名。

**解决方案**:
```bash
# 1. 使用 JADX-GUI 的重命名功能
# 右键 -> Rename -> 输入有意义的名称

# 2. 搜索字符串常量定位关键代码
# JADX-GUI -> Text Search -> 搜索关键字符串

# 3. 查找交叉引用
# 右键某个方法/类 -> Find Usage

# 4. 导出为 Gradle 项目
jadx -d output --export-gradle app.apk
# 在 Android Studio 中打开，使用代码导航

# 5. 结合动态调试
# 使用 Frida Hook 关键函数，查看运行时数据

# 6. 使用反混淆工具
# Simplify: https://github.com/CalebFenton/simplify
# 对 DEX 文件进行反混淆预处理

# 7. 分析字符串解密函数
# 很多混淆会加密字符串，找到解密函数并提取
```

---

### Q5.4: Frida 无法附加到应用

**错误信息**:
```
Failed to attach: unable to access process with pid XXX
```

**解决方案**:
```bash
# 1. 确认 Frida Server 运行
adb shell su -c "ps | grep frida-server"

# 2. 重启 Frida Server
adb shell su -c "killall frida-server"
adb shell su -c "/data/local/tmp/frida-server &"

# 3. 确认设备已 root
adb shell su -c "id"

# 4. 检查 SELinux
adb shell getenforce
# 如果是 Enforcing:
adb shell su -c "setenforce 0"

# 5. 使用正确的附加方式
# 通过包名:
frida -U -f com.example.app -l script.js --no-pause

# 通过进程名:
frida -U -n "App Name" -l script.js

# 通过 PID:
frida -U -p 12345 -l script.js

# 6. 绕过反调试
# 某些应用会检测 Frida，需要使用反反调试脚本

# 7. 检查 Frida 版本兼容性
frida --version
# 确保 Frida 工具和 Frida Server 版本一致
```

---

### Q5.5: 无法绕过 SSL Pinning

**解决方案**:
```bash
# 方法 1: 使用 objection
objection -g com.example.app explore
android sslpinning disable

# 方法 2: 使用通用 Frida 脚本
# ssl-unpinning.js (多种框架)
frida -U -f com.example.app -l ssl-unpinning.js --no-pause

# 方法 3: Hook 自定义 Pinning
# 反编译找到 Pinning 实现，编写针对性 Hook 脚本

# 方法 4: 修改 APK
# 1. 反编译 APK
apktool d app.apk
# 2. 删除 network_security_config.xml 中的 pinning 配置
# 3. 或修改代码移除证书校验
# 4. 重打包签名

# 方法 5: 使用 Magisk 模块
# TrustMeAlready: 系统级 SSL Pinning 绕过
# https://github.com/ViRb3/TrustMeAlready

# 方法 6: 使用 Xposed 模块
# JustTrustMe
# SSLUnpinning
```

---

## 6. 工具使用问题

### Q6.1: Android Studio 非常卡顿

**解决方案**:
```bash
# 1. 增加 IDE 内存
# 编辑 studio.vmoptions
# Linux: ~/.config/Google/AndroidStudio*/studio64.vmoptions
-Xms2g
-Xmx8g
-XX:ReservedCodeCacheSize=512m

# 2. 禁用不需要的插件
# File -> Settings -> Plugins

# 3. 排除不需要的文件夹
# File -> Settings -> Editor -> File Types -> Ignore files and folders
# 添加: *.o;*.so;out;.repo

# 4. 禁用自动同步
# File -> Settings -> Build -> Compiler
# 取消勾选 "Compile independent modules in parallel"

# 5. 使用 Power Save Mode
# File -> Power Save Mode

# 6. 清理缓存
# File -> Invalidate Caches / Restart

# 7. 升级硬件
# - 使用 SSD
# - 增加 RAM 到 16GB+
# - 使用更快的 CPU
```

---

### Q6.2: ccache 不生效

**解决方案**:
```bash
# 1. 确认已启用
export USE_CCACHE=1
export CCACHE_DIR=~/.ccache

# 2. 检查 ccache 状态
ccache -s

# 3. 设置合适的大小
ccache -M 50G

# 4. 查看配置
ccache -p

# 5. 清空并重建缓存
ccache -C
ccache -z  # 重置统计

# 6. 添加到环境变量
echo 'export USE_CCACHE=1' >> ~/.bashrc
echo 'export CCACHE_DIR=~/.ccache' >> ~/.bashrc
source ~/.bashrc

# 7. 验证是否使用 ccache
# 编译时查看命令是否包含 ccache
make -n | grep ccache
```

---

## 7. 性能优化

### Q7.1: 如何加速 AOSP 编译？

**优化方案**:
```bash
# 1. 使用 ccache
export USE_CCACHE=1
ccache -M 100G

# 2. 增加并行编译数
make -j$(nproc)  # 使用所有 CPU 核心
make -j16        # 或指定数量

# 3. 使用 SSD 存储源码
# 编译速度提升 3-5 倍

# 4. 增加内存
# 至少 16GB，推荐 32GB+

# 5. 使用预编译的工具链
export USE_PREBUILT_TOOLCHAIN=1

# 6. 只编译需要的模块
mmm packages/apps/Settings
mma  # 编译当前目录

# 7. 使用增量编译
# 不要每次都 make clean

# 8. 禁用 Jack 服务器内存限制
export JACK_SERVER_VM_ARGUMENTS="-Xmx4g"

# 9. 使用分布式编译 (高级)
# distcc, icecc
```

---

### Q7.2: 如何减小编译输出大小？

**优化方案**:
```bash
# 1. 使用 user 版本 (移除调试符号)
lunch aosp_arm64-user

# 2. 剥离调试信息
# 在 BoardConfig.mk 中:
BOARD_SYSTEMIMAGE_PARTITION_SIZE := 2147483648
TARGET_COPY_OUT_SYSTEM := system

# 3. 移除不需要的应用
# device.mk:
PRODUCT_PACKAGES_REMOVE := \
    Browser Email Gallery2

# 4. 使用压缩
# system.img 使用 squashfs
BOARD_SYSTEMIMAGE_FILE_SYSTEM_TYPE := squashfs

# 5. 精简语言资源
PRODUCT_LOCALES := en_US zh_CN

# 6. 移除不需要的架构
# 只编译 ARM64:
TARGET_ARCH := arm64
```

---

## 8. 进阶问题

### Q8.1: 如何添加自定义系统服务？

**步骤**:
```bash
# 1. 创建服务类
# frameworks/base/services/core/java/com/android/server/MyService.java

package com.android.server;

public class MyService extends SystemService {
    public MyService(Context context) {
        super(context);
    }
    
    @Override
    public void onStart() {
        // 服务启动逻辑
    }
}

# 2. 注册服务
# frameworks/base/services/java/com/android/server/SystemServer.java

mSystemServiceManager.startService(MyService.class);

# 3. 添加 SELinux 权限
# system/sepolicy/private/myservice.te

type myservice, domain;
type myservice_exec, exec_type, file_type;

# 4. 编译
mmm frameworks/base/services

# 5. 刷入测试
adb root
adb remount
adb push out/target/product/.../services.jar /system/framework/
adb reboot
```

---

### Q8.2: 如何移植 ROM 到新设备？

**步骤**:
```bash
# 1. 获取设备树 (Device Tree)
# 从官方 ROM 提取或参考类似设备

device/
└── manufacturer/
    └── device_codename/
        ├── Android.mk
        ├── AndroidProducts.mk
        ├── BoardConfig.mk
        ├── device.mk
        ├── lineage_device.mk
        └── proprietary-files.txt

# 2. 提取专有文件
adb pull /system/vendor vendor/
adb pull /system/lib lib/
adb pull /system/lib64 lib64/

# 3. 配置内核
# 使用厂商提供的内核或自己编译

# 4. 配置 Vendor
# 提取 vendor.img 或从官方 ROM 获取

# 5. 编译
lunch lineage_device-userdebug
mka bacon

# 6. 测试
fastboot flashall

# 7. 调试
# 根据日志解决兼容性问题
adb logcat | grep -E "ERROR|FATAL"
```

---

### Q8.3: 如何编译 Android 内核？

**步骤**:
```bash
# 1. 下载内核源码
git clone https://android.googlesource.com/kernel/common -b android-4.4-o
cd common

# 2. 配置交叉编译工具链
export ARCH=arm64
export CROSS_COMPILE=aarch64-linux-android-

# 3. 加载设备配置
make device_defconfig

# 或自定义:
make menuconfig

# 4. 编译
make -j$(nproc)

# 5. 输出
# arch/arm64/boot/Image.gz-dtb

# 6. 打包 boot.img
# 使用 mkbootimg:
mkbootimg \
    --kernel arch/arm64/boot/Image.gz-dtb \
    --ramdisk ramdisk.img \
    --cmdline "..." \
    --base 0x80000000 \
    --pagesize 4096 \
    --os_version 8.1.0 \
    --os_patch_level 2018-12 \
    --output boot.img

# 7. 刷入测试
fastboot flash boot boot.img
fastboot reboot
```

---

## 附录: 快速诊断脚本

```bash
#!/bin/bash
# android_diagnostic.sh - Android 开发环境诊断脚本

echo "========== Android Development Environment Diagnostic =========="
echo ""

# 检查 Java
echo "=== Java ==="
if command -v java &> /dev/null; then
    java -version 2>&1 | head -n 3
    echo "JAVA_HOME: ${JAVA_HOME:-NOT SET}"
else
    echo "❌ Java not installed"
fi
echo ""

# 检查 Python
echo "=== Python ==="
if command -v python2 &> /dev/null; then
    echo "Python 2: $(python2 --version 2>&1)"
else
    echo "❌ Python 2 not installed"
fi
if command -v python3 &> /dev/null; then
    echo "Python 3: $(python3 --version 2>&1)"
else
    echo "❌ Python 3 not installed"
fi
echo ""

# 检查构建工具
echo "=== Build Tools ==="
for tool in make git repo adb fastboot; do
    if command -v $tool &> /dev/null; then
        echo "✓ $tool: $($tool --version 2>&1 | head -n 1)"
    else
        echo "❌ $tool not installed"
    fi
done
echo ""

# 检查磁盘空间
echo "=== Disk Space ==="
df -h / | tail -n 1 | awk '{print "Root: " $4 " available (" $5 " used)"}'
echo ""

# 检查内存
echo "=== Memory ==="
free -h | grep "Mem:" | awk '{print "Total: " $2 ", Used: " $3 ", Available: " $7}'
echo ""

# 检查 ccache
echo "=== ccache ==="
if command -v ccache &> /dev/null; then
    ccache -s | grep -E "cache size|max cache size"
    echo "USE_CCACHE: ${USE_CCACHE:-NOT SET}"
else
    echo "❌ ccache not installed"
fi
echo ""

# 检查设备连接
echo "=== Connected Devices ==="
if command -v adb &> /dev/null; then
    adb devices
else
    echo "❌ ADB not available"
fi
echo ""

echo "========== Diagnostic Complete =========="
```

---

## 总结

本 FAQ 涵盖了 Android 8.1 构建、刷机、调试和逆向工程中最常见的问题。遇到问题时：

1. ��先查阅本文档
2. 查看详细的错误日志
3. 搜索官方文档和社区论坛
4. 使用诊断脚本定位问题

**记住**: 大部分问题都有解决方案，保持耐心和细心！

**祝您开发顺利！** 🚀
