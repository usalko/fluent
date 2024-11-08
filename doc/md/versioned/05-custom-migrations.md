---
title: Custom migrations
id: custom-migrations
---
:::info Supporting repository

The change described in this section can be found in
[PR #7](https://github.com/rotemtam/fluent-versioned-migrations-demo/pull/7/files)
in the supporting repository.

:::

## Custom migrations
In some cases, you may want to write custom migrations that are not automatically
generated by Atlas. This can be useful in cases where you want to perform changes
to your database that aren't currently supported by Ent, or if you want to seed
the database with data. 

In this section, we will learn how to add custom migrations to our project. For the
purpose of this guide, let's assume we want to seed the users table with some data.

## Create a custom migration

Let's start by adding a new migration file to our project:

```shell
atlas migrate new seed_users --dir file://fluent/migrate/migrations
```

Observe that a new file named `20221115102552_seed_users.sql` was created in the
`ent/migrate/migrations` directory. 

Continue by opening the file and adding the following SQL statements:

```sql
INSERT INTO `users` (`name`, `email`, `title`)
VALUES ('Jerry Seinfeld', 'jerry@seinfeld.io', 'Mr.'),
       ('George Costanza', 'george@costanza.io', 'Mr.')
```

## Recalculating the checksum file

Let's try to run our new custom migration:

```shell
atlas migrate apply --dir file://fluent/migrate/migrations --url mysql://root:pass@localhost:3306/db
```
Atlas fails with an error:
```text
You have a checksum error in your migration directory.
This happens if you manually create or edit a migration file.
Please check your migration files and run

'atlas migrate hash'

to re-hash the contents and resolve the error

Error: checksum mismatch
```
Atlas introduces the concept of [migration directory integrity](https://atlasgo.io/concepts/migration-directory-integrity)
as a means to enforce a linear migration history. This way, if multiple developers work on the 
same project in parallel, they can be sure that their merged migration history is correct.

Let's re-hash the contents of our migration directory to resolve the error:

```shell
atlas migrate hash --dir file://fluent/migrate/migrations
```

If we run `atlas migrate apply` again, we will see that the migration was successfully applied:
```text
atlas migrate apply --dir file://fluent/migrate/migrations --url mysql://root:pass@localhost:3306/db
```
Atlas reports:
```text
Migrating to version 20221115102552 from 20221115101649 (1 migrations in total):

  -- migrating version 20221115102552
    -> INSERT INTO `users` (`name`, `email`, `title`)
VALUES ('Jerry Seinfeld', 'jerry@seinfeld.io', 'Mr.'),
       ('George Costanza', 'george@costanza.io', 'Mr.')
  -- ok (9.077102ms)

  -------------------------
  -- 19.857555ms
  -- 1 migrations 
  -- 1 sql statements
```

In the next section, we will learn how to automatically verify the safety of our
schema migrations using Atlas's [Linting](https://atlasgo.io/versioned/lint) feature.