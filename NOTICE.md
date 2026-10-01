# NOTICE / 归属声明

Vigile 是 **Beszel** 的衍生作品（derivative work），基于 MIT 许可证使用。

## 上游项目

- 项目：Beszel
- 仓库：https://github.com/henrygd/beszel
- 作者：henrygd
- 许可证：MIT
- 本衍生作品基于的上游提交：`5b0952ff`

上游的 MIT 许可证与版权声明已**原样保留**在根目录 `LICENSE` 文件中，
其条款要求在所有副本或实质性部分中保留该版权声明与许可声明。

## 本衍生作品所做的修改

由 **Akilaea** 维护，主要改动包括：

1. **重命名与品牌化**
   - Go module 路径：`github.com/henrygd/beszel` → `github.com/Akilaea/vigile`
   - 产品名与二进制名：`beszel` / `beszel-agent` → `vigile` / `vigile-agent`
   - 数据目录：`beszel_data` → `vigile_data`
   - 前端全局对象：`globalThis.BESZEL` → `globalThis.VIGILE`
   - 环境变量前缀：`BESZEL_AGENT_*` / `BESZEL_HUB_*` → `VIGILE_AGENT_*` / `VIGILE_HUB_*`
   - 后端 API 路径：`/api/beszel/*` → `/api/vigile/*`
   - 界面文案与多语言词条

2. **功能性调整**
   - 自动更新指向改为 `Akilaea/vigile`
   - 页脚仓库链接改为本项目仓库

3. **后续计划的功能扩展**（详见 README）

## 有意保留的上游引用

以下内容**刻意未作修改**，以保证功能可用性与出处可追溯：

| 保留项 | 原因 |
|---|---|
| `beszel.dev` / `gh.beszel.dev` / `get.beszel.dev` / `static.beszel.dev` | 上游更新镜像与文档服务端点，改动会导致功能失效 |
| 源码注释中的 `github.com/henrygd/beszel/issues/NNNN` 链接 | 记录代码为何如此实现，改动即等于伪造出处 |
| `supplemental/CHANGELOG.md` 等变更日志 | 上游历史记录 |
| Docker 镜像名 `henrygd/beszel-agent`（安装说明中） | 本项目尚未发布自有镜像 |

## 免责声明

本衍生作品与上游作者 henrygd 无隶属关系，亦未获得其背书。
上游作者不对本衍生作品提供任何支持、担保或维护。

## 许可证

本衍生作品同样以 **MIT 许可证**发布，完整条款见根目录 `LICENSE`。
