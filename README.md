# gator
Blog aggregator following the guided project on Boot.Dev

# Requirements
* Go 1.24.3 (or later)
* PostgreSQL 16.9 (or later)
* goose 3.24.3 (or later)

# Installation

These instructions assume you are on Linux/WSL.

1. Install Postgres v16 or later. For example, on Linux/WSL:

```
sudo apt udpate
sudo apt install postgresql postgresql-contrib
```

2. Update the postgres password:

```
sudo passwd postgres
```

3. Start Postgres server in the background:

```
sudo service postgresql start
```

4. Create the database and update the user password:

```
sudo -u postgres psql -c 'CREATE DATABASE gator;'
sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"
```

5. If you've followed these steps to the letter, your database connection URL should be:

```
postgres://postgres:postgres@localhost:5432/gator
```

If not, you'll have to figure out your connection URL for your Postgres configuration. The general structure is:

```
protocol://username:password@host:port/database
```

6. Clone the repo, `cd` into the `sql\schmea` directory and run:

```
goose postgres <connection_URL> up
```

5 times (once for each `.sql` file in the directory). 

7. Add the following config JSON to your home directory as `~/.gatorconfig/json`:

```
{
    "db_url": <connection_URL>,
    "current_user_name": ""
}
```

8. Install gator:

```
go install github.com/pbojar/gator
```

# Usage

gator supports the following commands:

* `gator register <username>`: Registers user.
* `gator login <username>`: Logs user in.
* `gator users`: Lists users.
* `gator reset`: Resets database! Use with care!
* `gator addfeed <name> <url>`: Adds feed to db and follows feed for current user.
* `gator following`: Lists feeds followed by current user.
* `gator follow <url>`: Follows feed for current user.
* `gator unfollow <url>`: Unfollows feed for current user.
* `gator agg <time_between_requests>`: Fetches feeds every `<time_between_requests>`.
* `gator browse [limit]`: Displays `[limit]` (defaults to 2) items from latest feeds fetched.
