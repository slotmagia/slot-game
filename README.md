# Android 8.1 定制构建与逆向工程系统

> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **作者**: Android Development Team

这是一个完整的 Android 8.1 系统定制构建和逆向工程文档项目。项目提供了从零开始编译 AOSP、定制 ROM、刷机部署到逆向分析的完整指南。

---

## 🎯 项目概述

本项目包含两大核心内容：

### 1. Android 8.1 AOSP 定制构建
- ✅ 完整的 AOSP 8.1 编译流程
- ✅ 系统定制和 ROM 修改
- ✅ 设备刷机和部署
- ✅ 内核定制和优化

### 2. Android 逆向工程
- ✅ APK 反编译和分析
- ✅ 动态调试和 Hook 技术
- ✅ 加密解密和脱壳
- ✅ 安全研究和漏洞分析

---

## 📚 文档中心

所有详细文档位于 [`docs/`](./docs/) 目录：

| 文档 | 说明 |
|------|------|
| [**📖 文档导航**](./docs/README.md) | 文档总览和学习路径 |
| [**⚙️ 环境配置**](./docs/environment-setup.md) | 开发环境搭建（必读第一步） |
| [**🔨 AOSP 构建**](./docs/android-8.1-build-guide.md) | Android 8.1 完整编译流程 |
| [**🔍 逆向工具**](./docs/reverse-engineering-tools.md) | APK 分析、反编译、Hook |
| [**📱 刷机调试**](./docs/device-flashing-debug.md) | 设备刷机、ADB 调试技巧 |
| [**❓ FAQ**](./docs/faq-troubleshooting.md) | 常见问题和故障排除 |

---

## 🚀 快速开始

### 新手推荐路线

```
1️⃣ 阅读文档导航 (5 分钟)
   ↓
2️⃣ 搭建开发环境 (1-2 小时)
   ↓
3️⃣ 下载 AOSP 源码 (4-24 小时，取决于网络)
   ↓
4️⃣ 编译第一个 ROM (2-8 小时，取决于硬件)
   ↓
5️⃣ 刷入设备测试 (30 分钟)
```

### 快速命令

```bash
# 1. 查看文档
cd docs/
cat README.md

# 2. 环境检查 (如果已有 GoFrame 环境)
java -version    # 需要 Java 8
python3 --version
adb version

# 3. 开始学习
# 按照 docs/README.md 中的学习路径进行
```

---

## 💻 系统要求

### 最低配置
- **CPU**: 4 核心
- **内存**: 16 GB RAM
- **存储**: 250 GB (推荐 SSD)
- **系统**: Ubuntu 18.04+ / Debian 10+

### 推荐配置
- **CPU**: 8 核心+ (Intel i7/AMD Ryzen 7)
- **内存**: 32 GB+ RAM
- **存储**: 500 GB SSD
- **系统**: Ubuntu 20.04 LTS

### 编译时间参考
- 低配 (4核/16GB): 8-12 小时
- 推荐配置: 2-4 小时
- 高配 (16核/64GB): 1-2 小时

---

## 🛠️ 核心工具

### AOSP 构建工具
- Java 8
- Python 2.7 & 3.x
- Repo 工具
- Android SDK
- Make & ccache

### 逆向工程工具
- APKTool
- JADX
- dex2jar
- Frida
- Xposed Framework

### 刷机工具
- ADB & Fastboot
- TWRP Recovery
- Magisk (Root)

**详细安装说明**: 见 [环境配置指南](./docs/environment-setup.md)

---

## 📖 学习路径

### 🌟 初级 (0-1 个月)
**目标**: 成功编译并刷入第一个 ROM

- [x] 搭建开发环境
- [x] 下载 AOSP 源码
- [x] 编译系统
- [x] 解锁 Bootloader
- [x] 刷入测试

### 🌟🌟 中级 (1-3 个月)
**目标**: 能够定制 ROM 并进行基础逆向

- [x] 修改系统应用和属性
- [x] 自定义启动动画
- [x] 学习 ADB 和 Logcat
- [x] APK 反编译基础
- [x] 简单的 Frida Hook

