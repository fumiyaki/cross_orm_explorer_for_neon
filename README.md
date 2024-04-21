# このリポジトリについて

Neon を使用した drizzle と Prisma を比較するためのリポジトリです．

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
