# serveIt

This is a simple HTTPS server that serves a static files from a directory. 

## Setup

Download the binary from the [releases](https://github.com/Elixir-Craft/servefiles/releases) page or build it from source.

```bash
git clone https://github.com/Elixir-Craft/servefiles.git
cd servefiles
```

Build the binary
```bash
go build -o ./build ./servefiles
```

Run the binary or copy it to your bin directory to run it from anywhere.

Run 
```bash
./build/servefiles
```

Copy binary to bin directory (Linux)

```bash
cp ./build/servefiles /usr/local/bin
```

## Usage

Change port 
```bash
servefiles -p 8080
```

Regenerate certificate
```bash
servefiles -r
```

Set password
```bash
servefiles -P password
```

To set configuration file for certificate, create a file named `config.yaml` in the servefile configuration directory. 
on Linux, the configuration directory is `~/.ServeFiles/` 
    
```yaml
cert:
    organization: "Your Company Name"
    country: "Your Country"
    province: "Your Province"
    locality: "Your City"
    street_address: "Your Street Address"
    postal_code: "Your Postal Code"
```


## Development

Run 

```bash
docker compose up
```
or

```bash
air
```



