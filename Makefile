# .envを読み込む
include .env
export

.PHONY: dev migrate swag

# マイグレーション実行
migrate:
	go run migrations/migrations.go

# Swagger ドキュメント生成
swag:
	swag init -g cmd/main.go --parseDependency --parseInternal

# 開発サーバー起動（migrate -> swag -> air）
dev: migrate swag
	air -c .air.toml
