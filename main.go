package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// === ツール1: hello_world ===
	helloTool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)
	s.AddTool(helloTool, helloHandler)

	// === ツール2: list_directory ===
	listTool := mcp.NewTool("list_directory",
		mcp.WithDescription("List files and directories in the specified path"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Absolute path to the directory"),
		),
	)
	s.AddTool(listTool, listDirectoryHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// Handler for hello_world
func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

// Handler for list_directory
func listDirectoryHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := req.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError("Missing path"), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to read directory: %v", err)), nil
	}

	var result string
	for _, entry := range entries {
		kind := "📄"
		if entry.IsDir() {
			kind = "📁"
		}
		result += fmt.Sprintf("%s %s\n", kind, entry.Name())
	}

	if result == "" {
		result = "(empty)"
	}

	return mcp.NewToolResultText(result), nil
}
