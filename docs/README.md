# Android 8.1 定制构建与逆向工程文档中心

> **项目**: Android 8.1 Custom Build & Reverse Engineering  
> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **作者**: Android Development Team

---

## 📚 文档导航

本文档集提供了 Android 8.1 系统定制构建和逆向工程的完整指南，从零基础到高级实战。

### 🎯 核心文档

| 文档 | 说明 | 适合人群 |
|------|------|---------|
| [**环境配置指南**](./environment-setup.md) | 开发环境搭建和工具安装 | 所有开发者（必读） |
| [**AOSP 构建指南**](./android-8.1-build-guide.md) | 完整的 Android 8.1 编译流程 | 系统开发者 |
| [**逆向工程工具**](./reverse-engineering-tools.md) | APK 分析、反编译、动态调试 | 安全研究员、逆向工程师 |
| [**设备刷机调试**](./device-flashing-debug.md) | 刷机方法、ADB 调试技巧 | ROM 玩家、开发者 |
| [**FAQ 故障排除**](./faq-troubleshooting.md) | 常见问题和解决方案 | 所有用户（遇到问题必看） |

---

## 🚀 快速开始

### 新手入门路线

```
1. 环境配置
   ↓
2. 下载 AOSP 源码
   ↓
3. 编译第一个 ROM
   ↓
4. 刷入设备测试
   ↓
5. 学习定制修改
```

**推荐阅读顺序**:
1. [环境配置指南](./environment-setup.md) - 搭建开发环境
2. [AOSP 构建指南](./android-8.1-build-guide.md) - 编译系统
3. [设备刷机调试](./device-flashing-debug.md) - 刷机和测试
4. [FAQ 故障排除](./faq-troubleshooting.md) - 解决问题

### 逆向工程路线

```
1. 工具安装
   ↓
2. APK 基础分析
   ↓
3. 静态分析 (反编译)
   ↓
4. 动态分析 (Hook)
   ↓
5. 高级逆向技术
```

**推荐阅读顺序**:
1. [环境配置指南](./environment-setup.md) - 安装逆向工具
2. [逆向工程工具](./reverse-engineering-tools.md) - 掌握核心技术
3. [设备刷机调试](./device-flashing-debug.md) - ADB 调试技巧
4. [FAQ 故障排除](./faq-troubleshooting.md) - 逆向相关问题

---

## 📖 文档详细介绍

### 1. 环境配置指南

**[environment-setup.md](./environment-setup.md)**

完整的开发环境搭建教程，包括：

- ✅ 系统环境准备 (Ubuntu/Debian/macOS)
- ✅ Java、Python、构建工具安装
- ✅ Android SDK 配置
- ✅ 逆向工程工具安装
- ✅ IDE 配置 (Android Studio, VS Code)
- ✅ 环境验证和测试

**适合**: 所有开发者，必读第一步

**时长**: 1-2 小时

---

### 2. Android 8.1 AOSP 构建指南

**[android-8.1-build-guide.md](./android-8.1-build-guide.md)**

从零开始编译 Android 8.1 系统，包括：

- 📦 源码下载 (使用 Repo)
- ⚙️ 构建配置 (lunch 选择)
- 🔨 编译构建 (make 命令)
- 🎨 系统定制 (移除/添加应用)
- 📝 修改系统属性
- 🎬 自定义启动动画
- 🔧 内核定制
- 🚀 刷机部署

**适合**: 想要编译 AOSP、定制 ROM 的开发者

**时长**: 首次编译 4-8 小时 (根据硬件)

**硬件要求**:
- CPU: 8 核心+
- RAM: 32 GB+
- 存储: 500 GB SSD

---

### 3. 逆向工程工具完全指南

**[reverse-engineering-tools.md](./reverse-engineering-tools.md)**

全面的 Android 逆向工程教程，包括：

- 🔍 APK 结构分析
- 🛠️ 核心工具使用
  - APKTool (反编译/重打包)
  - JADX (Java 反编译)
  - dex2jar + JD-GUI
  - Frida (动态插桩)
  - Xposed Framework
- 📱 动态调试技术
- 🎯 代码注入与 Hook
- 🔐 加密与脱壳
- 💡 实战案例
  - 移除广告
  - 绕过验证
  - 提取密钥

**适合**: 安全研究员、逆向工程师、好奇的开发者

**⚠️ 法律提醒**: 仅用于学习和研究，请遵守法律法规

---

### 4. 设备刷机与调试指南

**[device-flashing-debug.md](./device-flashing-debug.md)**

完整的刷机和调试教程，包括：

- 🔓 Bootloader 解锁
- 💾 数据备份
- 🔄 Recovery 刷入 (TWRP)
- 📲 刷机方法
  - Fastboot 刷机
  - Recovery 刷机
  - ADB Sideload