### 🌟🌟🌟 高级 (3-6 个月)
**目标**: 深度定制和高级逆向分析

- [x] 移植 ROM 到新设备
- [x] 编译定制内核
- [x] 添加系统服务
- [x] 高级逆向 (脱壳、反混淆)
- [x] Native 代码分析

---

## 🎓 实战项目

### 项目 1: 编译纯净 AOSP ⭐
编译官方 Android 8.1 源码并刷入设备

### 项目 2: 定制精简 ROM ⭐⭐
移除系统应用，添加自定义功能

### 项目 3: 移除应用广告 ⭐⭐
使用逆向技术移除 APK 广告模块

### 项目 4: 绕过应用验证 ⭐⭐⭐
使用 Frida Hook 绕过会员验证

### 项目 5: 移植 LineageOS ⭐⭐⭐⭐
将第三方 ROM 移植到新设备

**详细步骤**: 见 [文档导航](./docs/README.md)

---

## 📂 项目结构

```
android-8.1-custom-build/
├── docs/                           # 📚 完整文档
│   ├── README.md                   # 文档导航中心
│   ├── environment-setup.md        # 环境配置指南
│   ├── android-8.1-build-guide.md  # AOSP 构建教程
│   ├── reverse-engineering-tools.md # 逆向工程指南
│   ├── device-flashing-debug.md    # 刷机调试教程
│   └── faq-troubleshooting.md      # FAQ 故障排除
│
├── config/                         # 配置文件
│   └── android-build.yaml          # Android 构建配置示例
│
├── internal/model/                 # 数据模型 (预留)
│   ├── entity/                     # 数据实体
│   ├── input/                      # 输入模型
│   └── output/                     # 输出模型
│
├── README.md                       # 本文件
└── LICENSE                         # MIT 开源协议
```

---

## 🔗 相关资源

### 官方文档
- [Android 开源项目 (AOSP)](https://source.android.com/)
- [Android 开发者文档](https://developer.android.com/)
- [AOSP 镜像源 (清华)](https://mirrors.tuna.tsinghua.edu.cn/help/AOSP/)

### 社区论坛
- [XDA Developers](https://forum.xda-developers.com/)
- [Reddit - r/Android](https://www.reddit.com/r/Android/)
- [看雪论坛](https://bbs.pediy.com/) (中文)

### 开源项目
- [LineageOS](https://lineageos.org/)
- [Magisk](https://github.com/topjohnwu/Magisk)
- [Frida](https://github.com/frida/frida)

---

## ⚠️ 重要提醒

### 法律声明
1. **仅供学习**: 本项目文档仅用于学习和研究
2. **尊重版权**: 不要用于破解商业软件
3. **遵守法律**: 请遵守当地法律法规
4. **负责任**: 发现安全漏洞应负责任地报告

### 风险提示
1. **刷机风险**: 可能导致设备变砖，请务必备份数据
2. **保修失效**: 解锁 Bootloader 会失去保修
3. **数据丢失**: 解锁和刷机会清除所有数据

---

## 🤝 贡献与反馈

欢迎提交 Issue 和 Pull Request！

如果本项目对您有帮助，请给个 ⭐ Star！

---

## 📜 许可证

本项目采用 [MIT License](./LICENSE) 开源协议。

---

## 📧 联系方式

- **问题反馈**: 提交 GitHub Issue
- **文档贡献**: 发起 Pull Request

---

**🎉 祝您学习愉快，构建顺利！**

**Happy Building & Happy Hacking!** 💻✨

---

## 附录: GoFrame 原项目信息

本项目基于 GoFrame v2 框架搭建，保留了原有的 Web 服务能力。如需开发 Android 构建管理系统的 Web API，可以扩展 `internal/` 目录下的代码。

### 技术栈
- **框架**: GoFrame v2.9.1
- **语言**: Go 1.24.2
- **配置**: YAML
- **文档**: Markdown

### 运行 Web 服务 (可选)
```bash
go mod tidy
go run main.go
# 服务器将在 http://localhost:8080 启动
```