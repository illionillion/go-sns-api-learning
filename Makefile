# .envを読み込む
include .env
export

.PHONY: dev migrate

# マイグレーション実行
migrate:
	go run migrations/migrations.go

# Air で開発サーバー起動（マイグレーション付き）
dev: migrate
	air -c .air.toml
