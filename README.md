# このリポジトリについて

DB: Neon
Migration: drizzle, Prisma
ORM: gorm, sqlx, sqlc
を使用して比較するためのリポジトリです．

# CLI

CLI で何が出来るかをまとめます．

## prisma

```shell
>>> prisma --help

◭  Prisma is a modern DB toolkit to query, migrate and model your database (https://prisma.io)

Usage

  $ prisma [command]

Commands

            init   Set up Prisma for your app
        generate   Generate artifacts (e.g. Prisma Client)
              db   Manage your database schema and lifecycle
         migrate   Migrate your database
          studio   Browse your data with Prisma Studio
        validate   Validate your Prisma schema
          format   Format your Prisma schema
         version   Displays Prisma version info
           debug   Displays Prisma debug info

Flags

     --preview-feature   Run Preview Prisma commands
     --help, -h          Show additional information about a command

Examples

  Set up a new Prisma project
  $ prisma init

  Generate artifacts (e.g. Prisma Client)
  $ prisma generate

  Browse your data
  $ prisma studio

  Create migrations from your Prisma schema, apply them to the database, generate artifacts (e.g. Prisma Client)
  $ prisma migrate dev

  Pull the schema from an existing database, updating the Prisma schema
  $ prisma db pull

  Push the Prisma schema state to the database
  $ prisma db push

  Validate your Prisma schema
  $ prisma validate

  Format your Prisma schema
  $ prisma format

  Display Prisma version info
  $ prisma version

  Display Prisma debug info
  $ prisma debug
```

## drizzle

```shell
>>> drizzle-kit --help
Usage: drizzle-kit [options] [command]

Options:
  --version, -v                output the version number
  -h, --help                   display help for command

Commands:
  generate:pg [options]
  generate:mysql [options]
  generate:sqlite [options]
  check:pg [options]
  check:sqlite [options]
  up:pg [options]
  up:mysql [options]
  up:sqlite [options]
  introspect:pg [options]
  introspect:mysql [options]
  drop [options]
  push:mysql [options]
  introspect:sqlite [options]
  push:sqlite [options]
  check:mysql [options]
  push:pg [options]
  studio [options]
  help [command]               display help for command
```

# ORM

CRUD 処理とテーブル JOIN を試し ORM 特有の書き心地と Raw SQL の使い心地を検証します．
ドキュメントの分かりやすさ．充実度も確認します．

## gorm

ドキュメントは分かりづらい．
`user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}`
の User どっから来たんだなど

説明も最低限で薄い気がする．

そして結局，色々触ったがうまく動かせなかった．無念…

## sqlx

## sqlc
