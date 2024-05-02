# SSH-Cloud-Agent

Create a `servers` file at the root of the directory and add a list of ssh connections like: `ssh user@host`

Then in the files/ folder add all the files you wish to distribute to the servers.

Edit `up.sh` with any command you wish to run...


# Running

Run using `go main.go`, this will copy all the files over and execute `up.sh`.
