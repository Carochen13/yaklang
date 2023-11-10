
<p align="center">
  <a href="https://yaklang.io/"><img src="imgs/yaklang-logo.png" style="width: 400px"/></a> 
 <h2 align="center">为网络安全而生的领域编程语言</h2>
<p align="center">
<img src="https://img.shields.io/github/issues-pr/yaklang/yaklang">
<a href="https://github.com/yaklang/yaklang/releases"><img src="https://img.shields.io/github/downloads/yaklang/yaklang/total">
<a href="https://github.com/yaklang/yaklang/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/yaklang/yaklang">
<a href="https://github.com/yaklang/yaklang/releases/"><img src="https://img.shields.io/github/release/yaklang/yaklang">
<a href="https://github.com/yaklang/yaklang/issues"><img src="https://img.shields.io/github/issues-raw/yaklang/yaklang">
<a href="https://github.com/yaklang/yaklang/discussions"><img src="https://img.shields.io/github/stars/yaklang/yaklang">
<a href="https://github.com/yaklang/yaklang/blob/main/LICENSE.md"><img src="https://img.shields.io/github/license/yaklang/yaklang">
</p>

<p align="center">
  <a href="#快速开始">快速开始</a> •
  <a href="https://yaklang.io/docs/intro">官方文档</a> •
  <a href="https://github.com/yaklang/yaklang/issues">问题反馈</a> •
  <a href="https://yaklang.io/docs/api/global_buildin_ops/">接口手册</a> •
  <a href="#贡献你的代码">贡献代码</a> •
  <a href="#社区 ">加入社区</a> •
  <a href="#项目架构">项目架构</a> 
</p>

<p align="center">
 :book:语言选择： <a href="https://github.com/yaklang/yaklang/blob/main/README—_EN.md">English</a> • 
  <a href="https://github.com/yaklang/yaklang/blob/main/README.md">中文</a> 
</p>

---
# CDSL-Yakang 简介

CDSL：Cybersecurity Domain Specific Language，全称网络安全领域编程语言。

Yaklang 团队综合“领域限定语言”的思想，构建了CDSL的概念，并以此为核心构建了Yak(又称Yaklang)语言来构建基础设施和语言生态。

Yak 是一门针对网络安全领域研发的易书写，易分发的高级计算机编程语言。Yak具备强类型、动态类型的经典类型特征，兼具编译字节码和解释执行的运行时特征。

Yak语言的运行时环境只依赖于YakVM，可以实现“一次编写，处处运行”的特性，只要有YakVM部署的环境，都可以快速执行Yak语言程序。

<h3 align="center">
  <img src="imgs/yaklang-cdsl.png" style="width: 800px" alt="yaklang-cdsl.png" ></a>
</h3>

Yak语言起初只作为一个“嵌入式语言”在宿主程序中存在，后在电子科技大学网络空间安全学院学术指导下，由 Yaklang.io 研发团队进行长达两年的迭代与改造，实现了YakVM虚拟机让语言可以脱离“宿主语言”独立运行，并与2023年完全开源。 支持目前主流操作系统：macOS，Linux，Windows。


## Yaklang 的优势

基于CDSL概念构建的网络安全领域编程语言Yak，具备了几乎DSL所有的优势，它被设计为针对安全能力研发领域的专用编程语言，实现了常见的大多数安全能力，可以让各种各样的安全能力彼此之间“互补，融合，进化”；提高安全从业人员的生产力。

CDSL在网络安全领域提供的能力具备很多优势：
- 简洁性：使用CDSL构建的安全产品更能实现业务和能力的分离，并且解决方案更加直观；

- 易用性：非专业的人员也可以使用CDSL构建安全产品，而避免安全产品工程化中的信息差；

- 灵活性：CDSL一般被设计为单独使用和嵌入式使用均可，用户可以根据自己的需求去编写DSL脚本以实现特定的策略和检测规则，这往往更能把用户的思路展示出来，而不必受到冗杂知识的制约；

