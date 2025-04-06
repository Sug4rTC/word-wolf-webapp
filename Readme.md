# Word Wolf 人狼 Webアプリ

## 概要
このプロジェクトは、オンラインでプレイ可能なワードウルフ人狼ゲームのWebアプリケーションです。

## 特徴
- 個別のお題表示機能：各ユーザーに対して異なるお題を表示
- リアルタイム通信：WebSocketを利用した即時同期
- ホットリロード：React + Vite により高速な開発体験
- コンテナ環境：Docker Compose により、フロントエンドとバックエンドが統合管理

## 技術スタック
- フロントエンド: React, TypeScript, Vite, Tailwind CSS
- バックエンド: Go, gorilla/websocket, Redis（予定）
- コンテナ: Docker, Docker Compose
- インフラ（将来的に）: AWS ECS (Fargate) など

## セットアップ
1. リポジトリをクローン
   ```bash
   git clone https://github.com/yourusername/word-wolf-webapp.git

2. プロジェクトルートに移動
   ```bach
   cd word-wolf-webapp

3. Docker Compose を使用して環境を起動
   ```bash
   docker-compose up --build

