# Android 8.1 开发环境配置完全指南

> **作者**: Android Build Team  
> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **适用系统**: Ubuntu 18.04/20.04, Debian 10+, macOS 10.13+

---

## 📋 目录

- [1. 系统环境准备](#1-系统环境准备)
- [2. 必备软件安装](#2-必备软件安装)
- [3. Android SDK 配置](#3-android-sdk-配置)
- [4. 构建工具安装](#4-构建工具安装)
- [5. IDE 配置](#5-ide-配置)
- [6. 逆向工具安装](#6-逆向工具安装)
- [7. 环境验证](#7-环境验证)
- [8. 常见问题](#8-常见问题)

---

## 1. 系统环境准备

### 1.1 Ubuntu/Debian 系统配置

#### 更新系统

```bash
# 更新软件包列表
sudo apt-get update

# 升级已安装的软件包
sudo apt-get upgrade -y

# 安装基础工具
sudo apt-get install -y \
    wget curl vim git build-essential \
    software-properties-common apt-transport-https
```

#### 配置系统参数

```bash
# 增加文件监视数量限制
echo "fs.inotify.max_user_watches=524288" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p

# 增加文件描述符限制
echo "* soft nofile 65536" | sudo tee -a /etc/security/limits.conf
echo "* hard nofile 65536" | sudo tee -a /etc/security/limits.conf

# 配置交换空间 (如果内存不足)
sudo fallocate -l 16G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

### 1.2 macOS 系统配置

#### 安装 Homebrew

```bash
# 安装 Homebrew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# 更新 Homebrew
brew update

# 安装基础工具
brew install wget curl git coreutils
```

#### 配置环境变量

```bash
# 编辑 ~/.zshrc 或 ~/.bash_profile
echo 'export PATH="/usr/local/opt/coreutils/libexec/gnubin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### 1.3 磁盘空间规划

| 用途 | 推荐空间 | 说明 |
|------|---------|------|
| AOSP 源码 | 250 GB | 完整源码树 |
| 构建输出 | 100 GB | 编译生成的文件 |
| ccache | 50-100 GB | 编译缓存 |
| 工具和 SDK | 50 GB | 各种开发工具 |
| **总计** | **450-500 GB** | 推荐使用 SSD |

---

## 2. 必备软件安装

### 2.1 安装 Java 开发环境

#### Ubuntu/Debian

```bash
# 安装 OpenJDK 8 (AOSP 8.1 必需)
sudo apt-get install -y openjdk-8-jdk openjdk-8-jre

# 安装 OpenJDK 11 (用于新工具)
sudo apt-get install -y openjdk-11-jdk

# 管理 Java 版本
sudo update-alternatives --config java
sudo update-alternatives --config javac

# 设置环境变量
echo 'export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64' >> ~/.bashrc
echo 'export PATH=$JAVA_HOME/bin:$PATH' >> ~/.bashrc
source ~/.bashrc

# 验证安装
java -version
javac -version
```

#### macOS

```bash
# 安装 OpenJDK 8
brew tap AdoptOpenJDK/openjdk
brew install --cask adoptopenjdk8

# 设置 JAVA_HOME
echo 'export JAVA_HOME=$(/usr/libexec/java_home -v 1.8)' >> ~/.zshrc
echo 'export PATH=$JAVA_HOME/bin:$PATH' >> ~/.zshrc
source ~/.zshrc

# 验证安装
java -version
```

### 2.2 安装 Python 环境

```bash
# Ubuntu/Debian
sudo apt-get install -y python2.7 python3 python3-pip

# 创建软链接
sudo ln -s /usr/bin/python2.7 /usr/bin/python2
sudo ln -s /usr/bin/python3 /usr/bin/python3

# 升级 pip
python3 -m pip install --upgrade pip

# macOS
brew install python@2.7 python@3.9

# 验证安装
python2 --version
python3 --version
pip3 --version
```

### 2.3 安装构建依赖

#### Ubuntu 18.04/20.04

```bash
sudo apt-get install -y \
    git-core gnupg flex bison gperf build-essential \
    zip curl zlib1g-dev gcc-multilib g++-multilib \
    libc6-dev-i386 lib32ncurses5-dev x11proto-core-dev \
    libx11-dev lib32z-dev libgl1-mesa-dev libxml2-utils \
    xsltproc unzip m4 lib32z1-dev ccache libssl-dev \
    fontconfig rsync schedtool lzop bc imagemagick \
    libncurses5 lib32ncurses5-dev
```

#### Debian 10+

```bash
sudo apt-get install -y \
    git-core gnupg flex bison build-essential zip curl \
    zlib1g-dev gcc-multilib g++-multilib libc6-dev-i386 \
    libncurses5 lib32ncurses5-dev x11proto-core-dev \
    libx11-dev lib32z1-dev libgl1-mesa-dev libxml2-utils \
    xsltproc unzip fontconfig rsync ccache libssl-dev \
    schedtool lzop bc imagemagick python-dev
```

#### macOS

```bash
brew install git gnupg coreutils findutils \
    gnu-tar gnu-sed grep make m4 bison \
    libxml2 ccache
```

### 2.4 安装 Repo 工具

```bash
# 创建 bin 目录
mkdir -p ~/bin

# 下载 Repo
curl https://storage.googleapis.com/git-repo-downloads/repo > ~/bin/repo

# 如果访问不了 Google，使用清华镜像
curl https://mirrors.tuna.tsinghua.edu.cn/git/git-repo > ~/bin/repo

# 设置权限
chmod a+x ~/bin/repo

# 添加到 PATH
echo 'export PATH=~/bin:$PATH' >> ~/.bashrc
source ~/.bashrc

# 验证安装
repo version
```

### 2.5 配置 Git

```bash
# 设置用户信息
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# 配置颜色输出
git config --global color.ui auto

# 配置别名
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit

# 配置代理 (如果需要)
# git config --global http.proxy http://proxy.example.com:8080
# git config --global https.proxy https://proxy.example.com:8080

# 查看配置
git config --list
```

---

## 3. Android SDK 配置

### 3.1 下载 Android SDK

#### 方法 1: 使用 Android Studio (推荐)

```bash
# 下载 Android Studio
# 访问: https://developer.android.com/studio

# Ubuntu
wget https://redirector.gvt1.com/edgedl/android/studio/ide-zips/2022.1.1.21/android-studio-2022.1.1.21-linux.tar.gz

# 解压
tar -xzf android-studio-*-linux.tar.gz -C ~/

# 运行
~/android-studio/bin/studio.sh

# SDK 路径通常在:
# ~/Android/Sdk
```

#### 方法 2: 仅安装命令行工具

```bash
# 下载 Command Line Tools
cd ~/
mkdir -p android-sdk/cmdline-tools
cd android-sdk/cmdline-tools

# Linux
wget https://dl.google.com/android/repository/commandlinetools-linux-9477386_latest.zip
unzip commandlinetools-linux-*_latest.zip
mv cmdline-tools latest

# macOS
wget https://dl.google.com/android/repository/commandlinetools-mac-9477386_latest.zip
unzip commandlinetools-mac-*_latest.zip
mv cmdline-tools latest

# 设置环境变量
echo 'export ANDROID_HOME=~/android-sdk' >> ~/.bashrc
echo 'export PATH=$ANDROID_HOME/cmdline-tools/latest/bin:$PATH' >> ~/.bashrc
echo 'export PATH=$ANDROID_HOME/platform-tools:$PATH' >> ~/.bashrc
echo 'export PATH=$ANDROID_HOME/emulator:$PATH' >> ~/.bashrc
source ~/.bashrc
```

### 3.2 安装 SDK 组件

```bash
# 查看可用的 SDK 包
sdkmanager --list

# 接受许可协议
sdkmanager --licenses

# 安装必需的组件
sdkmanager "platform-tools" "platforms;android-27" "build-tools;27.0.3"

# 安装额外组件
sdkmanager "emulator" "system-images;android-27;google_apis;x86_64"
sdkmanager "extras;android;m2repository" "extras;google;m2repository"

# 更新所有已安装的包
sdkmanager --update

# 验证安装
sdkmanager --list_installed
```

### 3.3 安装 Platform Tools (ADB/Fastboot)

```bash
# 如果已安装 SDK，Platform Tools 应该已包含
# 单独下载 Platform Tools:

# Linux
wget https://dl.google.com/android/repository/platform-tools-latest-linux.zip
unzip platform-tools-latest-linux.zip -d ~/
echo 'export PATH=~/platform-tools:$PATH' >> ~/.bashrc

# macOS
wget https://dl.google.com/android/repository/platform-tools-latest-darwin.zip
unzip platform-tools-latest-darwin.zip -d ~/
echo 'export PATH=~/platform-tools:$PATH' >> ~/.zshrc

source ~/.bashrc  # 或 source ~/.zshrc

# 验证安装
adb version
fastboot --version
```

### 3.4 配置 USB 访问权限 (Linux)

```bash
# 创建 udev 规则
sudo nano /etc/udev/rules.d/51-android.rules

# 添加以下内容 (根据设备厂商添加):
# Google
SUBSYSTEM=="usb", ATTR{idVendor}=="18d1", MODE="0666", GROUP="plugdev"
# Samsung
SUBSYSTEM=="usb", ATTR{idVendor}=="04e8", MODE="0666", GROUP="plugdev"
# Xiaomi
SUBSYSTEM=="usb", ATTR{idVendor}=="2717", MODE="0666", GROUP="plugdev"
# OnePlus
SUBSYSTEM=="usb", ATTR{idVendor}=="2a70", MODE="0666", GROUP="plugdev"
# 更多厂商 ID: https://developer.android.com/studio/run/device

# 设置权限
sudo chmod a+r /etc/udev/rules.d/51-android.rules

# 重新加载 udev 规则
sudo udevadm control --reload-rules
sudo udevadm trigger

# 将用户添加到 plugdev 组
sudo usermod -aG plugdev $LOGNAME

# 重启 ADB 服务
adb kill-server
adb start-server

# 测试连接
adb devices
```

---

## 4. 构建工具安装

### 4.1 配置 ccache

```bash
# 安装 ccache
# Ubuntu/Debian
sudo apt-get install ccache

# macOS
brew install ccache

# 设置 ccache 大小 (50-100GB)
ccache -M 50G

# 配置环境变量
echo 'export USE_CCACHE=1' >> ~/.bashrc
echo 'export CCACHE_DIR=~/.ccache' >> ~/.bashrc
source ~/.bashrc

# 查看 ccache 状态
ccache -s

# 清空 ccache
# ccache -C
```

### 4.2 安装 Make 工具

```bash
# Ubuntu/Debian (通常已包含)
sudo apt-get install make

# macOS
xcode-select --install

# 验证版本 (需要 3.81 或 3.82)
make --version
```

### 4.3 安装编译工具链

```bash
# Ubuntu/Debian
sudo apt-get install -y \
    gcc g++ gcc-multilib g++-multilib \
    clang llvm lld

# macOS (Xcode Command Line Tools 已包含)
xcode-select --install

# 验证安装
gcc --version
g++ --version
clang --version
```

---

## 5. IDE 配置

### 5.1 Android Studio 配置

#### 基础配置

```bash
# 启动 Android Studio
~/android-studio/bin/studio.sh

# 首次启动配置:
# 1. 选择 "Standard" 安装类型
# 2. 选择主题 (Darcula/Light)
# 3. 下载 SDK 组件
# 4. 完成设置
```

#### 性能优化

```bash
# 编辑 studio.vmoptions
# Linux: ~/.config/Google/AndroidStudio*/studio64.vmoptions
# macOS: ~/Library/Application Support/Google/AndroidStudio*/studio.vmoptions

# 修改 JVM 参数 (根据内存调整):
-Xms2g
-Xmx8g
-XX:ReservedCodeCacheSize=512m
-XX:+UseG1GC
-XX:SoftRefLRUPolicyMSPerMB=50
-Dsun.io.useCanonCaches=false
-Djdk.http.auth.tunneling.disabledSchemes=""
-Djdk.attach.allowAttachSelf=true
-Dkotlinx.coroutines.debug=off
```

#### 安装插件

推荐插件:
- **ADB Idea**: 快速 ADB 操作
- **Android Parcelable Generator**: 自动生成 Parcelable 代码
- **GsonFormat**: JSON 转 Java 类
- **Rainbow Brackets**: 彩虹括号
- **Key Promoter X**: 快捷键提示

### 5.2 Visual Studio Code 配置

```bash
# 安装 VS Code
# Ubuntu
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg
sudo install -o root -g root -m 644 packages.microsoft.gpg /etc/apt/trusted.gpg.d/
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/code stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt-get update
sudo apt-get install code

# macOS
brew install --cask visual-studio-code

# 推荐扩展:
code --install-extension vscjava.vscode-java-pack
code --install-extension vscjava.vscode-java-debug
code --install-extension redhat.java
code --install-extension GitHub.copilot
code --install-extension ms-python.python
```

---

## 6. 逆向工具安装

### 6.1 APKTool

```bash
# 创建工具目录
mkdir -p ~/android-tools
cd ~/android-tools

# 下载 APKTool
wget https://raw.githubusercontent.com/iBotPeaches/Apktool/master/scripts/linux/apktool
wget https://bitbucket.org/iBotPeaches/apktool/downloads/apktool_2.9.3.jar

# macOS
wget https://raw.githubusercontent.com/iBotPeaches/Apktool/master/scripts/osx/apktool

# 设置权限
mv apktool_2.9.3.jar apktool.jar
chmod +x apktool apktool.jar

# 安装到系统
sudo cp apktool.jar /usr/local/bin/
sudo cp apktool /usr/local/bin/

# 验证
apktool --version
```

### 6.2 dex2jar

```bash
cd ~/android-tools
wget https://github.com/pxb1988/dex2jar/releases/download/v2.4/dex2jar-2.4.zip
unzip dex2jar-2.4.zip
chmod +x dex2jar-2.4/*.sh

echo 'export PATH=$PATH:~/android-tools/dex2jar-2.4' >> ~/.bashrc
source ~/.bashrc

# 验证
d2j-dex2jar.sh --help
```

### 6.3 JADX

```bash
cd ~/android-tools
wget https://github.com/skylot/jadx/releases/download/v1.4.7/jadx-1.4.7.zip
unzip jadx-1.4.7.zip -d jadx

echo 'export PATH=$PATH:~/android-tools/jadx/bin' >> ~/.bashrc
source ~/.bashrc

# 验证
jadx --version

# 启动 GUI
jadx-gui
```

### 6.4 Frida

```bash
# 安装 Frida 工具
pip3 install frida-tools

# 验证
frida --version

# 安装 objection
pip3 install objection

# 验证
objection --help
```

### 6.5 其他工具

```bash
# JD-GUI (Java 反编译器)
cd ~/android-tools
wget https://github.com/java-decompiler/jd-gui/releases/download/v1.6.6/jd-gui-1.6.6.jar

# 运行
java -jar jd-gui-1.6.6.jar

# Ghidra (逆向工程平台)
cd ~/android-tools
wget https://github.com/NationalSecurityAgency/ghidra/releases/download/Ghidra_10.4_build/ghidra_10.4_PUBLIC_20230928.zip
unzip ghidra_10.4_PUBLIC_*.zip

# 运行
~/android-tools/ghidra_10.4_PUBLIC/ghidraRun

# Burp Suite (抓包工具)
# 下载社区版: https://portswigger.net/burp/communitydownload
```

---

## 7. 环境验证

### 7.1 验证构建环境

```bash
#!/bin/bash
# check_build_env.sh - 验证构建环境

echo "======== Android Build Environment Check ========"

# Java
echo -n "Java: "
if command -v java &> /dev/null; then
    java -version 2>&1 | head -n 1
else
    echo "NOT INSTALLED"
fi

# Python
echo -n "Python 2: "
python2 --version 2>&1 || echo "NOT INSTALLED"
echo -n "Python 3: "
python3 --version 2>&1 || echo "NOT INSTALLED"

# Make
echo -n "Make: "
make --version 2>&1 | head -n 1 || echo "NOT INSTALLED"

# Git
echo -n "Git: "
git --version 2>&1 || echo "NOT INSTALLED"

# Repo
echo -n "Repo: "
repo version 2>&1 | head -n 1 || echo "NOT INSTALLED"

# ADB
echo -n "ADB: "
adb version 2>&1 | head -n 1 || echo "NOT INSTALLED"

# Fastboot
echo -n "Fastboot: "
fastboot --version 2>&1 | head -n 1 || echo "NOT INSTALLED"

# ccache
echo -n "ccache: "
ccache --version 2>&1 | head -n 1 || echo "NOT INSTALLED"

# 磁盘空间
echo ""
echo "======== Disk Space ========"
df -h | grep -E "Filesystem|/$"

# 内存
echo ""
echo "======== Memory ========"
free -h

echo ""
echo "======== Environment Variables ========"
echo "JAVA_HOME: ${JAVA_HOME:-NOT SET}"
echo "ANDROID_HOME: ${ANDROID_HOME:-NOT SET}"
echo "USE_CCACHE: ${USE_CCACHE:-NOT SET}"
echo "PATH: $PATH"

echo ""
echo "======== Check Complete ========"
```

运行验证脚本:

```bash
chmod +x check_build_env.sh
./check_build_env.sh
```

### 7.2 测试 ADB 连接

```bash
# 启动 ADB 服务
adb start-server

# 检查设备连接
adb devices

# 应该看到类似输出:
# List of devices attached
# ABC123456789    device

# 测试命令
adb shell getprop ro.build.version.release  # Android 版本
adb shell getprop ro.product.model          # 设备型号
```

### 7.3 测试构建工具

```bash
# 创建测试项目
cd ~/test
repo init -u https://android.googlesource.com/platform/manifest -b android-build-tools
repo sync

# 初始化环境
source build/envsetup.sh

# 选择目标
lunch aosp_arm-eng

# 测试编译 (不实际构建)
make help
```

---

## 8. 常见问题

### 8.1 Java 版本问题

**问题**: 编译时报错 "Unsupported Java version"

**解决方案**:
```bash
# 确保使用 Java 8
sudo update-alternatives --config java
# 选择 java-8-openjdk

# 设置 JAVA_HOME
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
```

### 8.2 Repo 同步失败

**问题**: `repo sync` 失败或速度慢

**解决方案**:
```bash
# 使用清华镜像
export REPO_URL='https://mirrors.tuna.tsinghua.edu.cn/git/git-repo'
repo init -u https://mirrors.tuna.tsinghua.edu.cn/git/AOSP/platform/manifest -b android-8.1.0_r81

# 减少并行数
repo sync -j1

# 强制同步
repo sync --force-sync
```

### 8.3 ADB 设备未找到

**问题**: `adb devices` 显示 "unauthorized" 或不显示设备

**解决方案**:
```bash
# 1. 检查 USB 调试是否启用
# 设置 -> 开发者选项 -> USB 调试

# 2. 重启 ADB 服务
adb kill-server
adb start-server

# 3. 检查 udev 规则 (Linux)
lsusb  # 查看设备 Vendor ID
sudo nano /etc/udev/rules.d/51-android.rules
# 添加对应规则

# 4. 授权设备
# 在设备上点击 "允许 USB 调试"
```

### 8.4 磁盘空间不足

**问题**: 编译过程中提示空间不足

**解决方案**:
```bash
# 清理 ccache
ccache -C

# 清理构建输出
cd ~/aosp/android-8.1
make clean

# 删除 .repo 目录中的缓存
rm -rf .repo/repo
rm -rf .repo/project-objects

# 使用 du 查找大文件
du -sh * | sort -hr | head -20
```

### 8.5 编译内存不足

**问题**: 编译时出现 "Out of memory" 错误

**解决方案**:
```bash
# 减少并行编译数
make -j4  # 而不是 -j$(nproc)

# 增加交换空间
sudo fallocate -l 16G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile

# 查看内存使用
free -h
```

---

## 附录

### A. 环境变量完整配置

```bash
# ~/.bashrc 或 ~/.zshrc

# Java
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
export PATH=$JAVA_HOME/bin:$PATH

# Android SDK
export ANDROID_HOME=~/android-sdk
export PATH=$ANDROID_HOME/cmdline-tools/latest/bin:$PATH
export PATH=$ANDROID_HOME/platform-tools:$PATH
export PATH=$ANDROID_HOME/emulator:$PATH
export PATH=$ANDROID_HOME/tools:$PATH
export PATH=$ANDROID_HOME/tools/bin:$PATH

# AOSP Build
export USE_CCACHE=1
export CCACHE_DIR=~/.ccache
export ANDROID_JACK_VM_ARGS="-Xmx4g -Dfile.encoding=UTF-8 -XX:+TieredCompilation"

# Tools
export PATH=~/bin:$PATH
export PATH=~/android-tools/dex2jar-2.4:$PATH
export PATH=~/android-tools/jadx/bin:$PATH

# Proxy (如果需要)
# export http_proxy=http://proxy.example.com:8080
# export https_proxy=http://proxy.example.com:8080
```

### B. 快速安装脚本

```bash
#!/bin/bash
# install_android_env.sh - 快速安装 Android 开发环境

set -e

echo "======== Installing Android Development Environment ========"

# 更新系统
sudo apt-get update

# 安装基础依赖
sudo apt-get install -y \
    git-core gnupg flex bison gperf build-essential \
    zip curl zlib1g-dev gcc-multilib g++-multilib \
    libc6-dev-i386 lib32ncurses5-dev x11proto-core-dev \
    libx11-dev lib32z-dev libgl1-mesa-dev libxml2-utils \
    xsltproc unzip fontconfig openjdk-8-jdk python-dev \
    python3-dev rsync schedtool ccache lzop bc libssl-dev \
    imagemagick android-tools-adb android-tools-fastboot

# 安装 Repo
mkdir -p ~/bin
curl https://storage.googleapis.com/git-repo-downloads/repo > ~/bin/repo
chmod a+x ~/bin/repo

# 配置环境变量
cat >> ~/.bashrc << 'EOF'
export PATH=~/bin:$PATH
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
export USE_CCACHE=1
export CCACHE_DIR=~/.ccache
EOF

# 配置 ccache
ccache -M 50G

# 配置 USB 规则
sudo tee /etc/udev/rules.d/51-android.rules > /dev/null << 'EOF'
SUBSYSTEM=="usb", ATTR{idVendor}=="18d1", MODE="0666", GROUP="plugdev"
SUBSYSTEM=="usb", ATTR{idVendor}=="04e8", MODE="0666", GROUP="plugdev"
SUBSYSTEM=="usb", ATTR{idVendor}=="2717", MODE="0666", GROUP="plugdev"
EOF

sudo chmod a+r /etc/udev/rules.d/51-android.rules
sudo udevadm control --reload-rules
sudo usermod -aG plugdev $LOGNAME

echo ""
echo "======== Installation Complete ========"
echo "Please run: source ~/.bashrc"
echo "Then restart your computer for USB rules to take effect"
```

### C. 工具下载镜像源

```bash
# 中国用户推荐使用镜像源

# AOSP 源码 - 清华大学镜像
https://mirrors.tuna.tsinghua.edu.cn/help/AOSP/

# Android SDK - 阿里云镜像
https://mirrors.aliyun.com/android.googlesource.com/

# Maven - 阿里云镜像
https://maven.aliyun.com/mvn/guide

# Gradle - 腾讯云镜像
https://mirrors.cloud.tencent.com/gradle/
```

---

## 总结

本指南详细介绍了 Android 8.1 开发和逆向工程环境的完整配置过程。通过以上步骤，您应该已经拥有：

✅ 完整的系统开发环境  
✅ Android SDK 和构建工具  
✅ 逆向工程工具集  
✅ IDE 和开发工具  
✅ 优化的编译环境  

**下一步**: 查看 [Android 8.1 AOSP 构建指南](./android-8.1-build-guide.md) 开始构建您的定制系统！

**祝您配置顺利！** 🎉
