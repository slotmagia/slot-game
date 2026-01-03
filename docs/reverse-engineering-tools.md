# Android 逆向工程工具完全指南

> **作者**: Android Security Team  
> **版本**: 1.0.0  
> **更新时间**: 2025-01-03  
> **适用版本**: Android 8.1 及其他版本

---

## 📋 目录

- [1. 逆向工程概述](#1-逆向工程概述)
- [2. 工具安装配置](#2-工具安装配置)
- [3. APK 结构分析](#3-apk-结构分析)
- [4. 反编译工具](#4-反编译工具)
- [5. 动态调试](#5-动态调试)
- [6. 代码注入与Hook](#6-代码注入与hook)
- [7. 加密与脱壳](#7-加密与脱壳)
- [8. 实战案例](#8-实战案例)
- [9. 最佳实践](#9-最佳实践)

---

## 1. 逆向工程概述

### 1.1 什么是逆向工程

Android 逆向工程是指对 Android 应用进行反向分析的过程，目的包括：

- **安全审计**: 发现应用中的安全漏洞
- **学习研究**: 理解应用的实现机制
- **破解分析**: 移除限制、广告等（仅供学习）
- **兼容性测试**: 分析应用与系统的交互
- **恶意软件分析**: 识别和分析恶意行为

### 1.2 逆向工程工作流程

```
APK 获取 → 解包分析 → 反编译 → 代码分析 → 动态调试 → 修改重打包
```

### 1.3 法律与道德

⚠️ **重要提醒**:
- 仅对您拥有权限的应用进行逆向分析
- 不要用于非法用途
- 尊重软件版权和知识产权
- 遵守当地法律法规

---

## 2. 工具安装配置

### 2.1 核心工具列表

| 工具 | 用途 | 官网/下载 |
|------|------|----------|
| **APKTool** | APK 反编译/重打包 | https://ibotpeaches.github.io/Apktool/ |
| **dex2jar** | DEX 转 JAR | https://github.com/pxb1988/dex2jar |
| **JD-GUI** | Java 反编译查看器 | http://java-decompiler.github.io/ |
| **JADX** | DEX 直接反编译 | https://github.com/skylot/jadx |
| **Frida** | 动态插桩框架 | https://frida.re/ |
| **objection** | Frida 增强工具 | https://github.com/sensepost/objection |
| **Android Studio** | 开发和调试 | https://developer.android.com/studio |
| **Ghidra** | 逆向工程平台 | https://ghidra-sre.org/ |

### 2.2 安装 APKTool

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install -y openjdk-8-jdk

# 下载 APKTool
cd ~/tools
wget https://raw.githubusercontent.com/iBotPeaches/Apktool/master/scripts/linux/apktool
wget https://bitbucket.org/iBotPeaches/apktool/downloads/apktool_2.9.3.jar

# 重命名并设置权限
mv apktool_2.9.3.jar apktool.jar
chmod +x apktool
chmod +x apktool.jar

# 移动到系统路径
sudo mv apktool.jar /usr/local/bin/
sudo mv apktool /usr/local/bin/

# 验证安装
apktool --version
```

### 2.3 安装 dex2jar

```bash
# 下载 dex2jar
cd ~/tools
wget https://github.com/pxb1988/dex2jar/releases/download/v2.4/dex2jar-2.4.zip

# 解压
unzip dex2jar-2.4.zip
cd dex2jar-2.4

# 设置权限
chmod +x d2j-*.sh

# 添加到 PATH
echo 'export PATH=$PATH:~/tools/dex2jar-2.4' >> ~/.bashrc
source ~/.bashrc

# 验证安装
d2j-dex2jar.sh --help
```

### 2.4 安装 JADX

```bash
# 下载 JADX
cd ~/tools
wget https://github.com/skylot/jadx/releases/download/v1.4.7/jadx-1.4.7.zip

# 解压
unzip jadx-1.4.7.zip -d jadx

# 添加到 PATH
echo 'export PATH=$PATH:~/tools/jadx/bin' >> ~/.bashrc
source ~/.bashrc

# 验证安装
jadx --version

# 或使用图形界面
jadx-gui
```

### 2.5 安装 Frida

```bash
# 使用 pip 安装 Frida 工具
pip3 install frida-tools

# 验证安装
frida --version

# 下载 Frida Server (for Android)
# 访问 https://github.com/frida/frida/releases
# 下载对应架构的 frida-server

# 示例: ARM64
wget https://github.com/frida/frida/releases/download/16.1.4/frida-server-16.1.4-android-arm64.xz
xz -d frida-server-16.1.4-android-arm64.xz
mv frida-server-16.1.4-android-arm64 frida-server

# 推送到设备
adb push frida-server /data/local/tmp/
adb shell "chmod 755 /data/local/tmp/frida-server"

# 运行 Frida Server (需要 root)
adb shell "su -c '/data/local/tmp/frida-server &'"
```

### 2.6 配置 Android 调试环境

```bash
# 安装 Android SDK Platform Tools
sudo apt-get install android-tools-adb android-tools-fastboot

# 或从官方下载
wget https://dl.google.com/android/repository/platform-tools-latest-linux.zip
unzip platform-tools-latest-linux.zip
echo 'export PATH=$PATH:~/platform-tools' >> ~/.bashrc
source ~/.bashrc

# 配置 ADB 设备连接
adb devices

# 启用 USB 调试 (在设备上)
# 设置 -> 关于手机 -> 连续点击版本号 7 次
# 设置 -> 开发者选项 -> 启用 USB 调试
```

---

## 3. APK 结构分析

### 3.1 APK 文件结构

APK (Android Package) 本质上是一个 ZIP 压缩包：

```
app.apk
├── AndroidManifest.xml      # 应用清单文件（二进制XML）
├── classes.dex               # Dalvik 字节码
├── classes2.dex              # 额外的 DEX 文件（如果存在）
├── resources.arsc            # 编译后的资源文件
├── res/                      # 资源目录
│   ├── drawable/             # 图片资源
│   ├── layout/               # 布局文件（二进制XML）
│   ├── values/               # 值资源
│   └── ...
├── assets/                   # 资产文件（原始文件）
├── lib/                      # 本地库（.so 文件）
│   ├── armeabi-v7a/
│   ├── arm64-v8a/
│   ├── x86/
│   └── x86_64/
├── META-INF/                 # 签名信息
│   ├── MANIFEST.MF           # 文件清单
│   ├── CERT.SF               # 签名文件
│   └── CERT.RSA              # 证书文件
└── kotlin/                   # Kotlin 相关（如果使用）
```

### 3.2 解压 APK

```bash
# 方法 1: 使用 unzip
unzip app.apk -d app_extracted

# 方法 2: 使用 7z
7z x app.apk -o./app_extracted

# 查看文件结构
cd app_extracted
tree
```

### 3.3 查看 APK 基本信息

```bash
# 使用 aapt (Android Asset Packaging Tool)
aapt dump badging app.apk

# 输出信息包括:
# - package: 包名
# - versionCode: 版本号
# - versionName: 版本名称
# - sdkVersion: 最低 SDK 版本
# - targetSdkVersion: 目标 SDK 版本
# - uses-permission: 权限列表
# - application: 应用组件

# 查看 AndroidManifest.xml
aapt dump xmltree app.apk AndroidManifest.xml

# 查看资源列表
aapt list app.apk
```

### 3.4 分析 DEX 文件

```bash
# 查看 DEX 文件信息
dexdump classes.dex

# 查看类列表
dexdump -f classes.dex | grep "Class descriptor"

# 查看方法列表
dexdump -f classes.dex | grep "method_name"

# 统计信息
dexdump -l xml classes.dex
```

---

## 4. 反编译工具

### 4.1 使用 APKTool 反编译

APKTool 可以将 APK 反编译为 Smali 代码和资源文件。

```bash
# 基本反编译
apktool d app.apk

# 指定输出目录
apktool d app.apk -o app_decompiled

# 不反编译资源文件
apktool d app.apk --no-res

# 不反编译 DEX 文件
apktool d app.apk --no-src

# 保留调试信息
apktool d app.apk -d

# 查看反编译后的结构
cd app_decompiled
tree -L 2
```

反编译后的目录结构：

```
app_decompiled/
├── AndroidManifest.xml       # 可读的 XML 清单
├── apktool.yml               # APKTool 配置
├── smali/                    # Smali 代码（主 DEX）
│   └── com/example/app/      # 包结构
│       ├── MainActivity.smali
│       ├── Utils.smali
│       └── ...
├── smali_classes2/           # 第二个 DEX（如果存在）
├── res/                      # 资源文件（可读的XML）
│   ├── layout/
│   ├── values/
│   └── ...
├── assets/                   # 资产文件
├── lib/                      # 本地库
├── original/                 # 原始 META-INF
└── unknown/                  # 未知文件
```

### 4.2 使用 JADX 反编译为 Java

JADX 可以直接将 DEX 反编译为 Java 源代码（更易读）。

```bash
# 命令行反编译
jadx app.apk -d output_dir

# 保留注释
jadx app.apk -d output_dir --comments-level info

# 使用多线程加速
jadx app.apk -d output_dir -j 4

# 导出为 Gradle 项目
jadx app.apk -d output_dir --export-gradle

# 图形界面（推荐）
jadx-gui app.apk
```

**JADX-GUI 特性**:
- 代码浏览器
- 全文搜索
- 交叉引用
- 代码导出
- 资源查看器

### 4.3 使用 dex2jar + JD-GUI

```bash
# Step 1: 转换 DEX 到 JAR
d2j-dex2jar.sh app.apk -o app.jar

# 或从解压的 APK
d2j-dex2jar.sh classes.dex -o classes.jar

# Step 2: 使用 JD-GUI 查看
# 下载 JD-GUI: http://java-decompiler.github.io/
java -jar jd-gui.jar app.jar

# 或命令行导出
jd-cli app.jar -od output_src
```

### 4.4 Smali 代码基础

Smali 是 Dalvik 字节码的汇编语言表示形式。

**基本语法**:

```smali
# 类定义
.class public Lcom/example/MainActivity;
.super Landroid/app/Activity;

# 字段定义
.field private userName:Ljava/lang/String;

# 方法定义
.method public onCreate(Landroid/os/Bundle;)V
    .locals 2              # 局部变量数量
    
    # 调用父类方法
    invoke-super {p0, p1}, Landroid/app/Activity;->onCreate(Landroid/os/Bundle;)V
    
    # 设置布局
    const v0, 0x7f0c001d   # layout id
    invoke-virtual {p0, v0}, Lcom/example/MainActivity;->setContentView(I)V
    
    # 返回
    return-void
.end method
```

**数据类型**:
- `V` - void
- `Z` - boolean
- `B` - byte
- `S` - short
- `C` - char
- `I` - int
- `J` - long (64 bit)
- `F` - float
- `D` - double (64 bit)
- `L` - 对象类型 (例如: `Ljava/lang/String;`)
- `[` - 数组 (例如: `[I` 表示 int 数组)

### 4.5 重新打包 APK

```bash
# 修改 Smali 代码或资源后，重新打包
apktool b app_decompiled -o app_modified.apk

# 对齐 APK
zipalign -v 4 app_modified.apk app_aligned.apk

# 签名 APK (生成密钥)
keytool -genkey -v -keystore my-release-key.jks \
        -keyalg RSA -keysize 2048 -validity 10000 \
        -alias my-key-alias

# 使用 jarsigner 签名
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 \
          -keystore my-release-key.jks app_aligned.apk my-key-alias

# 或使用 apksigner (Android SDK)
apksigner sign --ks my-release-key.jks --out app_signed.apk app_aligned.apk

# 验证签名
apksigner verify app_signed.apk

# 安装到设备
adb install -r app_signed.apk
```

---

## 5. 动态调试

### 5.1 使用 Android Studio 调试

```bash
# 1. 将 APK 安装到设备
adb install -t app.apk

# 2. 启动应用（可调试模式）
# 在 AndroidManifest.xml 中添加: android:debuggable="true"

# 3. 附加调试器
# Android Studio -> Run -> Attach Debugger to Android Process

# 4. 设置断点并调试
```

### 5.2 使用 ADB 调试

```bash
# 查看正在运行的进程
adb shell ps | grep <package_name>

# 查看应用日志
adb logcat | grep <package_name>

# 过滤特定标签
adb logcat -s TAG_NAME

# 清除日志
adb logcat -c

# 导出日志到文件
adb logcat > app_log.txt

# 查看崩溃日志
adb logcat | grep "AndroidRuntime"
```

### 5.3 使用 Frida 进行动态插桩

Frida 是强大的动态插桩工具，可以在运行时修改应用行为。

#### 基本使用

```bash
# 列出设备上的进程
frida-ps -U

# 列出应用
frida-ps -Uai

# 附加到运行中的应用
frida -U -n "App Name"

# 通过包名附加
frida -U -f com.example.app --no-pause

# 加载脚本
frida -U -f com.example.app -l script.js --no-pause
```

#### Frida 脚本示例

**1. Hook 函数调用**

```javascript
// hook_example.js
Java.perform(function() {
    // 获取目标类
    var MainActivity = Java.use("com.example.app.MainActivity");
    
    // Hook onCreate 方法
    MainActivity.onCreate.implementation = function(savedInstanceState) {
        console.log("[*] MainActivity.onCreate called");
        console.log("[*] savedInstanceState: " + savedInstanceState);
        
        // 调用原始方法
        this.onCreate(savedInstanceState);
        
        console.log("[*] onCreate finished");
    };
    
    // Hook 自定义方法
    MainActivity.checkLicense.implementation = function() {
        console.log("[*] checkLicense called");
        // 绕过检查，直接返回 true
        return true;
    };
});
```

**2. 打印方法参数和返回值**

```javascript
Java.perform(function() {
    var Utils = Java.use("com.example.app.Utils");
    
    Utils.encrypt.implementation = function(data) {
        console.log("[*] encrypt called with data: " + data);
        
        // 调用原始方法
        var result = this.encrypt(data);
        
        console.log("[*] encrypt returned: " + result);
        return result;
    };
});
```

**3. 修改函数返回值**

```javascript
Java.perform(function() {
    var LoginActivity = Java.use("com.example.app.LoginActivity");
    
    // 绕过登录验证
    LoginActivity.isValidUser.implementation = function(username, password) {
        console.log("[*] Login attempt: " + username + " / " + password);
        // 总是返回 true
        return true;
    };
});
```

**4. 枚举类的所有方法**

```javascript
Java.perform(function() {
    var targetClass = Java.use("com.example.app.TargetClass");
    var methods = targetClass.class.getDeclaredMethods();
    
    methods.forEach(function(method) {
        console.log("[*] Method: " + method.getName());
    });
});
```

**5. Hook 构造函数**

```javascript
Java.perform(function() {
    var User = Java.use("com.example.app.User");
    
    User.$init.overload('java.lang.String', 'int').implementation = function(name, age) {
        console.log("[*] Creating User: " + name + ", age: " + age);
        // 修改参数
        return this.$init("Admin", 99);
    };
});
```

### 5.4 使用 objection

objection 是基于 Frida 的更高级工具。

```bash
# 安装 objection
pip3 install objection

# 启动 objection
objection -g com.example.app explore

# objection 交互式命令:

# 查看应用信息
env

# 列出所有类
android hooking list classes

# 搜索类
android hooking search classes keyword

# 列出类的方法
android hooking list class_methods com.example.app.MainActivity

# Hook 方法
android hooking watch class_method com.example.app.MainActivity.onCreate --dump-args --dump-return

# 查看内存中的对象
android heap search instances com.example.app.User

# 调用方法
android heap execute <instance> toString

# 绕过 SSL Pinning
android sslpinning disable

# 监控文件操作
android monitoring file_system

# 截图
android ui screenshot /tmp/screen.png
```

---

## 6. 代码注入与 Hook

### 6.1 Xposed 框架

Xposed 是 Android 上最流行的 Hook 框架之一。

**安装 Xposed**:
1. 解锁 Bootloader
2. 安装 TWRP Recovery
3. 下载 Xposed Installer
4. 在 TWRP 中刷入 Xposed 框架

**创建 Xposed 模块**:

```java
// XposedModule.java
public class XposedModule implements IXposedHookLoadPackage {
    @Override
    public void handleLoadPackage(XC_LoadPackage.LoadPackageParam lpparam) throws Throwable {
        // 只 Hook 目标应用
        if (!lpparam.packageName.equals("com.example.app"))
            return;
        
        // Hook 方法
        XposedHelpers.findAndHookMethod(
            "com.example.app.MainActivity",
            lpparam.classLoader,
            "onCreate",
            Bundle.class,
            new XC_MethodHook() {
                @Override
                protected void beforeHookedMethod(MethodHookParam param) throws Throwable {
                    XposedBridge.log("MainActivity onCreate called");
                }
                
                @Override
                protected void afterHookedMethod(MethodHookParam param) throws Throwable {
                    XposedBridge.log("MainActivity onCreate finished");
                }
            }
        );
        
        // 替换方法实现
        XposedHelpers.findAndHookMethod(
            "com.example.app.Utils",
            lpparam.classLoader,
            "isPremiumUser",
            new XC_MethodReplacement() {
                @Override
                protected Object replaceHookedMethod(MethodHookParam param) throws Throwable {
                    // 直接返回 true
                    return true;
                }
            }
        );
    }
}
```

### 6.2 Native Hook (so 库)

对于 Native 代码 (.so 文件)，可以使用：

- **Substrate (Cydia Substrate)**: Native Hook 框架
- **Frida**: 也支持 Native Hook
- **IDA Pro**: 静态分析工具

**Frida Native Hook 示例**:

```javascript
// 查找导出函数
var nativeLib = Module.findExportByName("libnative.so", "Java_com_example_app_Native_encrypt");

if (nativeLib != null) {
    Interceptor.attach(nativeLib, {
        onEnter: function(args) {
            console.log("[*] encrypt called");
            console.log("[*] JNIEnv*: " + args[0]);
            console.log("[*] jobject: " + args[1]);
            console.log("[*] jstring: " + args[2]);
            
            // 读取字符串参数
            var str = Java.vm.getEnv().getStringUtfChars(args[2], null).readCString();
            console.log("[*] Input string: " + str);
        },
        onLeave: function(retval) {
            console.log("[*] Return value: " + retval);
        }
    });
}

// Hook 任意地址
var baseAddr = Module.findBaseAddress("libnative.so");
var targetAddr = baseAddr.add(0x1234); // 偏移地址

Interceptor.attach(targetAddr, {
    onEnter: function(args) {
        console.log("[*] Function at " + targetAddr + " called");
    }
});
```

---

## 7. 加密与脱壳

### 7.1 常见加密方式

Android 应用常用的加密方式：

1. **DEX 加密**: 对 DEX 文件进行加密，运行时解密
2. **资源加密**: 加密资源文件和资产
3. **字符串混淆**: 对关键字符串进行编码
4. **代码混淆**: ProGuard/R8 混淆
5. **Native 保护**: 关键逻辑放在 SO 库中

### 7.2 常见加固厂商

- **360 加固保**
- **腾讯乐固**
- **梆梆加固**
- **爱加密**
- **阿里云加固**

### 7.3 脱壳方法

#### 方法 1: 内存 Dump (Frida)

```javascript
// dump_dex.js
Java.perform(function() {
    // Hook ClassLoader
    var DexFile = Java.use("dalvik.system.DexFile");
    
    DexFile.loadDex.overload(
        'java.lang.String',
        'java.lang.String',
        'int'
    ).implementation = function(sourcePathName, outputPathName, flags) {
        console.log("[*] DexFile.loadDex called");
        console.log("[*] Source: " + sourcePathName);
        console.log("[*] Output: " + outputPathName);
        
        var result = this.loadDex(sourcePathName, outputPathName, flags);
        
        // Dump DEX 到文件
        var File = Java.use("java.io.File");
        var dexFile = File.$new(sourcePathName);
        if (dexFile.exists()) {
            var dexBytes = readFile(sourcePathName);
            var dumpPath = "/sdcard/" + sourcePathName.split("/").pop();
            writeFile(dumpPath, dexBytes);
            console.log("[*] DEX dumped to: " + dumpPath);
        }
        
        return result;
    };
});

function readFile(path) {
    var FileInputStream = Java.use("java.io.FileInputStream");
    var fis = FileInputStream.$new(path);
    var available = fis.available();
    var buffer = Java.array('byte', new Array(available));
    fis.read(buffer);
    fis.close();
    return buffer;
}

function writeFile(path, bytes) {
    var FileOutputStream = Java.use("java.io.FileOutputStream");
    var fos = FileOutputStream.$new(path);
    fos.write(bytes);
    fos.close();
}
```

#### 方法 2: 使用专用脱壳工具

```bash
# FART (ART 环境下的自动脱壳)
# 需要刷入定制 ROM

# BlackDex (基于 VirtualApp)
# 在沙盒环境中运行应用并自动脱壳

# DexHunter
# 针对某些加固的脱壳工具
```

#### 方法 3: 手动分析

```bash
# 1. 找到加载 DEX 的时机
# 搜索关键函数: DexClassLoader, PathClassLoader, loadDex

# 2. Hook 加载点
# 在 DEX 加载后，从内存中 dump 出来

# 3. 使用 IDA Pro 或 Ghidra 分析 SO 库
# 找到解密逻辑，提取密钥

# 4. 编写解密脚本
# 根据加密算法编写相应的解密工具
```

### 7.4 绕过 SSL Pinning

```bash
# 使用 objection
objection -g com.example.app explore
android sslpinning disable

# 使用 Frida 脚本
frida -U -f com.example.app -l ssl-unpinning.js --no-pause
```

**ssl-unpinning.js**:

```javascript
Java.perform(function() {
    // 绕过 OkHttp3 Certificate Pinning
    var CertificatePinner = Java.use("okhttp3.CertificatePinner");
    CertificatePinner.check.overload('java.lang.String', 'java.util.List').implementation = function() {
        console.log("[*] OkHttp3 Certificate Pinning bypassed");
    };
    
    // 绕过 TrustManager
    var X509TrustManager = Java.use("javax.net.ssl.X509TrustManager");
    var SSLContext = Java.use("javax.net.ssl.SSLContext");
    
    var TrustManager = Java.registerClass({
        name: "com.sensepost.test.TrustManager",
        implements: [X509TrustManager],
        methods: {
            checkClientTrusted: function(chain, authType) {},
            checkServerTrusted: function(chain, authType) {},
            getAcceptedIssuers: function() { return []; }
        }
    });
    
    var TrustManagers = [TrustManager.$new()];
    var SSLContext_init = SSLContext.init.overload(
        '[Ljavax.net.ssl.KeyManager;',
        '[Ljavax.net.ssl.TrustManager;',
        'java.security.SecureRandom'
    );
    
    SSLContext_init.implementation = function(keyManager, trustManager, secureRandom) {
        console.log("[*] SSLContext.init bypassed");
        SSLContext_init.call(this, keyManager, TrustManagers, secureRandom);
    };
});
```

---

## 8. 实战案例

### 8.1 案例 1: 移除应用内广告

**目标**: 移除应用中的广告模块

**步骤**:

```bash
# 1. 反编译 APK
apktool d app.apk -o app_decompiled

# 2. 分析 AndroidManifest.xml
# 查找广告相关的 Activity、Service、Receiver
cat app_decompiled/AndroidManifest.xml | grep -i "ad"

# 3. 删除广告组件
# 编辑 AndroidManifest.xml，移除广告相关的声明

# 4. 分析 Smali 代码
# 查找广告 SDK 的初始化代码
grep -r "admob\|adview" app_decompiled/smali/

# 5. 注释掉广告初始化代码
# 编辑对应的 .smali 文件，注释掉或修改广告相关代码

# 示例: 注释掉方法调用
# invoke-static {}, Lcom/google/android/gms/ads/MobileAds;->initialize(...)
# 改为:
# # invoke-static {}, Lcom/google/android/gms/ads/MobileAds;->initialize(...)

# 6. 重新打包
apktool b app_decompiled -o app_no_ads.apk

# 7. 签名
apksigner sign --ks my-key.jks --out app_signed.apk app_no_ads.apk

# 8. 安装测试
adb install -r app_signed.apk
```

### 8.2 案例 2: 绕过应用验证

**目标**: 绕过应用的会员验证

**步骤**:

```bash
# 1. 使用 JADX 反编译查看 Java 代码
jadx-gui app.apk

# 2. 搜索验证相关的代码
# 搜索关键词: "premium", "vip", "subscribe", "purchase"

# 3. 定位验证方法
# 例如: isPremiumUser(), isSubscribed()

# 4. 使用 Frida Hook 修改返回值
```

**Frida 脚本**:

```javascript
// bypass_premium.js
Java.perform(function() {
    var UserManager = Java.use("com.example.app.UserManager");
    
    // Hook isPremiumUser 方法
    UserManager.isPremiumUser.implementation = function() {
        console.log("[*] isPremiumUser called - returning true");
        return true;
    };
    
    // Hook 订阅检查
    UserManager.getSubscriptionStatus.implementation = function() {
        console.log("[*] getSubscriptionStatus called");
        return "ACTIVE"; // 返回激活状态
    };
    
    // Hook 到期时间检查
    UserManager.getExpirationDate.implementation = function() {
        console.log("[*] getExpirationDate called");
        // 返回遥远的未来时间
        var Date = Java.use("java.util.Date");
        return Date.$new(253402300799000); // 9999-12-31
    };
});
```

**运行**:

```bash
frida -U -f com.example.app -l bypass_premium.js --no-pause
```

### 8.3 案例 3: 提取应用内的加密密钥

**目标**: 从 Native 库中提取加密密钥

**步骤**:

```bash
# 1. 解压 APK
unzip app.apk -d app_extracted

# 2. 找到 Native 库
cd app_extracted/lib/arm64-v8a/
ls -la *.so

# 3. 使用 strings 查看字符串
strings libnative.so | grep -i "key\|secret\|password"

# 4. 使用 IDA Pro 或 Ghidra 反编译
# 打开 libnative.so，找到加密相关函数

# 5. 使用 Frida Hook Native 函数
```

**Frida 脚本**:

```javascript
// extract_key.js
var nativeLib = Process.getModuleByName("libnative.so");

// 查找 encrypt 函数
var encryptFunc = nativeLib.getExportByName("Java_com_example_app_Native_encrypt");

if (encryptFunc) {
    Interceptor.attach(encryptFunc, {
        onEnter: function(args) {
            console.log("[*] Encrypt function called");
            
            // 读取密钥 (假设密钥在特定偏移)
            var keyAddr = this.context.x2; // ARM64: x2 寄存器
            var key = keyAddr.readUtf8String();
            console.log("[*] Encryption key: " + key);
            
            // 或者从内存搜索
            Memory.scan(nativeLib.base, nativeLib.size, "41 45 53 20 4B 65 79", {
                onMatch: function(address, size) {
                    console.log("[*] Found key at: " + address);
                    console.log("[*] Key: " + address.readUtf8String());
                },
                onComplete: function() {
                    console.log("[*] Memory scan complete");
                }
            });
        }
    });
}
```

---

## 9. 最佳实践

### 9.1 逆向分析流程

1. **信息收集**:
   - 查看 APK 基本信息
   - 分析 AndroidManifest.xml
   - 识别使用的第三方库和框架

2. **静态分析**:
   - 反编译查看 Java/Smali 代码
   - 分析关键类和方法
   - 识别加密和混淆机制

3. **动态分析**:
   - 使用 Frida 进行运行时分析
   - Hook 关键函数查看参数和返回值
   - 监控网络请求和文件操作

4. **深入分析**:
   - 分析 Native 库
   - 提取加密密钥
   - 绕过安全检测

5. **验证和测试**:
   - 修改代码重新打包
   - 测试修改后的功能
   - 确保应用稳定运行

### 9.2 工具选择建议

| 任务 | 推荐工具 | 备选工具 |
|------|---------|---------|
| 快速查看代码 | JADX-GUI | JD-GUI, Bytecode Viewer |
| 代码修改 | APKTool | MT Manager (Android) |
| 动态调试 | Frida | Xposed, Cydia Substrate |
| Native 分析 | IDA Pro | Ghidra, radare2 |
| 网络抓包 | Charles | Burp Suite, Fiddler |
| 自动化分析 | MobSF | AndroGuard |

### 9.3 安全提示

⚠️ **重要提醒**:

1. **仅用于学习**: 逆向工程应仅用于学习和研究目的
2. **尊重版权**: 不要破解或分发受版权保护的软件
3. **遵守法律**: 了解并遵守当地的法律法规
4. **负责任披露**: 发现安全漏洞应负责任地报告给开发者
5. **隐私保护**: 不要窃取或滥用用户隐私数据

### 9.4 常用资源

**学习资源**:
- [Android Security Wiki](https://github.com/ashishb/android-security-awesome)
- [OWASP Mobile Security](https://owasp.org/www-project-mobile-security/)
- [Frida CodeShare](https://codeshare.frida.re/)
- [Android Hacker's Handbook](https://www.wiley.com/en-us/Android+Hacker%27s+Handbook-p-9781118608647)

**社区和论坛**:
- XDA Developers
- Reddit r/ReverseEngineering
- Stack Overflow
- 看雪论坛 (中文)
- 吾爱破解 (中文)

**工具集合**:
- [MobSF](https://github.com/MobSF/Mobile-Security-Framework-MobSF)
- [Android Security Tools](https://github.com/ashishb/android-security-awesome#tools)
- [Awesome Frida](https://github.com/dweinstein/awesome-frida)

---

## 总结

本指南涵盖了 Android 逆向工程的核心工具和技术，从基础的 APK 分析到高级的动态调试和代码注入。通过掌握这些技能，您可以：

✅ 理解 Android 应用的内部结构  
✅ 分析和调试应用代码  
✅ 识别安全漏洞和恶意行为  
✅ 学习 Android 系统的深层机制  
✅ 提升自己的安全研究能力  

**记住**: 技术本身是中立的，关键在于如何使用。请始终以合法和道德的方式使用这些技术！

---

**下一步**: 查看[设备刷机和调试文档](./device-flashing-debug.md)，学习如何将定制系统刷入设备。

**祝您学习愉快！** 🚀
