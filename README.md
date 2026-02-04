# gator

Feed aggreGATOR

Repo for https://www.boot.dev/courses/build-blog-aggregator-golang

## Installation

1. Install PostgreSQL
1. Create a database named `gator` on your PostgreSQL server
1. Put the connection string in `~/.gatorconfig.json` in the following format
   ```json
   {"db_url":"postgres://username:password@server:5432/gator?sslmode=disable"}
   ```
1. Install Go
1. Install `gator` by running
   ```shell
   go install github.com/WadeGulbrandsen/gator@latest
   ```

## Commands
| Command                                | Description |
| :------------------------------------- | :---------- |
| `gator register <name>`                | Register and login a user with `name` |
| `gator login <name>`                   | Login as user with `name` |
| `gator users`                          | List all users on the system |
| `gator addfeed <feed_name> <feed_url>` | Current user adds and follows a new feed |
| `gator feeds`                          | List all feeds |
| `gator follow <feed_url>`              | Current user follows an existing feed |
| `gator unfollow <feed_url>`            | Current user unfollow a feed |
| `gator following`                      | List the feeds that the current user is following |
| `gator agg <duration>`                 | Fetch posts for feeds. Waiting `duration` between fetch requests. `duration` is specified like `1m`, `5s`, etc. Press `CTRL+C` to exit |
| `gator browse [num_posts]`             | Current user see the `num_posts` most recent posts from the feeds they are following. `num_posts` defaults to `2` |

Options in triangular brackets are required and options in square brackets are optional.
If the value has spaces enclose the entire value in quotes.

```shell
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"
```

## Reset the Database
To delete everthing in the database without confirmation run `gator reset`.
This really will delete everything in the database. You have been warned.