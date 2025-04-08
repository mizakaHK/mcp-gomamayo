package main

import (
    "context"
    "errors"
    "fmt"

    "github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
	"github.com/jugesuke/gomamayo"
)

func main() {
    // Create MCP server
    s := server.NewMCPServer(
        "Demo 🚀",
        "1.0.0",
    )

    // Add tool
    tool := mcp.NewTool("check_gomamayo",
        mcp.WithDescription("Check gomamayo in text"),
        mcp.WithString("text",
            mcp.Required(),
            mcp.Description("Text to check gomamayo in"),
        ),
    )

    // Add tool handler
    s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		text, ok := request.Params.Arguments["text"].(string)
		if !ok {
			return nil, errors.New("text must be a string")
		}

		g, err := gomamayo.Init()
		if err != nil {
			panic(err)
		}

		result := g.Analyze(text)

        var resultText string
        if result.IsGomamayo {
            resultText = fmt.Sprintf("%d項%d次のゴママヨです\n", result.Ary, result.Degree)
        } else {
            resultText = fmt.Sprintf("ゴママヨではありません")  
        }
	
		return mcp.NewToolResultText(resultText), nil
	})

    // Start the stdio server
    if err := server.ServeStdio(s); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}
