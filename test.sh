#!/bin/bash

echo "启动 GoFrame 应用..."
./goframe-app &
SERVER_PID=$!

# 等待服务器启动
sleep 2

echo "测试 Hello API..."
echo "1. 测试无参数请求:"
curl -s "http://localhost:8080/api/v1/hello" | jq .

echo -e "\n2. 测试带参数请求:"
curl -s "http://localhost:8080/api/v1/hello?name=GoFrame" | jq .

echo -e "\n3. 检查 OpenAPI 文档:"
echo "Swagger UI: http://localhost:8080/swagger/"
echo "OpenAPI JSON: http://localhost:8080/api.json"

# 停止服务器
echo -e "\n停止服务器..."
kill $SERVER_PID
wait $SERVER_PID 2>/dev/null

echo "测试完成！"