- 🌐 刷入 GApps
- 🔑 获取 Root (Magisk)
- 🐛 ADB 调试技巧
- 📊 Logcat 日志分析
- 🔧 系统诊断 (dumpsys)
- ❗ 问题排查 (Bootloop, 变砖)

**适合**: ROM 玩家、开发者、想要刷机的用户

**⚠️ 警告**: 刷机有风险，请务必备份数据

---

### 5. FAQ 和故障排除

**[faq-troubleshooting.md](./faq-troubleshooting.md)**

最全面的问题解决方案集，包括：

- 🐛 编译构建问题
  - Java 版本错误
  - 内存不足
  - 磁盘空间不足
  - 权限错误
- 🌐 源码同步问题
  - repo sync 失败
  - 网络中断
  - 代理配置
- 📱 刷机问题
  - Bootloader 解锁失败
  - Bootloop (循环重启)
  - 功能异常
  - Fastboot 无法识别
- 🐞 调试问题
  - ADB 连接问题
  - Logcat 乱码
  - Wi-Fi ADB
- 🔍 逆向工程问题
  - APKTool 失败
  - 重打包问题
  - Frida Hook 失败
  - SSL Pinning 绕过
- ⚡ 性能优化
- 🚀 进阶问题

**适合**: 所有用户 (遇到问题时必看)

**提示**: 使用 `Ctrl+F` 搜索关键词快速定位问题

---

## 🛠️ 工具清单

### 必备工具

| 工具 | 用途 | 安装 |
|------|------|------|
| **Java 8** | AOSP 编译 | `apt-get install openjdk-8-jdk` |
| **Python 2/3** | 构建脚本 | `apt-get install python python3` |
| **Repo** | 源码管理 | 见[环境配置](./environment-setup.md) |
| **ADB/Fastboot** | 设备调试 | `apt-get install android-tools-adb` |
| **ccache** | 编译加速 | `apt-get install ccache` |

### 逆向工具

