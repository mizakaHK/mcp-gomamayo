# mcp-gomamayo

<!-- バッジ (必要に応じて実際のバッジURLに置き換えてください) -->
[![Go Version](https://img.shields.io/badge/go-v1.24-blue.svg
)](https://golang.org/)
[![MCP](https://img.shields.io/badge/MCP-Server-6236FF?style=flat&logo=claude&logoColor=white)](https://github.com/modelcontextprotocol/mcp)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

テキスト中の「ゴママヨ」構文をチェックするための [Model Context Protocol (MCP)](https://github.com/modelcontextprotocol/specification) サーバーです。 

## 目次

- [セットアップ](#セットアップ)
- [Claude Desktop での利用](#claude-desktop-での利用)
  - [MCPサーバー設定](#mcpサーバー設定)
- [使い方](#使い方)

## セットアップ

```bash
git clone https://github.com/mizakahk/mcp-gomamayo.git
cd mcp-gomamayo
task build
```

## Claude Desktop での利用

### MCPサーバー設定

`claude_desktop_config.json` に以下の設定を追加します

1.  **WSL上で起動:**
    ```json
    {
      "mcpServers": {
          "mcp-gomamayo": {
              "command": "wsl",
              "args": [
                  "bash",
                  "-ic",
                  "~/mcp-gomamayo/bin/mcp-gomamayo"
              ],
              "env": {}
          }
      }
    }
    ```

設定ファイルを保存した後、Claude Desktop を再起動すると、`gomamayo-checker` が利用可能になります。

## 使い方

MCP サーバーが Claude Desktop で有効になると、`check_gomamayo` ツールを使用してテキストのゴママヨ構文をチェックできます。

**例 「福山雅治」の判定:**

![akakara](screenshot/fukuyamamasaharu.png)
