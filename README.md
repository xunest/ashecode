# AsheCode

> 运行在终端里的 AI 编程智能体,Go 语言编写。

AsheCode 把智能体化开发工作流带进你的命令行:基于 Bubble Tea 的 TUI 界面、与 Anthropic / OpenAI / OpenAI 兼容模型的流式对话,以及一整套面向真实工程任务的工具——文件编辑、Bash 执行、代码搜索、Diff 审阅等。

## 目录

- [功能特性](#功能特性)
- [环境要求](#环境要求)
- [安装](#安装)
- [快速开始](#快速开始)
- [使用方法](#使用方法)
- [配置](#配置)
- [许可证](#许可证)

## 功能特性

- 🤖 **多模型接入**:Anthropic、OpenAI 及 OpenAI 兼容端点
- 🔌 **MCP 支持**:模型上下文协议客户端,可接入外部工具服务器
- 🧩 **技能系统**:内置技能、技能目录与安装器
- 👥 **子智能体与智能体团队**:协调者模式,支持 tmux / iTerm 后端
- 🛡️ **安全隔离**:操作系统级沙箱(macOS Seatbelt / Linux)与细粒度权限控制
- 🗂️ **Git Worktree 集成**:任务并行互不干扰
- 🧠 **上下文管理**:长期记忆、会话持久化与上下文自动压缩
- 🪝 **钩子与计划模式**:生命周期钩子,Plan Mode 让改动可控、可审阅
- 🌐 **远程与脚本化**:WebSocket 服务 + Web 界面;非交互式 `-p` 打印模式

## 环境要求

- Go 1.25+
- macOS 或 Linux;沙箱在 macOS 基于系统自带的 Seatbelt(`sandbox-exec`),在 Linux 需要安装 bubblewrap(`bwrap`),其他平台自动禁用沙箱
- 至少一个模型服务的 API Key(`ANTHROPIC_API_KEY` 或 `OPENAI_API_KEY`)

## 安装

```bash
git clone https://github.com/xunest/ashecode.git
cd ashecode
go build -o ashecode ./cmd/ashecode
```

将生成的 `ashecode` 可执行文件放入 `PATH`,例如:

```bash
go install ./cmd/ashecode
```

## 快速开始

1. 在项目根目录(或用户目录)创建配置:

```bash
mkdir -p ~/.ashecode
cat > ~/.ashecode/config.yaml <<'EOF'
providers:
  - name: claude
    protocol: anthropic
    base_url: https://api.anthropic.com
    model: claude-sonnet-4
  - name: gpt
    protocol: openai
    base_url: https://api.openai.com/v1
    model: gpt-4.1
EOF
```

2. 设置 API Key 并启动:

```bash
export ANTHROPIC_API_KEY=sk-ant-...
ashecode
```

## 使用方法

### 交互模式(TUI)

直接运行,进入终端对话界面:

```bash
ashecode
```

### 非交互模式(`-p` / `--print`)

执行单次任务并输出结果,适合脚本与 CI:

```bash
ashecode -p "解释这个项目的构建流程"
echo "修复 main.go 中的拼写错误" | ashecode -p
ashecode -p "运行测试并总结失败原因" --output-format stream-json
```

### 远程模式(`--remote`)

启动 WebSocket 服务,通过 Web 界面远程使用:

```bash
ashecode --remote :18888
```

## 配置

AsheCode 按以下顺序读取配置文件并逐层合并(后者覆盖前者):

1. `~/.ashecode/config.yaml`(用户级)
2. `<项目>/.ashecode/config.yaml`(项目级)
3. `<项目>/.ashecode/config.local.yaml`(本地覆盖,建议加入 .gitignore)

API Key 可在 provider 的 `api_key` 字段直接配置;未配置时按协议读取环境变量:`anthropic` → `ANTHROPIC_API_KEY`,`openai` / `openai-compat` → `OPENAI_API_KEY`。

配置项涵盖:模型提供商、MCP 服务器、钩子(Hooks)、沙箱(Sandbox)、协调者模式等,详见 `internal/config`。

## 许可证

暂未添加许可证文件。开源发布前建议在仓库根目录添加 `LICENSE`(如 MIT)。
