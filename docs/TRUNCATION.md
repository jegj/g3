# Trucantion

## limitations

Per [Gist API](https://docs.github.com/en/rest/gists/gists?apiVersion=2022-11-28#truncation)
there are some limitations regarding files content and number:

- Each file inside a Gist can be returned by the API in the
`content` property if its content is up to 1MB.

- For bigger files, the property `truncated` will be true and
the only way to get the content is through `raw_url`.
Be aware that for files larger than 10MB the only property
available for get the content is `git_pull_url`

- Also, the entire list of files in the Gist can be truncated
if the total number exceeds 300 files. Again the only possible
way to get the whole content is using `git_pull_url`

## Approach

For use `raw_url` I need to consider

```txt
300 files × 10 MB/file = 3,000 MB (3 GB)
```

I can use the limit above to split the file in two categories:

### For files less than 3GB

1. Use only one gist and split the file in chunks of 10MB

### For files bigger than 3GB

1. Use more than one gist and split in chunks of 10MB. 
2. When file count is more than 300, use a new gist