| 工具 | 用途 | 官网 |
|------|------|------|
| **APKTool** | APK 反编译/重打包 | [ibotpeaches.github.io/Apktool](https://ibotpeaches.github.io/Apktool/) |
| **JADX** | DEX 转 Java | [github.com/skylot/jadx](https://github.com/skylot/jadx) |
| **dex2jar** | DEX 转 JAR | [github.com/pxb1988/dex2jar](https://github.com/pxb1988/dex2jar) |
| **Frida** | 动态插桩 | [frida.re](https://frida.re/) |
| **JD-GUI** | Java 反编译查看 | [java-decompiler.github.io](http://java-decompiler.github.io/) |

### 刷机工具

| 工具 | 用途 | 下载 |
|------|------|------|
| **TWRP** | 第三方 Recovery | [twrp.me](https://twrp.me/) |
| **Magisk** | Root 管理 | [github.com/topjohnwu/Magisk](https://github.com/topjohnwu/Magisk) |
| **Open GApps** | Google 服务 | [opengapps.org](https://opengapps.org/) |

---

## 📊 系统要求

### 最低配置

| 组件 | 要求 |
|------|------|
| **CPU** | 4 核心 |
| **内存** | 16 GB |
| **存储** | 250 GB |
| **系统** | Ubuntu 18.04+ / Debian 10+ |

### 推荐配置

| 组件 | 要求 |
|------|------|
| **CPU** | 8 核心+ (i7/Ryzen 7) |
| **内存** | 32 GB+ |
| **存储** | 500 GB SSD |
| **系统** | Ubuntu 20.04 LTS |

**编译时间参考** (AOSP 完整编译):
- 低配 (4核/16GB): 8-12 小时
- 推荐配置: 2-4 小时
- 高配 (16核/64GB): 1-2 小时

---

## 🎓 学习路径

### 初级 (0-1 个月)

**目标**: 成功编译并刷入第一个 ROM

1. ✅ 搭建开发环境
2. ✅ 下载 AOSP 源码
3. ✅ 编译通用版本 (aosp_arm64-eng)
4. ✅ 解锁 Bootloader
5. ✅ 刷入编译的系统
6. ✅ 学习 ADB 基础命令

**学习资源**:
- [环境配置指南](./environment-setup.md)
- [AOSP 构建指南](./android-8.1-build-guide.md) (第 1-6 章)
- [设备刷机调试](./device-flashing-debug.md) (第 1-4 章)

### 中级 (1-3 个月)

**目标**: 能够定制 ROM 并解决常见问题

1. ✅ 修改系统应用和属性
2. ✅ 自定义启动动画
3. ✅ 添加/移除系统功能
4. ✅ 使用 Logcat 分析问题
5. ✅ 学习 APK 反编译基础
6. ✅ 使用 Frida 进行简单 Hook

**学习资源**:
- [AOSP 构建指南](./android-8.1-build-guide.md) (第 7 章)
- [逆向工程工具](./reverse-engineering-tools.md) (第 1-4 章)
- [FAQ 故障排除](./faq-troubleshooting.md)

### 高级 (3-6 个月)

**目标**: 深度定制系统和高级逆向分析

1. ✅ 移植 ROM 到新设备
2. ✅ 编译和定制内核
3. ✅ 添加自定义系统服务
4. ✅ 高级逆向技术 (脱壳、反混淆)
5. ✅ Native 代码分析 (SO 库)
6. ✅ 绕过各种安全检测

**学习资源**:
- [AOSP 构建指南](./android-8.1-build-guide.md) (进阶部分)
- [逆向工程工具](./reverse-engineering-tools.md) (第 5-8 章)
- [FAQ 故障排除](./faq-troubleshooting.md) (第 8 章)

---

## 🌟 实战项目

### 项目 1: 编译纯净 AOSP

**难度**: ⭐

**目标**: 编译官方 AOSP 并刷入设备

**步骤**:
1. 按照[环境配置指南](./environment-setup.md)搭建环境
2. 下载 Android 8.1 源码
3. 编译 aosp_arm64-eng
4. 刷入设备测试

### 项目 2: 定制精简 ROM

**难度**: ⭐⭐

**目标**: 移除系统应用，精简系统

**步骤**:
1. 基于 AOSP 源码
2. 移除不需要的系统应用
3. 添加自定义启动动画
4. 修改系统属性
5. 编译并测试

### 项目 3: 移除应用广告

**难度**: ⭐⭐

**目标**: 使用逆向技术移除 APK 中的广告

**步骤**:
1. 使用 APKTool 反编译
2. 使用 JADX 分析代码
3. 找到广告模块并移除
4. 重打包签名
5. 测试验证

### 项目 4: 绕过应用验证

**难度**: ⭐⭐⭐

**目标**: 使用 Frida Hook 绕过会员验证

**步骤**:
1. 使用 JADX 找到验证函数
2. 编写 Frida Hook 脚本
3. 修改返回值
4. 测试功能是否解锁

### 项目 5: 移植 LineageOS

**难度**: ⭐⭐⭐⭐

**目标**: 将 LineageOS 移植到新设备

**步骤**:
1. 创建设备树 (Device Tree)
2. 提取专有文件
3. 配置内核
4. 编译测试
5. 解决兼容性问题

---

## 🔗 相关资源

### 官方文档

- [Android 开源项目](https://source.android.com/)
- [Android 开发者文档](https://developer.android.com/)
- [AOSP 镜像站 (清华)](https://mirrors.tuna.tsinghua.edu.cn/help/AOSP/)

### 社区论坛

- [XDA Developers](https://forum.xda-developers.com/)
- [Reddit - r/Android](https://www.reddit.com/r/Android/)
- [Reddit - r/ReverseEngineering](https://www.reddit.com/r/ReverseEngineering/)
- [看雪论坛](https://bbs.pediy.com/) (中文)
- [52pojie 吾爱破解](https://www.52pojie.cn/) (中文)

### 开源项目

- [LineageOS](https://lineageos.org/) - 最流行的第三方 ROM
- [Magisk](https://github.com/topjohnwu/Magisk) - Root 解决方案
- [Xposed Framework](https://repo.xposed.info/) - Hook 框架
- [Frida](https://github.com/frida/frida) - 动态插桩工具

### 学习资源

- [Android Internals](http://newandroidbook.com/) - 深入理解 Android
- [Android Security Awesome](https://github.com/ashishb/android-security-awesome)
- [OWASP Mobile Security](https://owasp.org/www-project-mobile-security/)

---

## ⚠️ 免责声明

1. **教育目的**: 本文档仅供学习和研究使用
2. **版权声明**: 请尊重软件版权，不要用于非法用途
3. **风险提示**: 
   - 刷机有风险，可能导致设备变砖
   - 解锁 Bootloader 会清除数据并失去保修
   - 逆向工程请遵守当地法律法规
4. **责任**: 使用本文档造成的任何损失，作者不承担责任

---

## 📝 更新日志

### v1.0.0 (2025-01-03)

**新增**:
- ✅ 完整的环境配置指南
- ✅ Android 8.1 AOSP 构建教程
- ✅ 逆向工程工具使用指南
- ✅ 设备刷机和调试教程
- ✅ FAQ 和故障排除文档

---

## 📧 反馈与贡献

如果您发现文档中的错误或有改进建议，欢迎：

- 提交 Issue
- 发起 Pull Request
- 联系作者

---

## 📜 许可证

本文档采用 [MIT License](../LICENSE) 开源协议。

---

**祝您学习愉快，构建顺利！** 🚀

**Happy Building & Hacking!** 💻✨