除此之外，作为一门专门为网络安全研发设计的语言，Yak语言除了满足一些基础的语言本身需要具备的特性之外，还具有很多特殊功能，可以帮助用户快速构建网络安全应用：

1中间人劫持库函数

2. 复杂端口扫描和服务指纹识别

3. 网络安全领域的加解密库

4. 支持中国商用密码体系：支持SM2椭圆曲线公钥密码算法，SM4分组密码算法，SM3密码杂凑算法等

<h3 align="center">
  <img src="imgs/yaklang-fix.jpg" style="width: 800px" alt="yaklang-fix.jpg" ></a>
</h3>

## 项目架构 

![yaklang-architecture](imgs/yaklang-arch.jpg)

## 快速开始

### 通过 Yakit 来使用 Yaklang

Yakit (https://github.com/yaklang/yakit) 是 Yaklang.io 团队官方出品的开源 Yaklang IDE，它可以帮助你快速上手 Yaklang 语言。

同时 Yakit 也能将绝大部分安全工程师需要的核心功能图形化。他是免费的，你可以通过 [下载安装 Yakit](https://www.yaklang.com/products/download_and_install)，来开始使用 Yaklang。

关于Yakit的更多内容可移步：官方文档


### 通过命令行来安装使用

通过命令行来安装使用 Yaklang 请遵循：**https://www.yaklang.com/** 或 **https://www.yaklang.io/** 的指引，或直接执行

#### MacOS / Linux

```bash
bash <(curl -sS -L http://oss.yaklang.io/install-latest-yak.sh)
```

#### Windows

```bash
powershell (new-object System.Net.WebClient).DownloadFile('https://yaklang.oss-cn-beijing.aliyuncs.com/yak/latest/yak_windows_amd64.exe','yak_windows_amd64.exe') && yak_windows_amd64.exe install && del /f yak_windows_amd64.exe
```

## 社区

1. 你可以在 yaklang 或者 yakit 的 issues 中添加你想讨论的内容或者你想表达的东西，英文或中文均可，我们会尽快回复
2. 国内用户可以使用 WeChat 加入群组
3. 国际用户可以使用 Discord 加入社区

## 贡献你的代码 

这是一个高级话题，在贡献你的代码之前，确保你对 Yaklang 整个项目结构有所了解。

在贡献代码时，如果你希望修改 Yaklang 或 YakVM 本身的核心语法部分，最好与研发团队取得联系。

如果您仅仅想要增加库的功能，或者修复一些库的 Bug，那么您可以直接提交 PR，当然 PR 中最好包含对应的单元测试，这很有助于提升我们的代码质量。

## 项目成员 

### Maintainer

[v1ll4n](https://github.com/VillanCh): Yak Project Maintainer.

### yaklang 核心开发者 / Active yaklang core developers

1. [z3](https://github.com/OrangeWatermelon)
2. [Longlone](https://github.com/way29)
3. [Go0p](https://github.com/Go0p)
4. [Matrix-Cain](https://github.com/Matrix-Cain)
5. [bcy2007](https://github.com/bcy2007)
6. [naiquan](https://github.com/naiquann)
7. [Rookie-is](https://github.com/Rookie-is)

## 开源许可证

本仓库代码版本使用 AGPL 开源协议，这是一个严格的开源协议，且具有传染性，如果您使用了本仓库的代码，那么您的代码也必须开源。

1. 强制开源网络服务:要求提供网络服务的源代码必须开源。保证开源理念在网络环境下的实践。
2. 其他条款与 GPL 相同:开源免费、开源修改、衍生开源等。

本项目开源仓库仅应该作为个人开源和学习使用。

## 鸣谢

本项目经由[电子科技大学](https://www.uestc.edu.cn)张小松([网络空间安全学院](https://www.scse.uestc.edu.cn/))教授学术指导。

<div class="thanksfor-member-wrapper"><div class="thanksfor-member-grid-wrapper"><div class="thanksfor-member-opt"><div class="member-info"><div class="member-info-img"><img src="/img/team/safety.jpg" class="img-style"></div><div class="member-info-name-tag"><div class="name-style">电子科技大学网络空间安全研究院</div></div></div><div class="tag-div-style red-tag-div" style="font-size: 12px;"><div class="tag-styl">长期系统开展网络安全研究和人才培养</div></div><div class="member-description"><div class="description-style">YAK 架构和思想的策源地，核心团队和人员的成长培养</div></div><div class="member-address-link"><div class="address-body"><span role="img" class="anticon address-icon"><svg width="1em" height="1em" viewBox="0 0 24 24" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><title>定位</title><defs><polygon id="path-1" points="0 0 24 0 24 24 0 24"></polygon></defs><g id="页面-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"><g id="官网/团队致谢" transform="translate(-404.000000, -1358.000000)"><g id="卡片/贡献者备份-4" transform="translate(384.000000, 1212.000000)"><g id="定位" transform="translate(20.000000, 146.000000)"><mask id="mask-2" fill="white"><use xlink:href="#path-1"></use></mask><g id="Clip-2"></g><g id="定位-4" mask="url(#mask-2)" fill="#C3C7CF"><g transform="translate(5.500000, 4.000000)"><path d="M6.50716394,1.2605942 C9.41610589,1.2605942 11.7619609,3.51620027 11.7619609,6.28547981 C11.7505395,7.30942305 11.4228459,8.3037731 10.8253089,9.12756049 L10.7597516,9.23156211 C10.7671802,9.22087831 10.7756302,9.21104542 10.7848231,9.20187437 L10.7398802,9.25359154 L6.49323534,14.739488 L2.24659053,9.25359154 L2.22244762,9.22456564 L2.17629752,9.15025175 L2.21251188,9.21322 L2.20944759,9.21009996 L2.2152976,9.21832554 C1.56000334,8.29488569 1.2381598,7.31566315 1.2381598,6.28547981 C1.25050982,3.51052746 3.59051484,1.2605942 6.50716394,1.2605942 M6.50716394,0 C2.92287055,0 0.0157857481,2.79537449 0,6.28226522 L0.00157857481,6.42730021 C0.0294357774,7.63050442 0.405508012,8.77395497 1.07993089,9.75591937 L1.15273104,9.85963735 L1.18578826,9.91419093 L1.21643118,9.96080257 C1.23723122,9.99077394 1.25914556,10.019138 1.28365989,10.0471239 L1.28923133,10.0534586 L5.52101897,15.5198783 C5.58128339,15.597785 5.65036925,15.6682225 5.7268837,15.7297725 L5.7812981,15.7707113 C6.31912783,16.1555173 7.05845798,16.0453702 7.46545171,15.5198783 L11.6972394,10.0534586 L11.667525,10.0862663 C11.7201751,10.0314291 11.7603823,9.97715919 11.800311,9.91447457 L11.8393111,9.8504663 C12.5735341,8.83314135 12.9846135,7.59712935 13.0000279,6.29332721 C13.0000279,2.80360007 10.0843073,0 6.50716394,0" id="Fill-1"></path><path d="M6.50004179,4.8199739 C7.35479362,4.81959587 8.04797367,5.52482141 8.04834525,6.39503134 C8.04871653,7.26533582 7.35609362,7.97112864 6.50143465,7.97150683 L6.50004179,7.97150683 C5.64528995,7.97150683 4.95238847,7.26599765 4.95238847,6.39578771 C4.95238847,5.52548324 5.64528995,4.8199739 6.50004179,4.8199739 M6.50004179,3.55937986 C4.96148849,3.55937986 3.71432153,4.82923965 3.71432153,6.39578771 C3.71432153,7.96224123 4.96148849,9.23219557 6.50004179,9.23219557 C8.03859508,9.23219557 9.28576204,7.96224123 9.28576204,6.39578771 C9.28576204,4.82923965 8.03859508,3.55937986 6.50004179,3.55937986" id="Fill-3"></path></g></g></g></g></g></g></svg></span><span class="address-style">四川成都</span></div><div class="link-body"><a class="link-jump" href="http://www.uestc.edu.cn" target="_blank"><span role="img" aria-label="home" class="anticon anticon-home icon-style icon-home-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="home" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M946.5 505L534.6 93.4a31.93 31.93 0 00-45.2 0L77.5 505c-12 12-18.8 28.3-18.8 45.3 0 35.3 28.7 64 64 64h43.4V908c0 17.7 14.3 32 32 32H448V716h112v224h265.9c17.7 0 32-14.3 32-32V614.3h43.4c17 0 33.3-6.7 45.3-18.8 24.9-25 24.9-65.5-.1-90.5z"></path></svg></span></a></div></div></div><div class="thanksfor-member-opt"><div class="member-info"><div class="member-info-img"><img src="/img/team/projectdiscovery.png" class="img-style"></div><div class="member-info-name-tag"><div class="name-style">ProjectDiscovery</div></div></div><div class="tag-div-style purple-tag-div"><div class="tag-styl">nuclei 集成</div></div><div class="member-description"><div class="description-style">无私的 MIT 开源精神</div><div class="description-style">为 Yak 提供 nuclei 漏洞检测能力</div></div><div class="member-address-link"><div class="address-body"><span role="img" class="anticon address-icon"><svg width="1em" height="1em" viewBox="0 0 24 24" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><title>定位</title><defs><polygon id="path-1" points="0 0 24 0 24 24 0 24"></polygon></defs><g id="页面-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"><g id="官网/团队致谢" transform="translate(-404.000000, -1358.000000)"><g id="卡片/贡献者备份-4" transform="translate(384.000000, 1212.000000)"><g id="定位" transform="translate(20.000000, 146.000000)"><mask id="mask-2" fill="white"><use xlink:href="#path-1"></use></mask><g id="Clip-2"></g><g id="定位-4" mask="url(#mask-2)" fill="#C3C7CF"><g transform="translate(5.500000, 4.000000)"><path d="M6.50716394,1.2605942 C9.41610589,1.2605942 11.7619609,3.51620027 11.7619609,6.28547981 C11.7505395,7.30942305 11.4228459,8.3037731 10.8253089,9.12756049 L10.7597516,9.23156211 C10.7671802,9.22087831 10.7756302,9.21104542 10.7848231,9.20187437 L10.7398802,9.25359154 L6.49323534,14.739488 L2.24659053,9.25359154 L2.22244762,9.22456564 L2.17629752,9.15025175 L2.21251188,9.21322 L2.20944759,9.21009996 L2.2152976,9.21832554 C1.56000334,8.29488569 1.2381598,7.31566315 1.2381598,6.28547981 C1.25050982,3.51052746 3.59051484,1.2605942 6.50716394,1.2605942 M6.50716394,0 C2.92287055,0 0.0157857481,2.79537449 0,6.28226522 L0.00157857481,6.42730021 C0.0294357774,7.63050442 0.405508012,8.77395497 1.07993089,9.75591937 L1.15273104,9.85963735 L1.18578826,9.91419093 L1.21643118,9.96080257 C1.23723122,9.99077394 1.25914556,10.019138 1.28365989,10.0471239 L1.28923133,10.0534586 L5.52101897,15.5198783 C5.58128339,15.597785 5.65036925,15.6682225 5.7268837,15.7297725 L5.7812981,15.7707113 C6.31912783,16.1555173 7.05845798,16.0453702 7.46545171,15.5198783 L11.6972394,10.0534586 L11.667525,10.0862663 C11.7201751,10.0314291 11.7603823,9.97715919 11.800311,9.91447457 L11.8393111,9.8504663 C12.5735341,8.83314135 12.9846135,7.59712935 13.0000279,6.29332721 C13.0000279,2.80360007 10.0843073,0 6.50716394,0" id="Fill-1"></path><path d="M6.50004179,4.8199739 C7.35479362,4.81959587 8.04797367,5.52482141 8.04834525,6.39503134 C8.04871653,7.26533582 7.35609362,7.97112864 6.50143465,7.97150683 L6.50004179,7.97150683 C5.64528995,7.97150683 4.95238847,7.26599765 4.95238847,6.39578771 C4.95238847,5.52548324 5.64528995,4.8199739 6.50004179,4.8199739 M6.50004179,3.55937986 C4.96148849,3.55937986 3.71432153,4.82923965 3.71432153,6.39578771 C3.71432153,7.96224123 4.96148849,9.23219557 6.50004179,9.23219557 C8.03859508,9.23219557 9.28576204,7.96224123 9.28576204,6.39578771 C9.28576204,4.82923965 8.03859508,3.55937986 6.50004179,3.55937986" id="Fill-3"></path></g></g></g></g></g></g></svg></span><span class="address-style">- -</span></div><div class="link-body"><a class="link-jump" href="https://projectdiscovery.io/#/" target="_blank"><span role="img" aria-label="home" class="anticon anticon-home icon-style icon-home-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="home" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M946.5 505L534.6 93.4a31.93 31.93 0 00-45.2 0L77.5 505c-12 12-18.8 28.3-18.8 45.3 0 35.3 28.7 64 64 64h43.4V908c0 17.7 14.3 32 32 32H448V716h112v224h265.9c17.7 0 32-14.3 32-32V614.3h43.4c17 0 33.3-6.7 45.3-18.8 24.9-25 24.9-65.5-.1-90.5z"></path></svg></span></a><a href="https://github.com/projectdiscovery/" target="_blank"><span role="img" aria-label="github" class="anticon anticon-github icon-style icon-github-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="github" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M511.6 76.3C264.3 76.2 64 276.4 64 523.5 64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c30.9-15.9 62.4 4 98.9 57.9 26.4 39.1 77.9 32.5 104 26 5.7-23.5 17.9-44.5 34.7-60.8-140.6-25.2-199.2-111-199.2-213 0-49.5 16.3-95 48.3-131.7-20.4-60.5 1.9-112.3 4.9-120 58.1-5.2 118.5 41.6 123.2 45.3 33-8.9 70.7-13.6 112.9-13.6 42.4 0 80.2 4.9 113.5 13.9 11.3-8.6 67.3-48.8 121.3-43.9 2.9 7.7 24.7 58.3 5.5 118 32.4 36.8 48.9 82.7 48.9 132.3 0 102.2-59 188.1-200 212.9a127.5 127.5 0 0138.1 91v112.5c.8 9 0 17.9 15 17.9 177.1-59.7 304.6-227 304.6-424.1 0-247.2-200.4-447.3-447.5-447.3z"></path></svg></span></a></div></div></div><div class="thanksfor-member-opt"><div class="member-info"><div class="member-info-img"><img src="/img/team/vulhub.png" class="img-style"></div><div class="member-info-name-tag"><div class="name-style">Vulhub</div></div></div><div class="tag-div-style blue-tag-div"><div class="tag-styl">漏洞靶场基础设施</div></div><div class="member-description"><div class="description-style">无私的漏洞靶场基础设施提供者</div></div><div class="member-address-link"><div class="address-body"><span role="img" class="anticon address-icon"><svg width="1em" height="1em" viewBox="0 0 24 24" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><title>定位</title><defs><polygon id="path-1" points="0 0 24 0 24 24 0 24"></polygon></defs><g id="页面-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"><g id="官网/团队致谢" transform="translate(-404.000000, -1358.000000)"><g id="卡片/贡献者备份-4" transform="translate(384.000000, 1212.000000)"><g id="定位" transform="translate(20.000000, 146.000000)"><mask id="mask-2" fill="white"><use xlink:href="#path-1"></use></mask><g id="Clip-2"></g><g id="定位-4" mask="url(#mask-2)" fill="#C3C7CF"><g transform="translate(5.500000, 4.000000)"><path d="M6.50716394,1.2605942 C9.41610589,1.2605942 11.7619609,3.51620027 11.7619609,6.28547981 C11.7505395,7.30942305 11.4228459,8.3037731 10.8253089,9.12756049 L10.7597516,9.23156211 C10.7671802,9.22087831 10.7756302,9.21104542 10.7848231,9.20187437 L10.7398802,9.25359154 L6.49323534,14.739488 L2.24659053,9.25359154 L2.22244762,9.22456564 L2.17629752,9.15025175 L2.21251188,9.21322 L2.20944759,9.21009996 L2.2152976,9.21832554 C1.56000334,8.29488569 1.2381598,7.31566315 1.2381598,6.28547981 C1.25050982,3.51052746 3.59051484,1.2605942 6.50716394,1.2605942 M6.50716394,0 C2.92287055,0 0.0157857481,2.79537449 0,6.28226522 L0.00157857481,6.42730021 C0.0294357774,7.63050442 0.405508012,8.77395497 1.07993089,9.75591937 L1.15273104,9.85963735 L1.18578826,9.91419093 L1.21643118,9.96080257 C1.23723122,9.99077394 1.25914556,10.019138 1.28365989,10.0471239 L1.28923133,10.0534586 L5.52101897,15.5198783 C5.58128339,15.597785 5.65036925,15.6682225 5.7268837,15.7297725 L5.7812981,15.7707113 C6.31912783,16.1555173 7.05845798,16.0453702 7.46545171,15.5198783 L11.6972394,10.0534586 L11.667525,10.0862663 C11.7201751,10.0314291 11.7603823,9.97715919 11.800311,9.91447457 L11.8393111,9.8504663 C12.5735341,8.83314135 12.9846135,7.59712935 13.0000279,6.29332721 C13.0000279,2.80360007 10.0843073,0 6.50716394,0" id="Fill-1"></path><path d="M6.50004179,4.8199739 C7.35479362,4.81959587 8.04797367,5.52482141 8.04834525,6.39503134 C8.04871653,7.26533582 7.35609362,7.97112864 6.50143465,7.97150683 L6.50004179,7.97150683 C5.64528995,7.97150683 4.95238847,7.26599765 4.95238847,6.39578771 C4.95238847,5.52548324 5.64528995,4.8199739 6.50004179,4.8199739 M6.50004179,3.55937986 C4.96148849,3.55937986 3.71432153,4.82923965 3.71432153,6.39578771 C3.71432153,7.96224123 4.96148849,9.23219557 6.50004179,9.23219557 C8.03859508,9.23219557 9.28576204,7.96224123 9.28576204,6.39578771 C9.28576204,4.82923965 8.03859508,3.55937986 6.50004179,3.55937986" id="Fill-3"></path></g></g></g></g></g></g></svg></span><span class="address-style">新加坡</span></div><div class="link-body"><a class="link-jump" href="https://vulhub.org/" target="_blank"><span role="img" aria-label="home" class="anticon anticon-home icon-style icon-home-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="home" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M946.5 505L534.6 93.4a31.93 31.93 0 00-45.2 0L77.5 505c-12 12-18.8 28.3-18.8 45.3 0 35.3 28.7 64 64 64h43.4V908c0 17.7 14.3 32 32 32H448V716h112v224h265.9c17.7 0 32-14.3 32-32V614.3h43.4c17 0 33.3-6.7 45.3-18.8 24.9-25 24.9-65.5-.1-90.5z"></path></svg></span></a><a href="https://github.com/vulhub/" target="_blank"><span role="img" aria-label="github" class="anticon anticon-github icon-style icon-github-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="github" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M511.6 76.3C264.3 76.2 64 276.4 64 523.5 64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c30.9-15.9 62.4 4 98.9 57.9 26.4 39.1 77.9 32.5 104 26 5.7-23.5 17.9-44.5 34.7-60.8-140.6-25.2-199.2-111-199.2-213 0-49.5 16.3-95 48.3-131.7-20.4-60.5 1.9-112.3 4.9-120 58.1-5.2 118.5 41.6 123.2 45.3 33-8.9 70.7-13.6 112.9-13.6 42.4 0 80.2 4.9 113.5 13.9 11.3-8.6 67.3-48.8 121.3-43.9 2.9 7.7 24.7 58.3 5.5 118 32.4 36.8 48.9 82.7 48.9 132.3 0 102.2-59 188.1-200 212.9a127.5 127.5 0 0138.1 91v112.5c.8 9 0 17.9 15 17.9 177.1-59.7 304.6-227 304.6-424.1 0-247.2-200.4-447.3-447.5-447.3z"></path></svg></span></a></div></div></div><div class="thanksfor-member-opt"><div class="member-info"><div class="member-info-img"><img src="/img/team/cnss.png" class="img-style"></div><div class="member-info-name-tag"><div class="name-style">CNSS</div></div></div><div class="tag-div-style orange-tag-div"><div class="tag-styl">idea inspired</div></div><div class="member-description"><div class="description-style">优秀的 CTF 战队</div><div class="description-style">为 Yak 安全能力构建提供灵感</div></div><div class="member-address-link"><div class="address-body"><span role="img" class="anticon address-icon"><svg width="1em" height="1em" viewBox="0 0 24 24" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><title>定位</title><defs><polygon id="path-1" points="0 0 24 0 24 24 0 24"></polygon></defs><g id="页面-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"><g id="官网/团队致谢" transform="translate(-404.000000, -1358.000000)"><g id="卡片/贡献者备份-4" transform="translate(384.000000, 1212.000000)"><g id="定位" transform="translate(20.000000, 146.000000)"><mask id="mask-2" fill="white"><use xlink:href="#path-1"></use></mask><g id="Clip-2"></g><g id="定位-4" mask="url(#mask-2)" fill="#C3C7CF"><g transform="translate(5.500000, 4.000000)"><path d="M6.50716394,1.2605942 C9.41610589,1.2605942 11.7619609,3.51620027 11.7619609,6.28547981 C11.7505395,7.30942305 11.4228459,8.3037731 10.8253089,9.12756049 L10.7597516,9.23156211 C10.7671802,9.22087831 10.7756302,9.21104542 10.7848231,9.20187437 L10.7398802,9.25359154 L6.49323534,14.739488 L2.24659053,9.25359154 L2.22244762,9.22456564 L2.17629752,9.15025175 L2.21251188,9.21322 L2.20944759,9.21009996 L2.2152976,9.21832554 C1.56000334,8.29488569 1.2381598,7.31566315 1.2381598,6.28547981 C1.25050982,3.51052746 3.59051484,1.2605942 6.50716394,1.2605942 M6.50716394,0 C2.92287055,0 0.0157857481,2.79537449 0,6.28226522 L0.00157857481,6.42730021 C0.0294357774,7.63050442 0.405508012,8.77395497 1.07993089,9.75591937 L1.15273104,9.85963735 L1.18578826,9.91419093 L1.21643118,9.96080257 C1.23723122,9.99077394 1.25914556,10.019138 1.28365989,10.0471239 L1.28923133,10.0534586 L5.52101897,15.5198783 C5.58128339,15.597785 5.65036925,15.6682225 5.7268837,15.7297725 L5.7812981,15.7707113 C6.31912783,16.1555173 7.05845798,16.0453702 7.46545171,15.5198783 L11.6972394,10.0534586 L11.667525,10.0862663 C11.7201751,10.0314291 11.7603823,9.97715919 11.800311,9.91447457 L11.8393111,9.8504663 C12.5735341,8.83314135 12.9846135,7.59712935 13.0000279,6.29332721 C13.0000279,2.80360007 10.0843073,0 6.50716394,0" id="Fill-1"></path><path d="M6.50004179,4.8199739 C7.35479362,4.81959587 8.04797367,5.52482141 8.04834525,6.39503134 C8.04871653,7.26533582 7.35609362,7.97112864 6.50143465,7.97150683 L6.50004179,7.97150683 C5.64528995,7.97150683 4.95238847,7.26599765 4.95238847,6.39578771 C4.95238847,5.52548324 5.64528995,4.8199739 6.50004179,4.8199739 M6.50004179,3.55937986 C4.96148849,3.55937986 3.71432153,4.82923965 3.71432153,6.39578771 C3.71432153,7.96224123 4.96148849,9.23219557 6.50004179,9.23219557 C8.03859508,9.23219557 9.28576204,7.96224123 9.28576204,6.39578771 C9.28576204,4.82923965 8.03859508,3.55937986 6.50004179,3.55937986" id="Fill-3"></path></g></g></g></g></g></g></svg></span><span class="address-style">四川 - 成都</span></div><div class="link-body"><a class="link-jump" href="https://cnss.io/" target="_blank"><span role="img" aria-label="home" class="anticon anticon-home icon-style icon-home-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="home" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M946.5 505L534.6 93.4a31.93 31.93 0 00-45.2 0L77.5 505c-12 12-18.8 28.3-18.8 45.3 0 35.3 28.7 64 64 64h43.4V908c0 17.7 14.3 32 32 32H448V716h112v224h265.9c17.7 0 32-14.3 32-32V614.3h43.4c17 0 33.3-6.7 45.3-18.8 24.9-25 24.9-65.5-.1-90.5z"></path></svg></span></a><a href="https://github.com/cnss/" target="_blank"><span role="img" aria-label="github" class="anticon anticon-github icon-style icon-github-style"><svg viewBox="64 64 896 896" focusable="false" data-icon="github" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M511.6 76.3C264.3 76.2 64 276.4 64 523.5 64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c30.9-15.9 62.4 4 98.9 57.9 26.4 39.1 77.9 32.5 104 26 5.7-23.5 17.9-44.5 34.7-60.8-140.6-25.2-199.2-111-199.2-213 0-49.5 16.3-95 48.3-131.7-20.4-60.5 1.9-112.3 4.9-120 58.1-5.2 118.5 41.6 123.2 45.3 33-8.9 70.7-13.6 112.9-13.6 42.4 0 80.2 4.9 113.5 13.9 11.3-8.6 67.3-48.8 121.3-43.9 2.9 7.7 24.7 58.3 5.5 118 32.4 36.8 48.9 82.7 48.9 132.3 0 102.2-59 188.1-200 212.9a127.5 127.5 0 0138.1 91v112.5c.8 9 0 17.9 15 17.9 177.1-59.7 304.6-227 304.6-424.1 0-247.2-200.4-447.3-447.5-447.3z"></path></svg></span></a></div></div></div></div></div>

<img src="imgs/lab-logo.png" style="width: 180px"/>

### 基础理论学科

1. Alonzo Church, "A set of postulates for the foundation of logic", Annals of Mathematics, 33(2), 346-366, 1932.
2. Dana Scott, Christopher Strachey, "Toward a mathematical semantics for computer languages", Proceedings of the Symposium on Computers and Automata, Microwave Research Institute Symposia Series Vol. 21, New York, 1971.
3. Henk Barendregt, Wil Dekkers, Richard Statman, lambda Calculus with Types, Perspectives in Logic. Cambridge University Press, 2013.

### 工程技术

1. Terence Parr, "The Definitive ANTLR 4 Reference", Pragmatic Bookshelf, 2013.
2. Terence Parr, "Simplifying Complex Networks Using Temporal Pattern Mining: The Case of AT&T's Observed Data Network", Dissertation, 1995.
3. Terence Parr, Russell Quong, "ANTLR: A Predicated-LL(k) Parser Generator", Journal of Software Practice and Experience, July 1995.
4. Google Ins, "Protocol Buffers", https://developers.google.com/protocol-buffers, 2020.
5. Google Ins, "gRPC", https://grpc.io/, 2020.
6. Microsoft Ins, "Monaco Editor", https://microsoft.github.io/monaco-editor/, 2020.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=yaklang/yakit&type=Date)](https://star-history.com/#yaklang/yakit&Date)

