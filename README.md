# G3

Gist Storage Service CLI. Just like AWS S3, G3 is lightweight and flexible
CLI that provides an easy way to store and access data with the reliability
of GitHub’s infrastructure. G3 is perfect for developers who need a quick,
reliable, and Git-native way to manage small to medium-sized files without
setting up complex storage systems.

## Limitation

Since G3 relies in Github Gist API, do not use for really heavy files.
Gist limits the file size for a file in a Gist to 100MB. There is not a
strict limit for the number of files in a gist but there are some limits
for hitting the Github API per key

## Installation

## Configuration

```json
{
  "GITHUB_TOKEN": "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
  "AES_KEY": "A37F1C5D92E46B88F021BD6734FAC95E127A45EF983DBC7609DD4F602A9183BF"
}
```

- `GITHUB_TOKEN`: A personal access token (PAT) used to authenticate requests
  to the GitHub API. This token should have the necessary permissions to
  read and write gists.

- `AES_KEY`: A 32-byte encryption key used for AES-256 encryption and
  decryption of stored data. This key ensures that sensitive information
  remains secure. Keep this key private and do not share it.

## Usage

```sh
g3 <command> [options]

Commands:
  g3 ls          Show all the files in your storage
  g3 cp <file>   Copy a file to your storage. If the file exists, it will be
                 overwritten.
  g3 get <file>  Get file from your storage
  g3 rm <file>   Delete file from your storage

Options:
      --version  Show version number                                   [boolean]
  -v, --verbose  Show debug logging                   [boolean] [default: false]
  -c, --config   Path to config file
                         [string] [default: "/home/jegj/.config/g3/config.json"]
  -h, --help     Show help                                             [boolean]

```
